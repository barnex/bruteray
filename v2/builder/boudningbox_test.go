package builder

import (
	"fmt"
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
)

func ExampleBoundingBox() {
	s := NewSphere(nil, 1)
	s.Translate(Vec{3, 0, 0})
	fmt.Println(MakeBoundingBox([]Builder{s}))

	//Output:
	//{[2.5 -0.5 -0.5] [3.5 0.5 0.5]}
}

func TestBoundingbox_Intersect(t *testing.T) {
	b := &BoundingBox{Min: Vec{-1, -2, -3}, Max: Vec{1, 2, 3}}

	cases := []struct {
		ray  *Ray
		want bool
	}{
		{ray(Vec{0, 0, 0}, Vec{1, 0, 0}), true},
		{ray(Vec{0, 0, 0}, Vec{0, 1, 0}), true},
		{ray(Vec{0, 0, 0}, Vec{0, 0, 1}), true},
		{ray(Vec{0, 0, 0}, Vec{-1, 0, 0}), true},
		{ray(Vec{0, 0, 0}, Vec{0, -1, 0}), true},
		{ray(Vec{0, 0, 0}, Vec{0, 0, -1}), true},
		{ray(Vec{2, 0, 0}, Vec{+1, 0, 0}), false},
		{ray(Vec{2, 0, 0}, Vec{-1, 0, 0}), true},
		{ray(Vec{2, 0, 0}, Vec{+1, 0, 0}), false},
		{ray(Vec{2, 0, 0}, Vec{-1, 0, 0}), true},
		{ray(Vec{-2, 0, 0}, Vec{+1, 0, 0}), true},
		{ray(Vec{-2, 0, 0}, Vec{-1, 0, 0}), false},
		{ray(Vec{0, 4, 0}, Vec{+1, 0, 0}), false},
		{ray(Vec{0, 4, 0}, Vec{-1, 0, 0}), false},
		{ray(Vec{0, 4, 0}, Vec{0, +1, 0}), false},
		{ray(Vec{0, 4, 0}, Vec{0, -1, 0}), true},
		{ray(Vec{0, -4, 0}, Vec{0, +1, 0}), true},
		{ray(Vec{0, -4, 0}, Vec{0, -1, 0}), false},
		{ray(Vec{0, 0, 6}, Vec{0, 0, 1}), false},
		{ray(Vec{0, 0, 6}, Vec{0, 0, -1}), true},
		{ray(Vec{0, 0, -6}, Vec{0, 0, 1}), true},
		{ray(Vec{0, 0, -6}, Vec{0, 0, -1}), false},
		{ray(Vec{99, 99, 99}, Vec{99, 99, 99}), false},

		// TODO: fininte-length rays

	}

	for i, c := range cases {
		got := (b.Intersect(c.ray) > 0)
		if got != c.want {
			t.Errorf("case %v: %v, got: %v, want: %v", i, c.ray, got, c.want)
		}
	}
}

func TestInfBox(t *testing.T) {
	b := infBox()

	cases := []*Ray{
		ray(Vec{0, 0, 0}, Vec{1, 0, 0}),
		ray(Vec{0, 0, 0}, Vec{0, 1, 0}),
		ray(Vec{0, 0, 0}, Vec{0, 0, 1}),
		ray(Vec{0, 0, 0}, Vec{-1, 0, 0}),
		ray(Vec{0, 0, 0}, Vec{0, -1, 0}),
		ray(Vec{0, 0, 0}, Vec{0, 0, -1}),
		ray(Vec{2, 0, 0}, Vec{+1, 0, 0}),
		ray(Vec{2, 0, 0}, Vec{-1, 0, 0}),
		ray(Vec{2, 0, 0}, Vec{+1, 0, 0}),
		ray(Vec{2, 0, 0}, Vec{-1, 0, 0}),
		ray(Vec{-2, 0, 0}, Vec{+1, 0, 0}),
		ray(Vec{-2, 0, 0}, Vec{-1, 0, 0}),
		ray(Vec{0, 4, 0}, Vec{+1, 0, 0}),
		ray(Vec{0, 4, 0}, Vec{-1, 0, 0}),
		ray(Vec{0, 4, 0}, Vec{0, +1, 0}),
		ray(Vec{0, 4, 0}, Vec{0, -1, 0}),
		ray(Vec{0, -4, 0}, Vec{0, +1, 0}),
		ray(Vec{0, -4, 0}, Vec{0, -1, 0}),
		ray(Vec{0, 0, 6}, Vec{0, 0, 1}),
		ray(Vec{0, 0, 6}, Vec{0, 0, -1}),
		ray(Vec{0, 0, -6}, Vec{0, 0, 1}),
		ray(Vec{0, 0, -6}, Vec{0, 0, -1}),
		ray(Vec{99, 99, 99}, Vec{99, 99, 99}),
	}

	for i, r := range cases {
		got := (b.Intersect(r) > 0)
		if got != true {
			t.Errorf("case %v: %v, got: %v, want: %v", i, r, got, true)
		}
	}
}

func BenchmarkBoundingboxMiss(b *testing.B) {
	box := &BoundingBox{Min: Vec{-1, -2, -3}, Max: Vec{1, 2, 3}}

	r := ray(Vec{0, -4, 0}, Vec{0, -1, 0})
	for i := 0; i < b.N; i++ {
		box.Intersect(r)
	}
}

func ray(start, dir Vec) *Ray {
	return &Ray{Start: start, Dir: dir, Len: Inf}
}
