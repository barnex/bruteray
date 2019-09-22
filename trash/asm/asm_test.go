package asm

import (
	"math"
	"testing"
)

func BenchmarkMathSqrt(b *testing.B) {
	var a float64
	for i := 0; i < b.N; i++ {
		a = math.Sqrt(a)
	}
	println(a)
}

func BenchmarkMathSqrtAsm(b *testing.B) {
	var a float64
	for i := 0; i < b.N; i++ {
		a = Sqrt(a)
	}
	println(a)
}

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
