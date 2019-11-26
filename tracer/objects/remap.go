package objects

import . "github.com/barnex/bruteray/tracer/types"

func Remap(orig Interface, remap func(Vec) Vec) Interface {
	return &remapped{orig, remap}
}

type remapped struct {
	orig  Interface
	remap func(Vec) Vec
}

func (o *remapped) Intersect(r *Ray) HitRecord {
	h := o.orig.Intersect(r)
	h.Local = o.remap(r.At(h.T))
	return h
}

func (o *remapped) Inside(v Vec) bool {
	return o.orig.Inside(v)
}

func (o *remapped) Bounds() BoundingBox {
	return o.orig.Bounds()
}
