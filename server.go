package bruteray

import (
	"bytes"
	"flag"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	port = flag.String("http", ":3700", "Port to serve HTTP")
)

var (
	env *Env
)

const (
	DefaultWidth  = 800
	DefaultHeight = 600
	DefaultRec    = 4
)

// Serve starts a web UI server
// at the port specified by flag --http.
func Serve(e *Env) {

	env = e

	http.HandleFunc("/preview", handlePreview)
	http.HandleFunc("/render", handleRender)
	http.HandleFunc("/", mainHandler)

	log.Fatal(http.ListenAndServe(*port, nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(mainHTML))
}

var (
	progressive *Loop
	pmu         sync.Mutex
)

func handleRender(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	W := parseInt(q.Get("w"), DefaultWidth)
	H := parseInt(q.Get("h"), DefaultHeight)
	R := parseInt(q.Get("rec"), DefaultRec)

	pmu.Lock()
	defer pmu.Unlock()

	if progressive == nil {
		progressive = RenderLoop(env, R, W, H)
	}
	img := progressive.Current()
	encode(w, img)
}

var preview struct {
	w, h int
	data bytes.Buffer
	sync.Mutex
}

func handlePreview(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	W := parseInt(q.Get("w"), DefaultWidth)
	H := parseInt(q.Get("h"), DefaultHeight)

	preview.Lock()
	defer preview.Unlock()

	if W != preview.w || H != preview.h {
		start := time.Now()

		e2 := env.Preview()
		img := MakeImage(W, H)
		Render(e2, 3, img)

		log.Println("preview", time.Since(start).Round(time.Millisecond))
		encode(&preview.data, img)
		preview.w, preview.h = W, H
	}

	w.Write(preview.data.Bytes())
}

func encode(w io.Writer, img Image) {
	Print(jpeg.Encode(w, img, &jpeg.Options{Quality: 95}))
}

func parseInt(s string, Default int) int {
	x, _ := strconv.Atoi(s)
	//Print(err)
	if x == 0 {
		return Default
	}
	return x
}

func Print(err error) {
	if err != nil {
		log.Println(err)
	}
}
