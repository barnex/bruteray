package main

type Obj interface {
	Inters(Ray) Inter
	Intensity(Ray, float64) Color
}
