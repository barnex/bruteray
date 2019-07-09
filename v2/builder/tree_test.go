package builder

import (
	"math/rand"
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestTree(t *testing.T) {
	rng := rand.New(rand.NewSource(123))
	rnd := func() float64 {
		return 2*rng.Float64() - 1
	}

	scene := NewSceneBuilder()
	var tree Tree

	for i := 0; i < 300; i++ {
		s := NewSphere(material.Normal(), (0.1*rng.Float64() + .2))
		s.Texture = material.Normal()
		s.Translate(Vec{rnd(), rnd(), rnd()})
		tree.Add(s)
	}
	//tree.NoDivide=true
	scene.Add(&tree)

	scene.Camera.Translate(Vec{0, 0, -4})
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
