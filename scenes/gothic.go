package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	white := Diffuse(WHITE.EV(-.9))
	//white = Checkboard(1, ShadeShape(WHITE), ShadeShape(WHITE.EV(-.3)))

	const (
		w      = 3.0            // central width
		h1     = 3.0            // pillar height
		pill   = 0.5            // pillar width
		d      = 3 * (w + pill) // central depth
		pointy = 0.8            //
		b      = 0.5
		//d2    = 2.0
		//brick = 0.05
	)

	//hull := NBox(Vec{}, d+1, d+1, d+1, white)

	all := multiOr(
		chestz(Vec{}, w, h1, d, pointy, white),
		chestz(Vec{Y: b}, w-b, h1-b, d+4, pointy, Flat(WHITE)),

		chestz(Vec{X: -w - pill}, w, h1, d, pointy, white),
		chestz(Vec{X: +w + pill}, w, h1, d, pointy, white),

		chestx(Vec{}, w, h1, d, pointy, white),
		chestx(Vec{Z: -w - pill}, w, h1, d, pointy, white),
		chestx(Vec{Z: +w + pill}, w, h1, d, pointy, white),
	)

	tileB := Shiny(WHITE.EV(-5), EV(-5))
	tileW := Shiny(WHITE.EV(-1), EV(-5))

	tileB = Flat(BLACK)
	tileW = Flat(WHITE)

	e.Add(
		Sheet(Ey, 0.01, Checkboard(1, tileW, tileB)),
		all,
	)

	lighth := h1/2 + w/4 + pointy/2
	e.AddLight(
		//PointLight(Vec{0.3, 2.5, 0}, WHITE.EV(5)),
		//PointLight(Vec{1, 5, -3}, WHITE.EV(7)),
		RectLight(Vec{0, lighth, d/2 + b}, w/2, lighth, 0, WHITE.EV(5.6)),
	)

	e.Camera = Camera(0.65).Transl(0, 2.2, -4.9).RotScene(0 * Deg).Transf(RotX4(-0 * Deg))
	//e.Camera.Aperture = 0.05
	//e.Camera.Focus = 8
	//e.Camera.AA = true
	//e.Recursion = 3
	//e.Fog = 8

	serve.Env(e)
}

func chestz(pos Vec, w, h1, d, pointy float64, m Material) Obj {
	const off = 1e-4
	c1 := Cyl(Z, Vec{pointy / 2, h1, 0}.Add(pos), w+pointy, d-off, m)
	c2 := Cyl(Z, Vec{-pointy / 2, h1, 0}.Add(pos), w+pointy, d+off, m)
	ceil := And(c1, c2)
	box := NBox(pos.Add(Vec{Y: h1 / 2}), w, h1, d+2*off, m)
	return Or(box, ceil)
}

func chestx(pos Vec, w, h1, d, pointy float64, m Material) Obj {
	const off = 1e-4
	c1 := Cyl(X, Vec{0, h1, pointy / 2}.Add(pos), w+pointy, d-off, m)
	c2 := Cyl(X, Vec{0, h1, -pointy / 2}.Add(pos), w+pointy, d+off, m)
	ceil := And(c1, c2)
	box := NBox(pos.Add(Vec{Y: h1 / 2}), d+2*off, h1, w, m)
	return Or(box, ceil)
}

func multiOr(o ...Obj) Obj {
	obj := o[0]
	for i := 1; i < len(o); i++ {
		obj = Or(obj, o[i])
	}
	return obj
}
