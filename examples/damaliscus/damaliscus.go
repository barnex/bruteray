package main

import (
	. "github.com/barnex/bruteray/api"
	"github.com/barnex/bruteray/post"
)

func main() {

	skull := PlyFile(Matte(C(1, 1, 0.9).EV(-.6)), "assets/damaliscus.ply").ScaleToSize(0.4).Rotate(Ez, -30*Deg).WithCenterBottom(O)
	sand := IsoSurface(Matte(C(1, 0.8, 0.5).EV(-.9)), 3, 0.007, 3, LoadHeightMap("assets/lava2.png").Scale(0.5, 0.5).HeightMap()).WithCenterBottom(O)

	Render(Spec{
		Recursion:    3,
		Width:        1920 / 1,
		Height:       1080 / 1,
		NumPass:      900,
		DebugNormals: 0,
		Objects: []Object{
			skull.Rotate(Ey, 20*Deg).Rotate(Ex, 5*Deg).WithCenterBottom(V(0, -0.18, 0)),
			sand,
			Backdrop(Flat(C(0.8, 0.8, 1.0).EV(-1.9))),
		},
		Lights: []Light{
			SunLight(C(1, 0.9, 0.7).EV(1.0), 0.8*Deg, 33*Deg, 25*Deg),
			//SunLight(C(1, 0.9, 0.7).EV(0.7), 1.2*Deg, -130*Deg, 24*Deg),
		},

		Media: []Medium{
			Fog(0.1, 1, 1),
			//ExpFog(1, White, 1),
		},

		Camera: ProjectiveAperture(75*Deg, 0.004, 0.50, V(-0.05, 0.15, 0.4), 0*Deg, -12*Deg),

		PostProcess: post.Params{
			Gaussian: post.BloomParams{
				Radius:    0.005,
				Threshold: 1.1,
				Amplitude: 1,
			},
		},
	})
}
