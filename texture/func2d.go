package texture

import (
	. "github.com/barnex/bruteray/color"
	. "github.com/barnex/bruteray/geom"
)

// Func2D adapts a function to a 2D texture.
type Func2D func(u, v float64) Color

// AtUV implements Texture2D
// by calling the function.
func (f Func2D) AtUV(u, v float64) Color {
	return f(u, v)
}

func (f Func2D) At(p Vec) Color {
	return f(p[0], p[1])
}

type Func func(p Vec) Color

func (f Func) At(p Vec) Color {
	return f(p)
}
