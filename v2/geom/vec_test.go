package geom

import (
	"math"
	"testing"
)

func TestVec(t *testing.T) {
	check := func(got, want Vec) {
		t.Helper()
		const tol = 1e-6
		fail :=
			(math.Abs(got[X]-want[X]) > tol) ||
				(math.Abs(got[Y]-want[Y]) > tol) ||
				(math.Abs(got[Z]-want[Z]) > tol)
		if fail {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	check(Vec{1, 2, 3}.Add(Vec{4, 5, 6}), Vec{5, 7, 9})
	check(Vec{1, 2, 3}.MAdd(2, Vec{4, 5, 6}), Vec{9, 12, 15})
	check(Vec{1, 2, 3}.Sub(Vec{1, 3, 2}), Vec{0, -1, 1})
	check(Vec{1, 2, 3}.Mul(2), Vec{2, 4, 6})
	check(Vec{2, 4, 6}.Div(2), Vec{1, 2, 3})
	check(Vec{0, 3, 4}.Normalized(), Vec{0, 3. / 5., 4. / 5.})
	check(Vec{0, 0, 0}.Normalized(), Vec{0, 0, 0})
	check(Vec{1, 0, 0}.Cross(Vec{0, 1, 0}), Vec{0, 0, 1})
	check(Vec{0, 1, 0}.Cross(Vec{1, 0, 0}), Vec{0, 0, -1})
	check(Vec{-1, 7, 4}.Cross(Vec{-5, 8, 4}), Vec{-4, -16, 27})
}
