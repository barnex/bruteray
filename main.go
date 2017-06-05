package main

import (
	"flag"
	"math"
)

const (
	W = 300
	H = 200
)

var (
	Focal = Vec{0, 0, -1}
	Horiz = 10.0
)

func main() {

	flag.Parse()

	img := make([][]float64, H)
	for i := range img {
		img[i] = make([]float64, W)
	}

	scene := Sphere(0.5).Transl(0, 0, 0.5)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {

			y0 := (float64(i) - H/2 + 0.5) / H
			x0 := (float64(j) - W/2 + 0.5) / H

			start := Vec{x0, y0, 0}
			r := Ray{start, start.Sub(Focal).Normalized()}

			inter, ok := Bisect(r, scene)
			if ok {
				img[i][j] = inter.Z * 2
			}

		}
	}

	Encode(img, "out.jpg")
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
