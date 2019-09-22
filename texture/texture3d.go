package texture

import (
	. "github.com/barnex/bruteray/geom"
)

type ConstScalar3D float64

func (c ConstScalar3D) At(_ Vec) float64 {
	return float64(c)
}

// TODO: rename
func ScalarFunc3D(f func(Vec) float64) Texturef {
	return &scalarFunc3D{f}
}

type scalarFunc3D struct {
	f func(Vec) float64
	//Frame
}

func (f *scalarFunc3D) At(v Vec) float64 {
	return f.f(v)
	//return f(f.Frame.TransformToFrame(v))
}
