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
	Init()

	scene := NewEnv()
	scene.amb = func(dir Vec) Color { return Color(0.4 * dir.Y) }
	scene.Add(Sheet(-1, Ey), Diffuse1(0.5))

	box1 := Box(Vec{0, 0, 0}, -1, -1, -1)
	box1 = Transf(box1, RotY4(pi/5).Mul(RotX4(pi/7)))

	cube := &object{box1, Diffuse1(0.8)}
	var die Obj = cube

	scene.AddObj(die)

	lp := Vec{4, 8, -6}
	scene.AddLight(PointLight(lp, 50))

	cam := Camera(*focalLen)
	cam.Transl(Vec{0, 4, -6})
	cam.Transf(RotX(-15 * deg))
	cam.AA = true

	Render(scene, cam, "out.jpg")
}

func dice() *Env {
	s := NewEnv()
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
	img := MakeImage(*width, *height)
	every := 1
	W, H := img.Size()
	for i := 0; i < *iters; i++ {
		start := time.Now()
		cam.Render(s, img)
		speed := float64((W+1)*(H+1)) / time.Since(start).Seconds()
		fmt.Printf("%.2f Mpixel/s\n", speed/1e6)
		if i%every == 0 {
			Encode(img, fname, 1/(float64(cam.N)), *overExp)
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
