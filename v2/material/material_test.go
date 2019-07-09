package material_test

import (
	"testing"

	. "github.com/barnex/bruteray/v2/builder"
	"github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestFlat(t *testing.T) {
	scene := NewSceneBuilder()

	sph := NewSphere(Flat(color.Red), 1)
	sph.Translate(Vec{0, 0, 2})
	scene.Add(sph)

	built := scene.Build()
	built.Camera.FocalLen = 1
	test.OnePass(t, built, test.DefaultTolerance)
}

func TestNormal(t *testing.T) {
	scene := NewSceneBuilder()

	sph := NewSphere(Normal(), 1)
	sph.Translate(Vec{0, 0, 2})
	scene.Add(sph)

	built := scene.Build()
	built.Camera.FocalLen = 1
	test.OnePass(t, built, test.DefaultTolerance)
}
