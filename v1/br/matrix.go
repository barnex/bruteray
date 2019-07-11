package br

import (
	"bytes"
	"fmt"
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

func (a *Matrix4) Inv() *Matrix4 {
	assert(a[3] == Vec4{0, 0, 0, 1})
	i := Matrix3{
		{a[0][0], a[1][0], a[2][0]},
		{a[0][1], a[1][1], a[2][1]},
		{a[0][2], a[1][2], a[2][2]},
	}
	b := Vec{a[0][3], a[1][3], a[2][3]}
	return &Matrix4{
		{i[0][0], i[0][1], i[0][2], -b.Dot(i[0])},
		{i[1][0], i[1][1], i[1][2], -b.Dot(i[1])},
		{i[2][0], i[2][1], i[2][2], -b.Dot(i[2])},
		{0, 0, 0, 1},
	}
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
	c := cos(θ)
	s := sin(θ)
	return &Matrix4{
		{1, 0, 0, 0},
		{0, c, -s, 0},
		{0, s, c, 0},
		{0, 0, 0, 1},
	}
}

func RotY4(θ float64) *Matrix4 {
	c := cos(θ)
	s := sin(θ)
	return &Matrix4{
		{c, 0, -s, 0},
		{0, 1, 0, 0},
		{s, 0, c, 0},
		{0, 0, 0, 1},
	}
}

func RotZ4(θ float64) *Matrix4 {
	c := cos(θ)
	s := sin(θ)
	return &Matrix4{
		{c, -s, 0, 0},
		{s, c, 0, 0},
		{0, 0, 1, 0},
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

type Matrix3 [3]Vec
