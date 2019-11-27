package main

import (
	. "github.com/barnex/bruteray/api"
	"github.com/barnex/bruteray/imagef/post"
)

func main() {

	porcelain := Shiny(White.EV(-.3), 0.05)
	picnic := LoadTexture("assets/picnicbw.png").Scale(0.05, 0.05)
	//picnic := White.EV(-.3)
	table := IsoSurface(Matte(picnic), 1.5, 0.0007, 1.5, LoadHeightMap("assets/canvas.png").Scale(0.25, 0.25).HeightMap()).WithCenterBottom(O)
	//table := Rectangle(Matte(picnic), 1.5, 1.2, O)
	teapot := PlyFile(porcelain, "assets/teapot.ply").ScaleToSize(0.38).RotateAt(O, Ex, -90*Deg)
	mug := PlyFile(porcelain, "assets/nescafemug.ply").ScaleToSize(0.11).Rotate(Ex, -90*Deg)

	Render(Spec{
		Recursion:    3,
		NumPass:      300,
		Width:        1920 / 1,
		Height:       1080 / 1,
		DebugNormals: 0,

		Lights: []Light{
			RectangleLight(C(18.4, 15.6, 8.0).EV(0.6), 0.8, 0.8, V(-2, 2.8, 0.1)),
			RectangleLight(C(18.4, 15.6, 8.0).EV(1.6), 0.3, 0.3, V(0, 1.5, -2.1)),
		},
		Objects: []Object{
			mug.WithCenterBottom(V(0.10, 0, 0)).Rotate(Ey, 25*Deg),
			mug.WithCenterBottom(V(0.23, 0, -0.04)).Rotate(Ey, 30*Deg),
			teapot.WithCenterBottom(V(-0.17, 0, 0.14)).Rotate(Ey, 7*Deg),
			table.WithCenterBottom(V(0, 0, 0.35)),
			box().Rotate(Ey, 180*Deg).WithCenterBottom(V(0, -1, 0)),
			Backdrop(Flat(C(0.95, 0.95, 1).EV(-2.9))), // TODO: segfault without
		},
		Camera: ProjectiveAperture(80*Deg, 0.004, 0.70).Translate(V(0, 0.25, 0.5)).YawPitchRoll(6*Deg, -24*Deg, 0),

		PostProcess: post.Params{
			Gaussian: post.BloomParams{
				Radius:    0.02,
				Threshold: 0.6,
				Amplitude: 1,
			},
			Airy: post.BloomParams{
				Radius:    0.001,
				Threshold: 0.8,
				Amplitude: 0.5,
			},
			//Star: post.BloomParams{
			//	Radius:    0.1,
			//	Threshold: 1,
			//	Amplitude: 11,
			//},
		},
	})
}

func box() Object {

	const (
		s = .500 * 10
		d = 0.020
		S = s + 2*d
		L = 0.18
	)

	var (
		white = Matte(Color{0.708, 0.743, 0.767}.EV(-1))
		green = Matte(Color{0.415, 0.541, 0.401}.EV(-1.6))
		red   = Matte(Color{0.651, 0.059, 0.061})
	)

	return Tree(
		//RectangleWithVertices(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, 0, s}), // floor
		//RectangleWithVertices(white, Vec{0, s, 0}, Vec{s, s, 0}, Vec{0, s, s}),
		RectangleWithVertices(white, Vec{0, 0, s}, Vec{s, 0, s}, Vec{0, s, s}),
		RectangleWithVertices(red, Vec{0, 0, 0}, Vec{0, 0, s}, Vec{0, s, 0}),
		RectangleWithVertices(green, Vec{s, 0, 0}, Vec{s, 0, s}, Vec{s, s, 0}),

		//RectangleWithVertices(white, Vec{0 - d, 0 - d, 0}, Vec{S - d, 0 - d, 0}, Vec{0 - d, 0 - d, S}),
		//RectangleWithVertices(white, Vec{0 - d, S - d, 0}, Vec{S - d, S - d, 0}, Vec{0 - d, S - d, S}),
		//RectangleWithVertices(white, Vec{0 - d, 0 - d, S}, Vec{S - d, 0 - d, S}, Vec{0 - d, S - d, S}),
		//RectangleWithVertices(white, Vec{0 - d, 0 - d, 0}, Vec{0 - d, 0 - d, S}, Vec{0 - d, S - d, 0}),
		//RectangleWithVertices(white, Vec{S - d, 0 - d, 0}, Vec{S - d, 0 - d, S}, Vec{S - d, S - d, 0}),

		//RectangleWithVertices(white, Vec{-d, s, 0}, Vec{-d, s + d, 0}, Vec{s + d, s, 0}),
		//RectangleWithVertices(white, Vec{-d, s, 0}, Vec{-d, 0, 0}, Vec{0, s, 0}),
		//RectangleWithVertices(white, Vec{s, s, 0}, Vec{s, 0, 0}, Vec{s + d, s, 0}),
		//RectangleWithVertices(white, Vec{-d, 0, 0}, Vec{-d, 0 - d, 0}, Vec{s + d, 0, 0}),
	)
}
