package builder

import (
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestSheet(t *testing.T) {
	scene := NewSceneBuilder()

	{
		s := NewSheet(material.Normal(), O, Ez, Ex)
		s.Translate(Vec{0, -1, 0})
		scene.Add(s)
	}
	{
		s := NewSheet(material.Normal(), O, Ez, Ex)
		s.Translate(Vec{0, 1, 0})
		scene.Add(s)
	}
	{
		s := NewSheet(material.Normal(), O, Ey, Ez)
		s.Translate(Vec{-1, 0, 0})
		scene.Add(s)
	}
	{
		s := NewSheet(material.Normal(), O, Ey, Ez)
		s.Translate(Vec{1, 0, 0})
		scene.Add(s)
	}
	{
		s := NewSheet(material.Normal(), O, Ex, Ey)
		s.Translate(Vec{0, 0, -1})
		scene.Add(s)
	}

	built := scene.Build()
	built.Camera.FocalLen = 1
	(&built.Camera).Translate(Vec{.01, .01, -0.1})
	test.OnePass(t, built, test.DefaultTolerance)
}
