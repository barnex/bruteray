package cameras_test

import (
	"math"
	"testing"

	"github.com/barnex/bruteray/imagef/colorf"
	. "github.com/barnex/bruteray/tracer/cameras"
	"github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/types"
)

func focalLenToFOV(l float64) float64 {
	return 2 * math.Atan(0.5/l)
}

var (
	red   = test.Flat(colorf.Red)
	green = test.Flat(colorf.Green)
	blue  = test.Flat(colorf.Blue)
	white = test.Flat(colorf.White)
)

func TestTranslate(t *testing.T) {
	cam := Isometric(Z, 3).Translate(Vec{1, 0, 0})
	radius := 0.10
	test.OnePass(t,
		NewScene(
			1,
			[]Light{},
			test.Sphere(white, radius, Vec{0, 0, 0}),
			test.Sphere(red, radius, Vec{1, 0, 0}),
			test.Sphere(green, radius, Vec{0, 1, 0}),
			test.Sphere(blue, radius, Vec{0, 0, 1}),
		),
		cam,
		test.DefaultTolerance,
	)
}

func TestTransform(t *testing.T) {
	t.Skip("TODO")
	radius := 0.10
	test.OnePass(t,
		NewScene(
			1,
			[]Light{},
			test.Sphere(white, radius, Vec{0, 0, 0}),
			test.Sphere(red, radius, Vec{1, 0, 0}),
			test.Sphere(green, radius, Vec{0, 1, 0}),
			test.Sphere(blue, radius, Vec{0, 0, 1}),
		),
		Projective(focalLenToFOV(1)).YawPitchRoll(-0.4, 0.2, 0.5).Translate(Vec{0, 0, 5}),
		test.DefaultTolerance,
	)
}

// Render a scene similar to TestProjective_Handedness:
// Red sphere at x=+1, green at y=+1, blue at z=+1.
//
// The camera is at (0,0,0) and should see all three spheres:
//   - the sphere at x=+1 should be to the right of image's center (right-handed space)
//   - the sphere at y=+1 should be stretched out along the top (heavily distorted)
//   - the sphere at z=+1 is behind the camera and thus should be at the image's edges
func TestEnvironmentMap(t *testing.T) {
	radius := 0.5
	test.NPassSize(t,
		NewScene(
			1,
			[]Light{},
			test.Sphere(red, radius, Vec{1, 0, 0}),
			test.Sphere(green, radius, Vec{0, 1, 0}),
			test.Sphere(blue, radius, Vec{0, 0, 1}),
		),
		EnvironmentMap(),
		1,
		400, 400,
		test.DefaultTolerance,
	)
}

func TestIsometricZ(t *testing.T) {
	cam := Isometric(Z, 3)
	radius := 0.10
	test.OnePass(t,
		NewScene(
			1,
			[]Light{},
			test.Sphere(white, radius, Vec{0, 0, 0}),
			test.Sphere(red, radius, Vec{1, 0, 0}),
			test.Sphere(green, radius, Vec{0, 1, 0}),
			test.Sphere(blue, radius, Vec{0, 0, 1}),
		),
		cam,
		test.DefaultTolerance,
	)
}

func TestIsometricX(t *testing.T) {
	cam := Isometric(X, 3)
	radius := 0.10
	test.OnePass(t,
		NewScene(
			1,
			[]Light{},
			test.Sphere(white, radius, Vec{0, 0, 0}),
			test.Sphere(red, radius, Vec{1, 0, 0}),
			test.Sphere(green, radius, Vec{0, 1, 0}),
			test.Sphere(blue, radius, Vec{0, 0, 1}),
		),
		cam,
		test.DefaultTolerance,
	)
}

func TestIsometricY(t *testing.T) {
	cam := Isometric(Y, 3)
	radius := 0.10
	test.OnePass(t,
		NewScene(
			1,
			[]Light{},
			test.Sphere(white, radius, Vec{0, 0, 0}),
			test.Sphere(red, radius, Vec{1, 0, 0}),
			test.Sphere(green, radius, Vec{0, 1, 0}),
			test.Sphere(blue, radius, Vec{0, 0, 1}),
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
			1,
			[]Light{},
			test.Sphere(white, radius, Vec{0, 0, 0}),
			test.Sphere(red, radius, Vec{1, 0, 0}),
			test.Sphere(green, radius, Vec{0, 1, 0}),
			test.Sphere(blue, radius, Vec{0, 0, 1}),
		),
		Projective(focalLenToFOV(1)).Translate(Vec{0, 1, 5}),
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
			1,
			[]Light{},
			test.Sphere(white, radius, Vec{0, 0, 0}),
			test.Sphere(red, radius, Vec{1, 0, 0}),
			test.Sphere(red, radius, Vec{-1, 0, 0}),
		),
		Projective(focalLenToFOV(2)).Translate(Vec{0, 0, 4}),
		test.DefaultTolerance,
	)
}
