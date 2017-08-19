package bruteray

import "math"

const (
	Pi  = math.Pi
	Deg = Pi / 180
)

var (
	inf = math.Inf(1)
)

func Sqr(x float64) float64 {
	return x * x
}

func Sqrt(x float64) float64 {
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

// Rectify: max(x, 0)
func Re(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
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

func Sort(t0, t1 float64) (float64, float64) {
	if t0 < t1 {
		return t0, t1
	}
	return t1, t0
}
