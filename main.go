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
)

// Scene:
var (
	objects []*Obj
	sources []*PointSource
)

const off = 1e-6 // anti-bleeding offset, intersection points moved this much away from surface

func main() {
	Init()
	start := time.Now()

	img := MakeImage(*width, *height)

	InitScene()

	Render(img)

	Encode(img, "out.jpg")

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
	objects = []*Obj{
		{HalfspaceY(-2), Diffuse1(0.3)},
		{Sphere(Vec{-2, -1, 6}, 1), Reflective(0.5)},
		{Sphere(Vec{0, -1, 8}, 1), Reflective(0.9)},
		{Sphere(Vec{2, -1, 6}, 1), Diffuse1(0.8)},
	}
	sources = []*PointSource{
		{Pos: Vec{3.0, 8, 4.0}, Flux: 3},
		{Pos: Vec{3.0, 8, 4.5}, Flux: 3},
		{Pos: Vec{3.0, 8.0, 4}, Flux: 3},
		{Pos: Vec{3.0, 8.5, 4}, Flux: 3},
		{Pos: Vec{3.0, 7.0, 5}, Flux: 3},
		{Pos: Vec{3.3, 8.5, 6}, Flux: 3},
		{Pos: Vec{2.5, 9.0, 4}, Flux: 3},
		{Pos: Vec{3, 1, 0}, Flux: 2},
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
			y0 := (-float64(i) + float64(H)/2 + 0.5) / float64(H)
			x0 := (float64(j) - float64(W)/2 + 0.5) / float64(H)

			start := Vec{x0, y0, 0}
			r := Ray{start, start.Sub(focal).Normalized()}

			v := Intensity(r, 0)

			if !*overExp {
				v = clip(v, 0, 1)
			}

			img[i][j] = v
		}
	}
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
