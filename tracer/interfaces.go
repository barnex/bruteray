package tracer

import (
	. "github.com/barnex/bruteray/color"
	. "github.com/barnex/bruteray/geom"
)

// A Camera maps pixel positions to Rays.
// Various implementations are in package cameras.
type Camera interface {
	// RayFrom constructs a ray starting from position u,v on the image sensor.
	// u, v are both in the interval [0,1].
	//
	// In case a non-square image is desired, u or v can simply be varied
	// over a rectangle inside the unit square.
	//
	// Ctx should be used to allocate the Ray,
	// and to generate random lens samples if needed.
	RayFrom(ctx *Ctx, u, v float64) *Ray
}

// TODO: rename surface?
type Object interface {
	Intersect(r *Ray) HitRecord
}

// A HitRecord records the position and material where a Ray intersected an Object.
// This allows for lazy evaluation of the Color seen by the Ray, after it has been
// established that the color is actually going to be used
// (i.e., the hit point is not occluded by another Object).
// Without this lazy evaluation, the complexity of recursive ray tracing would
// become exponential in the recursion depth.
//
// The structure is designed to avoid allocations:
// It is passed by value, and the Material is only allocated once at initialization time.
type HitRecord struct {
	T        float64 // Position along the ray
	Normal   Vec     // Normal vector at intersection (needs not have unit length)
	Local    Vec     // Local coordinates at intersection, chosen by the Object. E.g. U,V coordinates for objects where this makes sense
	Material Material
}

type Light interface {
	// Object is rendered when a viewing ray sees the light directly.
	// E.g. for a spherical yellow light source, this object is a
	// bright yellow sphere.
	//
	// Care must be taken that the properties of this object
	// (size, surface brightness) exactly match the light intensity
	// obtained by calling Sample(). I.e. The amount of light
	// that a scene receives from this object via unidirectional path
	// tracing must be exactly equal to the amout of light
	// received from Sample() using bidirectional path tracing.
	Object() Object

	// Sample returns a position on the light's surface
	// (used to determine shadows for extended sources),
	// and the intensity at given target position.
	Sample(ctx *Ctx, target Vec) (pos Vec, intens Color)
}

// A Material determines the color of a surface fragment.
type Material interface {

	// Eval returns the color of the seen by Ray r which intersects
	// the scene at the given hit coordinates.
	// If the implementation uses recursion by calling s.Eval,
	// it must pass the recursion depth recDepth unmodified.
	// Scene.Eval takes care to decrement the recursion depth
	// and terminate recursion.
	//
	//TODO: rename: Shade?
	Eval(ctx *Ctx, s *Scene, r *Ray, recDepth int, h HitCoords) Color
}

// HitCoords record where a Ray intersected an Object.
type HitCoords struct {
	T      float64 // Position along the ray
	Normal Vec     // Normal vector at intersection (needs not have unit length)
	Local  Vec     // Local coordinates at intersection, chosen by the Object. E.g. U,V coordinates for objects where this makes sense
}

type TransparentMaterial interface {
	Filter(r *Ray, h HitRecord, background Color) Color
}

type Medium interface {
	Filter(ctx *Ctx, s *Scene, r *Ray, tMax float64, original Color) Color
}
