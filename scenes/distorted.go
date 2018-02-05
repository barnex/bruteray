package main

import "github.com/barnex/bruteray/serve"
import . "github.com/barnex/bruteray"

func main() {

	e := NewEnv()

	//m := Shiny(RED, EV(-2))
	m := Blend(0.9, Refractive(1, 20), 0.1, Flat(BLACK))
	//m := Shiny(WHITE.EV(-.3), EV(-1))
	m = Distort(1, 30, Vec{200, 200, 200}, 0.01, m)

	e.Add(
		//Sheet(Ez, 10, Checkboard(1, Diffuse(WHITE.EV(-4)), Diffuse(WHITE.EV(-1)))),
		Sheet(Ey, 0, Checkboard(1, Diffuse(WHITE.EV(-4)), Diffuse(WHITE.EV(-1)))),
		Sphere(Vec{0.1, 1, 0.2}, 1, m),
	)

	e.AddLight(
		SphereLight(Vec{4, 6, -8}, 2, WHITE.EV(11)),
	)

	e.SetAmbient(Flat(WHITE.EV(-5)))

	e.Camera = Camera(1).Transl(0, 2, -3).Transf(RotX4(15 * Deg))
	e.Camera.AA = true
	e.Camera.Focus = 3.5
	e.Camera.Aperture = 0.07

	serve.Env(e)
}
