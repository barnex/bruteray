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
)

var (
	port       = flag.String("http", ":3700", "Port to serve HTTP")
	flagWidth  = flag.Int("w", 1920, "image width")
	flagHeight = flag.Int("h", 1080, "image height")
)

var (
	env *bruteray.Env
)

// Starts a web UI server
// at the port specified by flag --http.
func Env(e *bruteray.Env) {

	flag.Parse()

	env = e

	http.HandleFunc("/render", handleRender)
	http.HandleFunc("/", mainHandler)

	progressive = RenderLoop(env, *flagWidth, *flagHeight)

	log.Fatal(http.ListenAndServe(*port, nil))
}

type Loop struct {
	env  *bruteray.Env
	w, h int
	acc  bruteray.Image
	n    float64
	mu   sync.Mutex
}

func (l *Loop) run() {
	for {
		l.iter()
	}
}

func (l *Loop) iter() {
	img := bruteray.MakeImage(l.w, l.h)
	bruteray.Render(l.env, img)
	l.mu.Lock()
	l.acc.Add(img)
	l.n++
	l.mu.Unlock()
}

func (l *Loop) Current() bruteray.Image {
	l.mu.Lock()
	defer l.mu.Unlock()
	//log.Println("current")
	img := bruteray.MakeImage(l.w, l.h)
	for i := range img {
		for j := range img[i] {
			img[i][j] = l.acc[i][j].Mul(1 / l.n)
		}
	}
	return img
}

func RenderLoop(e *bruteray.Env, w, h int) *Loop {
	l := &Loop{env: e, w: w, h: h, acc: bruteray.MakeImage(w, h)}
	l.iter()   // make sure we have 1 pass at least
	go l.run() // refine in the background
	return l
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

func encode(w io.Writer, img bruteray.Image) {
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
