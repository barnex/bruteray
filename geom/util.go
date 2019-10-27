// Package geom provides geometric primitives: vectors, matrices and transformations.
package geom

import (
	"math"
)

const (
	// Deg is one degree in radians. Useful for conversions, e.g.: 90*Deg.
	Deg = Pi / 180

	// Pi is shorthand for math.Pi.
	Pi = math.Pi
)

// Vector components.
const (
	X = 0
	Y = 1
	Z = 2
)

// Ex is shorthand for the basis vector [1 0 0].
var Ex = Vec{1, 0, 0}

// Ey is shorthand for the basis vector [0 1 0].
var Ey = Vec{0, 1, 0}

// Ez is shorthand for the basis vector [0 0 1].
var Ez = Vec{0, 0, 1}

// O is shorthand for the zero vector [0 0 0].
var O = Vec{0, 0, 0}

var Inf = math.Inf(1)

// TriangleNormal returns the normal vector of the triangle
// with vertices a, b, c.
func TriangleNormal(a, b, c Vec) Vec {
	return ((b.Sub(a)).Cross(c.Sub(a))).Normalized()
}
