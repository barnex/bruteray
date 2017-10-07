package bruteray

import (
	"fmt"
	"math"
)

type Vec [3]float64

var Unit = Matrix3{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

var (
	Ex = Unit[X]
	Ey = Unit[Y]
	Ez = Unit[Z]
)

const (
	X = 0
	Y = 1
	Z = 2
	W = 3
)

func (a Vec) Add(b Vec) Vec {
	return Vec{a[X] + b[X], a[Y] + b[Y], a[Z] + b[Z]}
}

func (a Vec) MAdd(s float64, b Vec) Vec {
	return Vec{a[X] + s*b[X], a[Y] + s*b[Y], a[Z] + s*b[Z]}
}

func (a Vec) Sub(b Vec) Vec {
	return Vec{a[X] - b[X], a[Y] - b[Y], a[Z] - b[Z]}
}

func (a Vec) Dot(b Vec) float64 {
	return a[X]*b[X] + a[Y]*b[Y] + a[Z]*b[Z]
}

func (v Vec) Mul(a float64) Vec {
	return Vec{a * v[X], a * v[Y], a * v[Z]}
}

func (v Vec) Mul3(a Vec) Vec {
	return Vec{a[X] * v[X], a[Y] * v[Y], a[Z] * v[Z]}
}

// Scalar division.
func (v Vec) Div(a float64) Vec {
	return v.Mul(1 / a)
}

// Pointwise division.
//func (v Vec) Div3(a Vec) Vec {
//	return Vec{v[X] / a[X], v[Y] / a[Y], v[Z] / a[Z]}
//}

// Length (norm).
func (v Vec) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

// Length squared
func (v Vec) Len2() float64 {
	return v.Dot(v)
}

// Returns a copy of v, scaled to unit length.
func (v Vec) Normalized() Vec {
	l := 1 / math.Sqrt(v[X]*v[X]+v[Y]*v[Y]+v[Z]*v[Z])
	return Vec{v[X] * l, v[Y] * l, v[Z] * l}
}

// May invert v to assure it points towards direction d.
// Used to ensure normal vectors point outwards.
func (n Vec) Towards(d Vec) Vec {
	if n.Dot(d) > 0 {
		return n.Mul(-1)
	}
	return n
}

// Reflects v against the plane normal to n.
func (v Vec) Reflect(n Vec) Vec {
	return v.MAdd(-2*v.Dot(n), n)
}

func (a Vec) Cross(b Vec) Vec {
	x := a[Y]*b[Z] - a[Z]*b[Y]
	y := a[Z]*b[X] - a[X]*b[Z]
	z := a[X]*b[Y] - a[Y]*b[X]
	return Vec{x, y, z}
}

func (v Vec) check() Vec {
	for _, x := range v {
		if math.IsNaN(x) {
			panic(fmt.Sprintf("bad vector: %v", v))
		}
	}
	return v
}

type Vec4 [4]float64

func (a Vec4) Dot(b Vec4) float64 {
	return a[X]*b[X] + a[Y]*b[Y] + a[Z]*b[Z] + a[W]*b[W]
}
