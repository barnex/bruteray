package br

type Light interface {
	Sample(e *Env, target Vec) (pos Vec, intens Color)
	Obj
}
