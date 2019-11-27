// +build example_moebius

// Usage:
// 	go run moebius.go
// or
// 	bruteray-watch moebius.go
package main

import (
	"math"

	. "github.com/barnex/bruteray/api"
)

func main() {

	Render(Spec{
		//DebugNormals: true,
		//DebugIsometricFOV: 8,
		//DebugIsometricDir: Y,

		Recursion: 5,
		NumPass:   200,
		// Postprocess.Bloom.Gaussian.Radius = 0.008
		// Postprocess.Bloom.Gaussian.Amplitude = 0.5
		// Postprocess.Bloom.Gaussian.Threshold = 0.9

		Lights: []Light{
			RectangleLight(C(1, 0.9, 0.8).EV(3), 4, 2, Vec{1, 3, 3}),
		},

		Objects: []Object{
			Parametric(
				Shiny(
					LoadTexture("assets/moebius.png").Scale(1./32., 1./4),
					0.1,
				),
				256, 16, // numU, numV
				func(u, v float64) Vec {
					w := 0.2
					u *= 2 * Pi
					v = 2 * (v - 0.5)
					x := (1 + w*v*cos(0.5*u)) * cos(u)
					y := (1 + w*v*cos(0.5*u)) * sin(u)
					z := 0.5 * v * sin(0.5*u)
					return V(x, y, z)
				}).
				Rotate(Ex, 90*Deg).
				Rotate(Ey, -20*Deg).
				WithCenterBottom(O).
				Rotate(Ez, -7*Deg),

			Rectangle(Matte(White), 10, 10, V(0, -1, 0)),

			Backdrop(Flat(C(0.6, 0.8, 1.0).EV(-2))),
		},

		Camera: Projective(60*Deg, Vec{0, 2, 5}, 0, -25*Deg),
	})

	//tex := texture.Pan(
	//	texture.MustLoad("../../../assets/monalisa.jpg"),
	//	0, 0,
	//)
	//mat := Blend(
	//	0.9, Matte(tex),
	//	0.1, Reflective(White),
	//)
	//Add(strip)

	//white := Matte(Color{1, 0.9, 0.8}.EV(-.1))
	//floor := Sheet(white, Vec{}, Ex, Ez)
	//Add(floor)

	//Add(builder.Ambient(White.EV(-4)))

	//Camera.Translate(Vec{0, 3, -2.6}.Mul(0.8))
	//Camera.Pitch(-46 * Deg)
	//Camera.FocalLen = 1.1
	//Camera.Focus = 2.5
	//Camera.Aperture = 0.03
	//Render()
}

var sin = math.Sin
var cos = math.Cos
