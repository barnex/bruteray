// +build ignore

package main

import . "github.com/barnex/bruteray"
import "github.com/barnex/bruteray/serve"

func main() {

	e := NewEnv()

	e.Add(
		Sheet(Ey, -1.0, Diffuse1(WHITE.Mul(EV(-0.3)))),
		SurfaceAnd(
			Sphere(Vec{}, 1, Shiny(WHITE.EV(-8), EV(-1))),
			Slab(Ey, -0.3, 0.3, Flat(RED)),
		),
	)
	e.AddLight(
		SphereLight(Vec{3, 3, 1}, .5, WHITE.Mul(EV(8))),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-5))))

	e.Camera = Camera(1).Transl(0, 2.5, -5).Transf(RotX4(22 * Deg))
	e.Camera.AA = true

	e.Cutoff = EV(5)

	serve.Env(e)
}
