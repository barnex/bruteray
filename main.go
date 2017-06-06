package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

var (
	width       = flag.Int("w", 300, "canvas width")
	height      = flag.Int("h", 200, "canvas height")
	focalLen    = flag.Float64("f", 1, "focal length")
	progressive = flag.Int("p", 16, "progressively increase resolution")
)

var (
	Focal = Vec{0, 0, -1}
)

// stats
var (
	nShade = 1
)

const (
	Horiz = 20.0
	fine  = 0.02
	tol   = 1e-9
)

const deg = math.Pi / 180

func main() {
	Init()
	start := time.Now()

	s := &Scene{
		light: Vec{-2, 4, 1},
		amb:   0.2,
		shape: coolSphere().RotX(-30*deg).Transl(0, 0, 6),
	}

	img := MakeImage(*width, *height)

	Render(s, img)
	fmt.Println(nShade, "evals", time.Since(start))
}

func MakeImage(W, H int) [][]float64 {
	img := make([][]float64, H)
	for i := range img {
		img[i] = make([]float64, W)
	}
	return img
}

func Render(s *Scene, img [][]float64) {
	for sub := *progressive; sub > 0; sub /= 2 {
		refine(s, img, sub, sub == *progressive)
		Encode(img, "out.jpg")
	}
}

func refine(s *Scene, img [][]float64, sub int, first bool) {
	W := *width
	H := *height
	for i := 0; i < H; i += sub {
		fmt.Println(sub, i, "/", H)
		for j := 0; j < W; j += sub {
			if i%(2*sub) == 0 && j%(2*sub) == 0 && !first {
				continue
			}
			y0 := (-float64(i) + float64(H)/2 + 0.5) / float64(H)
			x0 := (float64(j) - float64(W)/2 + 0.5) / float64(H)
			start := Vec{x0, y0, 0}
			r := Ray{start, start.Sub(Focal).Normalized()}
			v := PixelShade(s, r)
			for I := i; I < i+sub && I < H; I++ {
				for J := j; J < j+sub && J < W; J++ {
					img[I][J] = v
				}
			}
		}
	}
}

type Scene struct {
	light Vec
	amb   float64
	shape Shape
}

func PixelShade(scene *Scene, r Ray) float64 {
	nShade++
	c, n, ok := Normal(r, scene.shape)
	if !ok {
		return 0
	}
	d := scene.light.Sub(c).Normalized()

	secondary := Ray{c.MAdd(0.01, d), d}
	v := scene.amb
	if !inters(secondary, scene.shape) {
		v = 0.8*n.Dot(d) + scene.amb
	}
	v = clip(v, 0, 1)
	return v
}

func inters(r Ray, s Shape) bool {
	_, ok := Inters(r, s)
	return ok
}

func Inters(r Ray, s Shape) (float64, bool) {
	for t := 0.0; t < Horiz; t += fine {
		if s(r.At(t)) {
			return t, true
		}
	}
	return 0, false
}

func Bisect(r Ray, s Shape) (Vec, bool) {

	in, ok := Inters(r, s)
	if !ok {
		return Vec{}, false
	}

	out := in - fine

	if s(r.At(out)) || !s(r.At(in)) {
		return Vec{}, false
	}

	for math.Abs(in-out)/(in+out) > tol {
		mid := (in + out) / 2
		if s(r.At(mid)) {
			in = mid
		} else {
			out = mid
		}
	}
	return r.At(in), true
}

func Normal(r Ray, s Shape) (Vec, Vec, bool) {
	c, ok := Bisect(r, s)
	if !ok {
		return Vec{}, Vec{}, false
	}

	ra := r
	ra.Dir = ra.Dir.Add(Vec{1e-5, 0, 0})
	a, okA := Bisect(ra, s)

	rb := r
	rb.Dir = rb.Dir.Add(Vec{0, 1e-5, 0})
	b, okB := Bisect(rb, s)

	if !okA || !okB {
		return Vec{}, Vec{}, false
	}

	a = a.Sub(c)
	b = b.Sub(c)

	n := b.Cross(a).Normalized()
	return c, n, true

}

func clip(v, min, max float64) float64 {
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	return v
}

func assert(t bool) {
	if !t {
		panic("assertion failed")
	}
}

func Init() {
	flag.Parse()
	Focal = Vec{0, 0, -*focalLen}
}
