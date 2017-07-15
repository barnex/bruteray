package r

import (
	"bytes"
	"fmt"
	"math"
)

type Matrix4 [4]Vec4

func (a *Matrix4) Mul(b *Matrix4) *Matrix4 {
	c := new(Matrix4)
	for i := range c {
		for j := range c[i] {
			for k := range b {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func (T *Matrix4) TransfPoint(v Vec) Vec {
	r := Vec4{v[X], v[Y], v[Z], 1}
	return Vec{r.Dot(T[0]), r.Dot(T[1]), r.Dot(T[2])}
}

func (T *Matrix4) TransfDir(v Vec) Vec {
	r := Vec4{v[X], v[Y], v[Z], 0}
	return Vec{r.Dot(T[0]), r.Dot(T[1]), r.Dot(T[2])}
}

func UnitMatrix4() *Matrix4 {
	return &Matrix4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func RotX4(θ float64) *Matrix4 {
	c := math.Cos(θ)
	s := math.Sin(θ)
	return &Matrix4{
		{1, 0, 0, 0},
		{0, c, -s, 0},
		{0, s, c, 0},
		{0, 0, 0, 1},
	}
}

func RotY4(θ float64) *Matrix4 {
	c := math.Cos(θ)
	s := math.Sin(θ)
	return &Matrix4{
		{c, 0, -s, 0},
		{0, 1, 0, 0},
		{s, 0, c, 0},
		{0, 0, 0, 1},
	}
}

func Transl4(d Vec) *Matrix4 {
	return &Matrix4{
		{1, 0, 0, d[X]},
		{0, 1, 0, d[Y]},
		{0, 0, 1, d[Z]},
		{0, 0, 0, 1},
	}
}

func (a *Matrix4) String() string {
	var b bytes.Buffer
	for _, v := range a {
		fmt.Fprintln(&b, v)
	}
	return b.String()
}
