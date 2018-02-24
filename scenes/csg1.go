package main

import (
	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/csg"
	"github.com/barnex/bruteray/light"
	"github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
	"github.com/barnex/bruteray/serve"
	"github.com/barnex/bruteray/shape"
)

func main() {
	e := NewEnv()

	white := mat.Diffuse(WHITE.EV(-.3))
	red := mat.Diffuse(RED.EV(-.6))
	blue := mat.Diffuse(BLUE.EV(-.6))
	//s = ShadeShape(WHITE)

	floor := shape.Sheet(Ey, 0, white)

	cube := shape.NBox(1, 1, 1, red).Transl(Dy(0.501))
	sphere := shape.NSphere(1, blue).Transl(cube.Corner(R, R, -R))

	e.Add(
		floor,
		csg.Minus(cube, sphere),
	)

	e.AddLight(
		light.Sphere(Vec{4, 5, -2.5}, 0.3, WHITE.EV(8)),
	)

	e.SetAmbient(mat.Flat(WHITE.EV(-2)))
	e.Recursion = 4
	e.Cutoff = EV(2.6)

	cam := raster.Camera(1).Transl(0, 2.1, -2.5).RotScene(30 * Deg).Transf(RotX4(30 * Deg))
	//cam.Aperture = 0.04
	cam.Focus = 4
	cam.AA = true
	//e.Fog = 10
	//e.IndirectFog = true

	serve.Env(cam, e)
}

func Dy(off float64) Vec {
	return Vec{0, off, 0}
}

const (
	L = -1
	R = 1
)
