package texture

import (
	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
)

type UVMapper interface {
	Map(p Vec) (u, v float64)
}

func Map(t Texture2D, m UVMapper) Texture {
	return &mapped{t, m}
}

type mapped struct {
	texture Texture2D
	mapper  UVMapper
}

func (m *mapped) At(p Vec) Color {
	return m.texture.AtUV(m.mapper.Map(p))
}

type UVProject struct{}

func (c UVProject) Map(p Vec) (u, v float64) {
	return p[0], p[1]
}

// UVAffine maps an affine coordinate system.
// Most suited to map textures on plane surfaces.
// 	P0 -> (0, 0)
// 	Pu -> (1, 0)
// 	Pv -> (0, 1)
// Pu and Pv should be chosen orthogonal.
type affineMap struct {
	frame *Frame
}

func (c *affineMap) Map(pos Vec) (u, v float64) {
	p2 := c.frame.TransformToFrame(pos)
	return p2[0], p2[1]
}

func UVAffine(f *Frame) UVMapper {
	return &affineMap{f}
}

//// UVCyl maps a cylindrical coordinate system.
//// 	P0: center
//// 	Pu: point on the equator
//// 	Pv: north pole
//type cylinderMap struct {
//	CtrlPoints [3]*Vec
//}
//
//func (c *cylinderMap) Bind(x Transformable) {
//	ctrl := x.CtrlPoints()
//	if len(ctrl) < 3 {
//		panic("not enough control points")
//	}
//	copy(c.CtrlPoints[:], ctrl)
//}
//
//func (c *cylinderMap) Map(pos Vec) (u, v float64) {
//	P0, Pu, Pv := *c.CtrlPoints[0], *c.CtrlPoints[1], *c.CtrlPoints[2]
//	p := pos.Sub(P0)
//
//	pv := Pv.Sub(P0).Normalized()
//	pu := Pu.Sub(P0).Normalized()
//	pw := pu.Cross(pv).Normalized()
//
//	x := p.Dot(pu)
//	y := p.Dot(pw)
//	u = 0.5 + (math.Atan2(y, x))/(2*math.Pi)
//	v = 0.5 + (p.Dot(pv)/pv.Len2())/2
//	//v = 0.5 + math.Asin((p.Dot(pv)/pv.Len2()))/(Pi) // for projection on sphere
//	return u, v
//}
//
