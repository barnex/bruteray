package main

type Obj interface {
	// Returns the object-ray intersection interval (possibly empty)
	// and a shader that will provide color.
	Intersect(Ray) (Inter, Shader)
}

// An object composed of the intersection ("and") of two objects.
type ObjAnd struct {
	a, b Obj
}

func (s ObjAnd) Intersect(r Ray) (Inter, Shader) {
	a, A := s.a.Intersect(r)
	if !a.OK() {
		return a, A
	}
	b, B := s.b.Intersect(r)

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
