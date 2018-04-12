package main

import (
	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/light"
	. "github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
	"github.com/barnex/bruteray/serve"
	. "github.com/barnex/bruteray/shape"
)

func main() {
	e := NewEnv()

	const R = .095
	const D = .1
	w := Diffuse(WHITE.EV(-1.3))
	b := Diffuse(Color{.3, .3, 1}.EV(-.3))
	r := Diffuse(RED.EV(-.9))
	eye := Shiny(WHITE.EV(-.3), EV(-4))

	e.Add(
		NewSheet(Ey, -R/2, Diffuse(WHITE.EV(-0.3))),
		NewBox(R, 1*R, R, w).Transl(Vec{0 * D, 0, 0 * D}),
		NewBox(R, 1*R, R, w).Transl(Vec{1 * D, 0, 0 * D}),
		NewBox(R, 1*R, R, w).Transl(Vec{2 * D, 0, 0 * D}),
		NewBox(R, 1*R, R, w).Transl(Vec{2 * D, 0, 1 * D}),
		NewBox(R, 1*R, R, w).Transl(Vec{2 * D, 0, 2 * D}),
		NewBox(R, 1*R, R, w).Transl(Vec{1 * D, 0, 2 * D}),
		NewBox(R, 1*R, R, w).Transl(Vec{0 * D, 0, 2 * D}),
		NewBox(R, 1*R, R, w).Transl(Vec{0 * D, 0, 1 * D}),
		Minus(NewBox(R, 1*R, R, r).Transl(Vec{-1 * D, 0, 1 * D}), NewSphere(R/2, r).Transl(Vec{-1 * D, D / 2, 1 * D})),
		NewBox(R, 1*R, R, w).Transl(Vec{-2 * D, 0, 1 * D}),
		NewBox(R, 1*R, R, w).Transl(Vec{-2 * D, 0, 0 * D}),
		NewBox(R, 1*R, R, w).Transl(Vec{-2 * D, 0, -2 * D}),

		//NewSphere(R, b).Transl(Vec{-2 * D, 0, 2 * D}),

		NewSphere(R, b).Transl(Vec{2 * D, 0, -1 * D}),
		NewSphere(R/2, eye).Transl(Vec{2 * D, 0, -1 * D}).Transl(Vec{R / 4, R / 2.3, R / 3.3}),
		NewSphere(R/2, eye).Transl(Vec{2 * D, 0, -1 * D}).Transl(Vec{R / 4, R / 2.3, -R / 3.3}),

		NewSphere(R/2, r).Transl(Vec{4 * D, -R / 2, -1 * D}),
	)

	e.AddLight(light.Sphere(Vec{15, 25, -20}, 13, WHITE.EV(13.3)))
	e.SetAmbient(WHITE.EV(-1))
	cam := raster.Camera(1).Transl(0, 10, 0).Transf(RotX4(90 * Deg))
	cam.AA = true
	cam.FocalLen = 0
	e.Recursion = 2
	serve.Env(cam, e)
}
