package objects

import (
	. "github.com/barnex/bruteray/tracer/types"
)

// Backdrop returns a "fake backdrop", a gigantic sphere near infinity,
// on which a background like a sky can be textured.
//
// Usually, a Flat material should be applied, (since no light
// can reach the infinitely far away backdrop).
// A flat material also makes the backdrop an indirect light source.
// E.g.: a backdrop with a Flat solid color is the canonical ambient light.
//
// The local coordinates are the ray direction.
func Backdrop(m Material) Interface {
	return &backdrop{mat: m}
}

type backdrop struct {
	hollowSurface
	mat Material
}

func (o *backdrop) Intersect(r *Ray) HitRecord {
	// although the surface is truly at infinity, we return a large finite T
	// to avoid complications like inf != inf.
	return HitRecord{T: 1e99, Material: o.mat, Local: r.Dir}
}

func (o *backdrop) Bounds() BoundingBox {
	return infBox
}
