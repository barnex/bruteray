// +build example_x

package main

import (
	. "github.com/barnex/bruteray/api"
)

func main() {
	Render(Spec{
		//DebugNormals:      true,
		//DebugIsometricFOV: 8,
		//DebugIsometricDir: Z,

		Recursion: 2,
		NumPass:   300,

		Lights: []Light{
			SunLight(White, 1*Deg, -120*Deg, 60*Deg),
		},
		Objects: []Object{
			Backdrop(Flat(White.EV(-3))),
			Rectangle(Matte(White.EV(-1)), 10, 10, O),
			ObjFile(map[string]Material{
				"Body":         Matte(C(0.5, 0.5, 1)),
				"SkinColor":    Matte(C(1, 0.5, 0.5)),
				"Tooth":        Shiny(White, 0.1),
				"Eye-White":    Shiny(White, 0.1),
				"NoseTop":      Shiny(C(1, 0.5, 0.5), 0.02),
				"Material":     Shiny(Gray(0.5), 0.1),
				"Material.001": Shiny(Gray(0.0), 0.1),
			}, "../../../../../assets/gopher.obj").ScaleToSize(1).Rotate(Ey, -80*Deg).WithCenterBottom(V(0, 0, 0)),
		},
		Camera: Projective(90*Deg, V(0, 0.5, 1.5), 0, 0),
		//Camera: EnvironmentMap(V(0, 0.5, 1.5)),
	})
}
