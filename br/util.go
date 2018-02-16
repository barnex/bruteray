package br

import "math"

const (
	Pi  = math.Pi
	Deg = Pi / 180
)

var (
	inf = math.MaxFloat64 // math.Inf(1)
)

func sqr(x float64) float64 {
	return x * x
}

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func assert(t bool) {
	if !t {
		panic("assertion failed")
	}
}

func sin(x float64) float64 {
	switch x {
	default:
		return math.Sin(x)
	case Pi, -Pi:
		return 0
	}
}

func cos(x float64) float64 {
	switch x {
	default:
		return math.Cos(x)
	case Pi / 2, -Pi / 2:
		return 0
	}
}
