package br

// Light is an arbitrary kind of light source.
// Implementations are in package light.
type Light interface {
	Sample(ctx *Ctx, target Vec) (pos Vec, intens Color)

	Obj
}
