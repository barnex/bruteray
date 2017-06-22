package main

import "math"

type Matrix [3]Vec

func (r Vec) Transf(T *Matrix) Vec {
	return Vec{r.Dot(T[0]), r.Dot(T[1]), r.Dot(T[2])}
}

func (a *Matrix) Mul(b *Matrix) *Matrix {
	return &Matrix{
		{a[0].Dot(b[0]), a[0].Dot(b[1]), a[0].Dot(b[2])},
		{a[1].Dot(b[0]), a[1].Dot(b[1]), a[1].Dot(b[2])},
		{a[2].Dot(b[0]), a[2].Dot(b[1]), a[2].Dot(b[2])},
	}
}

func RotX(θ float64) Matrix {
	c := math.Cos(θ)
	s := math.Sin(θ)
	return Matrix{
		{1, 0, 0},
		{0, c, s},
		{0, -s, c},
	}
}

func UnitMatrix() Matrix {
	return Matrix{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}
