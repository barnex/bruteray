package cameras

import (
	"fmt"

	. "github.com/barnex/bruteray/tracer/types"
)

type isometric struct {
	size float64
}

// Isometric returns a camera that performs an isometric projection
// (no perspective, sizes are conserved).
// dir is the view direction:
// 	0:X
// 	1:Y
//	2:Z
// size is the horizontal viewport size.
func Isometric(dir int, size float64) *WithTransform {
	c := &isometric{size}
	switch dir {
	default:
		panic(fmt.Sprintf("NewIsoMetric: dir must be 0,1, or 2, have: %v", dir))
	case X:
		return yawPitchRoll(c, 90*Deg, 0, 0)
	case Y:
		return yawPitchRoll(c, 0, -90*Deg, 0)
	case Z: // nothing to do
		return yawPitchRoll(c, 0, 0, 0)
	}
}

// RayFrom implements tracer.Camera.
func (c *isometric) RayFrom(ctx *Ctx, u, v float64) *Ray {
	// checkUV(u, v) // TODO
	r := ctx.Ray()
	s := c.size
	r.Start = Vec{(u - 0.5) * s, (v - 0.5) * s, isoOffset}
	r.Dir = Vec{0, 0, -1}
	return r
}

const isoOffset = 4096
