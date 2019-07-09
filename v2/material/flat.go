package material

import (
	. "github.com/barnex/bruteray/v2/color"
	"github.com/barnex/bruteray/v2/texture"
	. "github.com/barnex/bruteray/v2/tracer"
)

type flat struct {
	texture texture.Texture
}

// Flat returns a material with flat shading.
// I.e., returns colors as-is, disregarding any lighting, shadows, etc.
// Such materials emit light (in an indirect way). Suited for large,
// dimly luminous surfaces like computer screens, the sky, etc.
func Flat(t texture.Texture) Material { return &flat{t} }

func (m *flat) Eval(_ *Ctx, _ *Scene, r *Ray, recDepth int, h HitCoords) Color {
	return m.texture.At(r.At(h.T))
}
