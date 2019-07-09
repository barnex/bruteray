package builder

import (
	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
	. "github.com/barnex/bruteray/v2/util"
)

type CSGBuilder interface {
	Builder
	Inside(p Vec) bool
}

func And(a, b CSGBuilder) CSGBuilder {
	return &and{a, b}
}

type and struct {
	a, b CSGBuilder
}

var _ CSGBuilder = (*and)(nil)

func (o *and) Init() {
	o.a.Init()
	o.b.Init()
}

/*
		  *===============
	*==============
*/
func (o *and) Intersect(c *Ctx, r *Ray) HitRecord {
	// TODO: this is only correct for two convex shapes
	// need to march forward.
	f := o.a.Intersect(c, r)
	if f.T > 0 && o.b.Inside(r.At(f.T)) {
		return f
	}
	f = o.b.Intersect(c, r)
	if f.T > 0 && o.a.Inside(r.At(f.T)) {
		return f
	}
	return HitRecord{}
}

func (o *and) Bounds() BoundingBox {
	a := o.a.Bounds()
	b := o.b.Bounds()
	return BoundingBox{
		Min: Vec{
			Max(a.Min[0], b.Min[0]),
			Max(a.Min[1], b.Min[1]),
			Max(a.Min[2], b.Min[2]),
		},
		Max: Vec{
			Min(a.Max[0], b.Max[0]),
			Min(a.Max[1], b.Max[1]),
			Min(a.Max[2], b.Max[2]),
		},
	}
}

func (o *and) Inside(p Vec) bool {
	return o.a.Inside(p) && o.b.Inside(p)
}
