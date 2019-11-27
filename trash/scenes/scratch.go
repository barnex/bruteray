// +build example

package main

import (
	//"github.com/barnex/bruteray/imagef/colorf"
	//"github.com/barnex/bruteray/tracer/cameras"
	//"github.com/barnex/bruteray/tracer/lights"
	//"github.com/barnex/bruteray/tracer/materials"
	//"github.com/barnex/bruteray/tracer/objects"
	. "github.com/barnex/bruteray/api"
	"github.com/barnex/bruteray/x"
)

func main() {
	s := .500
	L := .150
	dy := 0.0001
	white := Matte(Color{1, 1, 1}.EV(-.3))

	x.Display(Spec{
		Lights: []Light{
			RectangleLight(White.EV(4), L, L, Vec{s / 2, s - dy, s / 2}),
		},
		Objects: []Object{
			RectangleWithVertices(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, 0, s}),
			RectangleWithVertices(white, Vec{0, s, 0}, Vec{s, s, 0}, Vec{0, s, s}),
			RectangleWithVertices(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, s, 0}),
			RectangleWithVertices(white, Vec{0, 0, 0}, Vec{0, 0, s}, Vec{0, s, 0}),
			RectangleWithVertices(white, Vec{s, 0, 0}, Vec{s, 0, s}, Vec{s, s, 0}),
			PlyFile(white, "../assets/bunny_res4.ply"),
		},
		Camera:    Projective(60*Deg, Vec{0.25, 0.25, 0.9}, 0, 0*Deg, 0),
		NumPass:   100,
		Width:     600,
		Height:    600,
		Recursion: 3,
	})
}
