package texture

import (
	. "github.com/barnex/bruteray/v2/color"
)

// Func2D adapts a function to a 2D texture.
type Func2D func(u, v float64) Color

// AtUV implements Texture2D
// by calling the function.
func (f Func2D) AtUV(u, v float64) Color {
	return f(u, v)
}
