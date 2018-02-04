// +build ignore

package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	//green := Checkboard(.5, Diffuse0(Color{1, 1, 0.8}.EV(-1.6)), Diffuse0(WHITE.EV(-3)))
	//green := Diffuse0(Color{1, 1, 0.8}.EV(-1.6))

	a := Diffuse0(Color{1, 1, 0.8}.EV(-1.6))
	b := Shiny(Color{1, 1, 0.8}.EV(-2.6), EV(-2.3))
	f := func(v float64) Material {
		if v > 0 {
			return b
		} else {
			return a
		}
	}
	green := Waves(9, Vec{20, 20, 20}, f)

	e.Add(
		Sheet(Ey, 0, green),
	)

	e.Add(cross(-4, -2.7))
	e.Add(cross(0, -3))
	e.Add(cross(4, -3))

	e.Add(cross(0, 0))
	e.Add(cross(-6, 0))
	e.Add(cross(-3, 0))
	e.Add(cross(2.8, 0))
	//e.Add(cross(6, 0))

	//e.Add(cross(0, 3))
	e.Add(cross(-3, 3))
	//e.Add(cross(3, 3))
	e.Add(cross(7, 3))

	e.Add(cross(0, 6))
	//e.Add(cross(-3, 6))
	e.Add(cross(3, 6))
	e.Add(cross(7, 6))

	e.AddInvisibleLight(
		//SphereLight(Vec{5.03, 1.7, 6}.Mul(1.2), .2, WHITE.EV(10.6)),
		SphereLight(Vec{5.03, 1.7, 6}.Mul(1.2), .0, WHITE.EV(10.6)),
	)
	e.SetAmbient(Flat(BLACK))

	e.Camera = Camera(.9).Transl(-0, 2.2, -6.5).RotScene(10 * Deg).Transf(RotX4(20 * Deg))

	e.Camera.AA = true
	e.Recursion = 2
	e.Cutoff = EV(10)

	e.Camera.Focus = 3.8
	e.Camera.Aperture = 0.05
	e.Fog = 2.1
	serve.Env(e)
}

func cross(x, z float64) Obj {
	//white := Diffuse0(WHITE.EV(.6))

	a := Diffuse(WHITE.EV(0.6))
	b := Diffuse(WHITE.EV(-.03))
	f := func(v float64) Material {
		//v = v * v
		if v > 0.2 && v < 0.6 {
			return b
		} else {
			return a
		}
	}
	m1 := Waves(10, Vec{50, 50, 50}, f)

	const d = 0.2
	return Or(
		Box(Vec{x, 0.5, z}, 0.27, 1.5, d, m1),
		Box(Vec{x, 1.2, z}, 0.84, 0.27, 0.2001, m1),
	)
}
