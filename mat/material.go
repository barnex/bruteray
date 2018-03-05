// Package mat implements various types of materials.
package mat

import (
	"fmt"
	"math"

	. "github.com/barnex/bruteray/br"
)

const offset = 1e-6 // TODO: Ray.Offset

// ShadeDir returns a color based on the direction of a ray.
// Used for shading the ambient background, E.g., the sky.
type ShadeDir func(dir Vec) Color

func (s ShadeDir) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	pos := r.At(frag.T - offset)
	return s(pos)
}

func Skybox(tex Image) ShadeDir {
	return ShadeDir(
		func(dir Vec) Color {
			dir = dir.Normalized()
			u := 0.5*dir[X] + 0.5
			v := 0.5*dir[Z] + 0.5
			return tex.At(u, v)
		})
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

// DebugShape is a debug material that renders the object's shape only,
// even if no lighting is present. Useful while defining a scene before
// worrying about lighting.
func DebugShape(c Color) Material {
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
