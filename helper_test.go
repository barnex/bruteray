package main

import (
	"math"
	"testing"
)

func testVec(t *testing.T, have, want Vec, tol float64) {
	// go 1.9beta hack, TODO: remove
	var ti interface{} = t
	if h, ok := ti.(interface {
		Helper()
	}); ok {
		h.Helper()
	}

	fail :=
		(math.Abs(have.X-want.X) > tol) ||
			(math.Abs(have.Y-want.Y) > tol) ||
			(math.Abs(have.Z-want.Z) > tol)
	if fail {
		t.Errorf("have %v, want %v", have, want)
	}
}

func testFloat(t *testing.T, have, want float64, tol float64) {
	// go 1.9beta hack, TODO: remove
	var ti interface{} = t
	if h, ok := ti.(interface {
		Helper()
	}); ok {
		h.Helper()
	}

	if math.Abs(have-want) > tol {
		t.Errorf("have %v, want %v", have, want)
	}
}

//func testInter(t *testing.T, want Inter, wok bool) func(have Inter, hok bool) {
//
//	// go 1.9beta hack, TODO: remove
//	var ti interface{} = t
//	if h, ok := ti.(interface {
//		Helper()
//	}); ok {
//		h.Helper()
//	}
//
//	const tol = 1e-6
//	return func(have Inter, hok bool) {
//	var ti interface{} = t
//	if h, ok := ti.(interface {
//		Helper()
//	}); ok {
//		h.Helper()
//	}
//		fail :=
//			(math.Abs(have.Min-want.Min) > tol) ||
//				(math.Abs(have.Max-want.Max) > tol) ||
//				hok != wok
//		if fail {
//			t.Errorf("have %v,%v want %v,%v", have, hok, want, wok)
//		}
//	}
//}
