// +build ignore

package main

import . "github.com/barnex/bruteray"

func main() {
	e := NewEnv()
	e.Add(
		Sheet(Ey, -1.0, Diffuse1(WHITE.Mul(EV(-0.1)))),
	)
	e.AddLight(
		SphereLight(Vec{0, -0.4, 0}, 0.6, WHITE.Mul(EV(5.3))),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-2))))

	e.Camera = Camera(1).Transl(0, 2, -5).Transf(RotX4(27 * Deg))
	e.Camera.AA = true

	Serve(e)
}
