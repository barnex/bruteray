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
		Or(Box(Vec{0, 0, 0}, 3, 3, 3, wall), Box(Vec{2.8, 0, 2}, 1, 3, 1.2, wall)),
		Sheet(Ey, 0, Checkboard(0.75, white, black)),
	)

	e.AddLight(
		SphereLight(Vec{1, 2.1, 1}, .07, WHITE.EV(5)),
	)

	e.Camera = Camera(0.9).Transl(1, 1.7, -0).RotScene(0 * Deg).Transf(RotX4(0 * Deg))
	e.Camera.AA = true
	e.Recursion = 5
	e.Cutoff = EV(5)

}
