package texture

/*
import (
	"testing"

	"github.com/barnex/bruteray/builder"
	. "github.com/barnex/bruteray/tracer/geom"
	"github.com/barnex/bruteray/image"
	"github.com/barnex/bruteray/material"
	"github.com/barnex/bruteray/test"
)

func TestCylinderMap(t *testing.T) {
	scene := builder.NewSceneBuilder()

	uv := &cylinderMap{}
	tex2d := Pan(Nearest(image.MustLoad("../assets/earth.jpg")), 0.7, 0.0)
	tex := Map(tex2d, uv)
	s := builder.NewSphere(material.Flat(tex), 2)
	uv.Bind(s)

	s.Translate(Vec{-.5, 0, 4})
	//Roll(s, s.Origin(), 15*Deg)

	scene.Add(s)
	scene.Camera.FocalLen = 1
	//scene.Camera.Translate(Vec{.5, 0, -4})

	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
*/
