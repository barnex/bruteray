package builder

import (
	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
)

type Rectangle struct {
	Frame
	Texture Material

	ctrlVec [3]Vec
}

var _ Builder = (*Rectangle)(nil)

func NewRectangle(c Material, o, a, b Vec) *Rectangle {
	s := new(Rectangle)
	s.Frame = Frame{o, a, b, Vec{}}
	s.Texture = c
	return s
}

func (s *Rectangle) Init() {
	f := &s.Frame
	f.SetCtrlVec(2, f.CtrlVec(0).Cross(f.CtrlVec(1)))
	f.SetCtrlVec(2, f.CtrlVec(2).Normalized())
	s.Frame.Check()
	for i := range s.ctrlVec {
		s.ctrlVec[i] = s.CtrlVec(i).Mul(1 / s.CtrlVec(i).Len2())
	}
}

func (s *Rectangle) Bounds() BoundingBox {
	return BoundingBoxFromHull(append(s.Frame[:3], s.Origin().Add(s.CtrlVec(0)).Add(s.CtrlVec(1))))
}

func (s *Rectangle) Intersect(c *Ctx, r *Ray) HitRecord {
	origin := s.Frame[0]
	start := r.Start.Sub(origin)
	normal := s.ctrlVec[2]
	rs := start.Dot(normal)
	rd := r.Dir.Dot(normal)
	t := -rs / rd

	at := r.At(t).Sub(origin)
	u := at.Dot(s.ctrlVec[0])
	v := at.Dot(s.ctrlVec[1])
	if u < 0 || u > 1 || v < 0 || v > 1 {
		return HitRecord{}
	}

	return HitRecord{T: t, Normal: normal, Material: s.Texture}
}

func (s *Rectangle) Translate(detla Vec) *Rectangle {
	Translate(s, detla)
	return s
}
