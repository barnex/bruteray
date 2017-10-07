package main

import . "github.com/barnex/bruteray"
import "github.com/barnex/bruteray/serve"

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
		Rect(Vec{rWidth - off, 1.7, .75}, Ex, 0, 1, .7, Flat(WHITE.EV(1))),
		Rect(Vec{rWidth - off, 1.7, -.75}, Ex, 0, 1, .7, Flat(WHITE.EV(1))),
	)

	e.AddLight(
		SphereLight(Vec{0, ceil - 0.0, 0}, .1, WHITE.EV(1)),
	//RectLight(Rect(Vec{rWidth - off, ceil / 2, .6}, Ex, 0, 1, .5, Flat(BLUE)), WHITE.EV(2)),
	//RectLight(Rect(Vec{rWidth - off, ceil / 2, -.5}, Ex, 0, 1, .5, Flat(RED)), WHITE.EV(2)),
	)

	e.SetAmbient(Flat(WHITE.EV(-4)))

	e.Camera = Camera(0.8).Transl(0, 1.7, -3.6).RotScene(-30 * Deg).Transf(RotX4(0 * Deg))
	e.Camera.AA = true
	e.Recursion = 6
	e.Cutoff = EV(5)

	serve.Env(e)
}

const off = 1e-3
