package builder

import (
	"math"
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestParametric(t *testing.T) {
	scene := NewSceneBuilder()

	s := NewParametric(material.Normal(), 30, 20, func(u, v float64) Vec {
		phi := u * 2 * Pi
		theta := (2*v - 1) * Pi / 2
		x := math.Cos(theta) * math.Cos(phi)
		y := math.Cos(theta) * math.Sin(phi)
		z := math.Sin(theta)
		return Vec{x, y, z}
	})
	Translate(s, Vec{0, 1, 0})

	scene.Add(s)
	scene.Camera.FocalLen = 1
	scene.Camera.Translate(Vec{0, 1, -2.5})

	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
