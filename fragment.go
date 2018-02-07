package bruteray

// A Fragment calculates the color seen by a ray.
// Fregment shading is lazily evaluated:
// only when the frontmost shader has been determined
// will we call its Shade method. Shaders returned by
// objects hidden behind others will eventually not be used.
type Fragment struct {
	T        float64 // The t-value where the ray hit the object. Used to determine the frontmost Shader.
	Norm     Vec     // Surface normal where the ray hit the object. Passed to Material.
	Material         // Material.Shade will be called with the relevant position and normal to finally calculate the color.
	Object   Obj
}

// When a secondary ray is cast from a surface,
// we add this tiny offset to its starting position.
// This avoids that numerical round-off
// would cause the ray to start inside the surface.
const offset = 1. / (256 * 1024)

// Calculate the color seen by ray.
func (frag Fragment) Shade(e *Env, recursion int, r *Ray) Color {
	//pos := r.At(frag.T - offset)
	frag.Norm = frag.Norm.Towards(r.Dir()).Normalized()
	//pos = pos.MAdd(offset, norm)
	return frag.Material.Shade(e, recursion, r, frag)
}
