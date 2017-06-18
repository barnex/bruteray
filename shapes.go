package main

import "math"

type Sphere struct {
	C Vec
	R float64
	Color
}

func (s *Sphere) Intensity(ray Ray, t float64) Color {
	return s.Color
}

func (s *Sphere) Inters(ray Ray) (Inter, Obj) {
	v := ray.Start.Sub(s.C)
	r := s.R
	d := ray.Dir
	D := sqr(v.Dot(d)) - (v.Len2() - sqr(r))
	if D < 0 {
		return empty, nil
	}
	t1 := (-v.Dot(d) - math.Sqrt(D))
	t2 := (-v.Dot(d) + math.Sqrt(D))
	assert(t1 <= t2)

	return Inter{t1, t2}, s
}

func (s *Sphere) Transl(dx, dy, dz float64) Sphere {
	return Sphere{s.C.Add(Vec{dx, dy, dz}), s.R, s.Color}
}

type And struct {
	a, b Obj
}

func (s And) Inters(r Ray) (Inter, Obj) {
	a, A := s.a.Inters(r)
	if !a.OK() {
		return a, A
	}
	b, B := s.b.Inters(r)

	ival := a.And(b)

	if ival.Empty() {
		return ival, nil
	}
	// TODO: optimize
	if ival.Min == a.Min || ival.Min == a.Max {
		return ival, A
	}
	if ival.Min == b.Min || ival.Min == b.Max {
		return ival, B
	}
	panic("bug")
}

func (s *And) Intensity(r Ray, t float64) Color {
	panic("not supposed to be called, passed on to sub-objects")
}
