/*
Package objects provides concrete implementations of the tracer.Object interface.

In addition to the basic tracer.Object interface, the implementations in this package
provide accelerated ray-object intersection and constructive solid geometry (CSG) operations.
Both are hidden to the ray tracing algorithm. I.e., package tracer does not need to know
about this.
*/
package objects

import (
	. "github.com/barnex/bruteray/tracer/types"
)

// Interface satisfied by all tracer.Object implementations in this package.
// In addition to tracer.Object, this interface adds Bounds and Inside,
// which are used for efficient intersection tests
// and to construct solid geometry (CSG), respectively.
type Interface interface {
	Object

	// Bounds returns an axis-aligned bounding box completely enclosing the object.
	// This allows for accelerated ray-object intersection using a bounding volume hierarchy (BVH).
	Bounds() BoundingBox

	// Inside returns true if a point lies inside the Object's surface.
	// This allows for objects to be composed via constructive solid geometry (CSG).
	//
	// For hollow objects (pure surfaces that are empty on the inside),
	// Inside always returns false.
	Inside(point Vec) bool
}

// hollowSurface can be embedded to inherit an Inside method
// that always returns false.
type hollowSurface struct{}

// Inside on a hollow surface always returns false.
func (hollowSurface) Inside(Vec) bool {
	return false
}
