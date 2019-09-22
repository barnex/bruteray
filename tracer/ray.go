package tracer

import (
	. "github.com/barnex/bruteray/geom"
)

// A Ray is a half-line,
// starting at the Start point (exclusive) and extending in direction Dir.
type Ray struct {
	Start Vec
	Dir   Vec
}

// Returns point Start + t*Dir.
// t must be > 0 for the point to lie on the Ray.
func (r *Ray) At(t float64) Vec {
	return Vec{
		r.Start[0] + t*r.Dir[0],
		r.Start[1] + t*r.Dir[1],
		r.Start[2] + t*r.Dir[2],
	}
}

const Tiny = (1. / (1204 * 1024))
