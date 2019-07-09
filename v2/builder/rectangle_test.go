package builder

import (
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestRectangle(t *testing.T) {
	scene := NewSceneBuilder()

	grid := material.Grid()
	{
		s := NewRectangle(grid, O, Ex, Ey)
		s.Translate(Vec{0, 1, 0})
		s.Scale3(3, 2, 1)
		s.Pitch(-180 * Deg)
		scene.Add(s)
	}
	{
		s := NewRectangle(grid, O, Ex, Ey)
		s.Translate(Vec{-1, -0.5, 0})
		s.Scale(2)
		s.Pitch(-90 * Deg)
		scene.Add(s)
	}

	scene.Camera.Translate(Vec{.01, 1, -10})
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}

func TestRectangle_Ctrl(t *testing.T) {
	scene := NewSceneBuilder()
	//TODO: test bounding
	{
		s := NewSheetXZ()
		s.Texture = material.Grid()
		scene.Add(s)
	}
	{
		s := NewRectangle(material.Grid(), O, Ex, Ey)
		s.Frame[0] = Vec{0, 0, 0}
		s.Frame[1] = Vec{1, 0, 0}
		s.Frame[2] = Vec{0, 0, 1}
		s.Translate(Vec{0, 0.001, 0})

		scene.Add(s)
	}

	scene.Camera.Translate(Vec{.01, 3, -7})
	scene.Camera.Pitch(-20 * Deg)
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
