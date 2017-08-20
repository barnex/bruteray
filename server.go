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

	http.HandleFunc("/render", handleRender)
	http.HandleFunc("/", mainHandler)

	progressive = RenderLoop(env, DefaultRec, DefaultWidth, DefaultHeight)

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
	img := progressive.Current()
	encode(w, img)
}

var preview struct {
	w, h int
	data bytes.Buffer
	sync.Mutex
}

func encode(w io.Writer, img Image) {
	Print(jpeg.Encode(w, img, &jpeg.Options{Quality: 85}))
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
