package asm

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct{ x, ret float64 }{
		{9, 3},
	}
	for i, c := range cases {
		got := Sqrt(c.x)
		if got != c.ret {
			t.Errorf("case %v: %v: got: %v, want: %v", i, c.x, got, c.ret)
		}
	}
}
