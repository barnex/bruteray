package r

import (
	"math"
)

type Shape interface {
	Inters(r *Ray) Interval
	Normal(pos Vec) Vec
}

// -- sphere

func Sphere(center Vec, radius float64) *sphere {
	return &sphere{center, Sqr(radius)}
}

type sphere struct {
	c  Vec
	r2 float64
}

func (s *sphere) Normal(pos Vec) Vec {
	n := pos.Sub(s.c).Normalized()
	return n
}

func (s *sphere) Inters(r *Ray) Interval {
	v := r.Start.Sub(s.c)
	d := r.Dir
	vd := v.Dot(d)
	D := Sqr(vd) - (v.Len2() - s.r2)
	if D < 0 {
		return Interval{}
	}
	t1 := (-vd - math.Sqrt(D))
	t2 := (-vd + math.Sqrt(D))
	return Interv(t1, t2)
}

// -- half-space

func Sheet(dir Vec, off float64) *sheet {
	return &sheet{dir, off}
}

type sheet struct {
	dir Vec
	off float64
}

func (s *sheet) Normal(pos Vec) Vec {
	return s.dir
}

func (s *sheet) Inters(r *Ray) Interval {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t := (s.off - rs) / rd
	return Interval{t, t}
}