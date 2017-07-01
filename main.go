package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var (
	width    = flag.Int("w", 1920, "canvas width")
	height   = flag.Int("h", 1080, "canvas height")
	focalLen = flag.Float64("f", 1, "focal length")
	maxRec   = flag.Int("rec", 2, "maximum number of recursive rays")
	overExp  = flag.Bool("over", false, "highlight over/under exposed pixels")
	quality  = flag.Int("q", 90, "JPEG quality")
	useSRGB  = flag.Bool("srgb", true, "use sRGB color space")
	iters    = flag.Int("N", 10000, "number of iterations")
	pprof    = flag.String("pprof", ":9093", "pprof port")
)

func main() {
	//Init()

	//scene := &Env

	//cube := &object{Box(Vec{0, 0, 0}, -1, -1, -1), Diffuse1(0.9)}
	//pip = &object{Sphere(Vec{0, 0, -0.9}, r), Reflective(0.05)}
	//dice :=

	//die := cube
	//const r = 0.15
	//die = ShapeMinus(die, Sphere(Vec{0, 0, -0.9}, r))

	//cam := Camera(*width, *height, *focalLen)
	//cam.Pos = Vec{0, 4, -6}
	//cam.Transf = RotX(-15 * deg)
	//cam.AA = true

	//Render(scene, cam, "out.jpg")
}

func dice() *Env {
	s := &Env{}
	s.amb = func(Vec) Color { return 0.1 }
	cube := Box(Vec{0, 0, 0}, -1, -1, -1)

	die := cube
	const r = 0.15
	die = ShapeMinus(die, Sphere(Vec{0, 0, -0.9}, r))
	die = ShapeMinus(die, Sphere(Vec{0.5, 0.5, -0.9}, r))
	die = ShapeMinus(die, Sphere(Vec{-0.5, 0.5, -0.9}, r))
	die = ShapeMinus(die, Sphere(Vec{0.5, -0.5, -0.9}, r))
	die = ShapeMinus(die, Sphere(Vec{-0.5, -0.5, -0.9}, r))

	die = ShapeMinus(die, Sphere(Vec{0.4, 1.1, -0.4}, r))
	die = ShapeMinus(die, Sphere(Vec{-0.4, 1.1, 0.4}, r))

	die = ShapeAnd(die, Sphere(Vec{}, 0.98*math.Sqrt(2.)))

	s.Add(die, Diffuse2(0.9))

	s.Add(Sheet(-1, Ey), Diffuse2(0.5))

	s.AddLight(SmoothLight(Vec{2, 3, -3}, 15, 0.2))
	//s.AddLight(PointLight(Vec{1, 3, -3}, 15))

	return s
}

func Render(s *Env, cam *Cam, fname string) {
	every := 1
	W, H := cam.Size()
	for i := 0; i < *iters; i++ {
		start := time.Now()
		cam.iterate(s)
		speed := float64((W+1)*(H+1)) / time.Since(start).Seconds()
		fmt.Printf("%.2f Mpixel/s\n", speed/1e6)
		if i%every == 0 {
			Encode(cam.Img, fname, 1/(float64(cam.N)), *overExp)
			every++
		}
		if every > 20 {
			every = 20
		}
	}

}

func Init() {
	flag.Parse()
	if *pprof != "" {
		go func() {
			log.Fatal(http.ListenAndServe(*pprof, nil))
		}()
	}
}
