package random

import (
	"math"

	"github.com/barnex/bruteray/geom"
)

// UniformDisk transforms a point (u,v) from the unit square to the unit disk.
// If the original (u, v) is uniformly distributed, then the result is uniform too.
//
// This is used for sampling a disk-shaped light source.
func UniformDisk(u, v float64) (x, y float64) {
	theta := (2 * math.Pi) * u
	r := math.Sqrt(v)
	x = r * math.Cos(theta)
	y = r * math.Sin(theta)
	return x, y
}

// CosineSphere transforms a point (u,v) from the unit square to a vector
// on the heimsphere around the given normal, cosine weighted.
// I.e. the resulting vectors are distributed proportionally to the cosine of the angle with the normal,
// assuming that the original (u,v) is unifmormly distributed.
//
// This is used for cosine-weighted importance sampling. E.g. for Lambertian scattering (materials.Matte).
func CosineSphere(u, v float64, normal geom.Vec) geom.Vec {
	// Malley’s Method: project disk onto hemisphere.
	// http://www.pbr-book.org/3ed-2018/Monte_Carlo_Integration/2D_Sampling_with_Multidimensional_Transformations.html#fig:malley

	x, y := UniformDisk(u, v)
	z := math.Sqrt(1 - (x*x + y*y))
	S := geom.Vec{x, y, z}

	// Create basis with normal as z-axis
	// See Shirley, Fundamentals of Computer Graphids
	t := normal
	i := 0
	min := math.Abs(t[i])
	if math.Abs(t[1]) < min {
		i = 1
		min = math.Abs(t[1])
	}
	if math.Abs(t[2]) < min {
		i = 2
	}
	t[i] = 1

	ex := t.Cross(normal).Normalized()
	ey := ex.Cross(normal)
	ez := normal

	m := geom.Matrix{ex, ey, ez}

	r := m.MulVec(S)
	return r
}
