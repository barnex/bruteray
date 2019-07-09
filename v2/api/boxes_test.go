package api

import (
	"testing"

	"github.com/barnex/bruteray/v2/builder"
	"github.com/barnex/bruteray/v2/test"
)

var (
	white = Mate(Color{0.708, 0.743, 0.767})
	green = Mate(Color{0.115, 0.441, 0.101})
	red   = Mate(Color{0.651, 0.059, 0.061})
)

const (
	s = .500
	d = 0.020
	S = s + 2*d
	L = 0.01
)

func TestBoxes(t *testing.T) {
	reset()

	Camera.FocalLen = 1.30
	Camera.Translate(Vec{.278, .4501, -2})

	Camera.Pitch(0 * Deg)
	Recursion = 3
	NumPass = 300
	Postprocess.Bloom.Airy.Radius = 0.004
	Postprocess.Bloom.Airy.Amplitude = 0.04
	Postprocess.Bloom.Airy.Threshold = 0.76
	AntiAlias = false

	{
		box := Tree(box())
		Add(box)
	}

	{
		box := Tree(box())
		Translate(box, Vec{.55, 0, .02})
		Yaw(box, box.Bounds().Min, -18*Deg)
		Add(box)
	}

	{
		box := Tree(box())
		Translate(box, Vec{-.65, .55, .02})
		Yaw(box, box.Bounds().Max, 10*Deg)
		Add(box)
	}

	{
		tree := Tree(box())
		tree.Add(
			Rectangle(white, Vec{.130, .165, .065}, Vec{.082, .165, .225}, Vec{.290, .165, .114}),
			Rectangle(white, Vec{.130, .165, .065}, Vec{.082, .165, .225}, Vec{.290, .165, .114}),
			Rectangle(white, Vec{.290, .000, .114}, Vec{.290, .165, .114}, Vec{.240, .000, .272}),
			Rectangle(white, Vec{.130, .000, .065}, Vec{.130, .165, .065}, Vec{.290, .000, .114}),
			Rectangle(white, Vec{.082, .000, .225}, Vec{.082, .165, .225}, Vec{.130, .000, .065}),
			Rectangle(white, Vec{.240, .000, .272}, Vec{.240, .165, .272}, Vec{.082, .000, .225}),
			Rectangle(white, Vec{.423, .330, .247}, Vec{.265, .330, .296}, Vec{.472, .330, .406}),
			Rectangle(white, Vec{.423, .000, .247}, Vec{.423, .330, .247}, Vec{.472, .000, .406}),
			Rectangle(white, Vec{.472, .000, .406}, Vec{.472, .330, .406}, Vec{.314, .000, .456}),
			Rectangle(white, Vec{.314, .000, .456}, Vec{.314, .330, .456}, Vec{.265, .000, .296}),
			Rectangle(white, Vec{.265, .000, .296}, Vec{.265, .330, .296}, Vec{.423, .000, .247}),
		)
		Translate(tree, Vec{-.55, 0, 0})
		Add(tree)
	}

	{
		tree := Tree(box())
		Translate(tree, Vec{.55, .55, .02})
		Add(tree)
	}

	{
		box := Tree(box())
		Translate(box, Vec{.00, .55, .02})
		Add(box)
	}

	sky := &builder.Ambient{Color{0.5, 0.5, 0.8}.EV(-1)}
	Add(sky)

	floor := Sheet(Mate(White.EV(-2)), Vec{}, Ex, Ez)
	Translate(floor, Vec{0, -d, 0})
	Add(floor)

	test.OnePass(t, scene.Build(), 1)
}

func box() *builder.Tree {
	return Tree(
		Rectangle(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, 0, s}), // floor
		Rectangle(white, Vec{0, s, 0}, Vec{s, s, 0}, Vec{0, s, s}),
		Rectangle(white, Vec{0, 0, s}, Vec{s, 0, s}, Vec{0, s, s}),
		Rectangle(red, Vec{0, 0, 0}, Vec{0, 0, s}, Vec{0, s, 0}),
		Rectangle(green, Vec{s, 0, 0}, Vec{s, 0, s}, Vec{s, s, 0}),

		Rectangle(white, Vec{0, 0, 0}, Vec{S, 0, 0}, Vec{0, 0, S}).Translate(Vec{-d, -d, 0}),
		Rectangle(white, Vec{0, S, 0}, Vec{S, S, 0}, Vec{0, S, S}).Translate(Vec{-d, -d, 0}),
		Rectangle(white, Vec{0, 0, S}, Vec{S, 0, S}, Vec{0, S, S}).Translate(Vec{-d, -d, 0}),
		Rectangle(white, Vec{0, 0, 0}, Vec{0, 0, S}, Vec{0, S, 0}).Translate(Vec{-d, -d, 0}),
		Rectangle(white, Vec{S, 0, 0}, Vec{S, 0, S}, Vec{S, S, 0}).Translate(Vec{-d, -d, 0}),

		Rectangle(white, Vec{-d, s, 0}, Vec{-d, s + d, 0}, Vec{s + d, s, 0}),
		Rectangle(white, Vec{-d, s, 0}, Vec{-d, 0, 0}, Vec{0, s, 0}),
		Rectangle(white, Vec{s, s, 0}, Vec{s, 0, 0}, Vec{s + d, s, 0}),
		Rectangle(white, Vec{-d, 0, 0}, Vec{-d, 0 - d, 0}, Vec{s + d, 0, 0}),

		RectangleLight(Color{18.4, 15.6, 8.0}.EV(8), Vec{s/2 - L, s - .001, s/2 - L}, Vec{s/2 + L, s - .001, s/2 - L}, Vec{s/2 - L, s - .001, s/2 + L}),
	)
}
