package geom

import (
	"math"
)

// Vec is a 3-component vector.
// Used to represent either points in space or vectors.
type Vec [3]float64

// Add returns a + b.
func (a Vec) Add(b Vec) Vec {
	return Vec{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

// MAdd (multiply-add) returns a + s*b.
func (a Vec) MAdd(s float64, b Vec) Vec {
	return Vec{
		a[0] + s*b[0],
		a[1] + s*b[1],
		a[2] + s*b[2],
	}
}

// Sub returns a - b.
func (a Vec) Sub(b Vec) Vec {
	return Vec{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
}

// Dot returns the scalar product of a and b.
func (a Vec) Dot(b Vec) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// Mul returns a scaled vector: s * a.
func (a Vec) Mul(s float64) Vec {
	return Vec{
		s * a[0],
		s * a[1],
		s * a[2],
	}
}

// Len returns the length (norm).
func (a Vec) Len() float64 {
	return math.Sqrt(a[0]*a[0] + a[1]*a[1] + a[2]*a[2])
}

// Len2 returns the length squared
func (a Vec) Len2() float64 {
	return a[0]*a[0] + a[1]*a[1] + a[2]*a[2]
}

// Normalized returns a copy of a that is scaled to unit length.
func (a Vec) Normalized() Vec {
	vx, vy, vz := a[0], a[1], a[2]
	s := 1 / math.Sqrt(vx*vx+vy*vy+vz*vz)
	return Vec{
		a[0] * s,
		a[1] * s,
		a[2] * s,
	}
}

// Normalize scales a to unit length, overwriting the original.
func (a *Vec) Normalize() {
	vx, vy, vz := a[0], a[1], a[2]
	s := 1 / math.Sqrt(vx*vx+vy*vy+vz*vz)
	a[X] = vx * s
	a[Y] = vy * s
	a[Z] = vz * s
}

// Cross returns the cross product of a and b, assuming a right-handed space.
// See https://en.wikipedia.org/wiki/Cross_product#Matrix_notation
func (a Vec) Cross(b Vec) Vec {
	x := a[Y]*b[Z] - a[Z]*b[Y]
	y := a[Z]*b[X] - a[X]*b[Z]
	z := a[X]*b[Y] - a[Y]*b[X]
	return Vec{x, y, z}
}

// MakeBasis constructs an orthonormal basis.
// I.e. returns y and z so that x, y, and z are
// mutually orthogonal.
func MakeBasis(x Vec) (y, z Vec) {
	x.Normalize()

	//y = x
	//y[argMin(y)] = 1
	//y.Normalize()
	panic("todo")

}

func argMin(v Vec) int {
	a := 0
	min := math.Abs(v[a])
	for i := range v {
		if math.Abs(v[i]) < min {
			min = math.Abs(v[i])
			a = i
		}
	}
	return a
}

// IsNaN returns true if at least one componet is NaN.
func (a Vec) IsNaN() bool {
	return math.IsNaN(a[0]) || math.IsNaN(a[1]) || math.IsNaN(a[2])
}
