// Package mat implements various types of materials.
package mat

import (
	"fmt"
	"math"

	. "github.com/barnex/bruteray/br"
)

// A Flat material always returns the same color.
// Useful for debugging, or for rare cases like
// a computer screen or other extended, dimly luminous surfaces.
func Flat(c Color) Material {
	return &flat{c}
}

type flat struct {
	c Color
}

func (s *flat) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	return s.c
}

// A Diffuse material appears perfectly mate,
// like paper or plaster.
// See https://en.wikipedia.org/wiki/Lambertian_reflectance.
func Diffuse(c Color) Material {
	return &diffuse{diffuse0{c}}
}

type diffuse struct {
	diffuse0
}

const offset = 1e-6 // TODO: Ray.Offset

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
	acc = acc.Add(s.refl.Mul3(e.ShadeNonLum(ctx, sec, N))) // does not include explicit lights

	return acc
}

// Diffuse material with direct illumination only (no interreflection).
// Intended for debugging or rapid previews. Diffuse is much more realistic.
func Diffuse0(c Color) Material {
	return &diffuse0{c}
}

type diffuse0 struct {
	refl Color
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
		return s.refl.Mul(re(norm.Dot(secundary.Dir()))).Mul3(intens)
	}
}

// Rectify: max(x, 0)
func re(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}

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

// ShadeDir returns a color based on the direction of a ray.
// Used for shading the ambient background, E.g., the sky.
type ShadeDir func(dir Vec) Color

func (s ShadeDir) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	pos := r.At(frag.T - offset)
	return s(pos)
}

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
	if !frag.Object.(CSGObj).Inside(posAhead) { // TODO: avoid cast?
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

func sqr(x float64) float64 {
	return x * x
}

// Blend mixes two materials with certain weights. E.g.:
// 	Blend(0.9, Diffuse(WHITE), 0.1, Reflective(WHITE))  // 90% mate + 10% reflective, like a shiny billiard ball.
func Blend(a float64, matA Material, b float64, matB Material) Material {
	return &blend{a, matA, b, matB}
}

// Shiny is shorthand for Blend-ing diffuse + reflection, e.g.:
// Shiny(WHITE, 0.1) // a white billiard ball, 10% specular reflection
func Shiny(c Color, reflectivity float64) Material {
	return Blend(1-reflectivity, Diffuse(c), reflectivity, Reflective(WHITE))
}

type blend struct {
	a    float64
	matA Material
	b    float64
	matB Material
}

func (s *blend) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	ca := s.matA.Shade(ctx, e, N, r, frag)
	cb := s.matB.Shade(ctx, e, N, r, frag)

	return ca.Mul(s.a).MAdd(s.b, cb)
}

// ShadeNormal is a debug shader that colors according to the normal vector projected on dir.
func ShadeNormal(dir Vec) Material {
	return &shadeNormal{dir}
}

type shadeNormal struct {
	dir Vec
}

func (s *shadeNormal) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	v := frag.Norm.Dot(s.dir)
	if v < 0 {
		return RED.Mul(-v) // towards cam
	} else {
		return BLUE.Mul(v) // away from cam
	}
}

// ShadeShape is a debug material that renders the object's shape only,
// even if no lighting is present. Useful while defining a scene before
// worrying about lighting.
func ShadeShape(c Color) Material {
	return &shadeShape{c}
}

type shadeShape struct{ c Color }

func (s *shadeShape) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	v := -frag.Norm.Dot(r.Dir())
	if v < 0 || v > 1 {
		panic(fmt.Sprintf("norm=%v, norm.len=%v, dir=%v, dir.len=%v, dot=%v", frag.Norm, frag.Norm.Len(), r.Dir(), r.Dir().Len(), v))
	}
	return s.c.Mul(v)
}
