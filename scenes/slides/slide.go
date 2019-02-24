package main

import (
	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/light"
	. "github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
	"github.com/barnex/bruteray/serve"
	. "github.com/barnex/bruteray/shape"
	//. "github.com/barnex/bruteray/transf"
)

func main() {
	e := NewEnv()

	const U = 666

	//glass := Refractive(1, 1.38)

	img := MustLoad("sslide1.jpg").Mul(1.8)
	p0 := Vec{-150, 0, -100}
	pu := Vec{150, 0, -100}
	pv := Vec{-150, 0, 100}
	tex := NewImgTex(img, &UVAffine{p0, pu, pv})
	e.Add(
		Rect(Vec{0, 10, 0}, Ey, 150, U, 100, Diffuse(tex)),
		NewSheet(Ey, -0, Checkboard(30, Diffuse(WHITE.EV(-.6)), Diffuse(WHITE.EV(-1)))),

		NewSphere(120, ReflectFresnel(3, BLACK)).Transl(Vec{-40, 60, 140}),
		NewSphere(100, Refractive(1, 1.5)).Transl(Vec{120, 0, -120}),

		NewSphere(60, Diffuse(Color{.2, .2, .8})).Transl(Vec{-230, 30, 100}),
		NewSphere(60, Shiny(Color{.2, .8, .2}, 0.1)).Transl(Vec{-190, 30, 0}),
		NewSphere(60, Shiny(Color{.8, .2, .2}, 0.2)).Transl(Vec{-190, 30, 190}),

		NewCylinder(Y, Vec{170, 30, 145}, 100, 80, Shiny(WHITE.EV(-1), 0.2)),
		NewCube(80, Diffuse(WHITE.EV(-1))).Transl(Vec{220, 20, 0}),
	)

	e.AddLight(
		//light.RectLight(Vec{850, 300, -300}, 180, 800, 0, WHITE.EV(20.3)),
		//light.RectLight(Vec{400, 300, -300}, 180, 800, 0, WHITE.EV(20.3)),
		light.Sphere(Vec{800, 700, 600}, 150, WHITE.Mul(EV(24.3))),
		//light.Sphere(Vec{-190, 30, 150}, 30, WHITE.Mul(EV(15))),
	)

	sky := MustLoad("../../assets/sky1.jpg").Mul(1.5)
	e.SetAmbient(SkyDome(sky, -90*Deg))

	focalLen := 1.1

	cam := raster.Camera(focalLen).Transl(0, 300, -150).Transf(RotX4(60 * Deg))
	cam.AA = true

	e.Recursion = 5
	e.Cutoff = EV(2)

	cam.Focus = 280
	cam.Aperture = 15
	cam.Diaphragm = DiaHex

	serve.Env(cam, e)
}
