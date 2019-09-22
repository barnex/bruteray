package media

import (
	"math"

	. "github.com/barnex/bruteray/tracer/types"
	"github.com/barnex/bruteray/util"
)

func ExpFog(density float64, ambient Color, height float64) Medium {
	return &expFog{
		density: density,
		ambient: ambient,
		height:  height,
	}
}

type expFog struct {
	density float64
	ambient Color
	height  float64
}

// TODO: not correct when camera is in fog
func (m *expFog) Filter(ctx *Ctx, s *Scene, r *Ray, tMax float64, orig Color) Color {
	ten, tex := intersectHalfSpace(m.height, r)
	start, end := intersectIntervals(0, tMax, ten, tex)
	len := end - start
	util.Assert(len >= 0)
	trans := math.Exp(-len * m.density)
	return orig.Mul(trans).MAdd(1-trans, m.ambient)
}

func intersectIntervals(start1, end1, start2, end2 float64) (start, end float64) {
	start = util.Max(start1, start2)
	end = util.Min(end1, end2)
	if !(end > start) { // handles NaN
		return 0, 0
	}
	return start, end
}

func intersectHalfSpace(height float64, r *Ray) (ten, tex float64) {
	rs := r.Start[1] - height
	rd := r.Dir[1]
	ten = -rs / rd
	tex = -r.Dir[1] / 0 // inf with same same as
	if ten > tex {
		ten, tex = tex, ten
	}
	return ten, tex
}
