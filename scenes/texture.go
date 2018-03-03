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

	tex := mat.MustLoad("assets/monalisa.jpg")
	cube := shape.NBox(1, tex.Aspect(), .1, nil)
	cube.Transl(Vec{0, 0.5, 0})

	p0 := cube.Corner(-1, -1, -1)
	pu := cube.Corner(+1, -1, -1)
	pv := cube.Corner(-1, +1, -1) //.Add(Vec{0, -.8, 0})
	cube.Mat = mat.Diffuse(mat.NewImgTex(tex, p0, pu, pv))

	e.Add(
		shape.Sheet(Ey, cube.Min[Y], mat.Diffuse(WHITE.EV(-2))),
		shape.NSphere(.2, mat.ShadeShape(WHITE)).Transl(p0),
		shape.NSphere(.2, mat.ShadeShape(RED)).Transl(pu),
		shape.NSphere(.2, mat.ShadeShape(GREEN)).Transl(pv),
		cube,
	)

	e.AddLight(
		light.Sphere(Vec{0, cube.Min[Y] + .1, -.5}, .1, WHITE.EV(5)),
	)

	e.SetAmbient(mat.Flat(WHITE.EV(-4)))
	cam := raster.Camera(1).Transl(0, 0.5, -3).RotScene(9 * Deg).Transf(RotX4(5 * Deg))
	cam.AA = true
	serve.Env(cam, e)
}
