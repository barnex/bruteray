package media

import (
	"math"
	"testing"

	"github.com/barnex/bruteray/imagef/colorf"
	"github.com/barnex/bruteray/tracer/cameras"
	"github.com/barnex/bruteray/tracer/lights"
	"github.com/barnex/bruteray/tracer/materials"
	. "github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/types"
)

// Infinite exponential fog
func TestExpFog(t *testing.T) {
	OnePass(t,
		NewSceneWithMedia(
			1,
			[]Medium{
				ExpFog(0.1, Color{1, 1, 1}, math.Inf(1)),
			},
			[]Light{
				PointLight(Vec{1, 2, 0}),
			},
			Sphere(Checkers(White, Magenta), 1, Vec{0, 0, 0}),
			Sphere(Checkers(White, Green), 1, Vec{1, 0, -1}),
			Sphere(Checkers(White, Green), 1, Vec{-1, 0, -5}),
			Sheet(Checkers(White, Cyan), -0.2),
		),
		cameras.Projective(90*Deg).Translate(Vec{0, 1, 2}).YawPitchRoll(0, -10*Deg, 0),
		DefaultTolerance,
	)
}

// Finite height exponential fog, camera outside.
func TestExpFog_Height(t *testing.T) {
	OnePass(t,
		NewSceneWithMedia(
			1,
			[]Medium{
				ExpFog(2, Color{1, 1, 1}, 0.5),
			},
			[]Light{
				PointLight(Vec{1, 2, 0}),
			},
			Sphere(Checkers(White, Magenta), 1, Vec{0, 0.5, 0}),
			Sphere(Checkers(White, Green), 2, Vec{1, 1, -1}),
			Sphere(Checkers(White, Green), 3, Vec{-3, 1.5, -2}),
			Sheet(Checkers(White, Cyan), 0),
		),
		cameras.Projective(90*Deg).Translate(Vec{0, 1, 2}).YawPitchRoll(0, -10*Deg, 0),
		DefaultTolerance,
	)
}

// Finite height exponential fog, camera inside.
func TestExpFog_Height2(t *testing.T) {
	OnePass(t,
		NewSceneWithMedia(
			1,
			[]Medium{
				ExpFog(0.1, Color{1, 1, 1}, 1.7),
			},
			[]Light{
				PointLight(Vec{1, 2, 0}),
			},
			Sphere(Checkers(White, Magenta), 1, Vec{0, 0.5, 0}),
			Sphere(Checkers(White, Green), 2, Vec{1, 1, -1}),
			Sphere(Checkers(White, Green), 3, Vec{-3, 1.5, -2}),
			Sheet(Checkers(White, Cyan), 0),
		),
		cameras.Projective(90*Deg).Translate(Vec{0, 1, 2}).YawPitchRoll(0, 10*Deg, 0),
		DefaultTolerance,
	)
}

// Infinite pysical fog.
func TestFog(t *testing.T) {
	if testing.Short() {
		t.Skip("Slow")
	}
	w := materials.Matte(colorf.White)
	numLights := 1
	NPassSize(t,
		NewSceneWithMedia(
			1,
			[]Medium{
				Fog(0.1, math.Inf(1), numLights),
			},
			[]Light{
				lights.PointLight(Color{1, 1, 1}, Vec{1, 0.5, 0}),
			},
			Sphere(w, 1, Vec{0, 0, 0}),
			Sphere(w, 1, Vec{1, 0, -1}),
			Sphere(w, 1, Vec{-1, 0, -5}),
			Sheet(w, -0.2),
		),
		cameras.Projective(90*Deg).Translate(Vec{0, 1, 2}).YawPitchRoll(0, -10*Deg, 0),
		1000, // numPass
		150,  // width
		100,  // height
		0.5,  // tolerance
	)
}

// Finite height physical fog, camera outside.
func TestFog_Height(t *testing.T) {
	if testing.Short() {
		t.Skip("Slow")
	}
	w := materials.Matte(colorf.White)
	numLights := 1
	NPassSize(t,
		NewSceneWithMedia(
			1,
			[]Medium{
				Fog(0.3, 1.5, numLights),
			},
			[]Light{
				lights.PointLight(Color{1, 1, 1}, Vec{1, 1.5, 0}),
			},
			Sphere(w, 1, Vec{0, 0.5, 0}),
			//Sphere(w, 2, Vec{1, 1, -1}),
			//Sphere(w, 3, Vec{-3, 1.5, -2}),
			//Sheet(w, 0),
		),
		cameras.Projective(90*Deg).Translate(Vec{0, 2, 2}).YawPitchRoll(0, -30*Deg, 0),
		1000, // numPass
		150,  // width
		100,  // height
		0.5,  // tolerance (noisy)
	)
}

// Finite height exponential fog, camera inside.
func TestFog_Height2(t *testing.T) {
	if testing.Short() {
		t.Skip("Slow")
	}
	w := materials.Matte(colorf.White)
	numLights := 1
	NPassSize(t,
		NewSceneWithMedia(
			1,
			[]Medium{
				Fog(0.3, 1.5, numLights),
			},
			[]Light{
				lights.PointLight(Color{1, 1, 1}, Vec{1, 1.5, 0}),
			},
			Sphere(w, 1, Vec{0, 0.5, 0}),
			//Sphere(w, 2, Vec{1, 1, -1}),
			//Sphere(w, 3, Vec{-3, 1.5, -2}),
			//Sheet(w, 0),
		),
		cameras.Projective(90*Deg).Translate(Vec{0, 1, 2}).YawPitchRoll(0, -30*Deg, 0),
		1000, // numPass
		150,  // width
		100,  // height
		0.5,  // tolerance (noisy)
	)
}
