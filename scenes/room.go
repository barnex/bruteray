// +build ignore

package main

import . "github.com/barnex/bruteray"
import "github.com/barnex/bruteray/serve"

func main() {
	e := NewEnv()

	wall := Diffuse1(WHITE.EV(-.3))
	white := Shiny(WHITE.EV(-.6), EV(-3))
	black := Shiny(WHITE.EV(-6), EV(-3))

	e.Add(
		SurfaceAnd(
			Box(Vec{0, 0, 0}, 3, 3, 4, wall),
			Inverse(Or(
				Box(Vec{3, 0, 1.7}, 1, 2.2, .5, wall),
				Box(Vec{3, 1.3, 0}, 1, .7, .5, wall),
			)),
		),
		Sheet(Ey, 0, Checkboard(0.75, white, black)),
	)

	e.AddLight(
		SphereLight(Vec{6, 2, 1}, 1, WHITE.EV(10)),
	)

	e.SetAmbient(Flat(WHITE.EV(-4)))

	e.Camera = Camera(0.8).Transl(0, 1.7, -3).RotScene(-30 * Deg).Transf(RotX4(0 * Deg))
	e.Camera.AA = true
	e.Recursion = 5
	e.Cutoff = EV(5)

	serve.Env(e)
}
