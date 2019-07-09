package main

import (
	. "github.com/barnex/bruteray/v2/api"
)

func main() {
	Camera.FocalLen = 1.35
	Camera.Translate(Vec{.278, .273, -.800})
	Recursion = 3
	NumPass = 300
	Postprocess.Bloom.Gaussian.Radius = 0.01
	Postprocess.Bloom.Gaussian.Amplitude = 0.003
	Postprocess.Bloom.Gaussian.Threshold = 0.5

	light := RectangleLight(
		Color{18.4, 15.6, 8.0}.EV(1.3),
		Vec{.213, .548, .227}, Vec{.343, .548, .227}, Vec{.213, .548, .332},
	)
	Add(light)

	x, y, z := .550, .549, .560
	white := Mate(Color{0.708, 0.743, 0.767})
	green := Mate(Color{0.115, 0.441, 0.101})
	red := Mate(Color{0.651, 0.059, 0.061})
	floor := Rectangle(white, Vec{0, 0, 0}, Vec{x, 0, 0}, Vec{0, 0, z})
	Add(
		floor,
		Rectangle(white, Vec{0, y, 0}, Vec{x, y, 0}, Vec{0, y, z}),
		Rectangle(white, Vec{0, 0, z}, Vec{x, 0, z}, Vec{0, y, z}),
		Rectangle(red, Vec{0, 0, 0}, Vec{0, 0, z}, Vec{0, y, 0}),
		Rectangle(green, Vec{x, 0, 0}, Vec{x, 0, z}, Vec{x, y, 0}),
	)

	//bcol := Shiny(White.EV(-2), 0.3)
	bcol := ReflectFresnel(1.6, Mate(White.EV(-1.3)))
	bunny := PlyFile(bcol, "../../assets/bunny_res1.ply")
	Yaw(bunny, bunny.Bounds().Center(), 180*Deg)
	Scale(bunny, bunny.Bounds().Center(), 2.2)
	bbb := bunny.Bounds()
	bunnyC := bbb.Center()
	bunnyC[Y] = bbb.Min[Y]
	floorC := floor.Frame[1].Add(floor.Frame[2]).Mul(0.5)
	Translate(bunny, floorC.Sub(bunnyC))
	Translate(bunny, Vec{0, -.005, .05})
	Add(bunny)

	Render()
}
