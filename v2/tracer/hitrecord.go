package tracer

import (
	. "github.com/barnex/bruteray/v2/geom"
)

// A HitRecord records the position and material where a Ray intersected an Object.
// This allows for lazy evaluation of the Color seen by the Ray, after it has been
// established that this value is actually needed
// (i.e., the hit point is not occluded by another Object).
//
// The structure is designed to avoid allocations:
// It is passed by value, and the Material is only allocated once at Init time.
type HitRecord struct {
	T        float64 // Position along the ray
	Normal   Vec     // Normal vector at intersection (needs not have unit length)
	Local    Vec     // Local coordinates at intersection, chosen by the Object. E.g. U,V coordinates for objects where this makes sense
	Material Material
}
