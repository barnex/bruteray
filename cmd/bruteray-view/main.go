// Package x provides an interactive viewer for BruteRay.
package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/barnex/bruteray/api"
	"github.com/barnex/bruteray/geom"
	imagef "github.com/barnex/bruteray/image"

	. "github.com/barnex/bruteray/tracer/types"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

var (
	flagAddr = flag.String("addr", "localhost:37273", "HTTP server address")
)

func main() {
	flag.Parse()

	s := &state{
		view: api.View{
			Width:  32,
			Height: 32,
		},
		addr:        *flagAddr,
		bakery:      make(chan imagef.Image),
		lastImg:     dummyImage(),
		windowDirty: true,
		renderDirty: true,
		winSize:     image.Pt(32, 32), // dummy
		abort:       make(chan struct{}),
	}

	s.initWindow()
	s.run()
}

const maxDownScale = 8

type state struct {
	view api.View
	addr string

	// modified by user events (fast)
	mouseDown              mouse.Button
	lastMouseX, lastMouseY float32
	freeLooking            bool
	renderDirty            bool
	paused                 bool

	// background image refinement
	bakery         chan imagef.Image
	bakeResolution int
	isBaking       bool

	lastImg imagef.Image

	// modified, indirectly, by user events or renderer
	windowDirty bool // whether window is dirty, img needs to be drawn on window
	scr         screen.Screen
	win         screen.Window
	winEv       chan event
	winSize     image.Point
	winBuf      screen.Buffer
	winTex      screen.Texture

	// close to cleanup and return from Display()
	abort chan (struct{})
}

func (s *state) run() {
	for {
		runtime.Gosched() // give event loop the opportunity to send event

		// (1) handle user events,
		// but do not block if none are queued.
		select {
		case <-s.abort:
			return
		case e := <-s.winEv:
			s.handleEventBacklog(e)
		default:
		}
		s.repaintIfNeeded()

		// (2) there's no user events queued and we are not dragging the mouse:
		// start baking high-res image
		if !s.isBaking && !s.freeLooking && !s.renderDirty && !s.paused {
			s.goBakeAndSend()
		}

		// (3) wait for user event (which cancels baking)
		// OR baked image to become available. Blocking
		select {
		case <-s.abort:
			return
		case e := <-s.winEv:
			s.handleEventBacklog(e)
		case img := <-s.bakery:
			s.handleBakery(img)
		}
		s.repaintIfNeeded()
	}
}

func (s *state) handleAllInteractive() {
}

// ---- Rendering ----

func (s *state) repaintIfNeeded() {
	if s.renderDirty {
		s.cancelBaking()

		img := s.renderFreeLooking()
		fmt.Println("run: lastImg:", s.lastImg.Bounds())

		if img.Bounds() != image.ZR {
			s.lastImg = img
			s.renderDirty = false
			s.windowDirty = true
		}
	}

	if s.windowDirty {
		fmt.Println("run: repaint")
		s.repaint()
	}
}

func (s *state) renderFreeLooking() imagef.Image {
	defer trace()()

	//spec := s.modSpec
	//w := s.winSize.X / maxDownScale
	//h := s.winSize.Y / maxDownScale
	//return sampler.Uniform(spec.ImageFunc(), 1, w, h)

	downscaled := s.view // copy
	downscaled.Width /= maxDownScale
	downscaled.Height /= maxDownScale
	img, err := s.fetchHTTP(downscaled)
	if err != nil {
		logErr(err)
		//return imagef.MakeImage(32, 32) // dummy
		return nil
	}
	return img
}

func (s *state) handleBakery(img imagef.Image) {
	defer trace()()
	fmt.Println("run: <-bakery:", img.Bounds())

	if img.Bounds() != image.ZR && s.isBaking && !s.renderDirty && !s.freeLooking {
		s.lastImg = img
		s.windowDirty = true
	}
	s.isBaking = false
	s.bakeResolution++ // next one will be higher-res
}

func (s *state) cancelBaking() {
	defer trace()()
	if s.isBaking {
		resp, err := http.Get("http://" + s.addr + "/cancel")
		if err != nil {
			fmt.Println(err)
		}else{
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("status %v: %s\n", resp.StatusCode, readBody(resp.Body))
		}
	}
	}
	s.bakeResolution = 0
}

func (s *state) goBakeAndSend() {
	s.isBaking = true

	div := maxDownScale
	for i := 0; i < s.bakeResolution; i++ {
		div /= 2
		if div <= 1 {
			div = 1
			break
		}
	}

	v := s.view // copy
	v.Width = s.winSize.X / div
	v.Height = s.winSize.Y / div
	v.AntiAlias = (div == 1) // AA only at full resolution

	go func() {
		img, err := s.fetchHTTP(v)
		if err != nil {
			logErr(err)
			s.bakery <- nil
		} else {
			s.bakery <- img
		}
	}()
}

func (s *state) fetchHTTP(v api.View) (imagef.Image, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(v); err != nil {
		panic(err) // BUG
	}
	url := "http://" + s.addr + "/gob"
	req, err := http.NewRequest("GET", url, &body)
	if err != nil {
		panic(err) // BUG
	}
	resp, err := http.DefaultClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %v: %s", resp.StatusCode, readBody(resp.Body))
	}

	var img imagef.Image
	err = gob.NewDecoder(resp.Body).Decode(&img)
	if err != nil {
		return nil, err
	}
	if len(img) == 0 {
		img = nil // hack because gob decodes nil into empty slice
	}
	return img, nil
}

func readBody(r io.Reader) string {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return "error reading error response: " + err.Error()
	}
	return string(bytes)
}

//func (s *state) goBakeAndSendLocal() {
//	s.isBaking = true
//	s.cancelBakingCh = make(chan struct{})
//
//	div := maxDownScale
//	for i := 0; i < s.bakeResolution; i++ {
//		div /= 2
//		if div <= 1 {
//			div = 1
//			break
//		}
//	}
//
//	w := s.winSize.X / div
//	h := s.winSize.Y / div
//
//	if div == 1 {
//		if s.sampler == nil {
//			s.sampler = sampler.NewAdaptive(s.modSpec.ImageFunc(), w, h, true)
//		}
//		go func() {
//			s.sampler.SampleNumCPUWithCancel(runtime.NumCPU(), s.bakeResolution, s.cancelBakingCh)
//			s.bakery <- s.sampler.StoredImage()
//		}()
//	} else {
//		s.sampler = nil
//		go func() { s.bakery <- sampler.UniformWithCancel(s.modSpec.ImageFunc(), 1, w, h, s.cancelBakingCh) }()
//	}
//}

// ---- Event Handling ----

type event interface{}

// handleEventBacklog handles all currently queued events
// before returning
func (s *state) handleEventBacklog(e event) {
	s.handleEvent(e)
	for {
		runtime.Gosched()
		select {
		case e := <-s.winEv:
			s.handleEvent(e)
		default:
			return
		}
	}
}

func (s *state) handleEvent(e event) {
	switch e := e.(type) {
	case lifecycle.Event:
		s.handleLifecycleEvent(e)
	case size.Event:
		s.handleSizeEvent(e)
	case key.Event:
		s.handleKeyEvent(e)
	case mouse.Event:
		s.handleMouseEvent(e)
	case paint.Event:
		s.handlePaintEvent(e)
	}
}

func (s *state) handleKeyEvent(e key.Event) {
	fmt.Printf("%v, %#v\n", e, e)

	if e.Direction == key.DirPress {
		switch e.Code {
		case key.CodeLeftArrow, key.CodeS:
			s.handleMoveCam(Vec{-1, 0, 0}) // ?
		case key.CodeRightArrow, key.CodeF:
			s.handleMoveCam(Vec{1, 0, 0}) // ?
		case key.CodeUpArrow, key.CodeE:
			s.handleMoveCam(Vec{0, 0, -1})
		case key.CodeDownArrow, key.CodeD:
			s.handleMoveCam(Vec{0, 0, 1})
		case key.CodeSpacebar:
			s.handleMoveCam(Vec{0, 1, 0})
		case key.CodeZ:
			s.handleMoveCam(Vec{0, -1, 0})
		case key.CodeP:
			s.paused = !s.paused
		case key.CodeN:
			s.handleToggleNormals()
		}
	}
}

func (s *state) handleMouseEvent(e mouse.Event) {
	if e.Direction == mouse.DirPress {
		s.mouseDown = e.Button
		s.freeLooking = (e.Button == mouse.ButtonLeft)
		//fmt.Println("free looking:", s.freeLooking)
	}
	if e.Direction == mouse.DirRelease {
		s.mouseDown = 0
		if e.Button == mouse.ButtonLeft {
			s.freeLooking = false
		}
		//fmt.Println("free looking:", s.freeLooking)
	}
	if s.mouseDown == mouse.ButtonLeft {
		dx := e.X - s.lastMouseX
		dy := e.Y - s.lastMouseY
		s.handleLook(dx, dy)
	}

	s.lastMouseX = e.X
	s.lastMouseY = e.Y
}

func (s *state) handleToggleNormals() {
	s.view.DebugNormals = !s.view.DebugNormals
	s.renderDirty = true
}

func (s *state) handleSizeEvent(e size.Event) {
	newSize := image.Pt(e.WidthPx, e.HeightPx)
	// Release and remove buffers, but do not yet re-initialize them.
	// Re-initialization happens lazily on repaint.
	// Otherwise, resizing a window becomes too slow.
	if s.winSize != newSize {
		s.winSize = newSize
		if s.winBuf != nil {
			s.winBuf.Release()
		}
		s.winBuf = nil
		if s.winTex != nil {
			s.winTex.Release()
		}
		s.winTex = nil
	}
	s.view.Width = s.winSize.X
	s.view.Height = s.winSize.Y

	s.renderDirty = true
	s.windowDirty = true
}

func (s *state) handlePaintEvent(e paint.Event) {
	s.windowDirty = true
}

func (s *state) handleLook(dx, dy float32) {
	//defer trace(dx, dy)()
	const sens = 0.005
	s.view.CamYaw += float64(dx) * sens
	s.view.CamPitch += float64(dy) * sens

	// clamp camera pitch to +/- 90Deg so that we can't see the world upside down
	if s.view.CamPitch < -Pi/2 {
		s.view.CamPitch = -Pi / 2
	}
	if s.view.CamPitch > Pi/2 {
		s.view.CamPitch = Pi / 2
	}
	s.renderDirty = true
}

func (s *state) handleMoveCam(dir Vec) {
	//defer trace(dir)()

	v := &s.view
	const sens = 0.05
	viewDir := geom.YawPitchRoll(v.CamYaw, v.CamPitch, 0).A
	s.view.CamPos = v.CamPos.MAdd(sens, viewDir.MulVec(dir))
	s.renderDirty = true
}

// copy s.img to screen
func (s *state) repaint() {
	//fmt.Println("state:repaint")
	defer trace()()
	if s.winBuf == nil {
		s.initWinBufs()
	}

	buf := s.winBuf.RGBA()
	draw.Draw(buf, buf.Bounds(), s.lastImg, image.Pt(0, 0), draw.Src)

	s.winTex.Upload(image.Pt(0, 0), s.winBuf, s.winBuf.Bounds())
	s.win.Scale(s.winTex.Bounds(), s.winTex, s.lastImg.Bounds(), screen.Over, nil)
	s.win.Publish()
	s.windowDirty = false
}

func (s *state) initWinBufs() {
	defer trace()()
	s.winBuf = s.newBuffer(s.winSize)
	s.winTex = s.newTexture(s.winSize)
}

func (s *state) handleLifecycleEvent(e lifecycle.Event) {
	if e.To == lifecycle.StageDead {
		s.cancelBaking()
		close(s.abort)
	}
	s.windowDirty = true
}

func (s *state) initWindow() {
	done := make(chan struct{})
	go driver.Main(func(scr screen.Screen) {
		win, err := scr.NewWindow(&screen.NewWindowOptions{Title: "BruteRay"})
		check(err)

		s.scr = scr
		s.win = win
		s.winEv = make(chan event, 64) // buffering for handleBacklog to be usefull
		close(done)

		defer win.Release()
		for {
			s.winEv <- s.win.NextEvent()
			select {
			case <-s.abort:
				return
			default:
			}
		}
	})
	<-done
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (s *state) newBuffer(size image.Point) screen.Buffer {
	b, err := s.scr.NewBuffer(size)
	check(err)
	// new buffer contains rubbish, clear it.
	draw.Draw(b.RGBA(), b.Bounds(), &image.Uniform{color.Black}, image.Pt(0, 0), draw.Src)
	return b
}

func (s *state) newTexture(size image.Point) screen.Texture {
	t, err := s.scr.NewTexture(size)
	check(err)
	return t
}

var cnt int32

func trace(args ...interface{}) func() {
	cnt := atomic.AddInt32(&cnt, 1)
	pc, _, _, _ := runtime.Caller(1)
	name := path.Ext(runtime.FuncForPC(pc).Name())
	fmt.Println(cnt, " >> ", name, args)
	start := time.Now()
	return func() {
		fmt.Println(cnt, " << ", time.Since(start).Round(100*time.Microsecond))
	}
}

func nop() {}

func logErr(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func dummyImage() imagef.Image {
	img := imagef.MakeImage(32, 32)
	for i := range img {
		for j := range img[i] {
			img[i][j].B = 1
		}
	}
	return img
}
