// +build ignore

package main

import . "github.com/barnex/bruteray"

func main() {
	e := NewEnv()

	color := Shiny(RED, EV(-4))

	base := ObjAnd(
		Cube(Vec{}, 1, color),
		Sphere(Vec{}, Sqrt(2)-0.02, color),
	)

	p := 0.15
	d := 0.05

	//black := Reflective(WHITE.EV(-5))
	black := Diffuse0(WHITE.EV(-4))
	dice := base

	// 3
	dice = ObjMinus(dice, Sphere(Vec{0, 0, -(1 + d)}, p, black))
	dice = ObjMinus(dice, Sphere(Vec{-.4, -.4, -(1 + d)}, p, black))
	dice = ObjMinus(dice, Sphere(Vec{.4, .4, -(1 + d)}, p, black))

	// 4
	dice = ObjMinus(dice, Sphere(Vec{-0.35, -0.35, (1 + d)}, p, black))
	dice = ObjMinus(dice, Sphere(Vec{+0.35, +0.35, (1 + d)}, p, black))
	dice = ObjMinus(dice, Sphere(Vec{+0.35, -0.35, (1 + d)}, p, black))
	dice = ObjMinus(dice, Sphere(Vec{-0.35, +0.35, (1 + d)}, p, black))

	// 1
	dice = ObjMinus(dice, Sphere(Vec{-(1 + d), 0, 0}, p, black))

	// 5
	dice = ObjMinus(dice, Sphere(Vec{+0.0, (1 + d), +0.0}, p, black))
	dice = ObjMinus(dice, Sphere(Vec{-0.4, (1 + d), -0.4}, p, black))
	dice = ObjMinus(dice, Sphere(Vec{+0.4, (1 + d), +0.4}, p, black))
	dice = ObjMinus(dice, Sphere(Vec{+0.4, (1 + d), -0.4}, p, black))
	dice = ObjMinus(dice, Sphere(Vec{-0.4, (1 + d), +0.4}, p, black))

	// 2
	dice = ObjMinus(dice, Sphere(Vec{+0.3, -(1 + d), -0.3}, p, black))
	dice = ObjMinus(dice, Sphere(Vec{-0.3, -(1 + d), +0.3}, p, black))

	e.Add(
		Sheet(Ey, -1, Diffuse1(WHITE.EV(-.3))),
		Transf(dice, Transl4(Vec{0, 0.5, 0}).Mul(RotX4(-30*Deg)).Mul(RotY4(40*Deg))),
		Transf(dice, Transl4(Vec{2, 4, -0.2}).Mul(RotX4(+20*Deg)).Mul(RotY4(30*Deg))),
		Transf(dice, Transl4(Vec{8, 7, -2.0}).Mul(RotX4(+200*Deg)).Mul(RotY4(-130*Deg))),
	)

	e.AddLight(
		SphereLight(Vec{10, 25, -4}, 10, WHITE.Mul(EV(9))),
	)

	e.SetAmbient(Flat(WHITE.Mul(EV(-4))))

	e.Camera = Camera(1).Transl(5, 5, -15).Transf(RotX4(5 * Deg))
	e.Camera.AA = true

	Serve(e)
}
