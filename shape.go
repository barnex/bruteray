package main

type Shape interface {
	Inters(r *Ray) Inter
	Normal(r *Ray, t float64) Vec
}

type shape struct {
	inters func(r *Ray) Inter
	normal func(r *Ray, t float64) Vec
}

func (s *shape) Inters(r *Ray) Inter {
	return s.inters(r)
}

func (s *shape) Normal(r *Ray, t float64) Vec {
	return s.normal(r, t)
}

// Numerical approximation of normal vector
func NumNormal(s Shape, r *Ray, t float64) Vec {
	t0 := s.Inters(r).Min
	if t0 < 0 {
		return Vec{}
	}
	c := r.At(t0)

	perp1 := Vec{-r.Dir.Z, 0, r.Dir.X}.Normalized()
	perp2 := Vec{0, -r.Dir.Z, r.Dir.Y}.Normalized()

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
