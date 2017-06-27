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

	scene := chessboard()
	//scene := spheresInARoom()

	cam := Camera(*width, *height, *focalLen)
	cam.Pos = Vec{0, 5.5, -.5}
	cam.Transf = RotX(-17 * deg)

	Render(scene, cam, "out.jpg")
}

func chessboard() *Env {
	s := &Env{}
	s.amb = func(dir Vec) Color { return 0.5 }

	s.Add(Sheet(0, Ey), Diffuse2(0.7)) // floor

	s.Add(Box(Vec{0, 0, 8}, 5, 0.45, 5), Diffuse2(0.1)) // base

	s.Add(Rect(Vec{0, 0.5, 8}, Ey, 4, inf, 4), CheckBoard(Reflective(0.05), Diffuse2(0.9))) // checkboard

	// walls
	s.Add(Sheet(20, Ez), Diffuse2(0.7)) // back
	//s.Add(Sheet(20, Ex), Diffuse2(0.7))  // left
	//s.Add(Sheet(-20, Ex), Diffuse2(0.7)) // right
	s.Add(Sheet(20, Ey), Diffuse2(0.6)) // ceiling

	slab := Slab(0, 0.7)

	for j := 0.; j < 2; j++ {
		for i := j; i < 8; i += 2 {
			cyl := Cylinder(Vec{j - 3.5, -1.5, i + 4.5}, 0.37)
			s.Add(ShapeAnd(cyl, slab), Reflective(0.05))

			cyl = Cylinder(Vec{j + 2.5, -1.5, i + 4.5}, 0.37)
			s.Add(ShapeAnd(cyl, slab), Diffuse2(0.95))
		}
	}

	s.AddLight(SmoothLight(Vec{3, 12, 6}, 130, 2))
	//s.AddLight(PointLight(Vec{3, 12, 6}, 150))

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
