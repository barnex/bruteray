// Package mat implements various types of materials.
package mat

import (
	"fmt"

	. "github.com/barnex/bruteray/br"
)

const offset = 1e-6 // TODO: Ray.Offset

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
