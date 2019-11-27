package cameras

import (
	"fmt"
	"math"

	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/tracer/sequence"
	. "github.com/barnex/bruteray/tracer/types"
)

// Projective constructs a camera like ProjectiveAperture,
// but with zero aperture. I.e. a pinhole camera.
func Projective(FOV float64) *WithTransform {
	return ProjectiveAperture(FOV, 0, 1)
}

// A projective Camera projects onto a flat image sensor.
// FOV is the horizontal Field Of View, in radians.
//
// This camera optionally has a finite-size lens aperture
// which creates depth of field. If a non-zero Aperture is set,
// focusDist should be set to the distance from the camera to focus on.
//
// The camera is located at (0,0,0) and looks along the -Z direction.
// It can be rotated and translated if desired.
func ProjectiveAperture(FOV, aperture, focusDist float64) *WithTransform {
	if aperture == 0 && focusDist == 0 {
		focusDist = 1 // irrelevant as long as > 0
	}
	return Transform(
		&projective{
			focalLen:  fovToFocalLen(FOV),
			focusDist: focusDist,
			aperture:  aperture,
			diaphragm: sequence.UniformDisk,
		},
		geom.YawPitchRoll(180*Deg, 0, 0).A, // hack: historically camera looks along -z.
		O,
	)

}

type projective struct {
	focalLen  float64                           // lens focal length, determines Field Of View
	focusDist float64                           // distance from lens to focal plane. Irrelevant if aperture == 0.
	aperture  float64                           // radius of lens opening. 0 means pinhole camera
	diaphragm func(u, v float64) (x, y float64) // aperture shape. transforms lens samples [0..1] to positions on the lens [-1..1]
}

// fovToFocalLen converts a Field Of View (in radians) to focal length
// corresponding to a sensor of size 1.
//
// 	FOV = 2*atan(f/2)
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
	if c.aperture > 0 {
		xs, ys := c.diaphragm(ctx.GenerateLens())
		r.Start[0] += xs * c.aperture
		r.Start[1] += ys * c.aperture
	}

	end := Vec{
		-(u - 0.5),
		+(v - 0.5),
		c.focalLen,
	}
	if c.focusDist != 0 {
		end = end.Mul(c.focusDist)
	}
	r.Dir = end.Sub(r.Start).Normalized()
	return r
}

func checkUV(u, v float64) {
	if u < 0 || u > 1 || v < 0 || v > 1 {
		panic(fmt.Sprintf("Camera: illegal argument {u,v}={%v,%v}, want 0..1", u, v))
	}
}
