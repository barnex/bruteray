package cameras

import (
	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/tracer"
	. "github.com/barnex/bruteray/tracer/types"
)

// WithTransform wraps an affine transformation around a camera,
// allowing for the camera position and view direction to be set.
type WithTransform struct {
	orig   Camera
	pos    Vec
	matrix geom.Matrix
}

// Transform wraps an affine transformation around a camera,
// allowing for the camera position and view direction to be set.
//
// The original camera is not mutated.
//
// Rotation is always with respect to the camera's current position.
// Thus, the rotation and translation are independent of each other.
// (they even commute). Note that this is different from
// affine transformations, who do not commute.
func Transform(c tracer.Camera, rotate geom.Matrix, translate Vec) *WithTransform {
	if c, ok := c.(*WithTransform); ok {
		return &WithTransform{
			orig:   c.orig,
			matrix: rotate.Mul(&c.matrix),
			pos:    translate.Add(c.pos),
		}
	}
	return &WithTransform{
		orig:   c,
		matrix: rotate,
		pos:    translate,
	}
}

// RayFrom implements tracer.Camera.
func (c *WithTransform) RayFrom(ctx *Ctx, u, v float64) *Ray {
	r := c.orig.RayFrom(ctx, u, v)
	r.Start = c.matrix.MulVec(r.Start).Add(c.pos)
	r.Dir = c.matrix.MulVec(r.Dir)
	return r
}

// Translate returns an instance of this camera whose position has been translated.
// The original is not affected.
// The delta is in absolute coordinates, unaffected by the camera's view direction.
func (c *WithTransform) Translate(delta Vec) *WithTransform {
	return translate(c, delta)
}

// YawPitchRoll returns an instance of this camera whose view direction has been rotated
// by yaw, pitch, roll radians around the Y, X, and Z axes respectively.
// Rotation is around the camera's own postition.
func (c *WithTransform) YawPitchRoll(yaw, pitch, roll float64) *WithTransform {
	return yawPitchRoll(c, yaw, pitch, roll)
}

// position returns the camera's location.
// I.e., the point where its rays start from.
func position(c Camera) Vec {
	if c, ok := c.(*WithTransform); ok {
		return c.pos
	} else {
		return Vec{} // by convention, all non-transformed camears start at the origin
	}
}

// translate is like Transform, but only translates
// the camera's position.
func translate(c Camera, delta Vec) *WithTransform {
	return Transform(c, geom.UnitMatrix(), delta)
}

// yawPitchRoll is like Transform, but only rotates
// the camera's view direction.
func yawPitchRoll(c Camera, yaw, pitch, roll float64) *WithTransform {
	return Transform(c, geom.YawPitchRoll(yaw, pitch, roll).A, Vec{})
}
