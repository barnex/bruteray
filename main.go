// +build ignore

package main

import . "github.com/barnex/bruteray"

func main() {
	e := NewEnv()

	white := Blend(EV(-0.3), Diffuse1(WHITE), EV(-3), Reflective(WHITE))
	black := Blend(EV(-7), Diffuse1(WHITE), EV(-5), Reflective(WHITE))
	green := Blend(EV(-1), Diffuse1(GREEN), EV(-7), Reflective(WHITE))
	pink := Blend(EV(-1), Diffuse1(MAGENTA.MAdd(EV(-0.3), WHITE)), EV(-4), Reflective(WHITE))
	red := Blend(EV(-1), Diffuse1(RED), EV(-5), Reflective(WHITE))
	yellow := Blend(EV(-0.3), Diffuse1(YELLOW), EV(-4), Reflective(WHITE))
	blue := Blend(EV(-1), Diffuse1(BLUE.MAdd(EV(-2), WHITE)), EV(-7), Reflective(WHITE))
	purple := Blend(EV(-1), Diffuse1(BLUE.MAdd(EV(-0.3), RED)), EV(-4), Reflective(WHITE))

	e.Add(
		Sheet(Ey, -1.0, Diffuse1(WHITE.Mul(EV(-1)))),

		Sphere(Vec{0, 0.2, 0.9}, 1, black),
		Sphere(Vec{2, 0, -1}, 1, green),
		Sphere(Vec{4, 0, -1}, 1, pink),
		Sphere(Vec{2, 0, 4}, 1, red),
		Sphere(Vec{-2, 0, -1}, 1, blue),
		Sphere(Vec{-4, 0, -1}, 1, purple),
		Sphere(Vec{-2, 0, 4}, 1, yellow),
		Sphere(Vec{0, -0.3, -2}, 0.7, white),
	)

	e.AddLight(
		SmoothLight(Vec{7, 7, -5}, 1, WHITE.Mul(EV(7))),
	)

	e.SetAmbient(Flat(WHITE.Mul(EV(-3))))

	e.Camera = Camera(1).Transl(0, 3, -7).Transf(RotX4(20 * Deg))
	e.Camera.AA = true

	Serve(e)
}
