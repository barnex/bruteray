// +build ignore

package main

import (
	. "github.com/barnex/bruteray/v1/br"
	"github.com/barnex/bruteray/v1/csg"
	"github.com/barnex/bruteray/v1/light"
	"github.com/barnex/bruteray/v1/mat"
	"github.com/barnex/bruteray/v1/raster"
	"github.com/barnex/bruteray/v1/serve"
	"github.com/barnex/bruteray/v1/shape"
)

func main() {
	e := NewEnv()

	white := mat.Diffuse(WHITE.EV(-.3))
	red := mat.Diffuse(RED.EV(-.6))
	blue := mat.Diffuse(BLUE.EV(-.6))
	//s = ShadeShape(WHITE)

	floor := shape.Sheet(Ey, 0, white)

	cube := shape.NBox(1, 1, 1, red).Transl(Dy(0.501))
	//sphere := shape.NSphere(1, blue).Transl(cube.Corner(R, R, -R))
	cyl := shape.NCyl(Z, 0.7, blue).Transl(cube.Center())

	e.Add(
		floor,
		csg.Minus(cube, cyl),
	)

	e.AddLight(
		light.Sphere(Vec{4, 5, -2.5}, 0.3, WHITE.EV(8)),
	)

	e.SetAmbient(mat.Flat(WHITE.EV(-2)))
	e.Recursion = 4
	e.Cutoff = EV(2.6)

	cam := raster.Camera(1).Transl(0, 1.1, -2.5).RotScene(30 * Deg).Transf(RotX4(20 * Deg))
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
