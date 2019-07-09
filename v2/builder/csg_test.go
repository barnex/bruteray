package builder

import (
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestAnd(t *testing.T) {
	scene := NewSceneBuilder()

	s1 := NewSphere(material.Normal(), 1)
	Translate(s1, Vec{-.1, 0.0})
	s2 := NewSphere(material.Normal(), 1)
	Translate(s2, Vec{+.1, 0.0})
	and := NewTree(And(s1, s2))
	scene.Add(and)

	scene.Camera.FocalLen = 1
	scene.Camera.Translate(Vec{0, 0, -1.5})
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
