// +build ignore

package main

import (
	"math"

	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	dice := dice()

	a := Diffuse(WHITE.EV(-1))
	b := Diffuse(WHITE.EV(-.3))
	f := func(v float64) Material {
		if v > 0 {
			return b
		} else {
			return a
		}
	}
	m1 := Waves(10, Vec{1, 9, 60}, f)

	e.Add(
		Sheet(Ey, -1, m1),
		//CapCyl(Vec{0, -1, 190}, 15, 0.001, m1),
		Transf(dice, Transl4(Vec{0, 0.5, 0}).Mul(RotX4(-30*Deg)).Mul(RotY4(40*Deg))),
		Transf(dice, Transl4(Vec{2, 4, -0.2}).Mul(RotX4(+20*Deg)).Mul(RotY4(30*Deg))),
		Transf(dice, Transl4(Vec{8, 7, -2.0}).Mul(RotX4(+200*Deg)).Mul(RotY4(-130*Deg))),
	)

	e.AddLight(
		SphereLight(Vec{10, 25, -4}, 12, WHITE.Mul(EV(12))),
	)

	//e.SetAmbient(Flat(WHITE.Mul(EV(-4))))
	e.SetAmbient(ShadeDir(func(v Vec) Color {
		if v[Z] < 0 && v[X] < 0 {
			return WHITE.EV(-1)
		} else {
			return BLACK
		}
	}))

	e.Camera = Camera(1).Transl(4.5, 5, -12).Transf(RotX4(5 * Deg))
	e.Camera.AA = true
	e.Camera.Focus = 12.5
	e.Camera.Aperture = 0.2

	serve.Env(e)
}

func dice() Obj {
	color := Shiny(WHITE, EV(-4))

	base := And(
		Cube(Vec{}, 1, color),
		Sphere(Vec{}, math.Sqrt(2)-0.02, color),
	)

	p := 0.15
	d := 0.05

	//black := Reflective(WHITE.EV(-5))
	black := Reflective(WHITE.EV(-3))
	dice := base

	// 3
	dice = Minus(dice, Sphere(Vec{0, 0, -(1 + d)}, p, black))
	dice = Minus(dice, Sphere(Vec{-.4, -.4, -(1 + d)}, p, black))
	dice = Minus(dice, Sphere(Vec{.4, .4, -(1 + d)}, p, black))

	// 4
	dice = Minus(dice, Sphere(Vec{-0.35, -0.35, (1 + d)}, p, black))
	dice = Minus(dice, Sphere(Vec{+0.35, +0.35, (1 + d)}, p, black))
	dice = Minus(dice, Sphere(Vec{+0.35, -0.35, (1 + d)}, p, black))
	dice = Minus(dice, Sphere(Vec{-0.35, +0.35, (1 + d)}, p, black))

	// 1
	dice = Minus(dice, Sphere(Vec{-(1 + d), 0, 0}, p, black))

	// 5
	dice = Minus(dice, Sphere(Vec{+0.0, (1 + d), +0.0}, p, black))
	dice = Minus(dice, Sphere(Vec{-0.4, (1 + d), -0.4}, p, black))
	dice = Minus(dice, Sphere(Vec{+0.4, (1 + d), +0.4}, p, black))
	dice = Minus(dice, Sphere(Vec{+0.4, (1 + d), -0.4}, p, black))
	dice = Minus(dice, Sphere(Vec{-0.4, (1 + d), +0.4}, p, black))

	// 2
	dice = Minus(dice, Sphere(Vec{+0.3, -(1 + d), -0.3}, p, black))
	dice = Minus(dice, Sphere(Vec{-0.3, -(1 + d), +0.3}, p, black))

	return dice
}
