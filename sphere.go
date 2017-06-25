package main

import "math"

type sphere struct {
	c  Vec
	r2 float64
}

func Sphere(center Vec, radius float64) *sphere {
	return &sphere{center, sqr(radius)}
}

func (s *sphere) Normal(r *Ray, t float64) Vec {
	return r.At(t).Sub(s.c).Normalized()
}

func (s *sphere) Hit(r *Ray) float64 {
	v := r.Start.Sub(s.c)
	d := r.Dir
	vd := v.Dot(d)
	D := sqr(vd) - (v.Len2() - s.r2)
	if D < 0 {
		return 0
	}
	if t1 := (-vd - math.Sqrt(D)); t1 > 0 {
		return t1
	}
	if t2 := (-vd + math.Sqrt(D)); t2 > 0 {
		return t2
	}

	return 0
}

func (s *sphere) Inters(r *Ray) Inter {
	v := r.Start.Sub(s.c)
	d := r.Dir
	vd := v.Dot(d)
	D := sqr(vd) - (v.Len2() - s.r2)
	if D < 0 {
		return empty
	}
	t1 := (-vd - math.Sqrt(D))
	t2 := (-vd + math.Sqrt(D))
	t1, t2 = Sort(t1, t2)
	return Inter{t1, t2}
}

func (s *sphere) Transl(dx, dy, dz float64) *sphere {
	return &sphere{s.c.Add(Vec{dx, dy, dz}), s.r2}
}

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
	s0 := r.Start.X - s.c.X
	s1 := r.Start.Z - s.c.Z
	d0 := r.Dir.X
	d1 := r.Dir.Z

	d02 := sqr(d0)
	d12 := sqr(d1)
	s02 := sqr(s0)
	s12 := sqr(s1)
	R := s.r2

	t0 := (-math.Sqrt(d02*R+d02*(-s12)+2*d0*d1*s0*s1+d12*R-d12*s02) - d0*s0 - d1*s1) / (d02 + d12)
	if t0 > 0 {
		return t0
	}

	t1 := (math.Sqrt(d02*R+d02*(-s12)+2*d0*d1*s0*s1+d12*R-d12*s02) - d0*s0 - d1*s1) / (d02 + d12)
	if t1 > 0 {
		return t1
	}
	return 0
}
