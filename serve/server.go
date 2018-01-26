package serve

// server severs an Env over HTTP,
// so we can see it while being rendered.

import (
	"bytes"
	"flag"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"

	"github.com/barnex/bruteray"
	"golang.org/x/image/tiff"
)

var (
	port       = flag.String("http", ":3700", "Port to serve HTTP")
	flagWidth  = flag.Int("w", 1920, "image width")
	flagHeight = flag.Int("h", 1080, "image height")
	qual       = flag.Int("q", 90, "jpeg quality")
)

var (
	env  *bruteray.Env
	peek = make(chan chan bruteray.Image)
)

// Starts a web UI server
// at the port specified by flag --http.
func Env(e *bruteray.Env) {

	flag.Parse()

	env = e

	http.HandleFunc("/render", handleRender)
	http.HandleFunc("/tiff", handleTiff)
	http.HandleFunc("/", mainHandler)

	go bruteray.RenderLoop(e, *flagWidth, *flagHeight, peek)

	log.Fatal(http.ListenAndServe(*port, nil))
}

func handleRender(w http.ResponseWriter, r *http.Request) {
	encode(w, peekImg())
}

func peekImg() bruteray.Image {
	resp := make(chan bruteray.Image)
	peek <- resp
	return <-resp
}

func handleTiff(w http.ResponseWriter, r *http.Request) {
	img := peekImg()
	printErr(tiff.Encode(w, img, &tiff.Options{Predictor: true, Compression: tiff.Deflate}))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(mainHTML))
}

var preview struct {
	w, h int
	data bytes.Buffer
	sync.Mutex
}

func encode(w io.Writer, img bruteray.Image) {
	printErr(jpeg.Encode(w, img, &jpeg.Options{Quality: *qual}))
}

func parseInt(s string, Default int) int {
	x, _ := strconv.Atoi(s)
	//Print(err)
	if x == 0 {
		return Default
	}
	return x
}

func printErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
