package materials

import (
	. "github.com/barnex/bruteray/color"
	"github.com/barnex/bruteray/texture"
	. "github.com/barnex/bruteray/tracer"
)

func BlendMap(t texture.Texture, a, b Material) Func {
	return func(ctx *Ctx, s *Scene, r *Ray, h HitCoords) Color {
		key1 := t.At(h.Local).R
		key2 := 1 - key1
		var ca, cb Color
		if key1 > 0 {
			ca = a.Eval(ctx, s, r, h)
		}
		if key2 > 0 {
			cb = b.Eval(ctx, s, r, h)
		}
		return ca.Mul(key1).MAdd(key2, cb)
	}
}

type Func func(ctx *Ctx, s *Scene, r *Ray, h HitCoords) Color

func (f Func) Eval(ctx *Ctx, s *Scene, r *Ray, h HitCoords) Color {
	return f(ctx, s, r, h)
}

// Blend mixes two materials with certain weights. E.g.:
// 	Blend(0.9, Mate(WHITE), 0.1, Reflective(WHITE))  // 90% mate + 10% reflective, like a shiny billiard ball.
func Blend(a float64, matA Material, b float64, matB Material) Material {
	return &blend{a, matA, b, matB}
}

// Shiny is shorthand for Blend-ing diffuse + reflection, e.g.:
// Shiny(WHITE, 0.1) // a white billiard ball, 10% specular reflection
func Shiny(c texture.Texture, reflectivity float64) Material {
	return Blend(1-reflectivity, Matte(c), reflectivity, Reflective(Color{1, 1, 1}))
}

type blend struct {
	a    float64
	matA Material
	b    float64
	matB Material
}

func (m *blend) Eval(ctx *Ctx, s *Scene, r *Ray, h HitCoords) Color {
	ca := m.matA.Eval(ctx, s, r, h)
	cb := m.matB.Eval(ctx, s, r, h)
	return ca.Mul(m.a).MAdd(m.b, cb)
}
