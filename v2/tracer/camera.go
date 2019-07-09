package tracer

import (
	"math"
	"math/rand"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/util"
)

type Camera struct {
	Frame
	FocalLen  float64
	Focus     float64
	Aperture  float64
	Diaphragm func(rng *rand.Rand) (x, y float64)
	AA        bool
}

// Init sets all required fields that are unset.
func (c *Camera) Init() {
	if c.Frame == (Frame{}) {
		c.Frame = XYZ
	}
	if c.FocalLen == 0 {
		c.FocalLen = 1
	}
	if c.Aperture != 0 && c.Diaphragm == nil {
		c.Diaphragm = diaCircle
	}
}

// diaCircle draws a point from the unit disk.
func diaCircle(rng *rand.Rand) (x, y float64) {
	x, y = diaSquare(rng)
	for math.Sqrt(x*x+y*y) > 1 {
		x, y = diaSquare(rng)
	}
	return x, y
}

// diaHex draws a point from the unit hexagon.
func diaHex(rng *rand.Rand) (x, y float64) {
	const sqrt3 = 1.73205080756888
	x, y = diaSquare(rng)
	for abs(y) > sqrt3/2 || abs(x+y/sqrt3) > 1 || abs(x-y/sqrt3) > 1 {
		x, y = diaSquare(rng)
	}
	return x, y
}

func diaSquare(rng *rand.Rand) (x, y float64) {
	x = 2*rng.Float64() - 1
	y = 2*rng.Float64() - 1
	return x, y
}

func abs(x float64) float64 {
	return math.Abs(x)
}

func (c *Camera) RayFrom(ctx *Ctx, x, y float64) *Ray {
	util.Assert(c.Frame != Frame{})
	util.Assert(c.FocalLen != 0)

	end := Vec{x, y, c.FocalLen}
	if c.Focus > 0 {
		end = end.Mul(c.Focus)
	}

	start := Vec{}
	if c.Aperture > 0 {
		xs, ys := c.Diaphragm(ctx.Rng)
		start[0] += xs * c.Aperture
		start[1] += ys * c.Aperture
	}

	start = c.Frame.TransformToAbsolute(start)
	end = c.Frame.TransformToAbsolute(end)

	r := ctx.Ray()
	r.Start = start
	r.Dir = (end.Sub(r.Start).Normalized())
	r.Len = Inf

	return r
}
