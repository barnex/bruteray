package main

import "math"

type cylinder struct {
	c  Vec
	r2 float64
}

func Cylinder(center Vec, radius float64) *cylinder {
	return &cylinder{center, sqr(radius)}
}

func (s *cylinder) Normal(r *Ray, t float64) Vec {
	return NumNormal(s, r, t)
}

func (s *cylinder) Hit(r *Ray) float64 {
	return s.Inters(r).Min
}

func (s *cylinder) Inters(r *Ray) Inter {
	s0 := r.Start.X - s.c.X
	s1 := r.Start.Z - s.c.Z
	d0 := r.Dir.X
	d1 := r.Dir.Z

	d02 := sqr(d0)
	d12 := sqr(d1)
	s02 := sqr(s0)
	s12 := sqr(s1)
	R := s.r2

	D := d02*R + d02*(-s12) + 2*d0*d1*s0*s1 + d12*R - d12*s02
	if D < 0 {
		return empty
	}
	t0 := (-math.Sqrt(D) - d0*s0 - d1*s1) / (d02 + d12)
	//if t0 > 0 {
	//	return t0
	//}

	t1 := (math.Sqrt(D) - d0*s0 - d1*s1) / (d02 + d12)
	//if t1 > 0 {
	//	return t1
	//}
	t0, t1 = Sort(t0, t1) // TODO, not needed, assert

	return Inter{t0, t1}
}
