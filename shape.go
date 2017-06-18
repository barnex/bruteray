package main

import "math"

//type Shape interface {
//	Inters(Ray, float64) (Inter, Shape)
//	//Normal(r Ray, t float64) Vec
//}

type Sphere struct {
	C Vec
	R float64
	Color
}

func (s *Sphere) Intensity(ray Ray, t float64) Color {
	return s.Color
}

func (s *Sphere) Inters(ray Ray) Inter {
	v := ray.Start.Sub(s.C)
	r := s.R
	d := ray.Dir
	D := sqr(v.Dot(d)) - (v.Len2() - sqr(r))
	if D < 0 {
		return empty
	}
	t1 := (-v.Dot(d) - math.Sqrt(D))
	t2 := (-v.Dot(d) + math.Sqrt(D))
	assert(t1 <= t2)

	return Inter{t1, t2}
}

func (s *Sphere) Transl(dx, dy, dz float64) Sphere {
	return Sphere{s.C.Add(Vec{dx, dy, dz}), s.R, s.Color}
}

type And struct {
	a, b Obj
}

func (s And) Inters(r Ray) Inter {
	a := s.a.Inters(r)
	if !a.OK() {
		return a
	}
	b := s.b.Inters(r)

	return a.And(b)

	//if ival.Empty() {
	//	return ival, nil
	//}
	//if ival.Min == a.Min || ival.Min == a.Max {
	//	return ival
	//}
	//if ival.Min == b.Min || ival.Min == b.Max {
	//	return ival
	//}
	//panic("bug")
}

func (s *And) Intensity(r Ray, t float64) Color {

	a := s.a.Inters(r)
	b := s.b.Inters(r)

	ival := a.And(b)

	if ival.Empty() {
		panic("bug")
	}
	if ival.Min == a.Min || ival.Min == a.Max {
		return s.a.Intensity(r, t)
	}
	if ival.Min == b.Min || ival.Min == b.Max {
		return s.b.Intensity(r, t)
	}
	panic("bug")
}
