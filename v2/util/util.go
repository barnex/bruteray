package util

import (
	"fmt"
	"math"
	"math/rand"
)

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

func Max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func Re(x float64) float64 {
	return Max(x, 0)
}

func Frac(x float64) float64 {
	return x - math.Floor(x)
}

func Dither(rng *rand.Rand, x float64) int {
	i, f := math.Modf(x)
	if rng.Float64() < f {
		i++
	}
	return int(i)
}

func Sqr(x float64) float64 {
	return x * x
}

func Assert(test bool) {
	if !test {
		panic("assertion failed")
	}
}

func IsBad(v float64) bool {
	return math.IsNaN(v) || math.IsInf(v, 1) || math.IsInf(v, -1)
}

func IsBadVec(v [3]float64) bool {
	return IsBad(v[0]) || IsBad(v[1]) || IsBad(v[2])
}

func CheckNaN(v float64) {
	if IsBad(v) {
		panic(fmt.Sprint("got bad number:", v))
	}
}

func CheckNaNVec(v [3]float64) {
	if math.IsNaN(v[0]) || math.IsNaN(v[1]) || math.IsNaN(v[2]) {
		panic(fmt.Sprint("got NaN:", v))
	}
}
