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

	scene := dies()

	cam := Camera(*width, *height, *focalLen)
	cam.Pos = Vec{0, 1, -5}
	cam.Transf = RotX(-10 * deg)

	Render(scene, cam, "out.jpg")
}

func dies() *Env {
	s := &Env{}
	cube := Box(Vec{0, 0, 0}, 1, 1, 1)
	dot := Sphere(Vec{1, 1, -1}, 1)
	//die := ShapeAnd(cube, dot)
	die := ShapeMinus(cube, dot)

	s.Add(die, Diffuse1(0.8))
	s.AddLight(PointLight(Vec{1, 5, -2}, 5))

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
