package cameras_test

import (
	"math"
	"testing"

	"github.com/barnex/bruteray/color"
	. "github.com/barnex/bruteray/tracer/cameras"
	"github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/types"
)

func focalLenToFOV(l float64) float64 {
	return 2 * math.Atan(0.5/l)
}

func TestTransform(t *testing.T) {
	radius := 0.10
	test.OnePass(t,
		NewScene(
			[]Light{},
			test.Sphere(test.Flat(color.White), radius, Vec{0, 0, 0}),
			test.Sphere(test.Flat(color.Red), radius, Vec{1, 0, 0}),
			test.Sphere(test.Flat(color.Green), radius, Vec{0, 1, 0}),
			test.Sphere(test.Flat(color.Blue), radius, Vec{0, 0, 1}),
		),
		NewProjective(focalLenToFOV(1), Vec{0, 0, 5}).YawPitchRoll(-0.4, 0.2, 0.5),
		test.DefaultTolerance,
	)
}

func TestIsometric(t *testing.T) {
	cam := NewIsometric(Z, 3)
	cam.Pitch(-90 * Deg)
	cam.Translate(Vec{1, 0, 0})
	radius := 0.10
	test.OnePass(t,
		NewScene(
			[]Light{},
			test.Sphere(test.Flat(color.White), radius, Vec{0, 0, 0}),
			test.Sphere(test.Flat(color.Red), radius, Vec{1, 0, 0}),
			test.Sphere(test.Flat(color.Green), radius, Vec{0, 1, 0}),
			test.Sphere(test.Flat(color.Blue), radius, Vec{0, 0, 1}),
		),
		cam,
		test.DefaultTolerance,
	)
}

// Test the Projective camera by rendering 4 small spheres
// positioned at the origin, Ex, Ey and Ez.
// The resulting image should correspond to a right-handed coordinate system.
func TestProjective_Handedness(t *testing.T) {
	const radius = 0.10
	test.OnePass(t,
		NewScene(
			[]Light{},
			test.Sphere(test.Flat(color.White), radius, Vec{0, 0, 0}),
			test.Sphere(test.Flat(color.Red), radius, Vec{1, 0, 0}),
			test.Sphere(test.Flat(color.Green), radius, Vec{0, 1, 0}),
			test.Sphere(test.Flat(color.Blue), radius, Vec{0, 0, 1}),
		),
		NewProjective(focalLenToFOV(1), Vec{0, 1, 5}),
		test.DefaultTolerance,
	)
}

// Test Projective camera field of view.
// Place two red spheres at unit distances from the orign.
// Set a non-trival focal length (2) and position (z=4) so that both spheres
// are exactly at the edges of the image:
//
//      1      1
//  x------o------x  object plane
//   \     |2    /
//    \ .5 | .5 /
//     x---o---x     image plane
//      \  |2 /
//       \ | /
//        cam
func TestProjective_FOV(t *testing.T) {
	const radius = 0.10
	test.OnePass(t,
		NewScene(
			[]Light{},
			test.Sphere(test.Flat(color.White), radius, Vec{0, 0, 0}),
			test.Sphere(test.Flat(color.Red), radius, Vec{1, 0, 0}),
			test.Sphere(test.Flat(color.Red), radius, Vec{-1, 0, 0}),
		),
		NewProjective(focalLenToFOV(2), Vec{0, 0, 4}),
		test.DefaultTolerance,
	)
}
