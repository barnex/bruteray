package materials

import . "github.com/barnex/bruteray/tracer/types"

// A TwoSided material consists of a front-facing material
// (seen when looking towards the surface normal) and a
// back-facing material (seen when looking at the back side
// of the surface).
func TwoSided(front, back Material) Material {
	return &twoSided{front, back}
}

type twoSided struct {
	front, back Material
}

func (m *twoSided) Eval(ctx *Ctx, s *Scene, r *Ray, recDepth int, h HitCoords) Color {
	if r.Dir.Dot(h.Normal) > 0 {
		return m.front.Eval(ctx, s, r, recDepth, h)
	} else {
		return m.back.Eval(ctx, s, r, recDepth, h)
	}
}
