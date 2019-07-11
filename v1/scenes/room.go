//+build ignore

package main

import . "github.com/barnex/bruteray"
import "github.com/barnex/bruteray/v1/serve"

func main() {
	e := NewEnv()

	wall := Diffuse(WHITE.EV(-.3))
	white := Shiny(WHITE.EV(-.6), EV(-4))
	black := Shiny(WHITE.EV(-6), EV(-4))

	const (
		rWidth = 3.
		rDepth = 4.
		ceil   = 3.
	)
	e.Add(
		Box(Vec{0, 0, 0}, rWidth, ceil, rDepth, wall),
		Sheet(Ey, 0, Checkboard(rWidth/5, white, black)),
		Sphere(Vec{0, .5, 1}, .5, Reflective(WHITE.EV(-.3))),
		//Rect(Vec{rWidth - off, 1.7, .75}, Ex, 0, 1, .7, Flat(WHITE.EV(1))),
		//Rect(Vec{rWidth - off, 1.7, -.75}, Ex, 0, 1, .7, Flat(WHITE.EV(1))),
	)

	e.AddLight(
		SphereLight(Vec{0, ceil - 0.4, 0}, .1, WHITE.EV(3)),
		RectLight(Vec{rWidth - off, ceil / 2, .6}, 0, 1, .45, WHITE.EV(2)),
		RectLight(Vec{rWidth - off, ceil / 2, -.5}, 0, 1, .45, WHITE.EV(2)),
	)

	e.SetAmbient(Flat(WHITE.EV(-4)))

	e.Camera = Camera(0.7).Transl(0, 1.7, -3.0).RotScene(-30 * Deg).Transf(RotX4(0 * Deg))
	e.Camera.AA = true
	e.Recursion = 4
	e.Cutoff = EV(4)

	serve.Env(e)
}

const off = 1e-3
