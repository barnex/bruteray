package tracer_test

import (
	"testing"

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
		cameras.NewProjective(90*Deg, Vec{0, 0, 2}, 0, 0),
		8,
		test.DefaultTolerance,
	)
}
