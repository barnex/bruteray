package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

var (
	width        = flag.Int("w", 1024, "canvas width")
	height       = flag.Int("h", 768, "canvas height")
	focalLen     = flag.Float64("f", 1, "focal length")
	overExp      = flag.Bool("over", false, "highlight over/under exposed pixels")
	maxRecursion = flag.Int("r", 2, "maximum number of recursive rays")
)

var (
	objects []*Obj
)

func main() {
	Init()
	start := time.Now()

	img := MakeImage(*width, *height)

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

			v := Intensity(r)

			if !*overExp {
				v = clip(v, 0, 1)
			}

			img[i][j] = v
		}
	}
}

func Intensity(r Ray) float64{
	return 0.5
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

//const (
//	Horiz = 20.0
//)
//
//const deg = math.Pi / 180
//
//var (
//	Focal = Vec{0, 0, -1}
//	scene *Scene
//)
//
//type Scene struct {
//	light Vec
//	amb   float64
//	objs  []Obj
//}
//
//
//
//var nShade int
//
//func Render(s *Scene, img [][]float64) {
//	nShade = 0
//	for sub := *progressive; sub > 0; sub /= 2 {
//		refine(s, img, sub, sub == *progressive)
//		Encode(img, "out.jpg")
//	}
//}
//
//func refine(sc *Scene, img [][]float64, sub int, first bool) {
//	W := *width
//	H := *height
//	for i := 0; i < H; i += sub {
//		fmt.Printf("%.1f%%\n\u001B[F", float64(100*nShade)/float64((W+1)*(H+1)))
//		for j := 0; j < W; j += sub {
//			if i%(2*sub) == 0 && j%(2*sub) == 0 && !first {
//				continue
//			}
//			nShade++
//			y0 := (-float64(i) + float64(H)/2 + 0.5) / float64(H)
//			x0 := (float64(j) - float64(W)/2 + 0.5) / float64(H)
//			start := Vec{x0, y0, 0}
//			r := Ray{start, start.Sub(Focal).Normalized()}
//
//			v, _, _ := PixelShade(sc, r, *maxRecursion)
//			v = clip(v, 0, 1)
//
//			for I := i; I < i+sub && I < H; I++ {
//				for J := j; J < j+sub && J < W; J++ {
//					img[I][J] = v
//				}
//			}
//		}
//	}
//}
//
//func PixelShade(sc *Scene, r Ray, N int) (float64, Vec, bool) {
//	if N == 0 {
//		return scene.amb, Vec{}, false
//	}
//
//	i, _ := Nearest(sc.objs, r)
//	if i == -1 {
//		return 0, Vec{}, false
//	}
//	obj := sc.objs[i]
//	shape := obj.Shape
//
//	t, norm, ok := shape.Normal(r)
//	if !ok {
//		return 0, Vec{}, false
//	}
//
//	v := obj.Shader(t, norm, r, N)
//
//	//v = clip(v, 0, 1)
//	return v, r.At(t), true
//}
//
//
//func inters(r Ray, s Shape) bool {
//	_, _, ok := s.Normal(r)
//	return ok
//}
//
//func intersAny(r Ray, s []Obj) bool {
//	for _, s := range s {
//		if inters(r, s.Shape) {
//			return true
//		}
//	}
//	return false
//}
//
//
//
//func assert(t bool) {
//	if !t {
//		panic("assertion failed")
//	}
//}
//
