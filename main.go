// +build ignore

package main

import . "."

func main() {
	e := NewEnv()

	e.Add(
		Sheet(Ey, -1, Diffuse1(WHITE.Mul(EV(-1)))),
		Sphere(Vec{}, 1, Diffuse1(WHITE.Mul(EV(-0.3)))),
		Sphere(Vec{2, 0, -1}, 1, Diffuse1(WHITE.Mul(EV(-0.3)))),
		Sphere(Vec{-2, 0, -1}, 1, Diffuse1(WHITE.Mul(EV(-0.3)))),
		Sphere(Vec{0, 0, -2}, 1, Diffuse1(WHITE.Mul(EV(-0.3)))),
	)

	e.AddLight(
		PointLight(Vec{8, 6, -5}, WHITE.Mul(EV(7))),
	)

	e.Camera = Camera(1).Transl(0, 3, -5).Transf(RotX4(40 * Deg))

	Serve(e)
}
