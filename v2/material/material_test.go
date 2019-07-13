package material_test

import (
	"testing"

	. "github.com/barnex/bruteray/v2/builder"
	"github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
	"github.com/barnex/bruteray/v2/texture"
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

func TestRefractive(t *testing.T) {
	t.Skip("TODO")
	scene := NewSceneBuilder()
	//mat := Refractive(1.5)
	mat := Normal()

	sph := NewSphere(mat, 1)
	sph.Translate(Vec{0, 1, 2})
	scene.Add(sph)

	{
	tex := texture.Map(texture.Checkers(1,1, color.White, color.Blue), texture.UVProject{})
	floor := NewSheet(Flat(tex), O, Ex, Ez)
	scene.Add(floor)
	}

	built := scene.Build()
	built.Camera.FocalLen = 1
	built.Camera.Translate(Vec{0,1,-1})
	test.NPass(t, built, 1, 3, test.DefaultTolerance)
}
