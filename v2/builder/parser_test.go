package builder

import (
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestBunny(t *testing.T) {
	scene := NewSceneBuilder()
	f := PlyFile(material.Normal2(), "../../assets/bunny_res4.ply")
	scene.Add(f)
	scene.Camera.Translate(Vec{.0, .11, -.3})
	scene.Camera.FocalLen = 1
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
