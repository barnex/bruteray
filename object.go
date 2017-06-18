package main

// An object has a shape (provided by Inters) and color (provided by Intensity).
type Obj interface {
	Inters(Ray) (Inter, Obj)      // TODO: -> Intersect
	Intensity(Ray, float64) Color // TODO: -> Color
}

// An object composed of shape and color
type Shaded struct {
	s Shape
	c Color
}

func (s *Shaded) Intensity(ray Ray, t float64) Color {
	return s.c
}

func (s *Shaded) Inters(ray Ray) (Inter, Obj) {
	return s.s.Inters(ray), s
}

// An object composed of the intersection ("and") of two objects.
type ObjAnd struct {
	a, b Obj
}

func (s ObjAnd) Inters(r Ray) (Inter, Obj) {
	a, A := s.a.Inters(r)
	if !a.OK() {
		return a, A
	}
	b, B := s.b.Inters(r)

	ival := a.And(b)

	if ival.Empty() {
		return ival, nil
	}
	// TODO: optimize
	if ival.Min == a.Min || ival.Min == a.Max {
		return ival, A
	}
	if ival.Min == b.Min || ival.Min == b.Max {
		return ival, B
	}
	panic("bug")
}

func (s *ObjAnd) Intensity(r Ray, t float64) Color {
	panic("not supposed to be called, passed on to sub-objects")
}
