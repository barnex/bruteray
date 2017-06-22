package main

import "testing"

func Benchmark1(b *testing.B) {

	s := &Scene{}
	const h = 2
	ground := Diffuse1(s, Slab(-h, -h-100), 0.5)
	sp := Sphere(Vec{-0.5, -1, 8}, 2)
	die := &ShapeAnd{sp, Slab(-h+.2, -.2)}
	dice := Diffuse1(s, die, 0.95)
	s.objs = []Obj{
		ground,
		dice,
		Reflective(s, Sphere(Vec{3, -1, 10}, 1), 0.9),
	}
	s.sources = []Source{
		&PointSource{Vec{6, 10, 2}, 180},
	}
	s.amb = func(Vec) Color { return 1 }

	const (
		W = 1920
		H = 1080
	)
	cam := Camera(W, H, 1)
	cam.Transf = RotX(-10 * deg)

	b.ResetTimer()
	b.SetBytes((W + 1) * (H + 1)) // actually pixels
	for i := 0; i < b.N; i++ {
		cam.Render(s)
	}
}
