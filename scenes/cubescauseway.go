// +build ignore

package main

import (
	"math/rand"

	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/serve"
)

func main() {
	e := NewEnv()

	const (
		W = 0.6
	)
	w := Diffuse(WHITE.EV(-.3))
	g := Flat(WHITE.EV(1.0))

	e.Add(
		Sheet(Ey, 0, Diffuse(WHITE.Mul(EV(-1)))),
	)

	rand.Seed(5)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			u := float64(i) - 3.5 + rand.Float64()*0.4
			v := float64(j) - 3.5 + rand.Float64()*0.4
			h := rand.Float64()*0.9 - u*0.1 + v*0.3
			if (i+j)%2 == 0 {
				e.Add(Box(Vec{u, 0, v}, W, h, W, w))
			} else {
				e.Add(
					Box(Vec{u, 0, v}, W, h, W, g),
				)
			}
		}
	}

	e.AddLight(
	//SphereLight(Vec{18, 17, 18}.Mul(2), 20, WHITE.EV(13.0)),
	//SphereLight(Vec{-6, 6, -25}.Mul(10), 160, WHITE.EV(18)),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-5))))

	e.Camera = Camera(.9).Transl(0, 3.0, -7).RotScene(20 * Deg).Transf(RotX4(30 * Deg))

	e.Camera.AA = true
	e.Recursion = 3
	e.Cutoff = EV(40)
	e.Fog = 8

	serve.Env(e)
}
