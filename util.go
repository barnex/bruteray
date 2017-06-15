package main

import "math"

const deg = math.Pi / 180

var inf = math.Inf(1)

func sqr(x float64) float64 {
	return x * x
}

func assert(t bool) {
	if !t {
		panic("assertion failed")
	}
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Min3(x, y, z float64) float64 {
	return Min(Min(x, y), z)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Max3(x, y, z float64) float64 {
	return Max(Max(x, y), z)
}
