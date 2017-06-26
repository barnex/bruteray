package main

type Shape interface {
	Hit(r *Ray) float64
	Normal(r *Ray, t float64) Vec
}

type shape struct {
	hit    func(r *Ray) float64
	normal func(r *Ray, t float64) Vec
}

func (s *shape) Hit(r *Ray) float64 {
	return s.hit(r)
}

func (s *shape) Normal(r *Ray, t float64) Vec {
	return s.normal(r, t)
}

// Numerical approximation of normal vector
func NumNormal(s Shape, r *Ray, t float64) Vec {
	t0 := s.Hit(r)
	c := r.At(t0)

	perp1 := Vec{-r.Dir.Z, 0, r.Dir.X}.Normalized()
	perp2 := Vec{0, -r.Dir.Z, r.Dir.Y}.Normalized()

	const diff = 1. / (1024 * 1024)
	ra := r
	ra.Dir = r.Dir.MAdd(t*diff, perp1).Normalized()
	t1 := s.Hit(ra)
	a := ra.At(t1)

	rb := r
	rb.Dir = r.Dir.MAdd(t*diff, perp2).Normalized()
	t2 := s.Hit(rb)
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

func (s *shapeAnd) Hit(r *Ray) float64 {
	a := s.a.Inters(r)
	if !a.OK() {
		return 0
	}
	b := s.b.Inters(r)
	if !b.OK() {
		return 0
	}
	t := a.And(b).Min
	return Max(0, t)
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
