package main

import (
	. "github.com/barnex/bruteray/v2/api"
	"github.com/barnex/bruteray/v2/builder"
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
	L = 0.075
)

func main() {
	Camera.FocalLen = 1.30
	Camera.Translate(Vec{.278, .45, -1.80})
	Camera.Pitch(0 * Deg)
	Recursion = 3
	NumPass = 300
	Postprocess.Bloom.Airy.Radius = 0.004
	Postprocess.Bloom.Airy.Amplitude = 0.04
	Postprocess.Bloom.Airy.Threshold = 0.76

	floorCenter := Vec{0.25, 0, 0.25}
	{
		box := Tree(box())
		col := ReflectFresnel(4.0, Mate(White.EV(-4)))
		dragon := PlyFile(col, "../../assets/dragon_res1.ply")
		Yaw(dragon, dragon.Bounds().Center(), 150*Deg)
		Scale(dragon, dragon.Bounds().Center(), 2.2)
		TranslateTo(dragon, dragon.Bounds().CenterBottom(), floorCenter)
		Translate(dragon, Vec{0, -0.005, -0.05})
		box.Add(dragon)
		Add(box)
	}

	{

		// TODO: bug
		//box2 := box()
		//Scale(box2, floorCenter, 0.5)
		//Translate(box2, Vec{0, d/2, 0})
		//box.Add(box2)

		box := Tree(box())

		s1 := Sphere(ReflectFresnel(9, Flat(Black)), 0.19)
		TranslateTo(s1, s1.Bounds().CenterBottom(), floorCenter.Add(Vec{-.1, 0, .1}))
		s2 := Sphere(Shiny(White.EV(-5), 0.1), 0.15)
		TranslateTo(s2, s2.Bounds().CenterBottom(), floorCenter.Add(Vec{.1, 0, -.1}))
		box.Add(s1, s2)

		Translate(box, Vec{.55, 0, .02})
		Yaw(box, box.Bounds().Min, -18*Deg)
		Add(box)
	}

	{
		box := Tree(box())

		bcol := ReflectFresnel(1.6, Mate(White.EV(-1.6)))
		bunny := PlyFile(bcol, "../../assets/bunny_res1.ply")
		Yaw(bunny, bunny.Bounds().Center(), 180*Deg)
		Scale(bunny, bunny.Bounds().Center(), 2.2)
		TranslateTo(bunny, bunny.Bounds().CenterBottom(), floorCenter)
		Translate(bunny, Vec{0, -.005, .05})

		box.Add(bunny)
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

		col := Mate(White.EV(-.6))
		skull := PlyFile(col, "../../assets/damaliscus.ply")
		Scale(skull, skull.Bounds().Center(), 0.0010)
		Yaw(skull, skull.Bounds().Center(), 70*Deg)
		TranslateTo(skull, skull.Bounds().CenterBottom(), floorCenter)
		Translate(skull, Vec{0.072, 0, 0.05})
		tree.Add(skull)

		Translate(tree, Vec{.55, .55, .02})
		Add(tree)
	}

	{
		box := Tree(box())

		bcol := ReflectFresnel(1.6, Mate(White.EV(-1.0)))
		teapot := PlyFile(bcol, "../../assets/teapot.ply")
		Scale(teapot, teapot.Bounds().Center(), 0.027)
		Pitch(teapot, teapot.Bounds().Center(), -90*Deg)
		Yaw(teapot, teapot.Bounds().Center(), -20*Deg)
		TranslateTo(teapot, teapot.Bounds().CenterBottom(), floorCenter)
		Translate(teapot, Vec{0, -.000, .09})

		box.Add(teapot)
		Translate(box, Vec{.00, .55, .02})
		Add(box)
	}

	sky := &builder.Ambient{Color{0.5, 0.5, 0.8}.EV(-1)}
	Add(sky)

	floor := Sheet(Mate(White.EV(-2)), Vec{}, Ex, Ez)
	Translate(floor, Vec{0, -d, 0})
	Add(floor)

	Render()
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

		RectangleLight(Color{18.4, 15.6, 8.0}.EV(0.3), Vec{s/2 - L, s - .001, s/2 - L}, Vec{s/2 + L, s - .001, s/2 - L}, Vec{s/2 - L, s - .001, s/2 + L}),
	)
}
