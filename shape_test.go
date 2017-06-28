package main

import "testing"

func TestSphere(t *testing.T) {

	sp := Sphere(Vec{}, 1)

	cases := []struct {
		s  *sphere
		r  Ray
		w1 Vec
		ok bool
	}{
		{sp, ray(Vec{-2, 0, 0}, Ex), Vec{-1, 0, 0}, true},
		{sp, ray(Vec{2, 0, 0}, Ex), Vec{}, false},
		{sp, ray(Vec{0, -2, 0}, Ey), Vec{0, -1, 0}, true},
		{sp, ray(Vec{0, 2, 0}, Ey), Vec{}, false},
		{sp, ray(Vec{2, 2, 0}, Ey), Vec{}, false},
	}

	const tol = 1e-6

	for _, c := range cases {
		s := c.s
		r := c.r
		ok := c.ok
		have := c.s.Hit(&c.r)
		hok := have > 0
		if hok != ok {
			t.Errorf("intersect %v %v: have %v, want %v", s, r, hok, ok)
		}

		if !ok {
			continue
		}

		h1 := r.At(have)
		if h1.Sub(c.w1).Len() > tol {
			t.Errorf("intersect %v %v: point 1: have %v, want %v", s, r, h1, c.w1)
		}
	}
}

func TestSheet(t *testing.T) {

	s := Sheet(2, Ey)

	cases := []struct {
		s  Shape
		r  Ray
		w  Vec
		ok bool
	}{
		{s, ray(Vec{5, 3, 1}, Ey.Mul(-1)), Vec{5, 2, 1}, true},
		{s, ray(Vec{5, 3, 1}, Ey.Mul(1)), Vec{}, false},
		{s, ray(Vec{5, 1, 1}, Ey.Mul(1)), Vec{5, 2, 1}, true},
		{s, ray(Vec{5, 1, 1}, Ey.Mul(-1)), Vec{}, false},
		{s, ray(Vec{2, 4, 9}, Vec{1, -1, 0}.Normalized()), Vec{4, 2, 9}, true},
		{s, ray(Vec{2, 4, 9}, Vec{-1, -1, 0}.Normalized()), Vec{0, 2, 9}, true},
		{s, ray(Vec{2, 4, 9}, Vec{-1, 1, 0}.Normalized()), Vec{}, false},
		{s, ray(Vec{2, 4, 9}, Vec{1, 1, 0}.Normalized()), Vec{}, false},
	}

	const tol = 1e-6

	for _, c := range cases {
		s := c.s
		r := c.r
		ok := c.ok
		ival := c.s.Inters(&c.r)
		have := ival.Min
		hok := ival.Min > 0
		if hok != ok {
			t.Errorf("intersect %v %v: have %v, want %v", s, r, hok, ok)
		}

		if !ok {
			continue
		}

		h1 := r.At(have)
		if h1.Sub(c.w).Len() > tol {
			t.Errorf("intersect %v %v: point 1: have %v, want %v", s, r, h1, c.w)
		}
	}
}
