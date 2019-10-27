package main

import . "github.com/barnex/bruteray/api"

func main() {
	wall := Matte(White.Mul(0.6))
	mainW := 8.0
	mainH := 4.0
	mainD := 12.0


	central := Box(wall, mainW, mainH+mainW/2, mainD, O).WithCenterBottom(O)

	dome := Cylinder(wall, mainW, mainH, O).WithCenterBottom(O).Or(
		Sphere(wall, mainW, V(0, mainH, 0)),
	).Translate(V(0, 0, -mainW))

	Render(Spec{
		Recursion: 3,
		Objects: []Object{

			central.Or(dome).Remove(Box(nil, 2,4,2, dome.CenterBack())),

			Rectangle(wall, 100, 100, V(0, .01, 0)),
			Backdrop(Flat(C(0.8, 0.8, 1.0).EV(-.6))),
		},
		Lights: []Light{
			SunLight(White.EV(2), 5*Deg, -18*Deg, 18*Deg),
			PointLight(White.EV(4), dome.Center()),
		},

		Media: []Medium{
			//Fog(0.02, mainH*2),
		},
	})
}
