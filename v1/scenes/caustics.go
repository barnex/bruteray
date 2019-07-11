// +build ignore

package main

import (
	"math"

	. "github.com/barnex/bruteray/v1/br"
	. "github.com/barnex/bruteray/v1/mat"
	"github.com/barnex/bruteray/v1/raster"
	"github.com/barnex/bruteray/v1/serve"
	. "github.com/barnex/bruteray/v1/shape"
)

func main() {
	e := NewEnv()

	water := Refractive(1, 1.5)

	sea := func(p Vec) float64 {
		z := p[Z]
		x := p[X]
		y := p[Y]
		r := math.Sqrt(x*x + z*z)
		return -0.5*math.Sin(r)/r + 0.05*math.Sin(1.6432*z+.3*x) + 0.03*math.Sin(1.983*x+.6903*z) + y
	}

	e.Add(
		NewFunction(-.5, .5, water, sea),
		//bouy,
		//NewSheet(Ey, -1, Checkboard(1, Diffuse(WHITE), Diffuse(WHITE.EV(-1)))),
		NewSheet(Ey, -9, Diffuse(WHITE)),
		NewSphere(4, Flat(WHITE.EV(5))).Transl(Vec{10, 10, 10}),
	)

	e.SetAmbient(SkyDome(MustLoad("../assets/sky14d.jpg").Mul(EV(-1.3)), 180*Deg))
	cam := raster.Camera(1).Transl(4, 6, -20).RotScene(9 * Deg).Transf(RotX4(20 * Deg))
	cam.AA = true
	//cam.Aperture = 0.14
	//cam.Focus = 22
	serve.Env(cam, e)
}
