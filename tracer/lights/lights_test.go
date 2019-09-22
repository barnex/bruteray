package lights

import (
	"testing"

	"github.com/barnex/bruteray/color"
	"github.com/barnex/bruteray/tracer/cameras"
	"github.com/barnex/bruteray/tracer/materials"
	"github.com/barnex/bruteray/tracer/objects"
	"github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/types"
)

// TODO: meaningful test
//func TestPointLight(t *testing.T) {
//	brightness := color.White.EV(0)
//	s := .500
//	white := materials.Matte(Color{1, 1, 1}.EV(-1))
//
//	test.NPass(t,
//		NewScene(
//			[]Light{
//				PointLight(brightness, Vec{s / 2, 0.3, s / 2}),
//			},
//			//objects.Rectangle(materials.Flat(brightness), L, L, Vec{s / 2, s - dy, s / 2}),
//			objects.RectangleWithVertices(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, 0, s}),
//			objects.RectangleWithVertices(white, Vec{0, s, 0}, Vec{s, s, 0}, Vec{0, s, s}),
//			objects.RectangleWithVertices(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, s, 0}),
//			objects.RectangleWithVertices(white, Vec{0, 0, 0}, Vec{0, 0, s}, Vec{0, s, 0}),
//			objects.RectangleWithVertices(white, Vec{s, 0, 0}, Vec{s, 0, s}, Vec{s, s, 0}),
//		),
//		cameras.NewProjective(90*Deg, Vec{.250, .250001, 0.9}),
//		1, // recDepth
//		1, // nPass,
//		0.8,
//	)
//}

// This tests a rectangle light by comparing to a path-traced equivalent.
// The ray-traced test uses recursion depth 1, to exclude indirect lighting.
// The golden data replaces the light source by a flat white rectangle of
// the same brightness, and uses recursion depth 2.
func TestRectangleLight(t *testing.T) {
	t.Skip("Slow")
	brightness := color.White.EV(2)
	s := .500
	L := .350
	dy := 0.0001
	white := materials.Matte(Color{1, 1, 1}.EV(-1))

	test.NPass(t,
		NewScene(
			[]Light{
				RectangleLight(brightness, L, L, Vec{s / 2, s - dy, s / 2}),
			},
			//objects.Rectangle(materials.Flat(brightness), L, L, Vec{s / 2, s - dy, s / 2}),
			objects.RectangleWithVertices(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, 0, s}),
			objects.RectangleWithVertices(white, Vec{0, s, 0}, Vec{s, s, 0}, Vec{0, s, s}),
			objects.RectangleWithVertices(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, s, 0}),
			objects.RectangleWithVertices(white, Vec{0, 0, 0}, Vec{0, 0, s}, Vec{0, s, 0}),
			objects.RectangleWithVertices(white, Vec{s, 0, 0}, Vec{s, 0, s}, Vec{s, s, 0}),
		),
		cameras.NewProjective(90*Deg, Vec{.250, .250001, 0.9}),
		1,    // recDepth
		5000, // nPass,
		0.8,
	)
}
