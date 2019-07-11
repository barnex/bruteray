package shape

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

func sqr(x float64) float64 {
	return x * x
}
