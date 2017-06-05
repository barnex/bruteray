package main

import (
	"flag"
	"math"
)

var (
	width  = flag.Int("w", 300, "canvas width")
	height = flag.Int("h", 200, "canvas height")
)

var (
	Focal = Vec{0, 0, -1}
	Horiz = 10.0
)

func main() {

	flag.Parse()
	W := *width
	H := *height

	img := make([][]float64, H)
	for i := range img {
		img[i] = make([]float64, W)
	}

	scene := cubeFrame().RotY(-0.5).Transl(0, -0.2, 2)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {

			y0 := (-float64(i) + float64(H)/2 + 0.5) / float64(H)
			x0 := (float64(j) - float64(W)/2 + 0.5) / float64(H)

			start := Vec{x0, y0, 0}
			r := Ray{start, start.Sub(Focal).Normalized()}

			l := Vec{0.2, 0.8, -1}.Normalized().Mul(1.2)
			n, ok := Normal(r, scene)
			v := n.Dot(l) + 0.02
			if v < 0 {
				v = 0
			}
			if v > 1 {
				v = 1
			}
			if ok {
				img[i][j] = v
			}

		}
	}

	Encode(img, "out.jpg")
}

func cubeFrame() Shape {
	const (
		X = 1
		Y = 0.5
		Z = 1
		D = 0.2
	)
	return Slab(X, Y, Z).Sub(Slab(X, Y-D, Z-D)).Sub(Slab(X-D, Y, Z-D)).Sub(Slab(X-D, Y-D, Z))
}

const (
	fine = 0.01
	tol  = 1e-12
)

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

	assert(!s(r.At(out)))
	assert(s(r.At(in)))

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

func Normal(r Ray, s Shape) (Vec, bool) {
	c, ok := Bisect(r, s)
	if !ok {
		return Vec{}, false
	}

	ra := r
	ra.Dir = ra.Dir.Add(Vec{1e-3, 0, 0})
	a, okA := Bisect(ra, s)

	rb := r
	rb.Dir = rb.Dir.Add(Vec{0, 1e-3, 0})
	b, okB := Bisect(rb, s)

	if !okA || !okB {
		return Vec{}, false
	}

	a = a.Sub(c)
	b = b.Sub(c)

	return b.Cross(a).Normalized(), true

}

type Ray struct {
	Start Vec
	Dir   Vec
}

func (r *Ray) At(t float64) Vec {
	return r.Start.Add(r.Dir.Mul(t))
}

func assert(t bool) {
	if !t {
		panic("assertion failed")
	}
}
