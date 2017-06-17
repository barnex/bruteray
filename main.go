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

	cam := Camera{W: *width, H: *height, FocalLen: 0}

	scene := Scene1()
	img := cam.Render(scene)
	Encode(img, "001.jpg", *overExp)

	//Reset()
	//Scene2()
	//RenderScene(*width, *height, *iters, "002.jpg")

	//Reset()
	//Scene3()
	//RenderScene(*width, *height, *iters, "003.jpg")

	//Reset()
	//Scene4()
	//RenderScene(*width, *height, *iters, "004.jpg")

	fmt.Println("done,", time.Since(start))
}

//func Scene2() {
//	sp := Sphere(Vec{0, 0, -3}, 0.25)
//	objects = []Obj{
//		{Shape: sp},
//	}
//}
//
//func Scene3() {
//	const r = 0.25
//	sp := Sphere(Vec{0, 0, 3}, r)
//	objects = []*Obj{
//		{Shape: sp.Transl(r/2, 0, 0)},
//		{Shape: sp.Transl(-r/2, 0, 0)},
//	}
//}
//
//func Scene4() {
//	const r = 0.25
//	sp := Sphere(Vec{0, 0, 3}, r)
//	objects = []*Obj{
//		{Shape: And(sp.Transl(r/2, 0, 0), sp.Transl(-r/2, 0, 0))},
//	}
//}

func Init() {
	flag.Parse()
	if *pprof != "" {
		go func() {
			log.Fatal(http.ListenAndServe(*pprof, nil))
		}()
	}
}
