package bruteray

type Shape interface {
	// Hit returns the t-value where the ray first hits the surface.
	// I.e. r.At(t) lies on the surface.
	// If the ray intersects the surface multiple times,
	// the smallest positive t must be returned.
	// t<=0 is interpreted as the ray not intersecting the surface.
	Hit(r *Ray) float64

	// Normal vector at position.
	// Does not necessarily need to point outwards.
	Normal(pos Vec) Vec
}

// TODO: rename Volume
type SolidShape interface {
	Shape

	// Ray-shape intersection.
	// Inters must return ALL intersection point with r,
	// even those with t <= 0.
	Inters(r *Ray) []Interval
}

// TODO: rename Surface
type HollowShape interface {
	Shape

	// Ray-shape intersection.
	// Inters must return ALL intersection point with r,
	// even those with t <= 0.
	HitAll(r *Ray) []float64
}
