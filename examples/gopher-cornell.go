// +build example_gopher_obj

package main

import (
	. "github.com/barnex/bruteray/api"
)

const (
	s = .500
	d = 0.020
	S = s + 2*d
	L = 0.18
)

func main() {

	gopher := ObjFile(map[string]Material{
		"Body":         Matte(C(0.5, 0.5, 1)),
		"SkinColor":    Matte(C(1, 0.5, 0.5)),
		"Tooth":        Shiny(White, 0.1),
		"Eye-White":    Shiny(White, 0.1),
		"NoseTop":      Shiny(C(1, 0.5, 0.5), 0.02),
		"Material":     Shiny(Gray(0.5), 0.1),
		"Material.001": Shiny(Gray(0.0), 0.1),
	}, "/home/arne/assets/gopher.obj")
	Render(Spec{
		//DebugNormals:      true,
		//DebugIsometricFOV: 8,
		//DebugIsometricDir: Z,

		Recursion: 3,
		NumPass:   40,
		Width:     1024 / 4,
		Height:    1024 / 4,

		Lights: []Light{
			//	SunLight(White, 2*Deg, -120*Deg, 60*Deg),
			RectangleLight(Color{18.4, 15.6, 8.0}.EV(-1), L, L, V(0, s-.001, 0)),
		},
		Objects: []Object{
			Backdrop(Flat(White.EV(-3))),
			//Rectangle(Matte(White.EV(-1)), 10, 10, O),
			box().WithCenterBottom(V(0, 0, 0)).Rotate(Ey, 180*Deg),
			gopher.ScaleToSize(0.28).Rotate(Ey, -80*Deg).WithCenterBottom(V(-0.04, 0, -0.045)),
			//gopher.ScaleToSize(0.28).Rotate(Ey, -80*Deg).WithCenterBottom(V(0.04, 0, 0.1)),
		},
		Camera: Projective(45*Deg, V(0, 0.25, 0.88), 0, 0),
	})
}

func box() Object {

	var (
		white = Matte(Color{0.708, 0.743, 0.767})
		green = Matte(Color{0.115, 0.441, 0.101})
		red   = Matte(Color{0.651, 0.059, 0.061})
	)

	return Tree(
		RectangleWithVertices(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, 0, s}), // floor
		RectangleWithVertices(white, Vec{0, s, 0}, Vec{s, s, 0}, Vec{0, s, s}),
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
