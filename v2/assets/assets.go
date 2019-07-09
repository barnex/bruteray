//Package assets is a library of various resusable models.
package assets

import (
	. "github.com/barnex/bruteray/v2/api"
	"github.com/barnex/bruteray/v2/builder"
	"github.com/barnex/bruteray/v2/tracer"
)

func Box(t tracer.Material, wallThickness float64) *builder.Tree {
	var (
		d = wallThickness
		s = 1.0 - d
		S = s + 2*d
	)
	tr := Tree(
		Rectangle(t, Vec{0, 0, 0}, Vec{0, 0, s}, Vec{s, 0, 0}), // floor
		Rectangle(t, Vec{0, s, 0}, Vec{s, s, 0}, Vec{0, s, s}), // ceiling
		Rectangle(t, Vec{0, 0, s}, Vec{0, s, s}, Vec{s, 0, s}), // back
		Rectangle(t, Vec{0, 0, 0}, Vec{0, s, 0}, Vec{0, 0, s}), // left
		Rectangle(t, Vec{s, 0, 0}, Vec{s, 0, s}, Vec{s, s, 0}), //right

		Rectangle(t, Vec{0, 0, 0}, Vec{S, 0, 0}, Vec{0, 0, S}).Translate(Vec{-d, -d, 0}),
		Rectangle(t, Vec{0, S, 0}, Vec{0, S, S}, Vec{S, S, 0}).Translate(Vec{-d, -d, 0}),
		Rectangle(t, Vec{0, 0, S}, Vec{S, 0, S}, Vec{0, S, S}).Translate(Vec{-d, -d, 0}),
		Rectangle(t, Vec{0, 0, 0}, Vec{0, 0, S}, Vec{0, S, 0}).Translate(Vec{-d, -d, 0}),
		Rectangle(t, Vec{S, 0, 0}, Vec{S, 0, S}, Vec{S, S, 0}).Translate(Vec{-d, -d, 0}),

		Rectangle(t, Vec{-d, s, 0}, Vec{-d, s + d, 0}, Vec{s + d, s, 0}),
		Rectangle(t, Vec{-d, s, 0}, Vec{-d, 0, 0}, Vec{0, s, 0}),
		Rectangle(t, Vec{s, s, 0}, Vec{s, 0, 0}, Vec{s + d, s, 0}),
		Rectangle(t, Vec{-d, 0, 0}, Vec{-d, 0 - d, 0}, Vec{s + d, 0, 0}),
	)
	Translate(tr, Vec{0, d, 0})
	return tr
}

//func CornellBox(){
//	L = 0.075
//	var (
//		white = Mate(Color{0.708, 0.743, 0.767})
//		green = Mate(Color{0.115, 0.441, 0.101})
//		red   = Mate(Color{0.651, 0.059, 0.061})
//	)
//		RectangleLight(Color{18.4, 15.6, 8.0}.EV(-1), Vec{s/2 - L, s - .001, s/2 - L}, Vec{s/2 + L, s - .001, s/2 - L}, Vec{s/2 - L, s - .001, s/2 + L}),
//}
