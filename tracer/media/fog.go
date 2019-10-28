package media

import (
	"math"

	. "github.com/barnex/bruteray/tracer/types"
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

// TODO: not correct when camera is in fog?
func (m *fog) Filter(ctx *Ctx, s *Scene, r *Ray, tMax float64, orig Color) Color {
	if !ctx.IsInitial() {
		return orig
	}
	ten, tex := intersectHalfSpace(m.height, r)
	start, end := intersectIntervals(0, tMax, ten, tex)

	len := end - start
	u, _ := ctx.Generate2()
	tSampl, weight := randExpInterval(u, m.density, len)
	t := start + tSampl
	var acc Color
	p := r.At(t - Tiny)
	sec := ctx.Ray()
	//defer ctx.PutRay(sec)

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
		acc = acc.MAdd(weight, intens)
	}

	trans := math.Exp(-len * m.density)
	ctx.PutRay(sec)
	return orig.Mul(trans).Add(acc)
}

// randExpInterval draws a random number t from the probability distribution
// 	exp(-density * t), t = [0..max]
// Returns the random number and the itegral of the unnormalized distribution.
//
// This is used to draw a random sample from atmospheric where the line of sight
// is limited between t=0 and t=max. The sample must be weighted after use.
func randExpInterval(rand float64, density, max float64) (pos, weight float64) {
	y := rand
	N := 1 / (1 - math.Exp(-density*max))
	t := -1 / density * math.Log(1-y/N)
	return t, 1 / N
}
