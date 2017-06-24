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
	maxRec   = flag.Int("rec", 2, "maximum number of recursive rays")
	overExp  = flag.Bool("over", false, "highlight over/under exposed pixels")
	quality  = flag.Int("q", 85, "JPEG quality")
	useSRGB  = flag.Bool("srgb", true, "use sRGB color space")
	iters    = flag.Int("N", 10000, "number of iterations")
	pprof    = flag.String("pprof", ":9093", "pprof port")
)

func main() {
	Init()

	const h = 2
	const (
		G     = -2
		RoomW = 6
		RoomD = 10
		WallW = 0.1
		WallH = 5
		WallT = 0.02
	)

	scene := &Env{}
	scene.amb = func(v Vec) Color { return Color(0.2*v.Y + 0.2) }

	scene.Add(Sheet(-3, Ey), Diffuse1(0.8))
	scene.Add(Sphere(Vec{1, 0, 4}, 1), Diffuse1(0.9))
	scene.Add(Sphere(Vec{-1, 0, 5}, 1), Reflective(0.5))
	scene.Add(Sheet(20, Ez), Diffuse1(0.8))
	//scene.Add(Sheet(10, Ex), Diffuse1(0.8))
	//scene.Add(Sheet(-10, Ex), Diffuse2(0.8))
	scene.AddLight(PointLight(Vec{0, 8, 0}, 100))
	//scene.AddLight(SmoothLight(Vec{0, 5, 0}, 100, 2))

	//ground := Slab(G, -100)
	//die := ABox(Vec{-2, G, 4}, Vec{-1, G + 1, 5})
	//marble := Sphere(Vec{1, G + 1, 5}, 1)
	//walll := ABox(Vec{-RoomW / 2, G, 0}, Vec{-RoomW - WallT/2, G + WallH, 100})
	//wallr_ := SlabD(RoomW/2, RoomW+WallT/2, Vec{1, 0, 0})
	//wallr := wallr_
	//wallb := SlabD(RoomD, RoomD+WallT, Vec{0, 0, 1})
	//lightPos := Vec{0, 1, 4}
	//s.objs = []Obj{
	//	Diffuse2(s, ground, 0.8),
	//	Reflective(s, marble, 0.5),
	//	Diffuse2(s, die, 1),
	//	Diffuse2(s, walll, 0.8),
	//	Diffuse2(s, wallr, 0.8),
	//	Diffuse2(s, wallb, 0.8),
	//	Flat(Sphere(lightPos, 1), 10),
	//}
	//s.sources = []Source{
	//	&BulbSource{lightPos, 150, 2},
	//	//&PointSource{lightPos, 100},
	//}

	cam := Camera(*width, *height, *focalLen)
	//cam.Transf = RotX(-5 * deg)

	Render(scene, cam, "out.jpg")
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
		}
		every++
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
