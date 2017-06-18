package main

import "math"

type Shape interface {
	Inters(Ray) Inter // TODO: -> Intersect
}

type Sphere struct {
	C Vec
	R float64
}

func (s *Sphere) Inters(ray Ray) Inter {
	v := ray.Start.Sub(s.C)
	r := s.R
	d := ray.Dir
	D := sqr(v.Dot(d)) - (v.Len2() - sqr(r))
	if D < 0 {
		return empty
	}
	t1 := (-v.Dot(d) - math.Sqrt(D))
	t2 := (-v.Dot(d) + math.Sqrt(D))
	assert(t1 <= t2)

	return Inter{t1, t2}
}

func (s *Sphere) Transl(dx, dy, dz float64) Sphere {
	return Sphere{s.C.Add(Vec{dx, dy, dz}), s.R}
}

type ShapeAnd struct {
	a, b Shape
}

func (s ShapeAnd) Inters(r Ray) Inter {
	a := s.a.Inters(r)
	if !a.OK() {
		return a
	}
	b := s.b.Inters(r)

	return a.And(b)
}

type ShapeMinus struct {
	a, b Shape
}

func (s ShapeMinus) Inters(r Ray) Inter {
	a := s.a.Inters(r)
	if !a.OK() {
		return a
	}
	b := s.b.Inters(r)

	return a.Minus(b)
}
