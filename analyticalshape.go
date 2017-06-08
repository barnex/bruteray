package main

import "math"

type ShapeFunc func(r Ray) (t float64, normal Vec, intersect bool)

func (f ShapeFunc) Normal(r Ray) (t float64, normal Vec, intersect bool) {
	return f(r)
}

func ASphere(c Vec, r float64) Shape {
	return ShapeFunc(func(ray Ray) (float64, Vec, bool) {
		v := ray.Start.Sub(c)
		d := ray.Dir
		D := sqr(v.Dot(d)) - (v.Dot(v) - r*r)
		if D < 0 {
			return 0, Vec{}, false
		}
		t := (-v.Dot(d) - math.Sqrt(D)) / d.Len2()
		n := ray.At(t).Sub(c).Normalized()

		return t, n, true
	})
}
