// +build ignore

package main

import . "github.com/barnex/bruteray"

func main() {
	e := NewEnv()

	e.Add(
		Sheet(Ey, -1.0, Diffuse1(WHITE.Mul(EV(-0.3)))),

		Sphere(Vec{0, 0.5, 3}, 1.5, Shiny(RED, EV(-3))),
		Sphere(Vec{-2, 0.1, 0}, 1.1, Shiny(BLUE.EV(-0.3), EV(-3))),
		Sphere(Vec{2, 0, -1}, 1, Shiny(GREEN.EV(-1), EV(-3))),
		Sphere(Vec{0, -0.2, -2}, 0.8, Shiny(WHITE, EV(-2))),
		Rect(Vec{9, 9, -6}, Ex, 1, 2, 2, Flat(WHITE.EV(2.6))),
	)

	e.AddLight(
		SmoothLight(Vec{7, 7, -5}, 1, WHITE.Mul(EV(6.3))),
	)

	e.SetAmbient(Flat(WHITE.Mul(EV(-5))))

	e.Camera = Camera(1).Transl(0, 2.5, -6).Transf(RotX4(22 * Deg))
	e.Camera.AA = true

	Serve(e)
}
