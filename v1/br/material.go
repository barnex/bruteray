package br

// A Material determines the color of a surface fragment.
// E.g.: mate white, glossy red, ...
type Material interface {

	// Shade must return the color of the given surface fragment,
	// as seen by Ray r.
	// If Shade uses recursion, e.g., to calculate reflections,
	// it must pass N-1 as the new recursion depth, so that
	// recursion can eventually be terminated (by Env.Shade).
	Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color
}
