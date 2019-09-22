package objects

import (
	. "github.com/barnex/bruteray/tracer/types"
)

// Bounded returns the original object, wrapped with an Intersect method
// that first checks for intersection with the object's bounding box.
//
// This may speed up performace if the original Intersect is sufficiently expensive.
func Bounded(orig Interface) Interface {
	return &bounded{orig: orig, bounds: orig.Bounds()}
}

type bounded struct {
	bounds BoundingBox
	orig   Interface
}

func (b *bounded) Bounds() BoundingBox {
	return b.bounds
}

func (b *bounded) Inside(p Vec) bool {
	if !b.bounds.inside(p) {
		return false
	}
	return b.orig.Inside(p)
}

func (b *bounded) Intersect(r *Ray) HitRecord {
	if !b.bounds.intersects(r) {
		return HitRecord{}
	}
	return b.orig.Intersect(r)
}
