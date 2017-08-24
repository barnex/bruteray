package main

import . "github.com/barnex/bruteray"

func main() {
	e := NewEnv()
	e.Add(
		Sheet(Ey, -1.0, Diffuse1(WHITE.Mul(EV(-0.1)))),
		Sphere(Vec{0, 0.5, 2.8}, 1.5, Shiny(RED, EV(-2))),
		Sphere(Vec{-2, 0.1, -0.6}, 1.1, Shiny(BLUE.EV(-0.3), EV(-2))),
		Sphere(Vec{2, 0, -1}, 1, Shiny(GREEN.EV(-1), EV(-2))),
		Sphere(Vec{0, -0.2, -2}, 0.8, Shiny(WHITE, EV(-1))),
		//Sphere(Vec{7, 7, -5}, 1, Flat(WHITE.EV(-2))),
	)
	e.AddLight(
		SphereLight(Vec{0, -0.4, 0}, 0.6, WHITE.Mul(EV(5.3))),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-6))))

	e.Camera = Camera(1).Transl(0, 2, -5).Transf(RotX4(27 * Deg))
	e.Camera.AA = true

	Serve(e)
}
