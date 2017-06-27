package main

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

func ShapeMinus(a, b Convex) Shape {
	return &shapeMinus{a, b}
}

type shapeMinus struct {
	a, b Convex
}

func (s *shapeMinus) Inters(r *Ray) Inter {
	a := s.a.Inters(r)
	if !a.OK() {
		return empty
	}
	b := s.b.Inters(r)
	return a.Minus(b).Normalize()
}

func (s *shapeMinus) Normal(r *Ray, t float64) Vec {
	return NumNormal(s, r, t)
}
