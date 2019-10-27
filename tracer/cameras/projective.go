package cameras

import (
	"fmt"
	"math"
	"math/rand"

	. "github.com/barnex/bruteray/tracer/types"
)

// A projective Camera projects onto a flat image sensor.
// Rays go through a "lens" at distance FocalLen from the sensor.
// Thus, FocalLen determines the Field Of View (FOV):
//
// 	FOV = 2*atan(f/2)
//
// This camera optionally has a finite-size lens Apterture
// which creates depth of field. If a non-zero Aperture is set,
// Focus should be set to the distance from the camera to focus on.
type projective struct {
	FocalLen  float64
	Focus     float64
	Aperture  float64
	Diaphragm func(rng *rand.Rand) (x, y float64)
}

func NewProjective(fov float64, pos Vec, yaw, pitch float64) *Transformed {
	return Translate(YawPitchRoll(Projective(fov), yaw, pitch, 0), pos)
}

func Projective(fov float64) Camera {
	return YawPitchRoll(&projective{ // hack: historically camera looks along -z. TODO: rm (but do in api).
		FocalLen:  fovToFocalLen(fov),
		Focus:     1,
		Diaphragm: diaCircle,
	}, 180*Deg, 0, 0)
}

func fovToFocalLen(fov float64) float64 {
	if fov <= 0 || fov >= Pi {
		panic(fmt.Sprintf("camera: invalid field-of-view: %f (%f deg): need: 0 < fov < pi", fov, fov/Deg))
	}
	return 0.5 / math.Tan(fov/2)
}

// RayFrom implements tracer.Camera.
func (c *projective) RayFrom(ctx *Ctx, u, v float64) *Ray {
	//checkUV(u, v) TODO: fails with AA!

	r := ctx.Ray()

	r.Start = Vec{0, 0, 0}
	if c.Aperture > 0 {
		xs, ys := c.Diaphragm(ctx.Rng)
		r.Start[0] += xs * c.Aperture
		r.Start[1] += ys * c.Aperture
	}

	end := Vec{
		-(u - 0.5),
		+(v - 0.5),
		c.FocalLen,
	}
	if c.Focus != 0 {
		end = end.Mul(c.Focus)
	}
	r.Dir = end.Sub(r.Start).Normalized()
	return r
}

// diaCircle draws a point from the unit disk.
// TODO: ctx.SampleDisk
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
	for math.Abs(y) > sqrt3/2 || math.Abs(x+y/sqrt3) > 1 || math.Abs(x-y/sqrt3) > 1 {
		x, y = diaSquare(rng)
	}
	return x, y
}

func diaSquare(rng *rand.Rand) (x, y float64) {
	x = 2*rng.Float64() - 1
	y = 2*rng.Float64() - 1
	return x, y
}

func checkUV(u, v float64) {
	if u < 0 || u > 1 || v < 0 || v > 1 {
		panic(fmt.Sprintf("Camera: illegal argument {u,v}={%v,%v}, want 0..1", u, v))
	}
}
