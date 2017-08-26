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
	port       = flag.String("http", ":3700", "Port to serve HTTP")
	flagWidth  = flag.Int("w", 1920, "image width")
	flagHeight = flag.Int("h", 1080, "image height")
)

var (
	env *Env
)

const (
	DefaultRec = 6
)

// Serve starts a web UI server
// at the port specified by flag --http.
func Serve(e *Env) {

	flag.Parse()

	env = e

	http.HandleFunc("/render", handleRender)
	http.HandleFunc("/", mainHandler)

	progressive = RenderLoop(env, DefaultRec, *flagWidth, *flagHeight)

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
	printErr(jpeg.Encode(w, img, &jpeg.Options{Quality: 95}))
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
