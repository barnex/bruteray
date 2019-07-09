package builder

import (
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

// TODO: test occlusion, test respect ray Len, benchmark.
func TestNode(t *testing.T) {
	scene := NewSceneBuilder()
	//scene.Camera.FocalLen = 1

	sph1 := NewSphere(material.Normal(), 1)
	sph1.Translate(Vec{-1, 0, 4})

	sph2 := NewSphere(material.Normal(), 1)
	sph2.Translate(Vec{+1, 0, 4})

	scene.Add(NewTree(sph1, sph2))

	built := scene.Build()
	built.Camera.FocalLen = 1
	test.OnePass(t, built, test.DefaultTolerance)
}
