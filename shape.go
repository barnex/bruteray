package main

//type Shape interface {
//	Hit(Ray) float64
//	Intersect(Ray) Inter // TODO: -> Intersect
//	Normal(Ray, float64) Vec
//}
//
//type ShapeFunc func(Ray) Inter
//
//func (f ShapeFunc) Intersect(r Ray) Inter {
//	return f(r)
//}
//
//func (f ShapeFunc) Normal(r Ray, t float64) Vec {
//	return Normal(f, r, t)
//}
//
//func (f ShapeFunc) Hit(r Ray) float64 {
//	return f.Intersect(r).Min
//}

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

//type ShapeAnd struct {
//	a, b Shape
//}
//
//func (s ShapeAnd) Intersect(r Ray) Inter {
//	a := s.a.Intersect(r)
//	if !a.OK() {
//		return a
//	}
//	b := s.b.Intersect(r)
//
//	return a.And(b)
//}
//
//func (s ShapeAnd) Normal(r Ray, t float64) Vec {
//	return Normal(s, r, t)
//}
//
//func (s ShapeAnd) Hit(r Ray) float64 {
//	return s.Intersect(r).Min
//}
//
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
