package bruteray

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

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func min3(x, y, z float64) float64 {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

// Rectify: max(x, 0)
func re(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}

func max3(x, y, z float64) float64 {
	max := x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}

func sort2(t0, t1 float64) (float64, float64) {
	if t0 < t1 {
		return t0, t1
	}
	return t1, t0
}
