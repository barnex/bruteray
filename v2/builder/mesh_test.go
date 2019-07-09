package builder

import (
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

func TestBunnyMesh(t *testing.T) {
	scene := NewSceneBuilder()
	scene.Add(PlyFile(material.Normal(), "../assets/bunny_res4.ply"))
	scene.Camera.Translate(Vec{.0, .11, -.3})
	scene.Camera.FocalLen = 1
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}

// Shading normals are different from the real geometric normals,
// which may lead to subtle artifacts when a ray strikes under grazing incidence.
// E.g.: a ray may strike the geometric surface from the front
// but the shading surface from the back.
//
// This test zooms in on an area of a low-resolution mesh
// where such artifacts are visible.
func TestShadingNormals(t *testing.T) {
	t.Skip("TODO")
	scene := NewSceneBuilder()
	model := PlyFile(material.Normal(), "../assets/teapot.ply")
	ScaleToSize(model, 1)
	Pitch(model, model.Bounds().Center(), -90*Deg)
	scene.Add(model)
	scene.Camera.Translate(Vec{0.0, 0.0, -1})
	scene.Camera.FocalLen = 1
	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}
