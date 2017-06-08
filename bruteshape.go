package main

import "math"

// A Shape implemented as a function that returns true when a point lies inside.
// Intersection is calculated by brute force.
type BruteShape func(Vec) bool

const (
	fine = 0.02
	tol  = 1e-9
)

func (s BruteShape) Normal(r Ray) (float64, Vec, bool) {
	t, ok := s.bisect(r)
	c := r.At(t)
	if !ok {
		return 0, Vec{}, false
	}

	ra := r
	ra.Dir = ra.Dir.Add(Vec{1e-5, 0, 0})
	ta, okA := s.bisect(ra)
	a := ra.At(ta)

	rb := r
	rb.Dir = rb.Dir.Add(Vec{0, 1e-5, 0})
	tb, okB := s.bisect(rb)
	b := rb.At(tb)

	if !okA || !okB {
		return 0, Vec{}, false
	}

	a = a.Sub(c)
	b = b.Sub(c)

	n := b.Cross(a).Normalized()
	return t, n, true
}

func (s BruteShape) inters(r Ray) (float64, bool) {
	for t := 0.0; t < Horiz; t += fine {
		if s(r.At(t)) {
			return t, true
		}
	}
	return 0, false
}

func (s BruteShape) bisect(r Ray) (float64, bool) {
	in, ok := s.inters(r)
	if !ok {
		return 0, false
	}

	out := in - fine

	if s(r.At(out)) || !s(r.At(in)) {
		return 0, false
	}

	for math.Abs(in-out)/(in+out) > tol {
		mid := (in + out) / 2
		if s(r.At(mid)) {
			in = mid
		} else {
			out = mid
		}
	}
	return out, true
}

func Sphere(r float64) BruteShape {
	return Ellipsoid(r, r, r)
}

func Ellipsoid(rx, ry, rz float64) BruteShape {
	return func(r Vec) bool {
		return sqr(r.X/rx)+sqr(r.Y/ry)+sqr(r.Z/rz) <= 1
	}
}

func CylinderZ(radius, semiH float64) BruteShape {
	r2 := sqr(radius)
	return func(r Vec) bool {
		return r.Z > -semiH && r.Z < semiH && sqr(r.X)+sqr(r.Y) <= r2
	}
}

func sqr(x float64) float64 { return x * x }

func Slab(rx, ry, rz float64) BruteShape {
	return func(r Vec) bool {
		return r.X < rx && r.X > -rx && r.Y < ry && r.Y > -ry && r.Z < rz && r.Z > -rz
	}
}

func (s BruteShape) Transl(dx, dy, dz float64) BruteShape {
	return func(r Vec) bool {
		return s(Vec{r.X - dx, r.Y - dy, r.Z - dz})
	}
}

func (s BruteShape) Scale(sx, sy, sz float64) BruteShape {
	return func(r Vec) bool {
		return s(Vec{r.X / sx, r.Y / sy, r.Z / sz})
	}
}

func (s BruteShape) RotZ(θ float64) BruteShape {
	cos := math.Cos(θ)
	sin := math.Sin(θ)
	return func(r Vec) bool {
		x_ := r.X*cos + r.Y*sin
		y_ := -r.X*sin + r.Y*cos
		return s(Vec{x_, y_, r.Z})
	}
}

func (s BruteShape) RotY(θ float64) BruteShape {
	cos := math.Cos(θ)
	sin := math.Sin(θ)
	return func(r Vec) bool {
		x_ := r.X*cos - r.Z*sin
		z_ := r.X*sin + r.Z*cos
		return s(Vec{x_, r.Y, z_})
	}
}

func (s BruteShape) RotX(θ float64) BruteShape {
	cos := math.Cos(θ)
	sin := math.Sin(θ)
	return func(r Vec) bool {
		y_ := r.Y*cos + r.Z*sin
		z_ := -r.Y*sin + r.Z*cos
		return s(Vec{r.X, y_, z_})
	}
}

func (a BruteShape) Add(b BruteShape) BruteShape {
	return func(r Vec) bool {
		return a(Vec{r.X, r.Y, r.Z}) || b(Vec{r.X, r.Y, r.Z})
	}
}

func (a BruteShape) Intersect(b BruteShape) BruteShape {
	return func(r Vec) bool {
		return a(Vec{r.X, r.Y, r.Z}) && b(Vec{r.X, r.Y, r.Z})
	}
}

func (s BruteShape) Inverse() BruteShape {
	return func(r Vec) bool {
		return !s(Vec{r.X, r.Y, r.Z})
	}
}

func (a BruteShape) Sub(b BruteShape) BruteShape {
	return func(r Vec) bool {
		return a(Vec{r.X, r.Y, r.Z}) && !b(Vec{r.X, r.Y, r.Z})
	}
}
