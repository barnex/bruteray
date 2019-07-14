package main

import (
	. "github.com/barnex/bruteray/v2/api"
)

func main() {
	Camera.FocalLen = 1.2
	Camera.Translate(Vec{.278, .273, -.660})
	Recursion = 6
	NumPass = 100
	Postprocess.Bloom.Gaussian.Radius = 0.01
	Postprocess.Bloom.Gaussian.Amplitude = 0.03
	Postprocess.Bloom.Gaussian.Threshold = 0.5

	light := RectangleLight(
		Color{18.4, 15.6, 8.0}.EV(0.3),
		Vec{.213, .548, .227}, Vec{.343, .548, .227}, Vec{.213, .548, .332},
	)
	//light := Rectangle(
	//	Flat(Color{18.4, 15.6, 8.0}.EV(2)),
	//	Vec{.213, .548, .227}, Vec{.343, .548, .227}, Vec{.213, .548, .332},
	//)
	Scale(light, light.Bounds().Center(), 2.5)
	Add(light)

	x, y, z := .550, .549, .560
	white := Matte(Color{0.708, 0.743, 0.767})
	refl := Reflective(Color{0.8, 0.9, 0.9})
	floor := Rectangle(white, Vec{0, 0, 0}, Vec{x, 0, 0}, Vec{0, 0, z})
	Add(
		floor,
		Rectangle(white, Vec{0, y, 0}, Vec{x, y, 0}, Vec{0, y, z}),
		Rectangle(refl, Vec{0, 0, z}, Vec{x, 0, z}, Vec{0, y, z}),
		Rectangle(refl, Vec{0, 0, 0}, Vec{0, 0, z}, Vec{0, y, 0}),
		Rectangle(refl, Vec{x, 0, 0}, Vec{x, 0, z}, Vec{x, y, 0}),
	)

	//bcol := Shiny(White.EV(-2), 0.3)
	bcol := ReflectFresnel(1.6, Matte(White.EV(-1.0)))
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
