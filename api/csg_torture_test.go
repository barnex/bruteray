package api

import (
	"testing"

	"github.com/barnex/bruteray/tracer/test"
)

func BenchmarkCSG_Torture(b *testing.B) {
	//wall := Matte(White.EV(-1))
	wall := test.WithShadows(White)

	H1 := 6.0
	W1 := 6.0
	H2 := H1 + W1/2

	c := Cylinder(wall, W1, H1, V(0, H1/2, 0))
	s := Sphere(wall, W1, c.CenterTop())

	winW := 1.5
	winB := 1.0
	winH := H1 - winB
	winH2 := H1 - winB + winW

	win := Box(wall, winW, winH, 1, V(0, winH/2+winB, -W1/2)).Or(
		CylinderZ(wall, 2*winW, 1, V(winW/2, winH+winB, -W1/2)).And(
			CylinderZ(wall, 2*winW, 1, V(-winW/2, winH+winB, -W1/2)),
		),
	)

	glassT1 := Yellow
	glassT2 := Magenta
	glassT3 := Cyan

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

	piedestal := Box(wall, 1.7, 1.0, 1.7, O).WithCenterBottom(O)

	//indirect ray goes through window, forgets it's indirect and
	//uses scene.EvalAll instead of NonLuminous
	benchmark(b, Spec{
		Recursion: 3,
		Objects: []Object{

			head,

			Tree(
				vault.WithCenterBottom(V(0, H1, W1/2+0*W1)),
				wall1.RotateAt(O, Ey, -90*Deg).WithCenterBottom(V(+W1/2+wallTh/2, 0, W1/2)),
				wall1.RotateAt(O, Ey, +90*Deg).WithCenterBottom(V(-W1/2-wallTh/2, 0, W1/2)),
			),

			piedestal,

			Rectangle(wall, 100, 100, V(0, .01, 0)),
			Backdrop(Flat(C(0.8, 0.8, 1.0).EV(1))),
		},
		Lights: []Light{
			SunLight(C(1, 0.9, 0.7).EV(1), .8*Deg, 45*Deg, 36*Deg),
			//PointLight(White.EV(4), V(1, H1, 4)),
		},

		Media: []Medium{
			//Fog(0.07, H2),
		},

		Camera: Projective(90*Deg, O, 0, 0).Translate(V(0, 4, 8)),
	})
}
