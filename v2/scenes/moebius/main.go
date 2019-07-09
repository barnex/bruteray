package main

import (
	"math"

	. "github.com/barnex/bruteray/v2/api"
	"github.com/barnex/bruteray/v2/builder"
	"github.com/barnex/bruteray/v2/image"
	"github.com/barnex/bruteray/v2/texture"
)

func main() {
	Recursion = 5
	NumPass = 200
	Postprocess.Bloom.Gaussian.Radius = 0.008
	Postprocess.Bloom.Gaussian.Amplitude = 0.5
	Postprocess.Bloom.Gaussian.Threshold = 0.9

	l := RectangleLight(White.EV(5.1), Vec{}, Ex.Mul(5), Ey.Mul(3))
	Translate(l, Vec{1, 3, -2})
	Add(l)

	tex := texture.Pan(
		texture.Nearest(image.MustLoad("../../assets/moebius.png")),
		0, 0,
	)
	mat := Blend(
		0.9, Mate(texture.Map(tex, texture.UVProject{})),
		0.1, Reflective(White),
	)
	strip := Parametric(mat, 256, 16, func(u, v float64) Vec {
		u *= 2 * Pi
		v = 2 * (v - 0.5)
		x := (1 + 0.5*v*cos(0.5*u)) * cos(u)
		y := (1 + 0.5*v*cos(0.5*u)) * sin(u)
		z := 0.5 * v * sin(0.5*u)
		return Vec{x, y, z}
	})
	Pitch(strip, Vec{}, -90*Deg)
	Yaw(strip, Vec{}, 277*Deg)
	Roll(strip, Vec{}, 3*Deg)
	TranslateTo(strip, strip.Bounds().CenterBottom(), Vec{0, -.02, 0})
	Add(strip)

	white := Mate(Color{1, 0.9, 0.8}.EV(-.1))
	floor := Sheet(white, Vec{}, Ex, Ez)
	Add(floor)

	Add(&builder.Ambient{White.EV(-4)})

	Camera.Translate(Vec{0, 3, -2.6}.Mul(0.8))
	Camera.Pitch(-46 * Deg)
	Camera.FocalLen = 1.1
	Camera.Focus = 2.5
	Camera.Aperture = 0.03
	Render()
}

var sin = math.Sin
var cos = math.Cos
