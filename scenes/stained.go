package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	white := Diffuse0(WHITE.EV(-.6))
	//blue := Diffuse0(BLUE.EV(-.6))

	const (
		w     = 1.
		h     = 2.
		d     = 3.
		d2    = 2.
		brick = 0.05
	)
	main := chest(w, h, d, white)
	cross := Transf(chest(w, h, d2, white), RotY4(90*Deg).Mul(Transl4(Vec{1, 0, 0})))
	hull := Box(Vec{0, h / 2, 0}, d2/2+brick, h/2+brick, d/2+brick, white)

	e.Add(
		Sheet(Ey, 0, Checkboard(1, white, Diffuse0(RED))),
		Minus(hull, Or(main, cross)),
	)

	e.AddLight(
		PointLight(Vec{0, 0.5, 0}, WHITE.EV(1)),
		PointLight(Vec{1, 5, -3}, WHITE.EV(7)),
	)

	e.Camera = Camera(.7).Transl(0, 4, -3).RotScene(0 * Deg).Transf(RotX4(45 * Deg))
	//e.Camera.AA = true

	serve.Env(e)
}

func chest(w, h, d float64, m Material) Obj {
	return Or(Transf(CylZ(w/2, d, m), Transl4(Vec{0, h - w, 0})), Box(Vec{}, w/2, h/2, d/2-1e-6, m))
}
