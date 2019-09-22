package media

import (
	"math"

	. "github.com/barnex/bruteray/tracer/types"
	"github.com/barnex/bruteray/util"
)

func Fog(density float64, height float64) Medium {
	return &fog{
		density: density,
		height:  height,
	}
}

type fog struct {
	density float64
	height  float64
}

// TODO: not correct when camera is in fog
func (m *fog) Filter(ctx *Ctx, s *Scene, r *Ray, tMax float64, orig Color) Color {
	ten, tex := intersectHalfSpace(m.height, r)
	start, end := intersectIntervals(0, tMax, ten, tex)

	tSample := start + ctx.Rng.ExpFloat64()/m.density
	var acc Color
	if tSample < end {
		t := tSample
		p := r.At(t - Tiny)
		sec := ctx.Ray()
		defer ctx.PutRay(sec)

		//acc = Color{1, 1, 1} // RM !!
		for _, l := range s.Lights() {
			lpos, intens := l.Sample(ctx, p)
			if intens == (Color{}) {
				continue
			}

			lDelta := lpos.Sub(p)
			lDir := lDelta.Normalized()
			lDist := lDelta.Len()
			sec.Start = p
			sec.Dir = lDir

			intens = s.Occlude(sec, lDist, intens)
			acc = acc.Add(intens)
		}
	}

	len := end - start
	util.Assert(len >= 0)
	trans := math.Exp(-len * m.density)
	return orig.Mul(trans).Add(acc)
}
