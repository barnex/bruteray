package geom

import "math"

const (
	Pi  = math.Pi
	Deg = Pi / 180
)

var Inf = math.Inf(1)

func sqr(x float64) float64 {
	return x * x
}

func sqrt(x float64) float64 {
	return math.Sqrt(x)
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

func TriangleNormal(a, b, c Vec) Vec {
	return ((b.Sub(a)).Cross(c.Sub(a))).Normalized() // .Add(a)?
}
