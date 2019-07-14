package material

import (
	"math"

	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/tracer"
	. "github.com/barnex/bruteray/v2/util"
)

func Refractive(n float64) Material {
	return Refractive2(1, n)
}

// Refractive material with index of refraction n1 outside and n2 inside.
// E.g.:
// 	Refractive2(1, 1.5) // glass in air
// 	Refractive2(1.5, 1) // air in glass
func Refractive2(n1, n2 float64) Material {
	return &refractive{n1, n2}
}

type refractive struct {
	n1, n2 float64 // relative index of refraction outside and inside
}

// https://en.wikipedia.org/wiki/Fresnel_equations
func (s *refractive) Eval(ctx *Ctx, e *Scene, r *Ray, recDepth int, h HitCoords) Color {

	n := h.Normal.Normalized()
	i := r.Dir.Normalized() // incident direction

	// if we are exiting rather than entering the refractive material,
	// swap refractive indices and flip normal
	n1, n2 := s.n1, s.n2
	if i.Dot(n) > 0 {
		n = n.Mul(-1)
		n1, n2 = n2, n1
	}
	n12 := n1 / n2

	cosθi := -i.Dot(n)                      // cos of incident angle. Sign because ray points away from normal.
	sin2θt := n12 * n12 * (1 - cosθi*cosθi) // sin² of transsion angle, using Snell's law.

	// Total internal reflection?
	if sin2θt > 1 {
		r2 := ctx.Ray()
		r2.Start = r.At(h.T - Tiny) // start at same side of surface
		r2.Dir = r.Dir.Reflect(h.Normal)
		defer ctx.PutRay(r2)
		return e.Eval(ctx, r2, recDepth)
	}

	cosθt := math.Sqrt(1 - sin2θt)

	// Fresnel equations for reflected intensity:
	Rp := Sqr((n1*cosθi - n2*cosθt) / (n1*cosθi + n2*cosθt))
	Rs := Sqr((n1*cosθt - n2*cosθi) / (n1*cosθt + n2*cosθi))
	R := 0.5 * (Rp + Rs)
	T := 1 - R
	_=T

	// transmitted ray
	t := i.Mul(n12).MAdd((n12*cosθi - math.Sqrt(1-(sin2θt))), n)
	r2 := ctx.Ray()
	r2.Start = r.At(h.T + Tiny) // start at other side of surface
	r2.Dir = t
	defer ctx.PutRay(r2)
	cT := e.Eval(ctx, r2, recDepth).Mul(T)

	// reflected ray
	r3 := ctx.Ray()
	r3.Start = r.At(h.T - Tiny) // same side of surface
	r3.Dir = i.Reflect(n)
	defer ctx.PutRay(r3)
	cR := e.Eval(ctx, r3, recDepth).Mul(R)

	return cR.Add(cT)
}
