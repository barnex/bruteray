package main

import (
	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	//white := Diffuse(WHITE.EV(-.9))
	white := Bricks(0.3778, 0.05, Diffuse(WHITE.EV(-.9)), Diffuse(WHITE.EV(-1.3)))
	//white = Bricks(0.3778, 0.05, ShadeShape(WHITE.EV(-.9)), ShadeShape(WHITE.EV(-1.3)))

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

		Cutout(
			chestx(Vec{}, w, h1, d, pointy, white),
			chestx(Vec{X: 1, Y: b}, w-2*b, h1-b, d+1, pointy, white),
		),

		Cutout(
			chestx(Vec{Z: -w - pill}, w, h1, d, pointy, white),
			chestx(Vec{X: 1, Y: b, Z: -w - pill}, w-2*b, h1-b, d+1, pointy, white),
		),

		Cutout(
			chestx(Vec{Z: +w + pill}, w, h1, d, pointy, white),
			chestx(Vec{X: 1, Y: b, Z: +w + pill}, w-2*b, h1-b, d+1, pointy, white),
		),
	)

	tileB := speckled()
	tileW := marmer()

	barx := 2*w - 0.4
	barw := 0.099
	barc := Diffuse(WHITE.EV(-4))
	e.Add(
		Rect(Ey, Vec{0, 0.03, 0}, w+pill, 0.01, d, Checkboard(1, tileW, tileB)),
		all,

		Box(Vec{barx, 4, 0}, barw, barw, 20, barc),
		Box(Vec{barx, 4, 0}, 0.05, 20, barw, barc),
		Box(Vec{barx, 4, w + pill}, barw, 20, barw, barc),
		Box(Vec{barx, 4, -w - pill}, barw*3, 20, barw*3, barc),

		Rect(Vec{0, 1, 9.4}, Ez, 0.5, 2, 0.01, Flat(WHITE)),
		Rect(Vec{0.1, 1, 9.35}, Ez, 0.5, 2, 0.01, Flat(BLACK)),
	)

	//lighth := h1/2 + w/4 + pointy/2
	dst := 8.0
	e.AddLight(
		SphereLight(Vec{12, 3, -2}.Mul(dst), 0.04*dst*dst, WHITE.EV(13).Mul(dst*dst)),
	//PointLight(Vec{1, 5, -3}, WHITE.EV(7)),
	//RectLight(Vec{0, lighth, d/2 + b}, w/2, lighth, 0, WHITE.EV(5.6)),
	)

	e.Camera = Camera(0.65).Transl(0, 2.5, -7.0).RotScene(0 * Deg).Transf(RotX4(-0 * Deg))
	e.SetAmbient(Flat(WHITE.EV(0)))
	e.Camera.Aperture = 0.03
	e.Camera.Focus = 9
	e.Camera.AA = true
	e.Recursion = 4
	e.Cutoff = EV(4)
	e.Fog = 7

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

func marmer() Material {
	a := Shiny(WHITE.EV(-0.9), EV(-5))
	b := Shiny(WHITE.EV(-2), EV(-4))
	//a = Flat(WHITE)
	//b = Flat(BLACK)
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

func speckled() Material {
	//return Distort(123, 20, Vec{500, 500, 500}, 0.1, Reflective(WHITE.EV(-4)))
	//return Reflective(WHITE.EV(-4))
	a := Shiny(WHITE.EV(-6), EV(-5))
	b := Diffuse(WHITE.EV(-6))
	//a = Flat(WHITE)
	//b = Flat(BLACK)
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
