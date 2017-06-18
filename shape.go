package main

type Shape interface {
	Intersect(Ray) Inter // TODO: -> Intersect
}

func Normal(s Shape, r Ray) Vec {
	i0 := s.Intersect(r)
	c := r.At(i0.Min)

	ra := r
	ra.Dir = ra.Dir.MAdd(1e-4, RandVec(r.Dir)).Normalized()
	i1 := s.Intersect(ra)
	a := ra.At(i1.Min)

	rb := r
	rb.Dir = rb.Dir.MAdd(1e-4, RandVec(r.Dir)).Normalized()
	i2 := s.Intersect(rb)
	b := rb.At(i2.Min)

	a = a.Sub(c)
	b = b.Sub(c)
	return b.Cross(a).Normalized()
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
