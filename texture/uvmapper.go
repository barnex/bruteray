package texture

import (
	"math"

	. "github.com/barnex/bruteray/geom"
	. "github.com/barnex/bruteray/imagef/colorf"
)

// MapPlanar maps an affine coordinate system.
// Most suited to map textures on plane surfaces.
// 	P0 -> (0, 0)
// 	Pu -> (1, 0)
// 	Pv -> (0, 1)
// Pu and Pv should be chosen orthogonal.
//func MapPlanar(t Texture, f *Frame) Texture {
//	transf := f.Transform().Inverse()
//	return &mapped{t, func(p Vec) (u, v float64) {
//		p2 := transf.TransformPoint(p)
//		return p2[0], p2[1]
//	}}
//}

func MapFisheye(t Texture) Texture {
	return &mapped{t, func(dir Vec) (u, v float64) {
		dir = dir.Normalized()
		r := math.Sqrt(dir[Z]*dir[Z] + dir[X]*dir[X])
		r = math.Asin(r) / (math.Pi / 2)
		//dir = dir.Mul(r)
		th := math.Atan2(dir[Z], dir[X])
		x := r * math.Cos(th)
		y := r * math.Sin(th)
		u = 0.5 + x*0.5
		v = 0.5 + y*0.5
		return u, v
	}}
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

type mapped struct {
	texture Texture
	mapper  func(p Vec) (u, v float64)
}

func (m *mapped) At(p Vec) Color {
	u, v := m.mapper(p)
	return m.texture.At(Vec{u, v, 0})
}
