package builder

import (
	"testing"

	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestCamera(t *testing.T) {
	scene := NewSceneBuilder()

	radius := 0.10

	sph0 := NewSphere(Flat(White), radius)

	sph1 := NewSphere(Flat(Red), radius)
	sph1.Translate(Vec{1, 0, 0})

	sph2 := NewSphere(Flat(Green), radius)
	sph2.Translate(Vec{0, 1, 0})

	sph3 := NewSphere(Flat(Blue), radius)
	sph3.Translate(Vec{0, 0, 1})

	scene.Add(NewTree(sph0, sph1, sph2, sph3))

	scene.Camera.Translate(Vec{Y: 1, Z: -5})

	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
