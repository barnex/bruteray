package mat

import (
	. "github.com/barnex/bruteray/br"
)

// A Diffuse material appears perfectly mate,
// like paper or plaster.
// See https://en.wikipedia.org/wiki/Lambertian_reflectance.
func Diffuse(c Texture3D) Material {
	return &diffuse{diffuse0{c}}
}

type diffuse struct {
	diffuse0
}

func (s *diffuse) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	// This is the core of bi-directonial path tracing.

	pos, norm := r.At(frag.T-offset), frag.Norm

	// accumulates the result
	var acc Color

	// first sum over all explicit sources
	// (with fall-off)
	for _, l := range e.Lights {
		acc = acc.Add(s.lightIntensity(ctx, e, pos, norm, l))
	}

	// add one random ray to sample (only!) the indirect sources.
	// (no fall-off, the chance of hitting an object
	// automatically falls off correctly with distance).
	// Choose the random ray via importance sampling.
	sec := ctx.GetRay(pos.MAdd(offset, norm), RandVecCos(ctx.Rng, norm))
	defer ctx.PutRay(sec)
	c := s.refl.At(pos)
	acc = acc.Add(c.Mul3(e.ShadeNonLum(ctx, sec, N))) // does not include explicit lights

	return acc
}

// Diffuse material with direct illumination only (no interreflection).
// Intended for debugging or rapid previews. Diffuse is much more realistic.
func Diffuse0(c Texture3D) Material {
	return &diffuse0{c}
}

type diffuse0 struct {
	refl Texture3D
}

func (s *diffuse0) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	pos, norm := r.At(frag.T-offset), frag.Norm
	var acc Color
	for _, l := range e.Lights {
		acc = acc.Add(s.lightIntensity(ctx, e, pos, norm, l))
	}
	return acc
}

// TODO: Env method
func (s *diffuse0) lightIntensity(ctx *Ctx, e *Env, pos, norm Vec, l Light) Color {
	lpos, intens := l.Sample(ctx, pos)

	secundary := ctx.GetRay(pos, lpos.Sub(pos).Normalized())
	defer ctx.PutRay(secundary)

	//t := e.IntersectAny(secundary)

	lightT := lpos.Sub(pos).Len()

	if e.Occludes(ctx, secundary, lightT) {
		return Color{} // shadow
	} else {
		return s.refl.At(pos).Mul(re(norm.Dot(secundary.Dir()))).Mul3(intens)
	}
}

// Rectify: max(x, 0)
func re(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}
