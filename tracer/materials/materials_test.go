package materials_test

// TODO: these tests should not depend on builder

import (
	"testing"

	"github.com/barnex/bruteray/color"
	"github.com/barnex/bruteray/tracer/cameras"
	"github.com/barnex/bruteray/tracer/lights"
	"github.com/barnex/bruteray/tracer/test"

	. "github.com/barnex/bruteray/tracer/materials"
	. "github.com/barnex/bruteray/tracer/types"
)

// Place a point light exactly in between two spheres.
// The sphere behind the light should not cause a shadow
// on the sphere in front of the light, because the
// shadow ray hits the light first.
func TestMatte_Shadow(t *testing.T) {
	test.QuadView(t,
		NewScene(
			[]Light{
				lights.PointLight(color.White.EV(2), Vec{1, 1, 0}),
			},
			test.Sheet(Matte(color.White), -0.5),
			test.Sphere(Matte(color.White), 1, Vec{0, 0, 0}),
			test.Sphere(Matte(color.White), 1, Vec{2, 2, 0}),
		),
		cameras.NewProjective(90*Deg, Vec{0, 0.5, 3}, 0, 0),
		8,
		test.DefaultTolerance,
	)
}
func TestReflective(t *testing.T) {
	test.QuadViewN(t,
		NewScene(
			[]Light{
				lights.PointLight(color.White.EV(2), Vec{1, 1, 0}),
			},
			test.Sheet(test.Checkers4, -0.5),
			test.Sphere(Reflective(color.Color{1, 1, 0.5}), 1, Vec{0, 0, 0}),
			test.Sphere(Reflective(color.Color{1, 1, 1}), 1, Vec{1, 0, 0.5}),
			test.Sphere(test.Checkers2, 2, Vec{-2, 0.5, 1}),
		),
		cameras.NewProjective(90*Deg, Vec{0, 0.5, 1.8}, 0, -10*Deg, ),
		8, // isometric fov
		3, // recDepth
		1, // nPass
		test.DefaultTolerance,
	)
}

func TestRefractive(t *testing.T) {
	t.Skip("Looks wrong?")
	test.QuadViewN(t,
		NewScene(
			[]Light{
				lights.PointLight(color.White.EV(2), Vec{1, 1, 2}),
			},
			test.Sheet(test.Checkers4, -0.5),
			test.Sheet(test.Flat(color.Color{0.5, 0.5, 0.8}), 20000000),

			test.Sphere(Refractive(1.0001), 1, Vec{-3, 0, 0}),
			test.Sphere(Refractive(1.05), 1, Vec{-2, 0, 0}),
			test.Sphere(Refractive(1.20), 1, Vec{-1, 0, 0}),
			test.Sphere(Refractive(1.50), 1, Vec{+0, 0, 0}),
			test.Sphere(Refractive(1.97), 1, Vec{+1, 0, 0}), // looks wrong
			test.Sphere(Refractive(2.00), 1, Vec{+2, 0, 0}), // looks wrong
			test.Sphere(Refractive(10.0), 1, Vec{+3, 0, 0}),
		),
		cameras.NewProjective(90*Deg, Vec{0, 0.3, 3.5}, 0, -10*Deg, ),
		8, // isometric fov
		9, // recDepth
		1, // nPass
		test.DefaultTolerance,
	)
}

func TestTransparent(t *testing.T) {
	test.QuadViewN(t,
		NewScene(
			[]Light{
				test.PointLight(Vec{0, 2, 0}),
			},
			test.Sphere(Transparent(color.Color{0, 1, 0}, false), 1, Vec{0, 0.9, 0}),
			test.Sheet(Matte(color.White), 0), // floor
		),
		cameras.NewProjective(90*Deg, Vec{0, 1, 2.5}, 0, 0),
		6,   // isometric fov
		3,   // recDepth
		1,   // nPass
		0.7, // tolerance
	)
}

// TODO: Test texturing
//func TestFlat(t *testing.T) {
//	test.QuadView(t,
//		NewScene(
//			[]Light{},
//			test.Sphere(Flat(color.Red), 1, Vec{0, 0, 2}),
//		),
//		cameras.NewProjective(90*Deg, Vec{0, 0, 0}),
//		8,
//		test.DefaultTolerance,
//	)
//}
/*
func TestRefractive(t *testing.T) {
	scene := NewSceneBuilder()
	mat := Refractive(1.5)

	sph := NewSphere(mat, 2)
	Translate(sph, Vec{0, 1, 2})
	scene.Add(sph)

	{
		tex := texture.Checkers(1, 1, color.Blue, color.White)
		floor := NewSheet(Flat(tex), O, Ex, Ez)
		scene.Add(floor)
	}

	scene.Add(Ambient(color.Yellow.EV(-2)))

	scene.Camera.FocalLen = 1
	scene.Camera.Translate(Vec{0, 1.2, -1})
	built := scene.Build()
	test.NPass(t, built, 5, 1, test.DefaultTolerance)
}

*/
