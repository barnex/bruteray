package cameras

import (
	"github.com/barnex/bruteray/geom"
	. "github.com/barnex/bruteray/tracer/types"
)

type frame struct {
	frame geom.Frame
}

func newFrame(origin Vec, axes [3]Vec) frame {
	return frame{geom.NewFrame(origin, axes)}
}

func defaultFrame() frame {
	return newFrame(Vec{0, 0, 0}, [3]Vec{{-1, 0, 0}, {0, 1, 0}, {0, 0, -1}})
}

func (f *frame) transformRay(r *Ray) {
	r.Start = f.frame.TransformToGlobal(r.Start)
	r.Dir = f.frame.TransformDirToGlobal(r.Dir)
	r.Dir.Normalize()
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
func (c *frame) Yaw(angle float64) {
	c.rotate(geom.Ey, angle)
}

// Pitch tilts the camera upwards by the given angle (in radians). E.g.:
// 	camera.Pitch(  0*Deg) // look horizontally
// 	camera.Pitch(-90*Deg) // look at your feet
// 	camera.Pitch(+90*Deg) // look at the zenith
//
// Yaw, Pitch and Roll are not commutative.
// Use YawPitchRoll to apply them in canonical order.
func (c *frame) Pitch(angle float64) {
	c.rotate(geom.Ex, angle)
}

// Roll rotates the camera counterclockwise around the line of sight. E.g.:
// 	camera.Roll( 0*Deg) // horizon runs straight
// 	camera.Roll(45*Deg) // horizon runs diagonally, from top left to bottom right.
//
// Yaw, Pitch and Roll are not commutative.
// Use YawPitchRoll to apply them in canonical order.
func (c *frame) Roll(angle float64) {
	c.rotate(geom.Ez, angle)
}

func (c *frame) YawPitchRoll(yaw, pitch, roll float64) {
	c.Yaw(yaw)
	c.Pitch(pitch)
	c.Roll(roll)
}

func (c *frame) rotate(axis Vec, angle float64) {
	c.applyTransform(geom.Rotate(c.origin(), axis, angle).TransformPoint)
}

func (c *frame) Translate(delta Vec) {
	c.applyTransform(geom.Translate(delta).TransformPoint)
}

func (c *frame) origin() Vec {
	return c.frame.Origin()
}

func (c *frame) applyTransform(f func(Vec) Vec) {
	c.frame.ApplyTransform(f)
}
