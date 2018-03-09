package main

import (
	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/light"
	. "github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
	"github.com/barnex/bruteray/serve"
	. "github.com/barnex/bruteray/shape"
	. "github.com/barnex/bruteray/transf"
)

func main() {
	e := NewEnv()

	const U = 666

	glass := Refractive(1, 1.38)
	//glass := Diffuse(WHITE)
	plastic := Shiny(WHITE.EV(-7), EV(-5))

	const r = 75
	lens := And(
		NewSphere(2*r, glass).Transl(Vec{0, +0.8 * r, 0}),
		NewSphere(2*r, glass).Transl(Vec{0, -0.8 * r, 0}),
	)
	handle := Or(Minus(
		NewCylinder(Y, Vec{}, 45*2, 8, plastic),
		NewCylinder(Y, Vec{}, 40*2, 20, plastic),
	),
		And(
			Quad(Vec{}, Vec{1, 1, 0}, 25, plastic),
			Slab(Ez, -120, -45, plastic),
		),
	)
	mag := Or(lens, handle)
	//bmag := BoundBox(mag, Vec{-50, -10, -20}, Vec{50, 120, 20})
	tmag := Transf(mag, RotX4(-20*Deg).Mul(RotY4(35*Deg)).Mul(Transl4(Vec{5, 40, -29})))

	img := MustLoad("principia.png")
	p0 := Vec{-100, 0, -100}
	pu := Vec{100, 0, -100}
	pv := Vec{-100, 0, 100}
	//img.Mul(EV(-3))
	tex := NewImgTex(img, &UVAffine{p0, pu, pv})
	e.Add(
		Rect(Vec{0, 0, 0}, Ey, 100, U, 100, Diffuse(tex)),
		NewSheet(Ey, -10, Diffuse(WHITE.EV(-.6))),
		//NewSheet(Ey, -10, Checkboard(5, BLACK, WHITE)),
		tmag,
	)
	_ = tmag

	e.AddLight(
		light.RectLight(Vec{850, 300, -300}, 180, 800, 0, WHITE.EV(20.3)),
		light.RectLight(Vec{400, 300, -300}, 180, 800, 0, WHITE.EV(20.3)),
	)

	//e.SetAmbient(Flat(WHITE.EV(-3)))
	//pano := MustLoad("pano2.jpg")
	//pano.Mul(EV(.3))
	//e.SetAmbient(SkyCyl(pano, 0))
	e.SetAmbient(WHITE.EV(-.6))

	focalLen := 1.0

	cam := raster.Camera(focalLen).Transl(0, 150, -150).Transf(RotX4(50 * Deg))
	cam.AA = true

	e.Recursion = 7
	e.Cutoff = EV(40)

	cam.Focus = 210
	cam.Aperture = 1.5

	serve.Env(cam, e)
}
