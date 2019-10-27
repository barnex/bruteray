package main

import . "github.com/barnex/bruteray/api"

func main() {
	wall := Matte(White.Mul(0.6))
	mainW := 8.0
	mainH := 4.0

	cyl := Cylinder(wall, mainW, mainH, O).WithCenterBottom(O).Restrict(
				Box(nil, mainW, mainH+1, mainW, O).WithCenterBottom(V(0, -.1, 0)).Translate(V(0, 0, mainW/2)),
			)
	dome := Sphere(wall, mainW, V(0, mainH, 0)).Restrict(Box(nil, mainW, mainW/2, mainW, V(0,mainW/2+mainH/2,mainW/2)))


	Render(Spec{
		Recursion: 2,
		Objects: []Object{
			cyl,
			dome,
			Rectangle(wall, 100, 100, O),
			Backdrop(Flat(C(0.8, 0.8, 1.0).EV(-.6))),
		},
		Lights: []Light{
			SunLight(White, 5*Deg, -50*Deg, 50*Deg),
		},
	})
}
