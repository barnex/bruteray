package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	green := Diffuse0(Color{1, 1, 0.8}.EV(-1.6))

	e.Add(
		Sheet(Ey, 0, green),
	)

	e.Add(cross(-4, -2.7))
	e.Add(cross(0, -3))
	e.Add(cross(4, -3))

	e.Add(cross(0, 0))
	e.Add(cross(-6, 0))
	e.Add(cross(-3, 0))
	e.Add(cross(3, 0))
	//e.Add(cross(6, 0))

	//e.Add(cross(0, 3))
	e.Add(cross(-3, 3))
	//e.Add(cross(3, 3))
	e.Add(cross(7, 3))

	e.Add(cross(0, 6))
	//e.Add(cross(-3, 6))
	e.Add(cross(3, 6))
	e.Add(cross(6, 6))

	e.AddInvisibleLight(
		SphereLight(Vec{5.03, 1.7, 6}.Mul(1.2), .2, WHITE.EV(11.0)),
		//SphereLight(Vec{8, 1, 0}.Mul(1.2), .2, WHITE.EV(5.0)),
	)
	e.SetAmbient(Flat(BLACK))

	e.Camera = Camera(.9).Transl(0, 2.2, -6).RotScene(10 * Deg).Transf(RotX4(20 * Deg))

	e.Camera.AA = true
	e.Recursion = 3
	e.Cutoff = EV(10)

	e.Fog = 2.1

	serve.Env(e)
}

func cross(x, z float64) Obj {
	white := Diffuse0(WHITE.EV(-.1))

	const d = 0.2
	return Or(
		Box(Vec{x, 0.5, z}, 0.3, 1.5, d, white),
		Box(Vec{x, 1.2, z}, 0.84, 0.2, 0.2001, white),
	)
}
