// Package x provides an interactive viewer for BruteRay.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

// X opens a window that renders a scene specification
//func X(s *Spec) {
func main() {
	driver.Main(func(s screen.Screen) {
		w := NewWindow(s, "BruteRay")
		defer w.Release()

		var px, py int
		onMouse := func(e mouse.Event) {
			px = int(e.X)
			py = int(e.Y)
			w.Send(paint.Event{})
		}

		onRepaint := func(buf *image.RGBA) {
			draw.Draw(buf, buf.Bounds(), &image.Uniform{color.Black}, image.Pt(0, 0), draw.Src)
			draw.Draw(buf, image.Rect(px, py, px+256, py+256), &image.Uniform{color.White}, image.ZP, draw.Src)
		}

		EventLoop(s, w, onRepaint, nil, onMouse)
	})
}

func EventLoop(s screen.Screen, w screen.Window,
	onRepaint func(buf *image.RGBA), onKey func(key.Event), onMouse func(mouse.Event)) {
	if onKey == nil {
		onKey = func(key.Event) {}
	}
	if onMouse == nil {
		onMouse = func(mouse.Event) {}
	}

	var (
		winSize image.Point
		winBuf  screen.Buffer
		winTex  screen.Texture
	)

	for {
		e := w.NextEvent()
		printEvent(e)
		switch e := e.(type) {

		case lifecycle.Event:
			if e.To == lifecycle.StageDead {
				return
			}

		case size.Event:
			newSize := image.Pt(e.WidthPx, e.HeightPx)
			if winSize != newSize {
				winSize = newSize
				if winBuf != nil {
					winBuf.Release()
				}
				if winTex != nil {
					winTex.Release()
				}
				winBuf = newBuffer(s, winSize)
				winTex = newTexture(s, winSize)
			}
			w.Send(paint.Event{})

		case key.Event:
			if e.Code == key.CodeEscape {
				return
			}

		case mouse.Event:
			onMouse(e)

		case paint.Event:
			onRepaint(winBuf.RGBA())
			winTex.Upload(image.Pt(0, 0), winBuf, winBuf.Bounds())
			w.Copy(image.Pt(0, 0), winTex, winTex.Bounds(), screen.Over, nil)
			w.Publish()
		}
	}
}

func NewWindow(s screen.Screen, title string) screen.Window {
	w, err := s.NewWindow(&screen.NewWindowOptions{
		Title: title,
	})
	check(err)
	return w
}

func newBuffer(s screen.Screen, size image.Point) screen.Buffer {
	b, err := s.NewBuffer(size)
	check(err)
	return b
}

func newTexture(s screen.Screen, size image.Point) screen.Texture {
	t, err := s.NewTexture(size)
	check(err)
	return t
}

func printEvent(e interface{}) {
	format := "got %#v\n"
	if _, ok := e.(fmt.Stringer); ok {
		format = "got %v\n"
	}
	fmt.Printf(format, e)
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
