// +build ignore

package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	a := Diffuse(WHITE.EV(-.6))
	b := Shiny(WHITE.EV(-8), EV(-2))
	f := func(v float64) Material {
		//v = v * v
		if v > 0.3 && v < 0.5 {
			return b
		} else {
			return a
		}
	}
	m1 := Waves(10, Vec{30, 90, 30}, f)

	a2 := Diffuse(WHITE.EV(-2.3))
	b2 := Diffuse(WHITE.EV(-3))
	c2 := Diffuse(WHITE.EV(-3.3))
	f2 := func(v float64) Material {
		v = v * v
		if v > 0.1 && v < 0.9 {
			return a2
		} else if v < -0.5 {
			return b2
		} else {
			return c2
		}
	}
	m2 := Waves(3, Vec{50, 0, 1}, f2)

	e.Add(
		Sphere(Vec{0, 1, 0}, 1, m1),
		Sheet(Ey, 0, m2),
	)

	e.AddLight(
		SphereLight(Vec{7, 5, -6}, 2.5, WHITE.Mul(EV(10.6))),
	)

	e.SetAmbient(Flat(WHITE.Mul(EV(-5))))

	e.Camera = Camera(1).Transl(0, 2, -3).Transf(RotX4(20 * Deg))
	e.Camera.AA = true
	e.Camera.Aperture = 0.07
	e.Camera.Focus = 2.35

	serve.Env(e)
}
