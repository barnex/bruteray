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
	start := time.Now()

	const h = 2

	s := &Scene{}
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
	wallr_ := ABox(Vec{RoomW / 2, G, 0}, Vec{RoomW + WallT/2, G + WallH, 100})
	//win := ABox(Vec{RoomW/2 - 1, G + 2, 1}, Vec{RoomW/2 + 1, G + 3, 2})
	//win := ABox(Vec{2, G, 4}, Vec{4, G + 1, 5})
	//wallr := &ShapeMinus{wallr_, win}
	wallr := wallr_
	wallb := ABox(Vec{-RoomW / 2, G, RoomD}, Vec{RoomW, G + WallH, RoomD + WallT})

	s.objs = []Obj{
		Diffuse2(s, ground, 0.5),
		Reflective(s, marble, 0.9),
		Diffuse2(s, die, 1),
		Diffuse2(s, walll, 1),
		Diffuse2(s, wallr, 1),
		Diffuse2(s, wallb, 1),
		Flat(Sphere(Vec{2, 8, -2}, 1), 10),
	}
	s.sources = []Source{
		&BulbSource{Vec{2, 8, -2}, 150, 2},
		//&PointSource{Vec{2, 8, -2}, 100},
	}
	s.amb = func(v Vec) Color { return Color(0.1 * v.Y) }

	//ground := Diffuse2(s, ABox(Vec{-100, -h, 0}, Vec{100, -2 * h, 100}), 0.5)
	//sp := Sphere(Vec{-0.5, -1, 8}, 2)
	//die := &ShapeAnd{sp, Slab(-h+.2, -.2)}
	//dice := Diffuse2(s, die, 0.95)
	////dice := Flat(die, 0.95)
	//s.objs = []Obj{
	//	ground,
	//	dice,
	//	Reflective(s, Sphere(Vec{3, -1, 10}, 1), 0.9),
	//}

	cam := Camera(*width, *height, *focalLen)
	cam.Transf = RotX(-5 * deg)

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
