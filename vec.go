package main

import "math"

type Vec struct {
	X, Y, Z float64
}

func (a Vec) Add(b Vec) Vec {
	return Vec{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vec) MAdd(s float64, b Vec) Vec {
	return Vec{a.X + s*b.X, a.Y + s*b.Y, a.Z + s*b.Z}
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

func (v Vec) Div3(a Vec) Vec {
	return Vec{v.X / a.X, v.Y / a.Y, v.Z / a.Z}
}

func (v Vec) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec) Len2() float64 {
	return v.Dot(v)
}

func (v Vec) Normalized() Vec {
	l := v.Len()
	if l == 0 {
		l = 1
	}
	return v.Div(l)
}

func (a Vec) Cross(b Vec) Vec {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return Vec{x, y, z}
}

func (r Vec) RotX(θ float64) Vec {
	cos := math.Cos(θ)
	sin := math.Sin(θ)
	y_ := r.Y*cos + r.Z*sin
	z_ := -r.Y*sin + r.Z*cos
	return Vec{r.X, y_, z_}
}
