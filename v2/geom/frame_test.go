package geom

import (
	"testing"
)

func TestFrame(t *testing.T) {
	s := 2.0
	f := Frame{{1, 2, 3}, {s, s, 0}, {-s, s, 0}, {0, 0, 1}}
	p := Vec{4, 5, 6}
	p2 := f.TransformToFrame(p)
	p3 := f.TransformToAbsolute(p2)

	if p.Sub(p3).Len() > 1e-6 {
		t.Errorf("have: %v, want: %v", p3, p)
	}
}
