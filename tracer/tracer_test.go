package tracer_test

import (
	"fmt"
	"testing"

	"github.com/barnex/bruteray/tracer"
	"github.com/barnex/bruteray/tracer/cameras"
	"github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/types"
)

//func TestConvergence(t *testing.T) {
//	NPassSize(t,
//		NewScene(
//			1,
//			[]Light{
//				PointLight(Vec{1, 2, 0}),
//			},
//			Sphere(Checkers(White, Magenta), 1, Vec{0, 0, 0}),
//			Sheet(Checkers(White, Cyan), -0.2),
//			Sphere(Checkers(White, Green), 1, Vec{1, 0, -1}),
//		),
//		cameras.NewProjective(90*Deg, Vec{0, 0, 2}, 0, 0),
//		1,
//		60,
//		40,
//		test.DefaultTolerance,
//	)
//}

// Basic test of the core tracing logic:
//   - find frontmost intersection
//   - local coordinates and texturing
//   - point lighting
// Uses simple shapes and materials from package "test",
// so this test is not affected by real shapes, materials.
func TestTracer(t *testing.T) {
	QuadView(t,
		NewScene(
			1,
			[]Light{
				PointLight(Vec{1, 2, 0}),
			},
			Sphere(Checkers(White, Magenta), 1, Vec{0, 0, 0}),
			Sheet(Checkers(White, Cyan), -0.2),
			Sphere(Checkers(White, Green), 1, Vec{1, 0, -1}),
		),
		cameras.Projective(90*Deg).Translate(Vec{0, 0, 2}),
		8,
		test.DefaultTolerance,
	)
}

func ExampleIndexToCam() {
	for _, arg := range []struct {
		w, h   int
		ix, iy float64
	}{
		{1, 1, -.5, -.5},
		{1, 1, 0, 0},
		{1, 1, 0.5, 0.5},

		{4, 2, -.5, -.5},
		{4, 2, 1.5, 0.5},
		{4, 2, 3.5, 1.5},
	} {
		x, y := tracer.IndexToCam(arg.w, arg.h, arg.ix, arg.iy)
		fmt.Printf("[w: %2d, h: %2d]: (% 3.3f, % 3.3f)->(% 3.3f, % 3.3f)\n",
			arg.w, arg.h, arg.ix, arg.iy, x, y)
	}
	// Output:
	// [w:  1, h:  1]: (-0.500, -0.500)->( 0.000,  1.000)
	// [w:  1, h:  1]: ( 0.000,  0.000)->( 0.500,  0.500)
	// [w:  1, h:  1]: ( 0.500,  0.500)->( 1.000,  0.000)
	// [w:  4, h:  2]: (-0.500, -0.500)->( 0.000,  0.750)
	// [w:  4, h:  2]: ( 1.500,  0.500)->( 0.500,  0.500)
	// [w:  4, h:  2]: ( 3.500,  1.500)->( 1.000,  0.250)
}
