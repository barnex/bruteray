// +build ignore

package main

import . "github.com/barnex/bruteray"

func main() {
	e := NewEnv()

	e.Add(
		Sheet(Ey, -1, Diffuse1(WHITE.Mul(EV(-1)))),
		Sphere(Vec{}, 1, Diffuse1(WHITE.Mul(EV(-0.3)))),
		Sphere(Vec{2, 0, -1}, 1, Diffuse1(Color{R: 1, G: .3, B: 1}.Mul(EV(-0.3)))),
		Sphere(Vec{-2, 0, -1}, 1, Diffuse1(YELLOW.Mul(EV(-0.3)))),
		Sphere(Vec{0, 0, -2}, 1, Diffuse1(Color{G: 1, B: 1}.Mul(EV(-0.3)))),
	)

	e.AddLight(
		SmoothLight(Vec{8, 6, -5}, 1, WHITE.Mul(EV(7))),
	)

	e.Camera = Camera(1).Transl(0, 3, -5).Transf(RotX4(40 * Deg))

	Serve(e)
}
