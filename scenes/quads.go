package main

import . "github.com/barnex/bruteray"
import "github.com/barnex/bruteray/serve"

func main() {
	e := NewEnv()

	w := Diffuse1(WHITE.EV(-.6))

	e.Add(Sheet(Ey, -1, w))

	e.Add(SurfaceAnd(
		Quad(Vec{}, Vec{1, 0, 1}, .75, w),
		Slab(Ey, -1, 1, w),
	))
	e.Add(SurfaceAnd(
		Quad(Vec{}, Vec{1, 0, 1}, 1.5, w),
		Slab(Ey, -1, 1, w),
	))

	e.AddLight(
		SphereLight(Vec{18, 17, -8}.Mul(2), 16, WHITE.EV(15.3)),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-2))))

	e.Camera = Camera(1).Transl(0, 3, -2.2).RotScene(0 * Deg).Transf(RotX4(50 * Deg))
	e.Camera.AA = true

	serve.Env(e)
}
