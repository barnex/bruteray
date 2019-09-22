package objects

import (
	"math"
	"testing"

	. "github.com/barnex/bruteray/tracer/types"
)

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
	}

	for i, c := range cases {
		got := (intersectAABB(b, c.ray) > 0)
		if got != c.want {
			t.Errorf("case %v: %v, got: %v, want: %v", i, c.ray, got, c.want)
		}
	}
}

func TestInfBox(t *testing.T) {
	b := infBox

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
		got := (intersectAABB(&b, r) > 0)
		if got != true {
			t.Errorf("case %v: %v, got: %v, want: %v", i, r, got, true)
		}
	}
}

func TestIntersectAABB2(t *testing.T) {
	b := &BoundingBox{Min: Vec{-1, -1, -1}, Max: Vec{1, 1, 1}}

	cases := []struct {
		ray    *Ray
		t1, t2 float64
	}{
		{ray(Vec{0, 0, 0}, Vec{1, 0, 0}), -1, 1},
		{ray(Vec{1, 0, 0}, Vec{1, 0, 0}), -2, 0},
		{ray(Vec{-1, 0, 0}, Vec{-1, 0, 0}), -2, 0},
		{ray(Vec{2, 0, 0}, Vec{1, 0, 0}), -3, -1},
	}

	for i, c := range cases {
		const tol = 1e-8
		got1, got2 := intersectAABB2(b, c.ray)
		if math.Abs(got1-c.t1) > tol || math.Abs(got2-c.t2) > tol {
			t.Errorf("case %v: %v, got: %v,%v want: %v,%v", i, c.ray, got1, got2, c.t1, c.t2)
		}
	}
}

func BenchmarkBoundingboxMiss(b *testing.B) {
	box := &BoundingBox{Min: Vec{-1, -2, -3}, Max: Vec{1, 2, 3}}

	r := ray(Vec{0, -4, 0}, Vec{0, -1, 0})
	for i := 0; i < b.N; i++ {
		intersectAABB(box, r)
	}
}

func ray(start, dir Vec) *Ray {
	return &Ray{Start: start, Dir: dir}
}
