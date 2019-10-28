package materials

import (
	. "github.com/barnex/bruteray/color"
	"github.com/barnex/bruteray/texture"
	. "github.com/barnex/bruteray/tracer"
)

// TODO: always consume recursion
func Transparent(t texture.Texture, consumeRecursion bool) Material {
	return &transparent{t, consumeRecursion}
}

var _ TransparentMaterial = (*transparent)(nil)

type transparent struct {
	t      texture.Texture
	useRec bool
}

func (m *transparent) Eval(ctx *Ctx, s *Scene, r *Ray, h HitCoords) Color {
	pos := r.At(h.T + Tiny)
	r2 := ctx.Ray()
	r2.Start = pos
	r2.Dir = r.Dir
	defer ctx.PutRay(r2)
	// No caustics please
	return s.EvalMinusLights(ctx, r2).Mul3(m.t.At(h.Local))
}

func (m *transparent) Filter(r *Ray, h HitRecord, background Color) Color {
	return background.Mul3(m.t.At(h.Local))
}
