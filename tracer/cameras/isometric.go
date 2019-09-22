package cameras

import (
	"fmt"

	"github.com/barnex/bruteray/geom"
	. "github.com/barnex/bruteray/tracer/types"
)

type Isometric struct {
	frame
}

func NewIsometric(dir int, size float64) *Isometric {
	f := frame{geom.NewFrame(Vec{0, 0, 0}, [3]Vec{{-size, 0, 0}, {0, size, 0}, {0, 0, -size}})}
	switch dir {
	default:
		panic(fmt.Sprintf("NewIsoMetric: dir must be 0,1, or 2, have: %v", dir))
	case X:
		f.Yaw(90 * Deg)
	case Y:
		f.Pitch(-90 * Deg)
	case Z: // nothing to do
	}
	return &Isometric{f}
}

// RayFrom implements tracer.Camera.
func (c *Isometric) RayFrom(ctx *Ctx, u, v float64) *Ray {
	// checkUV(u, v) // TODO

	r := ctx.Ray()
	r.Start = Vec{-(u - 0.5), (v - 0.5), -isoOffset}
	r.Dir = Vec{0, 0, 1}
	c.frame.transformRay(r)
	return r
}

const isoOffset = 4096
