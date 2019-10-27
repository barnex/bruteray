// +build example_lucy

// Usage:
// 	go run lucy.go
// or
// 	bruteray-watch lucy.go
package main

import (
	"math"

	. "github.com/barnex/bruteray/api"
	"github.com/barnex/bruteray/geom"
)

func main() {

	Render(Spec{
		//DebugNormals: true,
		//		DebugIsometricFOV: 10,
		//DebugIsometricDir: Y,

		Recursion: 2,
		NumPass:   1000,

		Lights: []Light{
			SunLight(C(1, .95, .9).EV(-1), 1.53*Deg, -2.5*Deg, 24*Deg),
		},

		Media: []Medium{
			Fog(0.6, 2),
			//Fog(0.7, 0.5),
			//Fog(0.4, 2),
			////ExpFog(0.3, C(1, 1, 1), 0.9),
			////ExpFog(0.1, C(1, 1, 1), 2.2),
		},

		Objects: []Object{
			Rectangle(Matte(C(0.5, 0.7, 0.5).EV(-2)), 100, 100, O),
			Backdrop(Flat(C(0.8, 0.8, 0.9).EV(-0.6))),
			ObjFile(
				map[string]Material{"": Matte(White.EV(-.3))},
				"../../../../../assets/Alucy.obj",
				geom.Scale(O, 2./1000.),
				geom.Rotate(O, Ey, 90*Deg),
			).WithCenterBottom(O),
		},
		Camera: Projective(90*Deg, Vec{0.3, 0.9, 1.9}, 20*Deg, 0*Deg, ),
	})
}

func cross1(m Material) Object {
	w1 := 0.20
	H := 1.8
	W := 0.9
	d1 := 0.15

	d2 := 0.08
	r1 := 0.47
	r2 := 0.69

	vbeam := Box(m, w1, H, d1, V(0, H/2, 0))
	hbeam := Box(m, W, w1, d1, V(0, H-W/2, 0))
	ring := Difference(
		CylinderWithCaps(m, r2, d2, O),
		Cylinder(m, r1, math.Inf(1), O),
	).Rotate(Ex, 90*Deg).WithCenter(hbeam.Center())
	return Tree(vbeam, hbeam, ring)
}
