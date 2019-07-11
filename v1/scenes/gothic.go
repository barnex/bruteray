// +build ignore

package main

import (
	. "github.com/barnex/bruteray/v1/br"
	. "github.com/barnex/bruteray/v1/csg"
	. "github.com/barnex/bruteray/v1/light"
	. "github.com/barnex/bruteray/v1/mat"
	"github.com/barnex/bruteray/v1/raster"
	"github.com/barnex/bruteray/v1/serve"
	. "github.com/barnex/bruteray/v1/shape"
)

func main() {
	e := NewEnv()

	white := Bricks(0.3778, 0.05, ShadeShape(WHITE.EV(-1)), ShadeShape(WHITE.EV(-1.3)))

	//white = Bricks(0.3778, 0.05, brick(), cement())

	const (
		w      = 4.0              // central width
		h1     = 4.5              // pillar height
		pill   = 0.5              // pillar width
		d      = 3 * (w + 2*pill) // central depth
		pointy = 0.8              //
		b      = 0.6
		//d2    = 2.0
		//brick = 0.05
	)

	all := MultiOr(
		chestz(Vec{}, w, h1, d+w, pointy, white),
		//Cutout(
		//	chestz(Vec{}, w, h1, d+w, pointy, white),
		//	chestz(Vec{Y: b}, w-b, h1-b, d+5, pointy, Flat(BLUE)),
		//),

		chestz(Vec{X: -w - pill}, w, h1, d, pointy, white),
		chestz(Vec{X: +w + pill}, w, h1, d, pointy, white),

		//Cutout(
		chestx(Vec{}, w, h1, d, pointy, white),
		chestx(Vec{X: 1, Y: b}, w-2*b, h1-b, d+.1, pointy, white),
		//),

		//Cutout(
		chestx(Vec{Z: -w - pill}, w, h1, d, pointy, white),
		chestx(Vec{X: 1, Y: b, Z: -w - pill}, w-2*b, h1-b, d+.1, pointy, white),
		//),

		Cutout(
			chestx(Vec{Z: +w + pill}, w, h1, d, pointy, white),
			chestx(Vec{X: 1, Y: b, Z: +w + pill}, w-2*b, h1-b, d+1, pointy, white),
		),
	)

	tileB := speckled()
	tileW := marmer()

	tileB = Flat(BLACK)
	tileW = Flat(WHITE)

	barx := 2*w - 0.4
	barw := 0.099
	barc := Distort(5, 5, Vec{400, 400, 400}, 0.2, Reflective(WHITE.EV(-4)))
	e.Add(
		Box(Vec{0, 0.00, 0}, 2*w+2*pill, 0.01, d, Checkboard(1, tileW, tileB)),
		all,

		Box(Vec{barx, 4, 0}, barw, barw, 20, barc),
		Box(Vec{barx, 4, 0}, 0.05, 20, barw, barc),
		Box(Vec{barx, 4, w + pill}, barw, 20, barw, barc),
		Box(Vec{barx, 4, -w - pill}, barw*3, 20, barw*3, barc),

		Rect(Vec{0, 1, 9.4}, Ez, 0.5, 2, 0.01, Flat(WHITE)),
		Rect(Vec{0.1, 1, 9.35}, Ez, 0.5, 2, 0.01, Flat(BLACK)),
	)

	dst := 8.0
	e.AddLight(
		SphereLight(Vec{12, 2.7, -2.85}.Mul(dst), 0.04*dst*dst, WHITE.EV(11.9).Mul(dst*dst)),
		SphereLight(Vec{w + pill, 3, -w - pill - 0.5}, 0.1, WHITE.EV(0.3).Mul(dst*dst)),
		//SphereLight(Vec{w + pill - 1, 3, -w - pill + 0.5}, 0.2, WHITE.EV(3).Mul(dst*dst)),
	)
	e.SetAmbient(Flat(WHITE.EV(0)))
	e.Recursion = 3
	e.Cutoff = EV(2.6)

	cam := raster.Camera(0.65).Transl(0, 2.5, -6.9).RotScene(-0 * Deg).Transf(RotX4(-0 * Deg))
	cam.Aperture = 0.04
	cam.Focus = 9
	cam.AA = true
	//e.Fog = 10
	//e.IndirectFog = true

	serve.Env(cam, e)
}

func chestz(pos Vec, w, h1, d, pointy float64, m Material) CSGObj {
	const off = 1e-4
	c1 := Cyl(Z, Vec{pointy / 2, h1, 0}.Add(pos), w+pointy, d-off, m)
	c2 := Cyl(Z, Vec{-pointy / 2, h1, 0}.Add(pos), w+pointy, d+off, m)
	ceil := And(c1, c2)
	box := NBox(pos.Add(Vec{Y: h1 / 2}), w, h1, d+2*off, m)
	return Or(box, ceil)
}

func chestx(pos Vec, w, h1, d, pointy float64, m Material) CSGObj {
	const off = 1e-4
	c1 := Cyl(X, Vec{0, h1, pointy / 2}.Add(pos), w+pointy, d-off, m)
	c2 := Cyl(X, Vec{0, h1, -pointy / 2}.Add(pos), w+pointy, d+off, m)
	ceil := And(c1, c2)
	box := NBox(pos.Add(Vec{Y: h1 / 2}), d+2*off, h1, w, m)
	return Or(box, ceil)
}

func marmer() Material {
	a := Shiny(WHITE.EV(-0.9), EV(-5))
	b := Shiny(WHITE.EV(-1.6), EV(-3))
	f := func(v float64) Material {
		//v = v * v
		if v > 0.2 && v < 0.6 {
			return b
		} else {
			return a
		}
	}
	return Waves(20, 20, Vec{50, 50, 50}, f)
}

func brick() Material {
	m := Diffuse(WHITE.EV(-.9))
	return Distort(25, 25, Vec{40, 300, 40}, 0.03, m)
}

func cement() Material {
	m := Diffuse(WHITE.EV(-1.3))
	return Distort(10, 10, Vec{400, 400, 400}, 0.2, m)
}

func speckled() Material {
	a := Reflective(WHITE.EV(-4))
	b := Diffuse(WHITE.EV(-4))
	f := func(v float64) Material {
		//v = v * v
		if v*v > 0.9 {
			return b
		} else {
			return a
		}
	}
	return Waves(50, 20, Vec{200, 200, 200}, f)
}
