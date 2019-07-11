package br

// Light is an arbitrary kind of light source.
// Implementations are in package light.
type Light interface {

	// Sample returns a position on the light's surface
	// (used to determine shadows for extended sources),
	// and the intensity at given target position.
	Sample(ctx *Ctx, target Vec) (pos Vec, intens Color)

	// Obj is rendered when the camera looks at the light directly.
	Obj
}
