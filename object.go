package main

type Obj interface {
	Inters(Ray) (Inter, Obj)      // TODO: -> Intersect
	Intensity(Ray, float64) Color // TODO: -> Color
}
