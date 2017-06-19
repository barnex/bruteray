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
	overExp  = flag.Bool("over", true, "highlight over/under exposed pixels")
	quality  = flag.Int("q", 85, "JPEG quality")
	useSRGB  = flag.Bool("srgb", true, "use sRGB color space")
	iters    = flag.Int("N", 10000, "number of iterations")
	pprof    = flag.String("pprof", "", "pprof port")
)

//const off = 1e-6 // anti-bleeding offset, intersection points moved this much away from surface

func main() {
	Init()
	start := time.Now()

	const h = 2

	s := &Scene{}

	ground := Diffuse1(s, Slab(-h, -h-100), 0.95)
	sph := Diffuse1(s, Sphere(Vec{-1, -0.8, 4}, 1.2), 0.5)
	s.objs = []Obj{
		ground,
		sph,
	}
	s.sources = []Source{
		&BulbSource{Vec{3, 5, 1}, 100, 1},
	}

	cam := Camera(*width, *height, *focalLen)

	Encode(Stretch(cam.ZMap), "z.jpg", 1, true)
	every := 1
	for i := 0; i < *iters; i++ {
		cam.iterate(s)
		if i%every == 0 {
			Encode(cam.Img, "out.jpg", 1/(float64(cam.N)), *overExp)
		}
		every++
		if every > 20 {
			every = 20
		}
	}

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
