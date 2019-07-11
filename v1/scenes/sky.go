// +build ignore

package main

import (
	. "github.com/barnex/bruteray/v1/br"
	"github.com/barnex/bruteray/v1/mat"
	"github.com/barnex/bruteray/v1/raster"
	"github.com/barnex/bruteray/v1/serve"
	"github.com/barnex/bruteray/v1/shape"
)

func main() {
	e := NewEnv()

	img := mat.MustLoad("../assets/sky14e.jpg")
	img.Mul(EV(2.6))

	//tile := mat.MustLoad("../assets/monalisa.jpg")
	w := mat.Diffuse(WHITE.EV(-.9))
	//x := mat.Checkboard(0.1, WHITE, BLACK)

	e.Add(
		//shape.NewSheet(Ey, -.6, mat.Diffuse(mat.NewImgTex(tile, &mat.UVAffine{P0: Vec{-5, 0, -1.8}, Pu: Ex.Mul(10), Pv: Ez.Mul(10)}))),
		shape.NewSheet(Ey, -.6, mat.Checkboard(1, mat.Diffuse(WHITE.EV(-2)), mat.Reflective(WHITE.EV(-6)))),
		shape.NewCylinder(Y, Vec{0, -.5, .2}, 4, 0.1, w),
		shape.NewSphere(1, mat.Reflective(WHITE.EV(-1))).Transl(Vec{-.6, 0, .1}),
		shape.NewSphere(1, mat.Refractive(1, 1.4)).Transl(Vec{.6, 0, .6}),
	)

	e.AddLight()

	e.SetAmbient(mat.SkyDome(img, 270*Deg))
	cam := raster.Camera(1).Transl(0, .9, -2.6).Transf(RotX4(20 * Deg))
	cam.AA = true
	cam.Focus = 3.2
	cam.Diaphragm = DiaCircle
	cam.Aperture = 1. / 16.

	serve.Env(cam, e)
}
