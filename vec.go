package main

import "math"

type Vec struct {
	X, Y, Z float64
}

func (a Vec) Add(b Vec) Vec {
	return Vec{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vec) Sub(b Vec) Vec {
	return Vec{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vec) Dot(b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (v Vec) Mul(a float64) Vec {
	return Vec{a * v.X, a * v.Y, a * v.Z}
}

func (v Vec) Div(a float64) Vec {
	return v.Mul(1 / a)
}

func (v Vec) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec) Normalized() Vec {
	l := v.Len()
	if l == 0 {
		l = 1
	}
	return v.Div(l)
}
