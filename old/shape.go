package main

type Shape interface {
	Inters(r *Ray) Interval
	Normal(r *Ray, t float64) Vec
}

// Numerical approximation of normal vector
func NumNormal(s Shape, r *Ray, t float64) Vec {
	t0 := s.Inters(r).Min
	if t0 < 0 {
		return Vec{}
	}
	c := r.At(t0)

	perp1 := Vec{-r.Dir.Z, 0, r.Dir.X}.Normalized()
	perp2 := r.Dir.Cross(perp1) // Thanks trijnewijn.

	const diff = 1. / (1024 * 1024)
	ra := r
	ra.Dir = r.Dir.MAdd(t*diff, perp1).Normalized()
	t1 := s.Inters(ra).Min
	if t1 < 0 {
		return Vec{}
	}
	a := ra.At(t1)

	rb := r
	rb.Dir = r.Dir.MAdd(t*diff, perp2).Normalized()
	t2 := s.Inters(rb).Min
	if t2 < 0 {
		return Vec{}
	}
	b := rb.At(t2)

	a = a.Sub(c)
	b = b.Sub(c)
	n := b.Cross(a).Normalized()

	return n.Towards(r.Dir)
}

func (n Vec) Towards(d Vec) Vec {
	if n.Dot(d) > 0 {
		return n.Mul(-1)
	}
	return n
}

func Transf(s Shape, T *Matrix4) Shape {
	return &transShape{s, *T}
}

type transShape struct {
	orig   Shape
	transf Matrix4
}

var _ Shape = &transShape{}

func (s *transShape) Inters(r *Ray) Interval {
	r2 := transRay(r, &s.transf)
	return s.orig.Inters(&r2)
}

func (s *transShape) Normal(r *Ray, t float64) Vec {
	return NumNormal(s, r, t)
}

func transRay(r *Ray, T *Matrix4) Ray {
	return Ray{
		Start: T.TransfPoint(r.Start),
		Dir:   T.TransfDir(r.Dir),
	}
}