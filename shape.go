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

func ShapeAnd(a, b Convex) Shape {
	return &shapeAnd{a, b}
}

type shapeAnd struct {
	a, b Convex
}

type Convex interface {
	Inters(r *Ray) Inter
}

func (s *shapeAnd) Inters(r *Ray) Inter {
	a := s.a.Inters(r)
	if !a.OK() {
		return empty
	}
	b := s.b.Inters(r)
	if !b.OK() {
		return empty
	}
	return a.And(b).Normalize()
}

func (s *shapeAnd) Normal(r *Ray, t float64) Vec {
	return NumNormal(s, r, t)
}

//type ShapeMinus struct {
//	a, b Shape
//}
//
//func (s ShapeMinus) Intersect(r Ray) Inter {
//	a := s.a.Intersect(r)
//	if !a.OK() {
//		return a
//	}
//	b := s.b.Intersect(r)
//
//	return a.Minus(b)
//}
//
//func (s ShapeMinus) Normal(r Ray, t float64) Vec {
//	return Normal(s, r, t)
//}
//
//func (s ShapeMinus) Hit(r Ray) float64 {
//	return s.Intersect(r).Min
//}
