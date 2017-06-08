package main

type Shape interface {
	Normal(r Ray) (float64, Vec, bool)
}
