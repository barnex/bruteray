package main

type Shape interface {
	Inside(r Vec) bool
}

type ShapeFn func(Vec) bool

func (f ShapeFn) Inside(r Vec) bool {
	return f(r)
}

func Sphere(r float64) Shape {
	return ShapeFn(func(x Vec) bool {
		return x.Dot(x) < r*r
	})
}
