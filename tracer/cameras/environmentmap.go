package cameras

import (
	"math"

	. "github.com/barnex/bruteray/tracer/types"
)

// EnvironmentMap returns a camera that records a 360 degree view
// of the entire environment around the camera using a spherical projection.
//
// The camera's horizontal (U) axis spans 360 degrees in the horizontal (XZ) plane,
// with the center pixel looling along the -Z axis.
//
// The camera's veritical (V) axis spans 180 degrees,
// from looking along the -Y axis to looking along the +Y axis.
func EnvironmentMap(pos Vec) *Transformed {
	return Translate(envMap{}, pos)
}

type envMap struct{}

// RayFrom implements tracer.Camera.
func (envMap) RayFrom(ctx *Ctx, u, v float64) *Ray {
	//checkUV(u, v) TODO: fails with AA!

	phi := -u * 2 * Pi      // 0..2*pi. center pixel looks along -z. +x is to the right
	theta := (v - 0.5) * Pi // -pi/2..pi/2
	x := math.Sin(phi) * math.Cos(theta)
	z := math.Cos(phi) * math.Cos(theta)
	y := math.Sin(theta)

	r := ctx.Ray()
	r.Dir = Vec{x, y, z}

	return r
}
