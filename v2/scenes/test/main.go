package main

import (
	. "github.com/barnex/bruteray/v2/api"
	"github.com/barnex/bruteray/v2/assets"
)

func main() {
	Camera.FocalLen = 1
	Camera.Translate(Vec{0.5, 0.5, -1.2})
	Camera.Pitch(2 * Deg)
	Recursion = 3
	NumPass = 1
	Postprocess.Bloom.Airy.Radius = 0.002
	Postprocess.Bloom.Airy.Amplitude = 0.09
	Postprocess.Bloom.Airy.Threshold = 0.76

	Add(assets.Box(Normal, 0.05))

	Render()
}
