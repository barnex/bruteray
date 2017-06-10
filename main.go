package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

var (
	width    = flag.Int("w", 1024, "canvas width")
	height   = flag.Int("h", 768, "canvas height")
	focalLen = flag.Float64("f", 1, "focal length")
	maxRec   = flag.Int("rec", 2, "maximum number of recursive rays")
	overExp  = flag.Bool("over", false, "highlight over/under exposed pixels")
	quality  = flag.Int("q", 80, "JPEG quality")
	useSRGB  = flag.Bool("srgb", true, "use sRGB color space")
)

// Scene:
var (
	objects []*Obj
	sources []Source
)

const off = 1e-6 // anti-bleeding offset, intersection points moved this much away from surface

func main() {
	Init()
	start := time.Now()

	img := MakeImage(*width, *height)

	InitScene()

	for N := 0; N < 100; N++ {
		Render(img)
		Encode(img, "out.jpg", float64(N+1))
	}

	fmt.Println("done,", time.Since(start))
}

func Init() {
	flag.Parse()
}

func MakeImage(W, H int) [][]float64 {
	img := make([][]float64, H)
	for i := range img {
		img[i] = make([]float64, W)
	}
	return img
}

func InitScene() {
	lp := Vec{20, 60, -5}
	lr := 10.
	objects = []*Obj{
		{HalfspaceY(-2), Diffuse2(0.7)},
		{Sphere(Vec{-2, -1, 6}, 1), Reflective(0.8)},
		{Sphere(Vec{0, -1, 8}, 1), Reflective(0.8)},
		{Sphere(Vec{2, -1, 6}, 1), Diffuse2(0.9)},
		{Sphere(lp, lr/2), Flat(1)},
	}
	sources = []Source{
		//&BulbSource{Pos: Vec{3, 8, 4}, Flux: 30, R: 2},
		&BulbSource{Pos: lp, Flux: 80, R: lr},
	}
}

func Render(img [][]float64) {
	focal := Vec{0, 0, -*focalLen}
	W := *width
	H := *height
	nPix := 0
	for i := 0; i < H; i++ {
		fmt.Printf("%.1f%%\n\u001B[F", float64(100*nPix)/float64((W+1)*(H+1)))
		for j := 0; j < W; j++ {
			nPix++
			y0 := (-float64(i) + aa() + float64(H)/2) / float64(H)
			x0 := (float64(j) + aa() - float64(W)/2) / float64(H)

			start := Vec{x0, y0, 0}
			r := Ray{start, start.Sub(focal).Normalized()}

			v := Intensity(r, 0)

			img[i][j] += v
		}
	}
}

// Anti-aliasing jitter
func aa() float64 {
	return rand()
}

func Intensity(r Ray, N int) float64 {
	if N == *maxRec {
		return 0
	}
	t, n, obj := FirstIntersect(r)
	if obj != nil {
		return obj.Shader.Intensity(r, t, n, N)
	}
	return 0
}

func FirstIntersect(r Ray) (float64, Vec, *Obj) {
	var (
		nearestT        = math.Inf(1)
		nearestN        = Vec{}
		nearestObj *Obj = nil
	)

	for _, o := range objects {
		t, n, ok := o.Normal(r)
		if ok && t < nearestT {
			nearestT = t
			nearestN = n
			nearestObj = o
		}
	}
	return nearestT, nearestN, nearestObj
}
