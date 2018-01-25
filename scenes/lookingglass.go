// +build ignore

package main

import . "github.com/barnex/bruteray"
import "github.com/barnex/bruteray/serve"

func main() {
	e := NewEnv()

	const U = 666

	glass := Refractive(1, 1.38)
	//glass := Diffuse(WHITE)
	plastic := Shiny(WHITE.EV(-7), EV(-6))

	const r = 75
	lens := And(
		Sphere(Vec{0, +0.8 * r, 0}, r, glass),
		Sphere(Vec{0, -0.8 * r, 0}, r, glass),
	)
	handle := Or(Minus(
		CapCyl(Vec{}, 45, 8, plastic),
		CapCyl(Vec{}, 40, 20, plastic),
	),
		And(
			Quad(Vec{}, Vec{1, 1, 0}, 25, plastic),
			Slab(Ez, -120, -45, plastic),
		),
	)
	mag := Transf(Or(lens, handle), RotX4(-20*Deg).Mul(RotY4(35*Deg)).Mul(Transl4(Vec{5, 40, -30})))

	e.Add(
		Rect(Vec{0, 0, 0}, Ey, 100, U, 100, Checkboard(10, Diffuse(WHITE.EV(-.3)), Diffuse(WHITE.EV(-5)))),
		Sheet(Ey, -10, Diffuse(WHITE.EV(-.6))),
		mag,
	)

	e.AddLight(
		RectLight(Vec{850, 300, -300}, 180, 800, 0, WHITE.EV(20.6)),
		RectLight(Vec{400, 300, -300}, 180, 800, 0, WHITE.EV(20.6)),
	)

	e.SetAmbient(Flat(WHITE.EV(-3)))

	focalLen := 1.0
	e.Camera = Camera(focalLen).Transl(0, 150, -150).Transf(RotX4(50 * Deg))
	e.Camera.AA = true
	e.Recursion = 5
	e.Cutoff = EV(40)

	serve.Env(e)
}
