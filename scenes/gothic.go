package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	//white := Diffuse0(WHITE.EV(-.6))
	white := ShadeShape(WHITE)

	const (
		w  = 3.0 // central width
		h1 = 2.0 // pillar height
		d  = 6.0 // central depth
		//d2    = 2.0
		//brick = 0.05
	)

	xx

	e.Add(
		Sheet(Ey, 0, Checkboard(1, white, Flat(BLACK))),
		chestz(Vec{}, w, h1, d, white),
	)

	e.AddLight(
	//PointLight(Vec{0, 0.5, 0}, WHITE.EV(1)),
	//PointLight(Vec{1, 5, -3}, WHITE.EV(7)),
	)

	e.Camera = Camera(1).Transl(0, 4, -7).RotScene(0 * Deg).Transf(RotX4(30 * Deg))
	//e.Camera.AA = true

	serve.Env(e)
}

func chestz(pos Vec, w, h, d float64, m Material) Obj {
	return NBox(pos.Add(Vec{Y: h / 2}), w, h, d, m)
}
