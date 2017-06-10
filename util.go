package main

import (
	"math"
	random "math/rand"
)

const deg = math.Pi / 180

func sqr(x float64) float64 {
	return x * x
}

func clip(v, min, max float64) float64 {
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	return v
}

var rng = random.New(random.NewSource(1))

func rand() float64 {
	return rng.Float64()
}

func randNorm() float64 {
	return rng.NormFloat64()
}

func randVec(n Vec) Vec {
	v := Vec{randNorm(), randNorm(), randNorm()}.Normalized()
	if v.Dot(n) < 0 {
		v = v.Mul(-1)
	}
	return v
}

func assert(t bool) {
	if !t {
		panic("assertion failed")
	}
}
