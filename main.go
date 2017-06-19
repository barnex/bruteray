package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var (
	width    = flag.Int("w", 800, "canvas width")
	height   = flag.Int("h", 600, "canvas height")
	focalLen = flag.Float64("f", 1, "focal length")
	maxRec   = flag.Int("rec", 3, "maximum number of recursive rays")
	overExp  = flag.Bool("over", false, "highlight over/under exposed pixels")
	quality  = flag.Int("q", 85, "JPEG quality")
	useSRGB  = flag.Bool("srgb", true, "use sRGB color space")
	iters    = flag.Int("N", 1, "number of iterations")
	pprof    = flag.String("pprof", "", "pprof port")
)

// Scene:
var (
	objects []Obj // TODO: object sources, intersect([]obj), nearest([]obj)
	//sources []Source
	//ambient = func(v Vec) float64 {
	//	return 0.1 * math.Abs((v.Normalized().Y))
	//}
)

const off = 1e-6 // anti-bleeding offset, intersection points moved this much away from surface

func main() {
	Init()
	start := time.Now()

	s := &Scene{}
	cam := Camera(*width, *height, 0)
	Encode(cam.Render(s), "out.jpg", *overExp)
	Encode(Stretch(cam.ZMap), "z.jpg", true)

	fmt.Println("done,", time.Since(start))
}

func Init() {
	flag.Parse()
	if *pprof != "" {
		go func() {
			log.Fatal(http.ListenAndServe(*pprof, nil))
		}()
	}
}
