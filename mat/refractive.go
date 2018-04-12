package mat

import (
	. "github.com/barnex/bruteray/br"
	"math"
)

// Refractive material with index of refraction n1 outside and n2 inside.
// E.g.:
// 	Refractive(1, 1.5) // glass in air
// 	Refractive(1.5, 1) // air in glass
func Refractive(n1, n2 float64) Material {
	return &refractive{n1, n2}
}

type refractive struct {
	n1, n2 float64 // relative index of refraction
}

// https://en.wikipedia.org/wiki/Fresnel_equations
func (s *refractive) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {

	const offset = 1e-3

	posAhead := r.At(frag.T + offset)  // point in front of us, inside glass object
	posBehind := r.At(frag.T - offset) // point behind, outside of glass

	// if we are exiting rather than entering the refractive material,
	// swap refractive indices.
	n1, n2 := s.n1, s.n2
	if !frag.Object.(Insider).Inside(posAhead) { // TODO: avoid cast?
		n1, n2 = n2, n1
	}
	n12 := n1 / n2

	i := r.Dir().Normalized()               // incident direction
	n := frag.Norm.Normalized()             // normal direction
	cosθi := -i.Dot(n)                      // cos of incident angle. Sign because ray points away from normal.
	sin2θt := n12 * n12 * (1 - cosθi*cosθi) // sin² of transsion angle, using Snell's law.

	// Total internal reflection?
	if sin2θt > 1 {
		r2 := ctx.GetRay(posBehind, r.Dir().Reflect(frag.Norm))
		defer ctx.PutRay(r2)
		return e.ShadeAll(ctx, r2, N)
	}

	cosθt := math.Sqrt(1 - sin2θt)

	// Fresnel equations for reflected intensity:
	Rp := sqr((n1*cosθi - n2*cosθt) / (n1*cosθi + n2*cosθt))
	Rs := sqr((n1*cosθt - n2*cosθi) / (n1*cosθt + n2*cosθi))
	R := 0.5 * (Rp + Rs)
	T := 1 - R

	// transmitted ray
	t := i.Mul(n12).MAdd((n12*cosθi - math.Sqrt(1-(sin2θt))), frag.Norm)
	r2 := ctx.GetRay(posAhead, t)
	defer ctx.PutRay(r2)
	cT := e.ShadeAll(ctx, r2, N).Mul(T)

	// reflected ray
	r3 := ctx.GetRay(posBehind, i.Reflect(n))
	defer ctx.PutRay(r3)
	cR := e.ShadeAll(ctx, r3, N).Mul(R)

	return cR.Add(cT)
}
