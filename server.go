package bruteray

import (
	"bytes"
	"flag"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
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
)

func Serve(e *Env) {

	env = e

	http.HandleFunc("/preview", handlePreview)
	http.HandleFunc("/render", handleRender)
	http.HandleFunc("/", mainHandler)

	//log.Println("listen", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(mainHTML))
}

func handleRender(w http.ResponseWriter, r *http.Request) {

}

var preview struct {
	w, h int
	data bytes.Buffer
}

func handlePreview(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	W := parseInt(q.Get("w"), DefaultWidth)
	H := parseInt(q.Get("h"), DefaultHeight)

	if W != preview.w || H != preview.h {
		start := time.Now()

		e2 := env.Preview()
		img := MakeImage(W, H)
		Render(e2, 1, img)

		log.Println("preview", time.Since(start).Round(time.Millisecond))
		Print(jpeg.Encode(&(preview.data), img, &jpeg.Options{Quality: 95}))
		preview.w, preview.h = W, H
	}

	w.Write(preview.data.Bytes())
}

func parseInt(s string, Default int) int {
	x, err := strconv.Atoi(s)
	Print(err)
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
