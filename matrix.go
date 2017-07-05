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

type Matrix4 [4]Vec4

func (a *Matrix4) Mul(b *Matrix4) *Matrix4 {
	return &Matrix4{
		{a[0].Dot(b[0]), a[0].Dot(b[1]), a[0].Dot(b[2]), a[0].Dot(b[3])},
		{a[1].Dot(b[0]), a[1].Dot(b[1]), a[1].Dot(b[2]), a[1].Dot(b[3])},
		{a[2].Dot(b[0]), a[2].Dot(b[1]), a[2].Dot(b[2]), a[2].Dot(b[3])},
		{a[3].Dot(b[0]), a[3].Dot(b[1]), a[3].Dot(b[2]), a[3].Dot(b[3])},
	}
}

func (T *Matrix4) TransfPoint(v Vec) Vec {
	r := Vec4{v.X, v.Y, v.Z, 1}
	return Vec{r.Dot(T[0]), r.Dot(T[1]), r.Dot(T[2])}
}

func (T *Matrix4) TransfDir(v Vec) Vec {
	r := Vec4{v.X, v.Y, v.Z, 0}
	return Vec{r.Dot(T[0]), r.Dot(T[1]), r.Dot(T[2])}
}

func UnitMatrix4() Matrix4 {
	return Matrix4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func RotX4(θ float64) Matrix4 {
	c := math.Cos(θ)
	s := math.Sin(θ)
	return Matrix4{
		{1, 0, 0, 0},
		{0, c, s, 0},
		{0, -s, c, 0},
		{0, 0, 0, 1},
	}
}
