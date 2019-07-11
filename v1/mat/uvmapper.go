package mat

import (
	"math"

	. "github.com/barnex/bruteray/v1/br"
)

// A UVMapper maps 3D coordinates (x,y,z) on the surface of a shape
// onto 2D coordinates (u,v) suitable for indexing a texture.
// (u,v) coordinates typically lie within the range [0, 1].
type UVMapper interface {
	Map(pos Vec) (u, v float64)
}

// UVAffine maps an affine coordinate system.
// Most suited to map textures on plane surfaces.
// 	P0 -> (0, 0)
// 	Pu -> (1, 0)
// 	Pv -> (0, 1)
// Often, Pu and Pv are chosen orthogonally.
type UVAffine struct {
	P0, Pu, Pv Vec
}

func (c *UVAffine) Map(pos Vec) (u, v float64) {
	p := pos.Sub(c.P0)
	pu := c.Pu.Sub(c.P0)
	pv := c.Pv.Sub(c.P0)
	u = p.Dot(pu) / pu.Len2()
	v = p.Dot(pv) / pv.Len2()
	return u, v
}

// UVCyl maps a cylindrical coordinate system.
// 	P0: center
// 	Pu: point on the equator
// 	Pv: north pole
type UVCyl struct {
	P0, Pu, Pv Vec
}

func (c *UVCyl) Map(pos Vec) (u, v float64) {
	p := pos.Sub(c.P0)
	pv := c.Pv.Sub(c.P0)
	v = (p.Dot(pv) / pv.Len2())

	pu := c.Pu.Sub(c.P0).Normalized()

	pw := pu.Cross(pv).Normalized()
	x := p.Dot(pu)
	y := p.Dot(pw)
	u = 0.5 + (math.Atan2(y, x))/(2*math.Pi)
	return u, v
}
