package main

import (
	. "github.com/barnex/bruteray/api"
	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/imagef/post"
	"github.com/barnex/bruteray/texture"
)

func main() {
	wall := Matte(C(1, 0.9, 0.8).Mul(0.3))

	H1 := 6.0
	W1 := 6.0
	H2 := H1 + W1/2

	c := Cylinder(wall, W1, H1, V(0, H1/2, 0))
	s := Sphere(wall, W1, c.CenterTop())
	//b := Box(nil, W1, H2, W1/2, V(0, H2/2, -W1/4))

	winW := 1.5
	winB := 1.0
	winH := H1 - winB
	winH2 := H1 - winB + winW

	win := Box(wall, winW, winH, 1, V(0, winH/2+winB, -W1/2)).Or(
		CylinderZ(wall, 2*winW, 1, V(winW/2, winH+winB, -W1/2)).And(
			CylinderZ(wall, 2*winW, 1, V(-winW/2, winH+winB, -W1/2)),
		),
	)

	//glassT1 := White

	marbleT := LoadTexture("/home/arne/assets/marble2t.jpg")
	marble := ReflectFresnel(1.3, Matte(marbleT))
	_ = marble
	glassT1 := LoadTexture("/home/arne/assets/stained1.jpg")
	glassT2 := glassT1
	glassT3 := glassT1
	//glassT1 := LoadTexture("/home/arne/assets/gopherglass1.png")
	//glassT2 := LoadTexture("/home/arne/assets/gopherglass2.png")
	//glassT3 := LoadTexture("/home/arne/assets/gopherglass3.png")

	glass1 := RectangleWithVertices(Transparent(glassT1, true),
		V(0, 0, 0), V(winW, 0, 0), V(0, winH2, 0)).WithCenterBottom(win.CenterBottom().Add(V(0, 0, -.2)))
	glass2 := RectangleWithVertices(Transparent(glassT2, true),
		V(0, 0, 0), V(winW, 0, 0), V(0, winH2, 0)).WithCenterBottom(win.CenterBottom().Add(V(0, 0, -.2)))
	glass3 := RectangleWithVertices(Transparent(glassT3, true),
		V(0, 0, 0), V(winW, 0, 0), V(0, winH2, 0)).WithCenterBottom(win.CenterBottom().Add(V(0, 0, -.2)))

	headWalls := c.Or(s).ScaleAt(O, 1.1).AndNot(c.Or(s)).Restrict(
		BoxWithBounds(nil, V(-10, -10, -10), V(10, 10, 0)))
	head := Tree(
		headWalls.AndNot(
			Tree(win,
				win.RotateAt(O, Ey, 35*Deg),
				win.RotateAt(O, Ey, -35*Deg),
				win.RotateAt(O, Ey, 70*Deg),
				win.RotateAt(O, Ey, -70*Deg),
			),
		),
		glass1,
		glass2.RotateAt(O, Ey, 35*Deg),
		glass2.RotateAt(O, Ey, -35*Deg),
		glass3.RotateAt(O, Ey, 70*Deg),
		glass3.RotateAt(O, Ey, -70*Deg),
	)

	_ = head
	vault := Box(wall, W1, W1/2, W1, V(0, W1/4, 0)).AndNot(CylinderZ(wall, .99*W1, 99, V(0, 0, 0))).AndNot(CylinderX(wall, .99*W1, 99, V(0, 0, 0)))

	wallTh := 0.5
	wall1 := Tree(Box(wall, W1, H2, wallTh, V(0, H2/2, -W1/2)).AndNot(Tree(
		win.Translate(V(0, 0, 0)),
		win.Translate(V(-W1/4-winB/2, 0, 0)),
		win.Translate(V(+W1/4+winB/2, 0, 0)),
	)),
		glass1,
		glass2.Translate(V(-W1/4-winB/2, 0, 0)),
		glass3.Translate(V(+W1/4+winB/2, 0, 0)),
	)

	piedestal := Box(wall, 1.7, 0.6, 1.7, O).WithCenterBottom(O)

	//indirect ray goes through window, forgets it's indirect and
	//uses scene.EvalAll instead of NonLuminous
	Render(Spec{
		Recursion:    4,
		Width:        1920 / 2,
		Height:       1920 / 3,
		NumPass:      900,
		DebugNormals: 0,
		Objects: []Object{

			head,

			Tree(
				vault.WithCenterBottom(V(0, H1, W1/2+0*W1)),
				wall1.RotateAt(O, Ey, -90*Deg).WithCenterBottom(V(+W1/2+wallTh/2, 0, W1/2)),
				wall1.RotateAt(O, Ey, +90*Deg).WithCenterBottom(V(-W1/2-wallTh/2, 0, W1/2)),
			),

			piedestal,
			ObjFile(
				map[string]Material{"": marble},
				"/home/arne/assets/Alucy.obj",
				geom.Scale(O, 5./1000.),
				geom.Rotate(O, Ey, 90*Deg),
			).WithCenterBottom(piedestal.CenterTop()).Rotate(Ey, -15*Deg).Remap(
				func(v Vec) Vec { return v },
			),

			Rectangle(
				BlendMap(
					texture.Checkers(4, 4, White, Black),
					Matte(White.EV(-1)),
					Shiny(White.EV(-4), 0.06),
				),
				W1, W1, V(0, .01, 0)),
			Rectangle(Flat(White.EV(-2)), 100, 100, V(0, .001, 0)),
			//Backdrop(Flat(C(0.8, 0.8, 1.0).EV(-0.6))),
			Backdrop(Flat(C(0.8, 0.8, 1.0).EV(1.0))),
			Box(Flat(Black), 6, 6, 0.01, V(0, 3, 4)),
		},
		Lights: []Light{
			SunLight(C(1, 0.9, 0.7).EV(1.1), 0.4*Deg, 3.5*Deg, 4*Deg),
			SunLight(C(1, 0.9, 0.7).EV(0.8), 0.8*Deg, -126*Deg, 18*Deg),
			PointLight(C(1, 1, 1).EV(3.8), V(0.8, 4.5, 1.5)),
		},

		Media: []Medium{
			Fog(0.3, H2, 1),
		},

		Camera: ProjectiveAperture(50*Deg, 0.02, 2.30).Translate(V(0.4, 3.9, 2.8)).YawPitchRoll(0, -1*Deg, 0),

		PostProcess: post.Params{
			Gaussian: post.BloomParams{
				Radius:    0.03,
				Threshold: 1.5,
				Amplitude: 0.02,
			},
			//Airy: post.BloomParams{
			//	Radius:    0.002,
			//	Threshold: 5,
			//	Amplitude: 0.002,
			//},
		},
	})
}
