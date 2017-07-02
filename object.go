package main

type Obj interface {
	Shape
	Shade(e *Env, r *Ray, t float64, N int) Color
}

type object struct {
	Shape
	shader Shader
}

func (o *object) Shade(e *Env, r *Ray, t float64, N int) Color {
	if N == 0 {
		return e.Ambient(r.Dir)
	}
	n := o.Shape.Normal(r, t)
	return o.shader.Shade(e, r, t, n, N)
}

// An object composed of the intersection ("and") of two objects.
type objAnd struct {
	a, b Obj
}

func (s *objAnd) Inters(r *Ray) Inter {
	a := s.a.Inters(r)
	if !a.OK() {
		return empty
	}
	b := s.b.Inters(r)
	return a.And(b)
}

func (s *objAnd) Normal(r *Ray, t float64) Vec {
	return NumNormal(s, r, t)
}

func (s *objAnd) Shade(e *Env, r *Ray, t float64, N int) Color {

	a := s.a.Inters(r)
	assert(a.OK())

	b := s.b.Inters(r)
	ival := a.And(b)

	// TODO: optimize
	if ival.Min == a.Min || ival.Min == a.Max {
		return s.a.Shade(e, r, t, N)
	}
	if ival.Min == b.Min || ival.Min == b.Max {
		return s.b.Shade(e, r, t, N)
	}
	panic("bug")
}

// An object composed of the intersection ("and") of two objects.
type objMinus struct {
	a, b Obj
}

func (s *objMinus) Inters(r *Ray) Inter {
	a := s.a.Inters(r)
	if !a.OK() {
		return empty
	}
	b := s.b.Inters(r)
	return a.Minus(b)
}

func (s *objMinus) Normal(r *Ray, t float64) Vec {
	return NumNormal(s, r, t)
}

func (s *objMinus) Shade(e *Env, r *Ray, t float64, N int) Color {

	a := s.a.Inters(r)
	assert(a.OK())

	b := s.b.Inters(r)
	ival := a.Minus(b)

	// TODO: optimize
	if ival.Min == a.Min || ival.Min == a.Max {
		return s.a.Shade(e, r, t, N)
	}
	if ival.Min == b.Min || ival.Min == b.Max {
		return s.b.Shade(e, r, t, N)
	}
	panic("bug")
}
