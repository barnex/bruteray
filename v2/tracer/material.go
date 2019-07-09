package tracer

import (
	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
)

// A Material determines the color of a surface fragment.
type Material interface {

	// Eval returns the color of the given surface fragment, as seen by Ray r.
	// If Shade uses recursion, e.g., to calculate reflections,
	// it must pass recDepth-1 as the new recursion depth, so that
	// recursion can eventually be terminated (by Env.Shade).
	//TODO: other shapes set U,V,W? but get experience w/ textures first
	//TODO: rename: Shade?
	Eval(ctx *Ctx, s *Scene, r *Ray, recDepth int, h HitCoords) Color
}

// HitCoords record where a Ray intersected an Object.
type HitCoords struct {
	T      float64 // Position along the ray
	Normal Vec     // Normal vector at intersection (needs not have unit length)
	Local  Vec     // Local coordinates at intersection, chosen by the Object. E.g. U,V coordinates for objects where this makes sense
}
