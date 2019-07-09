package geom

import (
	"bytes"
	"fmt"
)

type Matrix [3]Vec

var UnitMatrix = Matrix{Ex, Ey, Ez}

// Mul performs a Matrix-Matrix multiplication
func (a *Matrix) Mul(b *Matrix) Matrix {
	var c Matrix
	for i := range c {
		for j := range c[i] {
			for k := range b {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

// MulVec performs a Matrix-Vector multiplication
// 	a . vT
func (a *Matrix) MulVec(v Vec) Vec {
	return Vec{v.Dot(a[0]), v.Dot(a[1]), v.Dot(a[2])}
}

func (a *Matrix) Mulf(f float64) Matrix {
	return Matrix{a[0].Mul(f), a[1].Mul(f), a[2].Mul(f)}
}

func (a *Matrix) Transpose() Matrix {
	return Matrix{
		{a[0][0], a[1][0], a[2][0]},
		{a[0][1], a[1][1], a[2][1]},
		{a[0][2], a[1][2], a[2][2]},
	}
}

func (m *Matrix) Inv() Matrix {
	a := m[0][0]
	b := m[1][0]
	c := m[2][0]
	d := m[0][1]
	e := m[1][1]
	f := m[2][1]
	g := m[0][2]
	h := m[1][2]
	i := m[2][2]

	A := e*i - f*h
	B := f*g - d*i
	C := d*h - e*g
	inv := Matrix{
		{e*i - f*h, f*g - d*i, d*h - e*g},
		{c*h - b*i, a*i - c*g, b*g - a*h},
		{b*f - c*e, c*d - a*f, a*e - b*d},
	}
	det := a*A + b*B + c*C
	return inv.Mulf(1 / det)
}

func (a *Matrix) String() string {
	var b bytes.Buffer
	for _, v := range a {
		fmt.Fprintln(&b, v)
	}
	return b.String()
}

func (a *Matrix) Sprintf(format string) string {
	return "[" +
		a[0].Sprintf(format) + "," +
		a[1].Sprintf(format) + "," +
		a[2].Sprintf(format) + "]"
}
