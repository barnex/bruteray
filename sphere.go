package main

import (
	"log"
	"math"
)

type sphere struct {
	C Vec
	R float64
}

func Sphere(center Vec, radius float64) *sphere {
	return &sphere{center, radius}
}

func (s *sphere) Normal(r Ray, t float64) Vec {
	return Normal(s, r, t)
}

func (s *sphere) Intersect(ray Ray) Inter {
	v := ray.Start.Sub(s.C)
	r := s.R
	d := ray.Dir
	D := sqr(v.Dot(d)) - (v.Len2() - sqr(r))
	if D < 0 {
		return empty
	}
	t1 := (-v.Dot(d) - math.Sqrt(D))
	t2 := (-v.Dot(d) + math.Sqrt(D))
	if t1 > t2 {
		log.Println("ERROR: sphere intersect t1=%v, t2=%v", t1, t2)
	}

	return Inter{t1, t2}
}

func (s *sphere) Transl(dx, dy, dz float64) *sphere {
	return &sphere{s.C.Add(Vec{dx, dy, dz}), s.R}
}
