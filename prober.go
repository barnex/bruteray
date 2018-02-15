package bruteray

type Prober interface {
	// Probe is a lightweight intersection test.
	// Probe must return the (positive) intersection distance
	// if the half-ray intersects the object.
	// If the half-ray does not intersect (including intersection at t < 0),
	// any non-positive number may be returned, including 0 and NaN.
	Probe(*Ray) float64
}

// Probe implementation for objects that only have Hit.
// TODO: remove.
func Probe(r *Ray, o CSGObj) float64 {
	//if p, ok := o.(Prober); ok {
	//	return p.Probe(r)
	//}

	T := inf
	hit := make([]Fragment, 0, 2)

	o.Hit1(r, &hit)

	for i := range hit {
		t := hit[i].T
		if t < T && t > 0 {
			T = t
		}
	}

	if T == inf {
		T = 0
	}
	return T
}
