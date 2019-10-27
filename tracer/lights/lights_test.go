package lights

import (
	"testing"

	"github.com/barnex/bruteray/color"
	"github.com/barnex/bruteray/geom"
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
		cameras.NewProjective(90*Deg, Vec{.250, .250001, 0.9}, 0, 0),
		1,    // recDepth
		5000, // nPass,
		0.8,
	)
}

func TestRectangleLight2(t *testing.T) {
	white := materials.Matte(Color{1, 1, 1})

	test.NPassSize(t,
		NewScene(
			[]Light{
				RectangleLight(Color{0, 1, 1}.EV(7), .2, .2, Vec{0, 2, 0}),
			},
			test.Sheet(white, 0),
			test.Sheet(white, 3),
		),
		cameras.NewProjective(90*Deg, Vec{0, 1, 4}, 0, 0),
		1, 50, // recDepth, nPass
		150, 100, // size
		0.06,
	)
}

func TestRectangleLight3(t *testing.T) {
	white := materials.Matte(Color{1, 1, 1})

	test.NPassSize(t,
		NewScene(
			[]Light{
				//DiskLight(Color{0, 1, 1}.EV(5), .5, Vec{0, 0.03, 0}),
				RectangleLight(Color{0, 1, 1}.EV(5), .5, .5, Vec{0.5, 0.03, 0}),
			},
			test.Sheet(white, 0),
		),
		cameras.NewProjective(90*Deg, Vec{0, 1, 0}, 0, -90*Deg),
		1, 50, // recDepth, nPass
		150, 100, // size
		3, // noisy
	)
}

func TestDiskLight(t *testing.T) {
	white := materials.Matte(Color{1, 1, 1})

	test.NPassSize(t,
		NewScene(
			[]Light{
				DiskLight(Color{0, 1, 1}.EV(5), .5, Vec{0.5, 0.03, 0}),
				//RectangleLight(Color{0, 1, 1}.EV(5), .5, .5, Vec{0, 0.03, 0}),
			},
			test.Sheet(white, 0),
		),
		cameras.NewProjective(90*Deg, Vec{0, 1, 0}, 0, -90*Deg),
		1, 50, // recDepth, nPass
		150, 100, // size
		3, // noisy
	)
}

func TestTransformed(t *testing.T) {
	white := materials.Matte(Color{1, 1, 1})

	test.NPassSize(t,
		NewScene(
			[]Light{
				Transformed(
					RectangleLight(Color{0, 1, 1}.EV(7), .2, .2, Vec{-2, 2, 0}),
					geom.Rotate(Vec{-2, 2, 0}, Vec{0, 0, 1}, 90*Deg),
				),
			},
			test.Sheet(white, 0),
			test.Sheet(white, 3),
		),
		cameras.NewProjective(90*Deg, Vec{0, 1, 4}, 0, 0),
		1, 50, // recDepth, nPass
		150, 100, // size
		0.06,
	)
}

// Compare bi-directionally traced scene with SunLight to
// path traced equivalent (recDepth+1, replace light by its Object()).
// The golden testdata is path traced.
func TestSunLight(t *testing.T) {
	white := materials.Matte(Color{1, 1, 1})
	test.NPassSize(t,
		NewScene(
			[]Light{
				SunLight(Color{1, 1, 1}.EV(1), 1.5, 0*Deg, 20*Deg),
			},
			//SunLight(Color{1, 1, 1}.EV(1), 1.5, 0*Deg, 20*Deg).Object(),
			test.Sphere(white, 2, Vec{1, 1, -1}),
			test.Sheet(white, 0),
		),
		cameras.NewProjective(90*Deg, Vec{0, 1, 2}, 0, 0),
		1, 1000, // recDepth, nPass
		60, 40, // size
		0.9, // noisy
	)
}

// Test that the sun in the zenith yields the desired brightness
// on a matte surface (as specified in the godoc).
// Compare to flat sphere of desired brighness so we can easily see it's correct.
func TestSunLight_Brightness(t *testing.T) {
	white := materials.Matte(Color{1, 1, 1})
	test.NPassSize(t,
		NewScene(
			[]Light{
				SunLight(Color{.2, .5, .7}, 0.1, 0*Deg, 90*Deg),
			},
			//SunLight(Color{2, 2, 2}, 0.5, Vec{0, 0.8, -1}).Object(),
			test.Sphere(materials.Flat(Color{.2, .5, .7}), 2, Vec{1, 1, -1}),
			test.Sheet(white, 0),
		),
		cameras.NewProjective(90*Deg, Vec{0, 1, 2}, 0, 0),
		1, 50, // recDepth nPass
		300, 200, // size
		0.2, // noisy edges
	)
}

// Place a 30 degree wide sun at various positions in the sky
// which are easy to recognize with an environment map camera.
func TestSunLight_Position(t *testing.T) {
	angDiam := 30 * Deg
	test.NPassSize(t,
		NewScene(
			[]Light{
				// Sun at various positions along the horizon
				SunLight(Color{1, 1, 1}.EV(1), angDiam, 0*Deg, 0*Deg),
				SunLight(Color{1, 1, 1}.EV(1), angDiam, -30*Deg, 0*Deg),
				SunLight(Color{1, 1, 1}.EV(1), angDiam, -60*Deg, 0*Deg),
				SunLight(Color{1, 1, 1}.EV(1), angDiam, -90*Deg, 0*Deg),
				SunLight(Color{1, 1, 1}.EV(1), angDiam, -120*Deg, 0*Deg),
				SunLight(Color{1, 1, 1}.EV(1), angDiam, -150*Deg, 0*Deg),

				// various positions at 30 deg pitch
				SunLight(Color{1, 1, 0}.EV(1), angDiam, 0*Deg, 30*Deg),
				SunLight(Color{1, 1, 0}.EV(1), angDiam, -60*Deg, 30*Deg),
				SunLight(Color{1, 1, 0}.EV(1), angDiam, -120*Deg, 30*Deg),

				// various positions at 60 deg pitch
				SunLight(Color{1, 0, 0}.EV(1), angDiam, 0*Deg, 60*Deg),
				SunLight(Color{1, 0, 0}.EV(1), angDiam, -120*Deg, 60*Deg),

				// zenith
				SunLight(Color{0, 0, 1}.EV(1), angDiam, 0*Deg, 90*Deg),
			},
			test.Sheet(test.Normal, 0),
		),
		cameras.EnvironmentMap(Vec{0, 1, 0}),
		1, 1, //recDepth, nPass
		400, 400, // size
		test.DefaultTolerance,
	)
}
