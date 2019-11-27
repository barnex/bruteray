//+build example_abstract_parametrics

// Usage:
// 	go run abstract-parametrics.go
// or
// 	bruteray-watch abstract-parametrics.go
package main

import (
	"math"

	. "github.com/barnex/bruteray/api"
)

func main() {
	Render(Spec{
		DebugNormals: false,
		Width:        1920 / 2,
		Height:       1080 / 2,
		Recursion:    5,
		NumPass:      1000,

		Lights: []Light{
			RectangleLight(C(1, 1, 0.8).EV(4), 2, 2, Vec{3, 5, 2}),
		},
		Objects: []Object{
			Rectangle(
				Matte(White.EV(-.4)),
				10, 10, V(0, 0, 0),
			),
			Backdrop(
				Flat(C(0.8, 0.8, 1.0).EV(-2)),
			),
			swirl(
				Shiny(White.EV(-.3), 0.01),
				0.15, 0.9, 4, 3).
				Translate(V(-0.8, 0, 0)),
			swirl(
				Shiny(C(0.1, 0.1, 0.15).EV(-4), 0.1),
				0.18, 0.5, 5, 1.5).
				Translate(V(0.6, 0, -0.9)),
			swirl(
				Shiny(C(1, 0.9, 0.8).EV(-.3), 0.03),
				0.15, 0.4, 1.2, 0.9).
				Translate(V(1.05, 0, 0.1)),
			swirl(
				Reflective(White.EV(-.6)),
				0.2, 0.3, 1.7, 1.2).
				Translate(V(-0.5, 0, -1.5)),
		},

		Camera: Projective(60*Deg, Vec{0.1, 1.3, 2.5}, 0, -15*Deg),
	})
}

func swirl(m Material, R, r, N, exp float64) Object {
	return Parametric(m, 96, 48, func(u, v float64) Vec {
		v *= 2 * Pi
		y := u
		R := R * math.Pow(1-u, 1/exp)
		r := r * R
		ph := u * 2 * Pi * N
		x := R*cos(v) + r*cos(ph)
		z := R*sin(v) + r*sin(ph)
		return Vec{x, y, z}
	})
}

var (
	sin  = math.Sin
	cos  = math.Cos
	sqrt = math.Sqrt
)
