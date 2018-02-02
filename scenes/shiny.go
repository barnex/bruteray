// +build ignore

package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	e.Add(
		Sheet(Ey, -1.0, Diffuse(WHITE.Mul(EV(-0.3)))),
		//Sheet(Ey, -1.0, Checkboard(1, Flat(WHITE), Flat(BLACK))),

		Sphere(Vec{0, 0.5, 4}, 1.5, Shiny(RED, EV(-2))),
		Sphere(Vec{-2, 0.1, 0.0}, 1.1, Shiny(BLUE.EV(-0.3), EV(-2))),
		Sphere(Vec{2, 0, 0.5}, 1, Shiny(GREEN.EV(-1), EV(-2))),
		Sphere(Vec{0, -0.2, -2.8}, 0.8, Shiny(WHITE, EV(-1.3))),
		Sphere(Vec{5, 1.0, 8}, 1, Shiny(YELLOW, EV(-1.3))),
		//Rect(Vec{9, 9, -6}, Ex, 1, 2, 2, Flat(WHITE.EV(2.6))),
	)

	e.AddLight(
		SphereLight(Vec{6, 5, -9}, 2, WHITE.Mul(EV(10.3))),
	)

	e.SetAmbient(Flat(WHITE.Mul(EV(-3))))

	e.Camera = Camera(1).Transl(0.2, 1.5, -5.5).RotScene(10 * Deg).Transf(RotX4(15 * Deg))
	e.Camera.AA = true
	e.Camera.Focus = 6
	e.Camera.Aperture = 0.3

	serve.Env(e)
}
