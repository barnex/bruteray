package main

import (
	"math"

	. "github.com/barnex/bruteray/br"
	. "github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
	"github.com/barnex/bruteray/serve"
	. "github.com/barnex/bruteray/shape"
)

func main() {
	e := NewEnv()

	//water := Reflective(WHITE.EV(-2))
	water := ReflectFresnel(1.3, GREEN.EV(-4.6))

	sea := func(p Vec) float64 {
		z := p[Z]
		x := p[X]
		y := p[Y]
		r := math.Sqrt(x*x + z*z)
		return 0.4*math.Sin(r)/r + 0.005*math.Sin(1.6432*z+.3*x) - y
	}

	e.Add(
		NewFunction(-.5, .5, water, sea),
	)

	e.SetAmbient(SkyDome(MustLoad("../assets/sky14c.jpg").Mul(EV(2)), 0))
	cam := raster.Camera(1).Transl(0.5, 4, -7).RotScene(9 * Deg).Transf(RotX4(20 * Deg))
	cam.AA = true
	serve.Env(cam, e)
}
