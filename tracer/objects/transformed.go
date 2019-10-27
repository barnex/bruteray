package objects

import (
	"github.com/barnex/bruteray/geom"
	. "github.com/barnex/bruteray/tracer/types"
)

// Transformed returns a transformed instance of the original object.
// The transform may be a rotation and/or scale (equal in all directions),
// but not an anistropic scale or shear, etc.
//
// Multiple transformed instances of the same object may be made,
// they efficiently share underlying state.
//
// If the original object is already a transformed instance,
// then the resulting chain of transforms is optimized into a single,
// equivalent transform. Thus, one may call
//  Transformed(Transformed(...))
// without negative performance implications.
func Transformed(object Interface, t *geom.AffineTransform) Interface {
	if object, ok := object.(*transformed); ok {
		return Transformed(object.orig, object.forward.Before(t))
	}
	return &transformed{
		bounds:  transformBounds(object.Bounds(), t),
		forward: *t,
		inverse: *t.Inverse(),
		orig:    object,
	}
}

type transformed struct {
	bounds  BoundingBox
	forward geom.AffineTransform
	inverse geom.AffineTransform
	orig    Interface
}

func (o *transformed) Intersect(r *Ray) HitRecord {
	if !o.bounds.intersects(r) {
		return HitRecord{}
	}
	prev := *r // backup ray

	// intersect original with inverse transformed ray
	r.Start = o.inverse.TransformPoint(prev.Start)
	r.Dir = o.inverse.TransformDir(prev.Dir)
	h := o.orig.Intersect(r)

	// forward transform normal
	h.Normal = o.forward.TransformDir(h.Normal)

	*r = prev // restore ray
	return h
}

func (o *transformed) Bounds() BoundingBox {
	return o.bounds
}

func (o *transformed) Inside(p Vec) bool {
	return o.orig.Inside(o.inverse.TransformPoint(p))
}

func transformBounds(orig BoundingBox, t *geom.AffineTransform) BoundingBox {
	h := orig.hull()
	for i := range h {
		h[i] = t.TransformPoint(h[i])
	}
	return boundingBoxFromHull(h)
}

func (B BoundingBox) hull() []Vec {
	a := B.Min
	b := B.Max
	return []Vec{
		Vec{a[X], a[Y], a[Z]},
		Vec{a[X], a[Y], b[Z]},
		Vec{a[X], b[Y], a[Z]},
		Vec{a[X], b[Y], b[Z]},
		Vec{b[X], a[Y], a[Z]},
		Vec{b[X], a[Y], b[Z]},
		Vec{b[X], b[Y], a[Z]},
		Vec{b[X], b[Y], b[Z]},
	}
}

func Translated(orig Interface, delta Vec) Interface {
	return Transformed(orig, geom.Translate(delta))
}

//type translated struct {
//	delta Vec
//	orig  Builder
//}
//
//func (o *translated) Bounds() BoundingBox {
//	bb := o.orig.Bounds()
//	bb.Min = bb.Min.Add(o.delta)
//	bb.Max = bb.Max.Add(o.delta)
//	return bb
//}
//
//func (o *translated) Intersect(ctx *Ctx, r *Ray) HitRecord {
//	r2 := ctx.Ray()
//	defer ctx.PutRay(r2)
//	*r2 = *r
//	r2.Start = r2.Start.Sub(o.delta)
//	return o.orig.Intersect(ctx, r2)
//}
//
//func (o *translated) Inside(p Vec) bool {
//	return o.orig.Inside(p.Sub(o.delta))
//}
//
