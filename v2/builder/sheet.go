package builder

import (
	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
)

type Sheet struct {
	Frame
	Texture Material
}

func NewSheet(m Material, o, a, b Vec) *Sheet {
	return &Sheet{Texture: m, Frame: Frame{o, a, b, TriangleNormal(o, a, b)}}
}

func NewSheetXZ() *Sheet {
	return &Sheet{Frame: Frame{Vec{}, Ex, Ez, Ey}}
}

func (s *Sheet) Init() {}

func (s *Sheet) Bounds() BoundingBox {
	return infBox()
}

func (s *Sheet) Intersect(c *Ctx, r *Ray) HitRecord {
	normal := s.CtrlVec(2)
	rs := r.Start.Sub(s.Origin()).Dot(normal)
	rd := r.Dir.Dot(normal)
	t := -rs / rd
	return HitRecord{T: t, Normal: normal, Material: s.Texture, Local: s.Frame.TransformToFrame(r.At(t))}

	//start := r.Start.Sub(s.Origin())
	//normal := s.CtrlVec(2)
	//rs := start.Dot(normal)
	//rd := r.Dir.Dot(normal)
	//t := rs / rd
	//return Fragment{T: t, Normal: normal, Material: s.Texture}
}
