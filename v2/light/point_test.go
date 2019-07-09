package light

import (
	"testing"

	"github.com/barnex/bruteray/v2/builder"
	. "github.com/barnex/bruteray/v2/builder"
	"github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestPoint(t *testing.T) {
	scene := NewSceneBuilder()
	{
		s := builder.NewSheetXZ()
		s.Texture = material.Mate(color.White)
		scene.Add(s)
	}
	{
		s := builder.NewSphere(material.Mate(color.White), 2)
		s.Translate(Vec{0, 1, 0})
		scene.Add(s)
	}
	{
		scene.Add(NewPointLight(color.White.EV(10), Vec{6, 6, -6}))
	}

	scene.Camera.Translate(Vec{.01, 1, -5})
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
