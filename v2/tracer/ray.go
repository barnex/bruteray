package tracer

import (
	. "github.com/barnex/bruteray/v2/geom"
)

// TODO: move to geom

// A Ray is a half-line,
// starting at the Start point (exclusive) and extending in direction Dir.
type Ray struct {
	Start Vec
	Dir   Vec
	Len   float64 //TODO: rm
}

// Returns point Start + t*Dir.
// t must be > 0 for the point to lie on the Ray.
func (r *Ray) At(t float64) Vec {
	// pprof shows this is where we spend most of our time.
	// manually inlined for ~10% overall performance improvement.
	return Vec{
		r.Start[X] + t*r.Dir[X],
		r.Start[Y] + t*r.Dir[Y],
		r.Start[Z] + t*r.Dir[Z],
	}
}
