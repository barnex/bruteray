package main

import "math"

type Shape func(Vec) bool

func Sphere(r float64) Shape {
	return Ellipsoid(r, r, r)
}

func Ellipsoid(rx, ry, rz float64) Shape {
	return func(r Vec) bool {
		return sqr(r.X/rx)+sqr(r.Y/ry)+sqr(r.Z/rz) <= 1
	}
}

func CylinderZ(radius, semiH float64) Shape {
	r2 := sqr(radius)
	return func(r Vec) bool {
		return r.Z > -semiH && r.Z < semiH && sqr(r.X)+sqr(r.Y) <= r2
	}
}

func sqr(x float64) float64 { return x * x }

func Slab(rx, ry, rz float64) Shape {
	return func(r Vec) bool {
		return r.X < rx && r.X > -rx && r.Y < ry && r.Y > -ry && r.Z < rz && r.Z > -rz
	}
}

func (s Shape) Transl(dx, dy, dz float64) Shape {
	return func(r Vec) bool {
		return s(Vec{r.X - dx, r.Y - dy, r.Z - dz})
	}
}

func (s Shape) Scale(sx, sy, sz float64) Shape {
	return func(r Vec) bool {
		return s(Vec{r.X / sx, r.Y / sy, r.Z / sz})
	}
}

func (s Shape) RotZ(θ float64) Shape {
	cos := math.Cos(θ)
	sin := math.Sin(θ)
	return func(r Vec) bool {
		x_ := r.X*cos + r.Y*sin
		y_ := -r.X*sin + r.Y*cos
		return s(Vec{x_, y_, r.Z})
	}
}

func (s Shape) RotY(θ float64) Shape {
	cos := math.Cos(θ)
	sin := math.Sin(θ)
	return func(r Vec) bool {
		x_ := r.X*cos - r.Z*sin
		z_ := r.X*sin + r.Z*cos
		return s(Vec{x_, r.Y, z_})
	}
}

func (s Shape) RotX(θ float64) Shape {
	cos := math.Cos(θ)
	sin := math.Sin(θ)
	return func(r Vec) bool {
		y_ := r.Y*cos + r.Z*sin
		z_ := -r.Y*sin + r.Z*cos
		return s(Vec{r.X, y_, z_})
	}
}

func (a Shape) Add(b Shape) Shape {
	return func(r Vec) bool {
		return a(Vec{r.X, r.Y, r.Z}) || b(Vec{r.X, r.Y, r.Z})
	}
}

func (a Shape) Intersect(b Shape) Shape {
	return func(r Vec) bool {
		return a(Vec{r.X, r.Y, r.Z}) && b(Vec{r.X, r.Y, r.Z})
	}
}

func (s Shape) Inverse() Shape {
	return func(r Vec) bool {
		return !s(Vec{r.X, r.Y, r.Z})
	}
}

func (a Shape) Sub(b Shape) Shape {
	return func(r Vec) bool {
		return a(Vec{r.X, r.Y, r.Z}) && !b(Vec{r.X, r.Y, r.Z})
	}
}
