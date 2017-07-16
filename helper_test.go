package bruteray

import (
	"math"
	"reflect"
	"testing"
)

type helper struct {
	*testing.T
}

func Helper(tst *testing.T) helper {
	tst.Parallel()
	return helper{tst}
}

func (t helper) Eq(a, b interface{}) {
	t.Helper()

	if !reflect.DeepEqual(a, b) {
		t.Errorf("have: %v, want: %v", a, b)
	}
}

func (t helper) EqVec(have, want Vec) {
	const tol = 1e-6
	fail :=
		(math.Abs(have[X]-want[X]) > tol) ||
			(math.Abs(have[Y]-want[Y]) > tol) ||
			(math.Abs(have[Z]-want[Z]) > tol)
	if fail {
		t.Errorf("have %v, want %v", have, want)
	}
}
