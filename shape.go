package main

type Shape interface {
	Intersect(Ray) Inter // TODO: -> Intersect
	Normal(Ray, float64) Vec
}

type ShapeFunc func(Ray) Inter

func (f ShapeFunc) Intersect(r Ray) Inter {
	return f(r)
}

func (f ShapeFunc) Normal(r Ray, t float64) Vec {
	return Normal(f, r, t)
}

// TODO: choose delta vectors perpendicular to ray
func Normal(s Shape, r Ray, t float64) Vec {
	i0 := s.Intersect(r)
	c := r.At(i0.Min)

	perp1 := Vec{-r.Dir.Z, 0, r.Dir.X}.Normalized()
	perp2 := Vec{0, -r.Dir.Z, r.Dir.Y}.Normalized()

	const diff = 1. / (1024 * 1024)
	ra := r
	ra.Dir = r.Dir.MAdd(t*diff, perp1).Normalized()
	i1 := s.Intersect(ra)
	a := ra.At(i1.Min)

	rb := r
	rb.Dir = r.Dir.MAdd(t*diff, perp2).Normalized()
	i2 := s.Intersect(rb)
	b := rb.At(i2.Min)

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

func (s ShapeAnd) Normal(r Ray, t float64) Vec {
	return Normal(s, r, t)
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

func (s ShapeMinus) Normal(r Ray, t float64) Vec {
	return Normal(s, r, t)
}
