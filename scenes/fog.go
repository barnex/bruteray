package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	green := Diffuse(GREEN.EV(-1))

	e.Add(
		Sheet(Ey, 0, green),
	)

	e.Add(cross(-3, -3))
	e.Add(cross(0, -3))
	e.Add(cross(4, -3))

	e.Add(cross(0, 0))
	e.Add(cross(-3, 0))
	e.Add(cross(3, 0))
	//e.Add(cross(6, 0))

	e.Add(cross(0, 3))
	e.Add(cross(-3, 3))
	e.Add(cross(3, 3))
	e.Add(cross(6, 3))

	e.Add(cross(0, 6))
	e.Add(cross(-3, 6))
	e.Add(cross(3, 6))
	e.Add(cross(6, 6))

	e.AddLight(
		SphereLight(Vec{4.8, 1.2, 6}.Mul(1.2), .2, WHITE.EV(10.3)),
		//SphereLight(Vec{8, 1, 0}.Mul(1.2), .2, WHITE.EV(5.0)),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-8))))

	e.Camera = Camera(.9).Transl(0, 3.0, -6).RotScene(10 * Deg).Transf(RotX4(20 * Deg))

	e.Camera.AA = true
	e.Recursion = 2
	e.Cutoff = EV(10)

	e.Fog = 5

	serve.Env(e)
}

func cross(x, z float64) Obj {
	white := Diffuse(WHITE.EV(-.3))

	const d = 0.2
	return Or(
		Box(Vec{x, 0.5, z}, 0.4, 1.3, d, white),
		Box(Vec{x, 1, z}, 1, 0.2, 0.2001, white),
	)
}
