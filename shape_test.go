package main

import "testing"

var (
	EX = Vec{1, 0, 0}
	EY = Vec{0, 1, 0}
	EZ = Vec{0, 0, 1}
)

func TestSphere(t *testing.T) {

	sp := Sphere(Vec{}, 1)

	cases := []struct {
		s      *sphere
		r      Ray
		w1, w2 Vec
		ok     bool
	}{
		{sp, ray(Vec{-2, 0, 0}, EX), Vec{-1, 0, 0}, Vec{1, 0, 0}, true},
		{sp, ray(Vec{-2, 0, 0}, EX.Mul(2)), Vec{-1, 0, 0}, Vec{1, 0, 0}, true},
		{sp, ray(Vec{0, -2, 0}, EY), Vec{0, -1, 0}, Vec{0, 1, 0}, true},
		{sp, ray(Vec{0, 2, 0}, EY), Vec{0, -1, 0}, Vec{0, 1, 0}, true},
		{sp, ray(Vec{2, 2, 0}, EY), Vec{}, Vec{}, false},
	}

	const tol = 1e-6

	for _, c := range cases {
		s := c.s
		r := c.r
		ok := c.ok
		have := c.s.Intersect(c.r)
		hok := have.OK()
		if hok != ok {
			t.Errorf("intersect %v %v: have %v, want %v", s, r, hok, ok)
		}

		h1 := r.At(have.Min)
		h2 := r.At(have.Max)
		if h1.Sub(c.w1).Len() > tol {
			t.Errorf("intersect %v %v: point 1: have %v, want %v", s, r, h1, c.w1)
		}
		if h1.Sub(c.w1).Len() > tol {
			t.Errorf("intersect %v %v: point 2: have %v, want %v", s, r, h2, c.w2)
		}
	}
}
