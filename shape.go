package main

type Shape interface {
	Intersect(Ray) Inter // TODO: -> Intersect
}

func Normal(s Shape, r Ray, t float64) Vec {
	i0 := s.Intersect(r)
	c := r.At(i0.Min)

	ra := r
	ra.Dir = ra.Dir.MAdd(1e-5, Vec{1, 0, 0}).Normalized()
	i1 := s.Intersect(ra)
	a := ra.At(i1.Min)

	rb := r
	rb.Dir = rb.Dir.MAdd(1e-5, Vec{0, 1, 0}).Normalized()
	i2 := s.Intersect(rb)
	b := rb.At(i2.Min)

	a = a.Sub(c)
	b = b.Sub(c)
	n := b.Cross(a).Normalized()

	//if n.Dot(r.Dir) > 0 {
	//	n = n.Mul(-1)
	//}
	return n
}

type ShapeAnd struct {
	a, b Shape
}

func (s ShapeAnd) Intersect(r Ray) Inter {
	a := s.a.Intersect(r)
	if !a.OK() {
		return a
	}
	b := s.b.Intersect(r)

	return a.And(b)
}

type ShapeMinus struct {
	a, b Shape
}

func (s ShapeMinus) Intersect(r Ray) Inter {
	a := s.a.Intersect(r)
	if !a.OK() {
		return a
	}
	b := s.b.Intersect(r)

	return a.Minus(b)
}
