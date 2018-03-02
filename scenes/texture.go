package main

import (
	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/light"
	"github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
	"github.com/barnex/bruteray/serve"
	"github.com/barnex/bruteray/shape"
)

func main() {
	e := NewEnv()

	ml := mat.MustLoad("assets/monalisa.jpg")
	m := mat.Texture(ml, Vec{-.4, -.5, 0}, Vec{.4, -.5, 0}, Vec{-.5, .5, 0})

	cube := shape.NBox(.8, 1, .1, m)

	e.Add(
		shape.Sheet(Ey, -.5, mat.Diffuse(WHITE)),
		cube,
	)

	e.AddLight(
		light.Sphere(Vec{4, 5, -3}, 1, WHITE.EV(10)),
	)

	e.SetAmbient(mat.Flat(WHITE.EV(-2)))
	cam := raster.Camera(1).Transl(0, 1, -3).RotScene(9 * Deg).Transf(RotX4(25 * Deg))
	cam.AA = true
	serve.Env(cam, e)
}
