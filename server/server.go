package server

import (
	"bytes"
	"flag"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/barnex/bruteray"
)

var (
	port = flag.String("http", ":3700", "Port to serve HTTP")
)

var (
	env *bruteray.Env
)

const (
	DefaultWidth  = 800
	DefaultHeight = 600
)

func Serve(e *bruteray.Env) {

	env = e

	http.HandleFunc("/render", renderHandler)
	http.HandleFunc("/", mainHandler)

	log.Println("listen", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(mainHTML))
}

var cache struct {
	w, h int
	data bytes.Buffer
}

func renderHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	W := parseInt(q.Get("w"), DefaultWidth)
	H := parseInt(q.Get("h"), DefaultHeight)

	if W != cache.w || H != cache.h {
		start := time.Now()
		img := bruteray.MakeImage(W, H)
		env.Camera.Render(env, 1, img)
		log.Println("rendered", time.Since(start).Round(time.Millisecond))
		Print(jpeg.Encode(&(cache.data), img, &jpeg.Options{Quality: 95}))
		cache.w, cache.h = W, H
	}

	w.Write(cache.data.Bytes())
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
