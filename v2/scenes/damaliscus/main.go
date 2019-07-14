package main

import (
	. "github.com/barnex/bruteray/v2/api"
)

func main() {
	Camera.FocalLen = 1.35
	Camera.Translate(Vec{.278, .273, -.800})
	Recursion = 4
	NumPass = 300
	Postprocess.Bloom.Gaussian.Radius = 0.01
	Postprocess.Bloom.Gaussian.Amplitude = 0.0005
	Postprocess.Bloom.Gaussian.Threshold = 0.5

	light := RectangleLight(
		Color{18.4, 15.6, 8.0}.EV(1.8),
		Vec{.213, .548, .227}, Vec{.343, .548, .227}, Vec{.213, .548, .332},
	)
	Add(light)

	x, y, z := .550, .549, .560
	white := Matte(Color{0.708, 0.743, 0.767}.EV(-1))
	green := Matte(Color{0.115, 0.441, 0.101}.EV(-1))
	red := Matte(Color{0.651, 0.059, 0.061}.EV(-1))
	floor := Rectangle(white, Vec{0, 0, 0}, Vec{x, 0, 0}, Vec{0, 0, z})
	Add(
		floor,
		Rectangle(white, Vec{0, y, 0}, Vec{x, y, 0}, Vec{0, y, z}),
		Rectangle(white, Vec{0, 0, z}, Vec{x, 0, z}, Vec{0, y, z}),
		Rectangle(red, Vec{0, 0, 0}, Vec{0, 0, z}, Vec{0, y, 0}),
		Rectangle(green, Vec{x, 0, 0}, Vec{x, 0, z}, Vec{x, y, 0}),
	)

	bcol := Matte(White.EV(-.6))
	bunny := PlyFile(bcol, "../../assets/damaliscus.ply")
	Scale(bunny, bunny.Bounds().Center(), 0.0010)
	bbb := bunny.Bounds()
	Yaw(bunny, bunny.Bounds().Center(), 70*Deg)
	bunnyC := bbb.Center()
	floorC := floor.Frame[1].Add(floor.Frame[2]).Mul(0.5)
	Translate(bunny, floorC.Sub(bunnyC))

	//bunnyC[Y] = bbb.Min[Y]
	Translate(bunny, Vec{0.075, 0.23, 0.05})
	Add(bunny)

	Render()
}
