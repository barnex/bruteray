package main

import "math"

type Shape interface {
	Inters(Ray, float64) (Inter, Shape)
	//Normal(r Ray, t float64) Vec
}

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

//func And(a, b Shape) Shape {
//	return and{a, b}
//}
//
//type and struct {
//	a, b Shape
//}
//
//func (s and) Inters(r Ray) (Inter, Shape) {
//	a, A := s.a.Inters(r)
//	b, B := s.b.Inters(r)
//	ival := a.And(b)
//
//	if ival.Empty() {
//		return ival, nil
//	}
//	if ival.Min == a.Min || ival.Min == a.Max {
//		return ival, A
//	}
//	if ival.Min == b.Min || ival.Min == b.Max {
//		return ival, B
//	}
//	panic("bug")
//}
