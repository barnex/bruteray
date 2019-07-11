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

	//water := Reflective(WHITE.EV(-2))
	water := ReflectFresnel(1.3, Diffuse(GREEN.EV(-3)))

	sea := func(p Vec) float64 {
		z := p[Z]
		x := p[X]
		y := p[Y]
		r := math.Sqrt(x*x + z*z)
		return 0.7*math.Sin(r)/r + 0.005*math.Sin(1.6432*z+.3*x) - y
	}

	plastic := ReflectFresnel(1.2, Diffuse(YELLOW))
	//plastic := Diffuse(YELLOW)
	d := 3.
	bouyH := 0.8
	bouypos := Vec{-0, 0.5, 0}
	bouy := Or(NewCylinder(Y, bouypos, d, bouyH, plastic), NewSphere(d, plastic).Transl(bouypos.Add(Vec{Y: bouyH / 3})))

	_ = bouy
	//_ = NewFunction(-.5, .5, water, sea)

	e.Add(
		NewFunction(-.5, .5, water, sea),
		bouy,
		//NewSheet(Ey, 0, Checkboard(1, BLACK, WHITE)),
	)

	e.SetAmbient(SkyDome(MustLoad("../assets/sky14d.jpg").Mul(EV(1)), 340*Deg))
	cam := raster.Camera(1).Transl(4, 4, -20).RotScene(19 * Deg).Transf(RotX4(20 * Deg))
	cam.AA = true
	cam.Aperture = 0.14
	cam.Focus = 22
	serve.Env(cam, e)
}
