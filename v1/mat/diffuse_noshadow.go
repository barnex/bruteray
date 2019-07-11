package mat

import (
	. "github.com/barnex/bruteray/v1/br"
)

// Diffuse material with direct illumination only and no shadows.
// Intended for the tutorial.
func Diffuse00(c Color) Material {
	return &diffuse00{c}
}

type diffuse00 struct {
	refl Color
}

func (s *diffuse00) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	pos, norm := r.At(frag.T-offset), frag.Norm
	var acc Color
	for _, l := range e.Lights {
		lpos, intens := l.Sample(ctx, pos)
		secundary := ctx.GetRay(pos, lpos.Sub(pos).Normalized())
		defer ctx.PutRay(secundary)
		i := s.refl.Mul(re(norm.Dot(secundary.Dir()))).Mul3(intens)
		acc = acc.Add(i)
	}
	return acc
}
