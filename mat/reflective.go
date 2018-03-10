package mat

import "math"
import . "github.com/barnex/bruteray/br"

// A Reflective surface. E.g.:
// 	Reflective(WHITE)        // perfectly reflective, looks like shiny metal
// 	Reflective(WHITE.EV(-1)) // 50% reflective, looks like darker metal
// 	Reflective(RED)          // Reflects only red, looks like metal in transparent red candy-wrap.
func Reflective(c Color) Material {
	return &reflective{c}
}

type reflective struct {
	c Color
}

func (s *reflective) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	pos, norm := r.At(frag.T-offset), frag.Norm
	r2 := ctx.GetRay(pos, r.Dir().Reflect(norm))
	defer ctx.PutRay(r2)
	return e.ShadeAll(ctx, r2, N).Mul3(s.c)
}

// ReflectFresnel is a transparent material with index of refraction n,
// on top of material transmitted. E.g. a wet or varnished material.
// This looks similar to simple reflection,
// but reflection is stronger under grazing incidence.
// E.g.:
//  ReflectFresnel(1.33, BLACK)           // a thin film of water on a black surface
//  ReflectFresnel(1.33, Diffuse(WHITE))  // milk
//  ReflectFresnel(20, BLACK)             // metal
func ReflectFresnel(n float64, transmitted Material) Material {
	return &reflectFresnel{n, transmitted}
}

type reflectFresnel struct {
	n     float64
	trans Material
}

func (s *reflectFresnel) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	pos, norm := r.At(frag.T-offset), frag.Norm
	r2 := ctx.GetRay(pos, r.Dir().Reflect(norm))
	defer ctx.PutRay(r2)
	R := fresnelReflection(1, s.n, math.Abs(frag.Norm.Dot(r.Dir())))
	T := 1 - R
	trans := s.trans.Shade(ctx, e, N-1, r, frag)
	return e.ShadeAll(ctx, r2, N).Mul(R).MAdd(T, trans)
}

func fresnelReflection(n1, n2, cosθi float64) float64 {
	n12 := n1 / n2
	sin2θt := n12 * n12 * (1 - cosθi*cosθi) // sin² of transsion angle, using Snell's law.
	cosθt := math.Sqrt(1 - sin2θt)
	Rp := sqr((n1*cosθi - n2*cosθt) / (n1*cosθi + n2*cosθt))
	Rs := sqr((n1*cosθt - n2*cosθi) / (n1*cosθt + n2*cosθi))
	return 0.5 * (Rp + Rs)
}
