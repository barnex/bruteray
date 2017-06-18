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
	pprof    = flag.String("pprof", ":6060", "pprof port")
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

	scenes := []func() *Scene{Scene1, Scene2, Scene3, Scene4, Scene5, Scene6}
	for _, s := range scenes {
		test(s())
	}

	fmt.Println("done,", time.Since(start))
}

var cnt = 0

func test(s *Scene) {
	cam := Camera(*width, *height, 0)
	cnt++
	name := fmt.Sprintf("%03d", cnt)
	Encode(cam.Render(s), name+".jpg", *overExp)
	Encode(Stretch(cam.ZMap), name+"-z.jpg", true)
}

func Init() {
	flag.Parse()
	if *pprof != "" {
		go func() {
			log.Fatal(http.ListenAndServe(*pprof, nil))
		}()
	}
}
