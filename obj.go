package bruteray

// Obj is an object that can be rendered as part of a scene.
// E.g., a red sphere, a blue cube, ...
type Obj interface {

	// Hit calculates the intersection between the object and a ray.
	// It appends to *f a surface fragment for each intersection with the ray.
	// The fragments do not need to be sorted.
	Hit(r *Ray, f *[]Fragment)
}

// CSGObj is an object that can be used with Constructive Solid Geometry.
type CSGObj interface {
	Obj

	// Inside returns true if point p lies inside the object.
	Inside(p Vec) bool
}

// embed noInside to get a hollow object.
type noInside struct{}

func (noInside) Inside(Vec) bool {
	return false
}

type Insider interface {
	Inside(pos Vec) bool
}
