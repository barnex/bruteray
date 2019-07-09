package builder

import (
	"math"

	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
	"github.com/barnex/bruteray/v2/util"
)

type Sphere struct {
	Frame
	Texture Material
	r2      float64
}

func NewSphere(m Material, diam float64) *Sphere {
	f := XYZ
	f.Scale(diam / 2)
	return &Sphere{
		Frame:   f,
		Texture: m,
	}
}

var _ Builder = (*Sphere)(nil)

func (s *Sphere) Init() {
	util.Assert(s.Frame != (Frame{}))
	s.r2 = s.Frame[0].Sub(s.Frame[1]).Len2()
}

func (s *Sphere) Bounds() BoundingBox {
	return BoundingBoxFromHull(s.Frame.MaxHull())
}

func (s *Sphere) Intersect(ctx *Ctx, r *Ray) (f HitRecord) {
	v := r.Start.Sub(s.Origin())
	d := r.Dir
	vd := v.Dot(d)
	D := vd*vd - (v.Len2() - s.r2)
	if D < 0 {
		return HitRecord{}
	}
	sqrtD := math.Sqrt(D)
	t1 := (-vd - sqrtD)
	t2 := (-vd + sqrtD)
	t := frontSolution(t1, t2, r.Len)
	n := r.At(t).Sub(s.Origin())
	if t > 0 {
		return HitRecord{T: t, Normal: n, Material: s.Texture}
	}
	return HitRecord{}
}

func (s *Sphere) Inside(p Vec) bool {
	v := p.Sub(s.Origin())
	return v.Len2() < s.r2
}
