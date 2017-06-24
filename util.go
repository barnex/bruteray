package main

import "math"

const (
	pi  = math.Pi
	deg = pi / 180
)

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
	if x < y {
		return x
	}
	return y
}

func Min3(x, y, z float64) float64 {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}

func Max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func Max3(x, y, z float64) float64 {
	max := x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}
