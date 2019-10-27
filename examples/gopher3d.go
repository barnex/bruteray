// +build example_gopher3d

// Usage:
// 	go run gopher3d.go
// or
// 	bruteray-watch gopher3d.go
package main

import (
	"fmt"
	"math/rand"

	. "github.com/barnex/bruteray/api"
	"github.com/barnex/bruteray/geom"
)

func main() {
	const numFrame = 300

	const numGopher = 5
	var jumps [numGopher][]float64
	for i := range jumps {
		jumps[i] = makeJumps(i, numFrame)
	}

	fmt.Println(jumps)

	gopher := PlyFile(
		Shiny(White, 0.02),
		"/home/arne/assets/gopher.ply",
		geom.Scale(O, 0.2),
		geom.Rotate(O, Ey, 90*Deg),
		geom.Rotate(O, Ex, -90*Deg),
	)

	grass := Matte(C(0.5, 0.9, 0.5).EV(-1))

	_ = grass
	_ = gopher

	th := 0.2
	white := Matte(White)
	cloud := Tree(
		CylinderWithCaps(white, 2, th, V(1, 0, 1)),
		CylinderWithCaps(white, 2, th, V(2, 0, 2.3)),
		CylinderWithCaps(white, 2, th, V(3, 0, 1.8)),
		CylinderWithCaps(white, 2, th, V(3, 0, 0.7)),
		CylinderWithCaps(white, 2, th, V(4, 0, 1)),
		Box(white, 3, th, 2, V(2.5, 0, 1)),
	).Rotate(Ex, -90*Deg)

	Animate(numFrame, func(i int) Spec {
		t := float64(i) / 30
		return Spec{
			//DebugNormals: true,
			//DebugIsometricFOV: 20,
			//DebugIsometricDir: Z,

			Recursion: 2,
			NumPass:   50,
			Width:     1920 / 4,
			Height:    1080 / 4,

			Lights: []Light{
				SunLight(C(1, .95, .9).EV(-.6), 7*Deg, 44*Deg, 25*Deg),
			},

			Media: []Medium{
				//Fog(0.01, 0.3),
			},

			Objects: []Object{
				cloud.WithCenter(V(-4+t, 6, 10)),
				cloud.WithCenter(V(6+t, 5.5, 11)),
				cloud.WithCenter(V(-11+t, 5, 10)),

				gopher.WithCenter(V(0, jumps[0][i], 0)),
				gopher.WithCenter(V(-1, jumps[1][i], 1)).WithMaterial(Shiny(C(0.5, 0.5, 1.0), 0.01)),
				gopher.WithCenter(V(1, jumps[2][i], 2)).WithMaterial(Shiny(C(.5, 1, .5), 0.01)),
				gopher.WithCenter(V(2, jumps[3][i], 1)).WithMaterial(Shiny(C(1, .5, .5), 0.01)),
				gopher.WithCenter(V(1, jumps[4][i], 1)).WithMaterial(Shiny(C(.4, .8, .8), 0.01)),

				//			Rectangle(Matte(C(0.5, 0.9, 0.5).EV(-1)), 100, 100, O),
				Object{And(
					Box(grass, 100, 1, 100, V(0, -.5, 0)),
					Not(Tree(
						Cylinder(grass, 1.2, 999, V(0, 0, 0)),
						Cylinder(grass, 1.2, 999, V(-1, 0, 1)),
						Cylinder(grass, 1.2, 999, V(1, 0, 2)),
						Cylinder(grass, 1.1, 999, V(2.1, 0, 1)),
						Cylinder(grass, 1.1, 999, V(1, 0, 1)),
					)))},
				Backdrop(Flat(C(0.6, 0.65, 1.0).EV(-1.0))),
			},

			//Media: []Medium{
			//	media.ExpFog(0.15, White),
			//},

			Camera: Projective(90*Deg, V(0.3, 0.5, -1.5), 180*Deg, 0*Deg, ),
		}
	})
}

func makeJumps(i int, numFrame int) []float64 {
	Y := make([]float64, numFrame)
	rng := rand.New(rand.NewSource(int64(i)))

	fps := 30.
	dt := 1. / fps

	vInit := 1 + rng.Float64()
	v := vInit
	a := -0.9
	y := rng.Float64()
	offset := -1.
	for i := range Y {
		y += v * dt
		v += a * dt
		if y < 0 {
			y = 0
			v = -v
		}
		Y[i] = y + offset
	}
	return Y
}
