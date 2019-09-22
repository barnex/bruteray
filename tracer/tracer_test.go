package tracer_test

import (
	"testing"

	"github.com/barnex/bruteray/tracer/cameras"
	"github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/types"
)

func TestTracer(t *testing.T) {
	QuadView(t,
		NewScene(
			[]Light{
				PointLight(Vec{1, 2, 0}),
			},
			Sphere(Checkers(White, Magenta), 1, Vec{0, 0, 0}),
			Sheet(Checkers(White, Cyan), -0.2),
			Sphere(Checkers(White, Green), 1, Vec{1, 0, -1}),
		),
		cameras.NewProjective(90*Deg, Vec{0, 0, 2}),
		8,
		test.DefaultTolerance,
	)
}
