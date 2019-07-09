package builder

import (
	"math"
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
	"github.com/barnex/bruteray/v2/texture"
)

func TestBumpmapConst(t *testing.T) {
	scene := NewSceneBuilder()

	s := NewSphere(material.Normal(), 2)
	s.Translate(Vec{-.5, 0, 4})
	scene.Add(BumpMapped(s, texture.ConstScalar3D(0)))

	scene.Camera.FocalLen = 1

	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}

func TestBumpmapFunc(t *testing.T) {
	scene := NewSceneBuilder()

	s := NewSphere(material.Normal(), 2)
	s.Translate(Vec{-.5, 0, 4})
	scene.Add(BumpMapped(s, texture.ScalarFunc3D(func(v Vec) float64 {
		return math.Sin(v[Y]*20) * 0.01
	})))
	scene.Camera.FocalLen = 1
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}

func TestBumpmapBleed(t *testing.T) {
	t.Skip("TODO")
	scene := NewSceneBuilder()

	s := NewSphere(material.Normal(), 2)
	s.Translate(Vec{-.5, 0, 4})
	scene.Add(BumpMapped(s, texture.ScalarFunc3D(func(v Vec) float64 {
		return math.Sin(v[Y]*40) * 0.05
	})))
	scene.Camera.FocalLen = 1
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
