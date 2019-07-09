package builder

import (
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestTriangle(t *testing.T) {
	scene := NewSceneBuilder()

	grid := material.Grid()
	{
		s := NewSheetXZ()
		s.Texture = material.Grid()
		//s.Translate(Vec{0, -1, 0})
		//s.Pitch(-180 * Deg)
		scene.Add(s)
	}
	{
		s := NewTriangle(grid, O, Ex, Ey)
		Translate(s, Vec{0, 0, 2})
		//s.Pitch(-180 * Deg)
		scene.Add(s)
	}
	{
		s := NewTriangle(grid, Vec{}, Vec{-1, 1, 0}, Vec{1, 1, 0})
		Translate(s, Vec{-1, 0, 0})
		scene.Add(s)
	}

	scene.Camera.Translate(Vec{.01, 1.5, -4})
	scene.Camera.Pitch(-10 * Deg)
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
