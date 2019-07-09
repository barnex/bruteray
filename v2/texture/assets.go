package texture

import (
	"math"

	. "github.com/barnex/bruteray/v2/color"
	"github.com/barnex/bruteray/v2/util"
)

/*
	This file provides various procedural texture implementations.
*/

// Checkers returns a checkboard pattern alternating between textures a and b.
// pitchU, pitchV is the number of repetition per unit length.
// E.g.:
// 	Checkers(1, 1, color.White, color.Black)
func Checkers(pitchU, pitchV float64, a, b Texture2D) Texture2D {
	pitchU *= 2
	pitchV *= 2
	return Func2D(func(u, v float64) Color {
		if (floor(u*pitchU)+floor(v*pitchV))%2 == 0 {
			return a.AtUV(u, v)
		} else {
			return b.AtUV(u, v)
		}
	})
}

func Grid(width, pitchU, pitchV float64, a, b Texture2D) Texture2D {
	return Func2D(func(u, v float64) Color {
		u2 := util.Frac(u * pitchU)
		v2 := util.Frac(v * pitchV)
		if u2 < width*pitchU || v2 < width*pitchV || u2 > 1-width*pitchU || v2 > 1-width*pitchV {
			return b.AtUV(u, v)
		} else {
			return a.AtUV(u, v)
		}
	})
}

func floor(x float64) int {
	return int(math.Floor(x))
}
