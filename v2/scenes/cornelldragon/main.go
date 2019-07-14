package main

import (
	. "github.com/barnex/bruteray/v2/api"
)

func main() {
	Camera.FocalLen = 1.35
	Camera.Translate(Vec{.278, .273, -.800})
	Recursion = 3
	NumPass = 300
	Postprocess.Bloom.Airy.Radius = 0.005
	Postprocess.Bloom.Airy.Amplitude = 0.055
	Postprocess.Bloom.Airy.Threshold = 0.7
	//Postprocess.Bloom.Star.Radius = 0.03
	//Postprocess.Bloom.Star.Amplitude = 0.03
	//Postprocess.Bloom.Star.Threshold = 0.7

	//Postprocess.Bloom.Gaussian.Radius = 0.01
	//Postprocess.Bloom.Gaussian.Amplitude = 0.02
	//Postprocess.Bloom.Gaussian.Threshold = 2

	light := RectangleLight(
		Color{18.4, 15.6, 8.0}.EV(1.3),
		Vec{.213, .548, .227}, Vec{.343, .548, .227}, Vec{.213, .548, .332},
	)
	Add(light)

	x, y, z := .550, .549, .560
	white := Matte(Color{0.708, 0.743, 0.767})
	green := Matte(Color{0.115, 0.441, 0.101})
	red := Matte(Color{0.651, 0.059, 0.061})
	floor := Rectangle(white, Vec{0, 0, 0}, Vec{x, 0, 0}, Vec{0, 0, z})
	Add(
		floor,
		Rectangle(white, Vec{0, y, 0}, Vec{x, y, 0}, Vec{0, y, z}),
		Rectangle(white, Vec{0, 0, z}, Vec{x, 0, z}, Vec{0, y, z}),
		Rectangle(red, Vec{0, 0, 0}, Vec{0, 0, z}, Vec{0, y, 0}),
		Rectangle(green, Vec{x, 0, 0}, Vec{x, 0, z}, Vec{x, y, 0}),
	)

	//bcol := Shiny(White.EV(-2), 0.3)
	//bcol := ReflectFresnel(4.0, Mate(White.EV(-4)))
	bcol := ReflectFresnel(4.0, Matte(White.EV(-4)))
	//bcol := Normal
	bunny := PlyFile(bcol, "../../assets/dragon_res1.ply")
	Yaw(bunny, bunny.Bounds().Center(), 150*Deg)
	Scale(bunny, bunny.Bounds().Center(), 2.5)
	bbb := bunny.Bounds()
	bunnyC := bbb.Center()
	bunnyC[Y] = bbb.Min[Y]
	floorC := floor.Frame[1].Add(floor.Frame[2]).Mul(0.5)
	Translate(bunny, floorC.Sub(bunnyC))
	Translate(bunny, Vec{0.00, -.005, .05})
	Add(bunny)

	Render()
}
