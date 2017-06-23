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
	iters    = flag.Int("N", 10000, "number of iterations")
	pprof    = flag.String("pprof", ":9093", "pprof port")
)

func main() {
	Init()

	const h = 2
	s := &Env{}
	const (
		G     = -2
		RoomW = 6
		RoomD = 10
		WallW = 0.1
		WallH = 5
		WallT = 0.02
	)

	ground := Slab(G, -100)
	die := ABox(Vec{-2, G, 4}, Vec{-1, G + 1, 5})
	marble := Sphere(Vec{1, G + 1, 5}, 1)
	walll := ABox(Vec{-RoomW / 2, G, 0}, Vec{-RoomW - WallT/2, G + WallH, 100})
	wallr_ := SlabD(RoomW/2, RoomW+WallT/2, Vec{1, 0, 0})
	wallr := wallr_
	wallb := SlabD(RoomD, RoomD+WallT, Vec{0, 0, 1})
	lightPos := Vec{0, 1, 4}
	s.objs = []Obj{
		Diffuse2(s, ground, 0.8),
		Reflective(s, marble, 0.5),
		Diffuse2(s, die, 1),
		Diffuse2(s, walll, 0.8),
		Diffuse2(s, wallr, 0.8),
		Diffuse2(s, wallb, 0.8),
		Flat(Sphere(lightPos, 1), 10),
	}
	s.sources = []Source{
		&BulbSource{lightPos, 150, 2},
		//&PointSource{lightPos, 100},
	}
	s.amb = func(v Vec) Color { return Color(0.1 * v.Y) }

	cam := Camera(*width, *height, *focalLen)
	cam.Transf = RotX(-5 * deg)

	Render(s, cam, "out.jpg")
}

func Render(s *Env, cam *Cam, fname string) {
	start := time.Now()
	every := 1
	for i := 0; i < *iters; i++ {
		cam.iterate(s)
		if i%every == 0 {
			Encode(cam.Img, fname, 1/(float64(cam.N)), *overExp)
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
