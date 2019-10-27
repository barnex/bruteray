package geom

import "github.com/barnex/bruteray/util"

/*
	TODO: Transform does not have an Origin(), so not usable as frame.
	Add frame here, or keep in builder?
	y = Ax + b
	o = Ao + b
	o - Ao = b
	o (I-A) = b
	o = (I-A)^-1 b
	frame and transform concats are in different order.
	for frame, linear operations and translations commute (translate camera, rotate, translate again)
	for trasform, not
*/

// AffineTransform represents an 3D affine transformation
//	y = A x + b
type AffineTransform struct {
	A Matrix
	B Vec
}

// ComposeLR composes affine transformations left-to-right.
// I.e., the leftmost argument is applied first.
func ComposeLR(t ...*AffineTransform) *AffineTransform {
	comp := UnitTransform()
	for _, t := range t {
		comp = comp.Before(t)
	}
	return comp
}

func UnitTransform() *AffineTransform {
	return &AffineTransform{
		A: UnitMatrix(),
		B: O,
	}
}

func UnitMatrix() Matrix {
	return Matrix{Ex, Ey, Ez}
}

// Scale returns a transform that scales by factor s,
// with origin as the fixed point.
func Scale(origin Vec, s float64) *AffineTransform {
	return (&AffineTransform{
		A: Matrix{{s, 0, 0}, {0, s, 0}, {0, 0, s}},
	}).WithOrigin(origin)
}

// Translate returns a Transform that translates by delta.
// Translation affects points (TransformPoint), but not directions (TransformDir).
func Translate(delta Vec) *AffineTransform {
	return &AffineTransform{
		A: Matrix{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		B: delta,
	}
}

// Rotate returns a Transform that rotates around an arbitrary axis,
// with origin as the fixed point.
// The rotation is counterclockwise in a right-handed space.
func Rotate(origin Vec, axis Vec, angle float64) *AffineTransform {
	return (&AffineTransform{A: rotationMatrix(axis, angle)}).WithOrigin(origin)
}

// Yaw rotates the camera counterclockwise by the given angle (in radians),
// while keeping it horizontal. E.g.:
// 	camera.Yaw(  0*Deg) // look North
// 	camera.Yaw(+90*Deg) // look West
// 	camera.Yaw(-90*Deg) // look East
// 	camera.Yaw(180*Deg) // look South
//
// Yaw, Pitch and Roll are not commutative.
// Use YawPitchRoll to apply them in canonical order.
func Yaw(angle float64) *AffineTransform {
	return Rotate(O, Ey, angle)
}

// Pitch tilts the camera upwards by the given angle (in radians). E.g.:
// 	camera.Pitch(  0*Deg) // look horizontally
// 	camera.Pitch(-90*Deg) // look at your feet
// 	camera.Pitch(+90*Deg) // look at the zenith
//
// Yaw, Pitch and Roll are not commutative.
// Use YawPitchRoll to apply them in canonical order.
func Pitch(angle float64) *AffineTransform {
	return Rotate(O, Ex, angle)
}

// Roll rotates the camera counterclockwise around the line of sight. E.g.:
// 	camera.Roll( 0*Deg) // horizon runs straight
// 	camera.Roll(45*Deg) // horizon runs diagonally, from top left to bottom right.
//
// Yaw, Pitch and Roll are not commutative.
// Use YawPitchRoll to apply them in canonical order.
func Roll(angle float64) *AffineTransform {
	return Rotate(O, Ez, angle)
}

func YawPitchRoll(yaw, pitch, roll float64) *AffineTransform {
	return Yaw(yaw).After(Pitch(pitch)).After(Roll(roll))
}

// https://en.wikipedia.org/wiki/Rotation_matrix#Rotation_matrix_from_axis_and_angle
func rotationMatrix(axis Vec, angle float64) Matrix {
	axis.Normalize()
	ux, uy, uz := axis[0], axis[1], axis[2]
	c := util.Cos(angle)
	s := util.Sin(angle)
	c1 := 1 - c
	return Matrix{
		{c + ux*ux*c1, uy*ux*c1 + uz*s, uz*ux*c1 - uy*s},
		{ux*uy*c1 - uz*s, c + uy*uy*c1, uz*uy*c1 + ux*s},
		{ux*uz*c1 + uy*s, uy*uz*c1 - ux*s, c + uz*uz*c1},
	}
}

func MapBaseTo(o, x, y, z Vec) *AffineTransform {
	return &AffineTransform{
		A: Matrix{
			x.Sub(o),
			y.Sub(o),
			z.Sub(o),
		},
		B: o,
	}
}

// TransformPoint applies affine transformation t to point x, returning
//	A x + b
// Use TransformDir to transform a direction (vector).
func (t *AffineTransform) TransformPoint(x Vec) Vec {
	return t.A.MulVec(x).Add(t.B)
}

// TransformDir applies affine transformation t to vector dir, returning
//	A x
// Directions are invariant against the the translation part of the transform.
// Use TransformPoint to transform a point, which does undergo translation.
func (t *AffineTransform) TransformDir(dir Vec) Vec {
	return t.A.MulVec(dir)
}

// Inverse returns the inverse affine transformation.
func (t *AffineTransform) Inverse() *AffineTransform {
	// if
	// 	y = Ax + b
	// then
	// 	x = (A^-1)y + (-A^-1)b
	// so the coefficients of the inverse transform are (A^-1), (-A^-1)b
	Ainv := t.A.Inverse()
	return &AffineTransform{
		A: Ainv,
		B: Ainv.MulVec(t.B).Mul(-1),
	}
}

// After returns a composite transform that applies s first, followed by t.
func (t *AffineTransform) After(s *AffineTransform) *AffineTransform {
	return s.Before(t)
}

// Before returns a composite transform that applies t first, followed by s.
func (t *AffineTransform) Before(s *AffineTransform) *AffineTransform {
	// y = TAx+Tb
	// z = SAy+Sb
	//   = SA(TAx+Tb)+Sb
	//   = (SA*TA)x + (SA*Tb+Sb)
	return &AffineTransform{
		A: s.A.Mul(&t.A),
		B: s.A.MulVec(t.B).Add(s.B),
	}
}

// WithOrigin returns a translated version of t so that o is the new origin
// (fixed point). E.g.:
// 	Rotate(Vec{}, Ez, θ).WithOrigin(Vec{1,2,0})
// rotates around [1, 2, 0] rather than [0, 0, 0]
func (t *AffineTransform) WithOrigin(o Vec) *AffineTransform {
	return ComposeLR(Translate(o.Mul(-1)), t, Translate(o))
}
