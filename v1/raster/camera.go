package raster

import (
	"math/rand"

	. "github.com/barnex/bruteray/v1/br"
)

// Camera renders a scene into a raw intensity image.
type Cam struct {
	FocalLen  float64
	Focus     float64
	Aperture  float64
	Diaphragm func(rng *rand.Rand) (x, y float64)
	transf    Matrix4
	AA        bool
}

// Constructs a camera with given focal length.
// Focal length 0 means infinity (orthogonal projection).
// Camera is at the origin, looking in the +z direction,
// and can be transformed later.
func Camera(focalLen float64) *Cam {
	return &Cam{
		FocalLen:  focalLen,
		transf:    *UnitMatrix4(),
		Diaphragm: DiaCircle,
	}
}

// RayFrom set r to a ray for pixel i,j.
func (c *Cam) RayFrom(ctx *Ctx, i, j int, W, H int, r *Ray) {

	r.Start = Vec{}

	if c.Aperture > 0 {
		xs, ys := c.Diaphragm(ctx.Rng)
		xs *= c.Aperture
		ys *= c.Aperture
		r.Start = Vec{xs, ys, 0}
	}

	// ray end point
	y0 := ((-float64(i) + c.aa(ctx.Rng) + float64(H)/2) / float64(H))
	x0 := ((float64(j) + c.aa(ctx.Rng) - float64(W)/2) / float64(H))
	z0 := c.FocalLen

	end := Vec{x0, y0, z0}
	if c.Focus > 0 {
		end = end.Mul(c.Focus)
	}

	// ray direction
	if c.FocalLen != 0 {
		r.SetDir(end.Sub(r.Start).Normalized())
	} else {
		r.SetDir(Vec{0, 0, 1})
		r.Start = Vec{x0, y0, 0}
	}

	// camera transform
	r.Transf(&(c.transf))

}

// Translates the camera.
func (c *Cam) Transl(dx, dy, dz float64) *Cam {
	c.Transf(Transl4(Vec{dx, dy, dz}))
	return c
}

// Transforms the camera direction,
// e.g. rotating the camera.
func (c *Cam) Transf(T *Matrix4) *Cam {
	c.transf = *((&c.transf).Mul(T))
	return c
}

func (c *Cam) RotScene(theta float64) *Cam {
	d := Vec{c.transf[X][W], c.transf[Y][W], c.transf[Z][W]}
	T := Transl4(d.Mul(-1)).Mul(RotY4(theta)).Mul(Transl4(d))
	c.transf = *((&c.transf).Mul(T))
	return c
}

// Anti-aliasing jitter
func (c *Cam) aa(rng *rand.Rand) float64 {
	if c.AA {
		return rng.Float64()
	} else {
		return 0.5
	}
}
