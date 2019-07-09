package builder

//TODO: use package util

func sort2(a, b float64) (float64, float64) {
	if a < b {
		return a, b
	}
	return b, a
}

func frontSolution(t1, t2 float64, max float64) float64 {
	if t1 > t2 {
		t1, t2 = t2, t1
	}
	if t1 > 0 && t1 < max {
		return t1
	}
	if t2 > 0 && t2 < max {
		return t2
	}
	return 0
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
