package texture

import (
	"math"
	"testing"
)

func TestWarp(t *testing.T) {
	cases := []struct{ input, want float64 }{
		{-1000, 0.0},
		{-8.1, 0.9},
		{-1.1, 0.9},
		{-2.0, 0.0},
		{-1.0, 0.0},
		{-0.9, 0.1},
		{-0.1, 0.9},
		{-0.0, 0.0},
		{0.0, 0.0},
		{0.1, 0.1},
		{0.9, 0.9},
		{1.0, 0.0},
		{1.1, 0.1},
		{1.9, 0.9},
		{2.0, 0.0},
		{8.9, 0.9},
		{1000.0, 0.0},
	}
	for _, c := range cases {
		got := warp(c.input)
		if math.Abs(got-c.want) > 1e-6 {
			t.Errorf("warp(%v): got: %v, want: %v", c.input, got, c.want)
		}
	}
}
