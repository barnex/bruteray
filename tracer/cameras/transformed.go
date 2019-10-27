package cameras

import (
	"github.com/barnex/bruteray/geom"
	. "github.com/barnex/bruteray/tracer/types"
)

func Translate(c Camera, delta Vec) *Transformed {
	return transform(c, geom.UnitMatrix(), delta)
}

func YawPitchRoll(c Camera, yaw, pitch, roll float64) *Transformed {
	return transform(c, geom.YawPitchRoll(yaw, pitch, roll).A, Vec{})
}

func Position(c Camera) Vec {
	if c, ok := c.(*Transformed); ok {
		return c.pos
	} else {
		return Vec{} // by convention, all non-transformed camears start at the origin
	}
}

func transform(c Camera, rotate geom.Matrix, translate Vec) *Transformed {
	if c, ok := c.(*Transformed); ok {
		return &Transformed{
			orig:   c.orig,
			matrix: rotate.Mul(&c.matrix),
			pos:    translate.Add(c.pos),
		}
	}
	return &Transformed{
		orig:   c,
		matrix: rotate,
		pos:    translate,
	}
}

type Transformed struct {
	orig   Camera
	pos    Vec
	matrix geom.Matrix
}

// RayFrom implements tracer.Camera.
func (c *Transformed) RayFrom(ctx *Ctx, u, v float64) *Ray {
	r := c.orig.RayFrom(ctx, u, v)
	//if r.Start != O {
	//	panic(fmt.Sprintln("start", r.Start))
	//}
	r.Start = c.matrix.MulVec(r.Start).Add(c.pos)
	r.Dir = c.matrix.MulVec(r.Dir)
	return r
}

func (c *Transformed) YawPitchRoll(yaw, pitch, roll float64) *Transformed {
	return YawPitchRoll(c, yaw, pitch, roll)
}

func (c *Transformed) Translate(delta Vec) *Transformed {
	return Translate(c, delta)
}
