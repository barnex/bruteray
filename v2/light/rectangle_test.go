package light

import (
	"testing"

	"github.com/barnex/bruteray/v2/builder"
	"github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
)

// This tests a rectangle light by comparing to a path-traced equivalent.
// The ray-traced test uses recursion depth 1, to exclude indirect lighting.
// The golden data replaces the light source by a flat white rectangle of
// the same brightness, and uses recursion depth 2.
func TestRectangleLight(t *testing.T) {
	t.Skip("slow")
	scene := builder.NewSceneBuilder()
	scene.Camera.FocalLen = 1.35
	scene.Camera.Translate(Vec{.278, .3, -1.0})

	white := material.Mate(Color{1, 1, 1}.EV(-1))
	//white := tracer.Normal2()
	s := .500
	L := .250
	dy := 0.0001

	box := builder.NewTree(
		builder.NewRectangle(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, 0, s}),
		builder.NewRectangle(white, Vec{0, s, 0}, Vec{s, s, 0}, Vec{0, s, s}),
		builder.NewRectangle(white, Vec{0, 0, s}, Vec{s, 0, s}, Vec{0, s, s}),
		builder.NewRectangle(white, Vec{0, 0, 0}, Vec{0, 0, s}, Vec{0, s, 0}),
		builder.NewRectangle(white, Vec{s, 0, 0}, Vec{s, 0, s}, Vec{s, s, 0}),
	)

	lc := color.White.EV(2)
	light := NewRectangleLight(lc, Vec{s/2 - L, s - dy, s/2 - L}, Vec{s/2 + L, s - dy, s/2 - L}, Vec{s/2 - L, s - dy, s/2 + L})
	//light := builder.NewRectangle(tracer.Flat(lc), Vec{s/2 - L, s - dy, s/2 - L}, Vec{s/2 + L, s - dy, s/2 - L}, Vec{s/2 - L, s - dy, s/2 + L})
	scene.Add(box, light)
	test.NPass(t, scene.Build(), 1, 10000, 0.5)
}
