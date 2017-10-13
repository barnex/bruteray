package main

import . "github.com/barnex/bruteray"
import "github.com/barnex/bruteray/serve"

func main() {
	e := NewEnv()

	white := Diffuse(WHITE.EV(-.3))
	green := Diffuse(GREEN.EV(-1.6))
	red := Diffuse(RED.EV(-.6))

	const (
		w = 540
		h = 560
		U = 66666
	)

	e.Add(
		Rect(Vec{0, 0, 0}, Ey, w/2, U, w/2, white),
		Rect(Vec{0, h, 0}, Ey, w/2, U, w/2, white),
		Rect(Vec{0, h / 2, w / 2}, Ez, w/2, h/2, U, white),
		Rect(Vec{w / 2, h / 2, 0}, Ex, U, h/2, w/2, green),
		Rect(Vec{-w / 2, h / 2, 0}, Ex, U, h/2, w/2, red),
		Transf(Box(Vec{120, 80, -80}, 80, 80, 80, white), RotY4(-18*Deg)),
		Transf(Box(Vec{-50, 180, 100}, 80, 180, 80, white), RotY4(20*Deg)),
		//Box(Vec{0, 0, 0}, rWidth, ceil, rDepth, wall),
		//Sheet(Ey, 0, Checkboard(rWidth/5, white, black)),
		//Sphere(Vec{0, .5, 1}, .5, Reflective(WHITE.EV(-.3))),
		//Rect(Vec{rWidth - off, 1.7, .75}, Ex, 0, 1, .7, Flat(WHITE.EV(1))),
		//Rect(Vec{rWidth - off, 1.7, -.75}, Ex, 0, 1, .7, Flat(WHITE.EV(1))),
	)

	e.AddLight(
		//PointLight(Vec{0, h - 100, 0}, WHITE.EV(10)),
		//SphereLight(Vec{0, h - 100, 0}, 50, WHITE.EV(30)),
		//RectLight(Vec{rWidth - off, ceil / 2, .6}, 0, 1, .45, WHITE.EV(2)),
		RectLight(Vec{0, h - 1e-4, 0}, 116/2, 0, 116/2, Color{0.8, 1.0, 0.7}.EV(20.3)),
	)

	e.SetAmbient(Flat(WHITE.EV(-6)))

	focalLen := 0.035 / 0.025
	e.Camera = Camera(focalLen).Transl(0, h/2, -1100)
	e.Camera.AA = true
	e.Recursion = 6
	e.Cutoff = EV(6)

	serve.Env(e)
}

const off = 1e-3
