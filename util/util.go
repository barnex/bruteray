// Package util provides miscellaneous maths functions.
package util

import (
	"fmt"
	"math"
	"math/rand"
)

func Max(x, y float64) float64 {
	var m float64
	if x > y {
		m = x
	} else {
		m = y
	}
	return m
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

func Min(x, y float64) float64 {
	var m float64
	if x < y {
		m = x
	} else {
		m = y
	}
	return m
}

func Re(x float64) float64 {
	return Max(x, 0)
}

// Frac returns x % 1, using python-like modulus.
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

// Sin is like math.Sin, but returns exactly 0 for arguments Pi, -Pi.
// this avoids annoying round-off errors for rotations of multiples of 90 degrees.
func Sin(x float64) float64 {
	switch x {
	default:
		return math.Sin(x)
	case pi, -pi:
		return 0
	}
}

// Cos is like math.Cos, but returns exactly 0 for arguments Pi/2, -Pi/2.
// this avoids annoying round-off errors for rotations of multiples of 90 degrees.
func Cos(x float64) float64 {
	switch x {
	default:
		return math.Cos(x)
	case pi / 2, -pi / 2:
		return 0
	}
}

// Sinc returns sin(x)/x, handling x == 0 gracefully.
func Sinc(x float64) float64 {
	if x == 0 {
		return 1
	}
	return math.Sin(x) / x
}

const pi = math.Pi

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
