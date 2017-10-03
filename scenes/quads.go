package main

import . "github.com/barnex/bruteray"

func main() {
	e := NewEnv()

	//w := Diffuse1(Color{1, 1, 1}.EV(-.6))
	//w := Flat(RED)
	w := Shiny(WHITE, EV(-4))

	e.Add(
		Sheet(Ey, -1, Diffuse1(WHITE.Mul(EV(-.6)))),
		//Sphere(Vec{}, 1, w),
		HAnd(
			Quad(Vec{}, Vec{1, -1, 1}, 1, w),
			Slab(Ey, 1, -1, w),
		),
	)

	e.AddLight(
		SphereLight(Vec{18, 17, -8}.Mul(2), 16, WHITE.EV(15.3)),
	//SphereLight(Vec{-6, 6, -25}.Mul(10), 160, WHITE.EV(18)),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-6))))

	e.Camera = Camera(1).Transl(0, 1, -3).RotScene(0 * Deg).Transf(RotX4(20 * Deg))
	e.Camera.AA = true

	Serve(e)
}
