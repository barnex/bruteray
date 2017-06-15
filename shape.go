package main

import "math"

type Shape interface {
	Inters(r Ray) Inter
}

func Sphere(c Vec, r float64) sphere {
	return sphere{c, r}
}

type sphere struct {
	c Vec
	r float64
}

func (s sphere) Inters(ray Ray) Inter {
	v := ray.Start.Sub(s.c)
	d := ray.Dir
	r := s.r
	D := sqr(v.Dot(d)) - (v.Len2() - sqr(r))
	if D < 0 {
		return empty
	}
	t1 := (-v.Dot(d) - math.Sqrt(D))
	t2 := (-v.Dot(d) + math.Sqrt(D))
	assert(t1 <= t2)

	return Inter{t1, t2}
}

func (s sphere) Transl(dx, dy, dz float64) sphere {
	return sphere{s.c.Add(Vec{dx, dy, dz}), s.r}
}

func And(a, b Shape) Shape {
	return and{a, b}
}

type and struct {
	a, b Shape
}

func (s and) Inters(r Ray) Inter {
	a := s.a.Inters(r)
	b := s.b.Inters(r)
	return a.And(b)
}
