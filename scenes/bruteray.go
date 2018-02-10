// +build ignore

package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	const w = 0.2
	const h = 0.1

	m := Diffuse(WHITE.EV(-.6))

	//B := Box(Vec{}, w, 1, h, m)
	B := Minus(
		CapCyl(Vec{}, 0.5, h, m),
		CapCyl(Vec{}, 0.2, h, m),
	)

	e.Add(
		B,
		//Sheet(Ey, 0, m),
	)

	e.AddLight(PointLight(Vec{0, 5, -3}, WHITE.EV(10)))

	e.Camera = Camera(1).Transl(0, 4, 0).Transf(RotX4(90 * Deg))

	serve.Env(e)

}
