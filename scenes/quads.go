package main

import . "github.com/barnex/bruteray"

func main() {
	e := NewEnv()

	w := Diffuse1(WHITE, EV(-.3))

	e.Add(Sheet(Ey, -1, w))

	HAnd(
		Quad(Vec{}, Vec{1, -1, 1}, 1, w),
		Slab(Ey, 1, -1, w),
	)

	e.AddLight(
		SphereLight(Vec{18, 17, -8}.Mul(2), 16, WHITE.EV(15.3)),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-6))))

	e.Camera = Camera(1).Transl(0, 1, -3).RotScene(0 * Deg).Transf(RotX4(20 * Deg))
	e.Camera.AA = true

	Serve(e)
}
