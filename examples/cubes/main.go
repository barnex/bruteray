package main

import (
	"math"

	. "github.com/barnex/bruteray/api"
	"github.com/barnex/bruteray/imagef/post"
	"github.com/barnex/bruteray/util"
)

func main() {
	Render(Spec{
		Width:        1920 / 1,
		Height:       1080 / 1,
		DebugNormals: 0,
		Recursion:    3,
		NumPass:      2000,

		Objects: append(
			cubes(),
			Backdrop(Flat(White.EV(-6))),
			Rectangle(Flat(White.EV(2)), 64, 64, V(0, -2, 0)),
		),

		Lights: []Light{
			PointLight(White.EV(8), V(0, 0.5, 0)),
		},

		Media: []Medium{
			Fog(0.04, 8, 1),
		},

		Camera: ProjectiveAperture(70*Deg, 0.03, 18).Translate(V(6, 3, 11)).YawPitchRoll(30*Deg, -10*Deg, 0),

		PostProcess: post.Params{
			Gaussian: post.BloomParams{
				Radius:    0.03,
				Amplitude: 0.1,
				Threshold: 1,
			},
		},
	})
}

func cubes() []Object {
	m1 := Matte(White.EV(-1))
	m2 := Flat(White.EV(0))

	var objs []Object

	N := 64
	Nf := float64(N)
	d := 0.5
	d2 := 0.8 * d
	delta := -(Nf * d) / 2
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			x := float64(i)*d + delta
			z := float64(j)*d + delta
			y := f(x/(d*Nf), z/(d*Nf))
			m := m1
			r := math.Sqrt(util.Sqr(x/(d*Nf)) + util.Sqr(z/(d*Nf)))
			if y > 0 && r < 0.06 {
				m = Blend(0.02*util.Max(y*y*y*y*y*y, 1), m2, 1, m1)
			}
			objs = append(objs, Box(m, d2, d2, d2, V(x, y, z)))
		}
	}
	return objs
}

func f(x, y float64) float64 {
	x += 0.02
	y += 0.05
	r := math.Sqrt(x*x + y*y)
	return 2 * util.Sinc(r*32)
}
