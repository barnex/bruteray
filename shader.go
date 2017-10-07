package bruteray

// A Shader calculates the color seen by a ray.
// Shaders are lazily evaluated:
// only when the frontmost shader has been determined
// will we call its Shade method. Shaders returned by
// objects hidden behind others will eventually not be used.
type Shader struct {
	T        float64 // The t-value where the ray hit the object. Used to determine the frontmost Shader.
	Norm     Vec     // Surface normal where the ray hit the object. Passed to Material.
	Material         // Material.Shade will be called with the relevant position and normal to finally calculate the color.
}

// When a secondary ray is cast from a surface,
// we add this tiny offset to its starting position.
// This avoids that numerical round-off
// would cause the ray to start inside the surface.
const offset = 1. / (256 * 1024)

// Calculate the color seen by ray.
func (s *Shader) Shade(e *Env, recursion int, r *Ray) Color {
	pos := r.At(s.T)
	norm := s.Norm.Towards(r.Dir)
	pos = pos.MAdd(offset, norm)
	return s.Material.Shade(e, r, recursion, pos, norm)
}
