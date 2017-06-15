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
	iters    = flag.Int("N", 1, "number of iterations")
	pprof    = flag.String("pprof", ":6060", "pprof port")
)

// Scene:
var (
	objects []*Obj // TODO: object sources, intersect([]obj), nearest([]obj)
	//sources []Source
	//ambient = func(v Vec) float64 {
	//	return 0.1 * math.Abs((v.Normalized().Y))
	//}
)

const off = 1e-6 // anti-bleeding offset, intersection points moved this much away from surface

func main() {
	Init()
	start := time.Now()

	*focalLen = 0

	Reset()
	Scene1()
	RenderScene(*width, *height, *iters, "001.jpg")

	Reset()
	Scene2()
	RenderScene(*width, *height, *iters, "002.jpg")

	Reset()
	Scene3()
	RenderScene(*width, *height, *iters, "003.jpg")

	Reset()
	Scene4()
	RenderScene(*width, *height, *iters, "004.jpg")

	fmt.Println("done,", time.Since(start))
}

func Scene1() {
	objects = []*Obj{
		{Shape: Sphere(Vec{0, 0, 3}, 0.25)},
	}
}

func Scene2() {
	sp := Sphere(Vec{0, 0, -3}, 0.25)
	objects = []*Obj{
		{Shape: sp},
	}
}

func Scene3() {
	const r = 0.25
	sp := Sphere(Vec{0, 0, 3}, r)
	objects = []*Obj{
		{Shape: sp.Transl(r/2, 0, 0)},
		{Shape: sp.Transl(-r/2, 0, 0)},
	}
}

func Scene4() {
	const r = 0.25
	sp := Sphere(Vec{0, 0, 3}, r)
	objects = []*Obj{
		{Shape: And(sp.Transl(r/2, 0, 0), sp.Transl(-r/2, 0, 0))},
	}
}

func Reset() {
	objects = nil
}

func RenderScene(w, h int, N int, outfname string) {
	img := MakeImage(w, h)

	for i := 0; i < N; i++ {
		fmt.Printf("%v/%v\n\u001B[F", i, N)
		Render(img)
		if i%5 == 0 {
			Encode(img, outfname, float64(i+1))
		}
	}

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
		for j := 0; j < W; j++ {
			nPix++
			y0 := (-float64(i) + aa() + float64(H)/2) / float64(H)
			x0 := (float64(j) + aa() - float64(W)/2) / float64(H)

			start := Vec{x0, y0, 0}
			dir := Vec{0, 0, 1}
			if *focalLen != 0 {
				dir = start.Sub(focal).Normalized()
			}
			r := Ray{start, dir}

			v := Intensity(r, 0, true)

			img[i][j] += v
		}
	}
}

// Anti-aliasing jitter
func aa() float64 {
	return 0
	//return Rand()
}

func Intensity(r Ray, N int, includeSources bool) float64 {
	if N == *maxRec {
		return 0
	}
	_, obj := FirstIntersect(r, objects)
	if obj != nil {
		return 1 //obj.Shader.Intensity(r, t, n, N)
	} else {
		return 0 //ambient(r.Dir)
	}
}

func FirstIntersect(r Ray, objs []*Obj) (float64, *Obj) {
	var (
		nearestT        = -inf
		nearestObj *Obj = nil
	)

	for _, o := range objs {
		ival := o.Inters(r)
		if ival.Min > nearestT && ival.Max > 0 {
			nearestT = ival.Min
			nearestObj = o
		}
	}
	return nearestT, nearestObj
}

func Init() {
	flag.Parse()
	if *pprof != "" {
		go func() {
			log.Fatal(http.ListenAndServe(*pprof, nil))
		}()
	}
}
