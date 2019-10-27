// +build example_celtic_cemetary

// Usage:
// 	go run celtic-cemetary.go
// or
// 	bruteray-watch celtic-cemetary.go
package main

import (
	"math"

	. "github.com/barnex/bruteray/api"
)

func main() {

	cross1 := cross1(Matte(White.Mul(0.7)))
	Render(Spec{
		//DebugNormals: true,
		//DebugIsometricFOV: 40,
		//DebugIsometricDir: Y,

		Recursion: 2,
		NumPass:   800,
		Width: 300,
		Height: 200,
		// Postprocess.Bloom.Gaussian.Radius = 0.008
		// Postprocess.Bloom.Gaussian.Amplitude = 0.5
		// Postprocess.Bloom.Gaussian.Threshold = 0.9

		Lights: []Light{
			SunLight(C(1, .95, .9).EV(-1), 0.53*Deg, -28*Deg, 9*Deg),
		},

		Media: []Medium{
			Fog(0.8, 0.6),
			Fog(0.4, 2),
			//ExpFog(0.3, C(1, 1, 1), 0.9),
			//ExpFog(0.1, C(1, 1, 1), 2.2),
		},

		Objects: []Object{
			//Sphere(Flat(White.EV(10)), 2, lp.Mul(1.1)),
			Tree(
				cross1.Translate(V(-2, 0, 0.1)),
				cross1.Translate(V(0, 0, -0.2)),
				cross1.Rotate(Ez, -6*Deg).WithCenterBottom(V(2, -0.1, 0)),

				cross1.Translate(V(-2, 0, -1.1)),
				cross1.Rotate(Ez, 5*Deg).WithCenterBottom(V(0.3, -0.1, -1.2)),
				cross1.WithCenterBottom(V(2, -0.1, -1)),

				cross1.Rotate(Ez, -5*Deg).WithCenterBottom(V(-2, -0.1, -2.3)),
				cross1.Rotate(Ez, 3*Deg).WithCenterBottom(V(0.3, -0.1, -2.1)),
				cross1.Rotate(V(0.1, 0, 1), -1*Deg).WithCenterBottom(V(2, -0.2, -2.1)),
			),

			Rectangle(Matte(C(0.5, 0.9, 0.5).EV(-1)), 100, 100, O),
			Backdrop(Flat(C(0.8, 0.8, 0.9).EV(-2.6))),
		},

		//Media: []Medium{
		//	media.ExpFog(0.15, White),
		//},

		Camera: Projective(100*Deg, Vec{0.3, 1.0, 2.15}, 0, -1*Deg, ),
	})
}

func cross1(m Material) Object {
	w1 := 0.20
	H := 1.8
	W := 0.9
	d1 := 0.15

	d2 := 0.08
	r1 := 0.47
	r2 := 0.69

	vbeam := Box(m, w1, H, d1, V(0, H/2, 0))
	hbeam := Box(m, W, w1, d1, V(0, H-W/2, 0))
	ring := Difference(
		CylinderWithCaps(m, r2, d2, O),
		Cylinder(m, r1, math.Inf(1), O),
	).Rotate(Ex, 90*Deg).WithCenter(hbeam.Center())
	return Tree(vbeam, hbeam, ring)
}
