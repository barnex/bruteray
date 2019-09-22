package objects

import (
	"github.com/barnex/bruteray/tracer"
	"github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/types"
)

// WithBounds returns the object wrapped in a semi-tranparent box
// the size of the object's Bounds().
// Intended to visualize the bounding box.
func WithBounds(object Interface) Object {
	// Need to enlarge box by 2*Tiny because secondary ray from transparent
	// material emerges at distance Tiny behind the trasparent surface.
	// If this roughly cooincides with the surface of the enclosed object, we miss it.
	return &withBounds{
		bounds: object.Bounds().withMargin(2 * Tiny),
		mat:    test.Transparent(Color{0.6, 0.7, 0.8}, Color{0.02, 0.03, 0.04}),
		orig:   object,
	}
}

type withBounds struct {
	hollowSurface
	bounds BoundingBox
	mat    Material
	orig   Interface
}

func (a *withBounds) Bounds() BoundingBox {
	return a.bounds
}

func (a *withBounds) Intersect(r *Ray) HitRecord {
	box := HitRecord{T: a.bounds.intersect(r), Material: a.mat, Normal: Vec{1, 1, 1}, Local: Vec{1, 1, 1}} // dummy normal
	orig := a.orig.Intersect(r)
	return tracer.Frontmost(&box, &orig)
}
