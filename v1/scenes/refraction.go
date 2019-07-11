// +build ignore

package main

import . "github.com/barnex/bruteray"
import "github.com/barnex/bruteray/v1/serve"

func main() {
	e := NewEnv()

	white := Diffuse(WHITE.EV(-.6))
	green := Diffuse(GREEN.EV(-1.6))
	red := Diffuse(RED.EV(-.6))

	const (
		w = 550
		h = 550
		U = 66666
	)

	e.Add(
		Rect(Vec{0, 0, 0}, Ey, w/2, U, w/2, white),
		Rect(Vec{0, h, 0}, Ey, w/2, U, w/2, white),
		Rect(Vec{0, h / 2, w / 2}, Ez, w/2, h/2, U, white),
		Rect(Vec{w / 2, h / 2, 0}, Ex, U, h/2, w/2, green),
		Rect(Vec{-w / 2, h / 2, 0}, Ex, U, h/2, w/2, red),
		Sphere(Vec{-120, 100, 100}, 100, Reflective(WHITE.EV(-.3))),
		Sphere(Vec{125, 100, -20}, 100, Blend(0.9, Refractive(1, 1.5), 0.05, Diffuse(WHITE))), // TODO: blend phong
	)

	e.AddLight(
		//PointLight(Vec{0, h - 100, 0}, WHITE.EV(10)),
		//SphereLight(Vec{0, h - 100, 0}, 50, WHITE.EV(30)),
		//RectLight(Vec{rWidth - off, ceil / 2, .6}, 0, 1, .45, WHITE.EV(2)),:w
		RectLight(Vec{0, h - 1e-4, 0}, 200/2, 0, 200/2, Color{1.0, 1.0, 1.0}.EV(17.6)),
	)

	e.SetAmbient(Flat(WHITE.EV(-10)))

	focalLen := 0.035 / 0.025
	e.Camera = Camera(focalLen).Transl(0, h/2, -1050)
	e.Camera.AA = true
	e.Recursion = 5
	e.Cutoff = EV(6)

	serve.Env(e)
}

const off = 1e-3
