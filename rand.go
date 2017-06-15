package main

import "math/rand"

var rng = rand.New(rand.NewSource(1))

func Rand() float64 {
	return rng.Float64()
}

func RandNorm() float64 {
	return rng.NormFloat64()
}

func RandVec(n Vec) Vec {
	v := Vec{RandNorm(), RandNorm(), RandNorm()}.Normalized()
	if v.Dot(n) < 0 {
		v = v.Mul(-1)
	}
	return v
}
