package sampler

import (
	"math"
	"math/rand"
	"testing"

	colorf "github.com/barnex/bruteray/color"
	"github.com/barnex/bruteray/tracer"
)

func TestStratifierVariance(t *testing.T) {
	f := func(ctx *tracer.Ctx, x, y float64) colorf.Color {
		c := float64(rand.NormFloat64()) + 2
		return colorf.Color{c, c, c}
	}

	ctx := tracer.NewCtx()
	s := NewAdaptive(f, 1, 1, false)
	s.samplePixel(ctx, 0, 0, 1, 1, 100000)
	v := s.variance3(0, 0)
	if math.Abs(float64(v.G-1)) > 0.01 {
		t.Errorf("variance: got %v, want: %v", v, 1)
	}
}
