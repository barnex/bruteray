// +build ignore

package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	e.Add(
		Rect(Vec{X: 3, Y: -1}, Ey, 12, 10, 30, Checkboard(1, Diffuse(WHITE.EV(-.4)), Diffuse(WHITE.Mul(EV(-0.3))))),
		Sheet(Ey, -1.5, Diffuse(WHITE.Mul(EV(-.6)))),
		//Sheet(Ey, -1.0, Checkboard(1, Flat(WHITE), Flat(BLACK))),

		Sphere(Vec{0, 0.5, 4}, 1.5, Shiny(RED, EV(-2))),
		Sphere(Vec{-2, 0.1, 0.0}, 1.1, Shiny(BLUE.EV(-0.3), EV(-2))),
		Sphere(Vec{2, 0, 0.5}, 1, Shiny(GREEN.EV(-1), EV(-2))),
		Sphere(Vec{0, -0.2, -2.8}, 0.8, Shiny(WHITE.EV(-.3), EV(-1.3))),
		Sphere(Vec{7, 0.0, 6}, 1, Shiny(WHITE.EV(-5), EV(-1.3))),
		//Rect(Vec{9, 9, -6}, Ex, 1, 2, 2, Flat(WHITE.EV(2.6))),
		Sphere(Vec{-37, 9, 35}, .5, Flat(YELLOW.EV(2))),
	)

	e.AddLight(
		SphereLight(Vec{6, 5, -9}, 1.5, WHITE.Mul(EV(10.3))),
	)

	e.SetAmbient(Flat(WHITE.Mul(EV(-4))))

	e.Camera = Camera(1).Transl(0.2, 1.5, -5.5).RotScene(10 * Deg).Transf(RotX4(15 * Deg))
	e.Camera.AA = true
	e.Camera.Focus = 6
	e.Camera.Aperture = 0.28
	e.Camera.Diaphragm = DiaHex
	e.Cutoff = EV(5)

	serve.Env(e)
}
