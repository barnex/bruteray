package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	//white := Diffuse0(WHITE.EV(-.6))
	white := ShadeShape()

	const (
		w     = 1.
		h     = 2.
		d     = 3.
		d2    = 2.
		brick = 0.05
	)

	e.Add(
		Sheet(Ey, 0, Checkboard(1, white, Diffuse0(RED))),
		chest(w, h, d, white),
	)

	e.AddLight(
	//PointLight(Vec{0, 0.5, 0}, WHITE.EV(1)),
	//PointLight(Vec{1, 5, -3}, WHITE.EV(7)),
	)

	e.Camera = Camera(1).Transl(0, .7, -1).RotScene(0 * Deg).Transf(RotX4(0 * Deg))
	//e.Camera.AA = true

	serve.Env(e)
}

func chest(w, h, d float64, m Material) Obj {
	return Or(Transf(CylZ(w/2, d, m), Transl4(Vec{0, h - w, 0})), Box(Vec{}, w/2, h/2, d/2-1e-6, m))
}
