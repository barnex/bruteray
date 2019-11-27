package materials

import (
	"github.com/barnex/bruteray/texture"
	. "github.com/barnex/bruteray/tracer/types"
)

type flat struct {
	texture texture.Texture
}

// Flat returns a material with flat shading.
// I.e., returns colors as-is, disregarding any lighting, shadows, etc.
// Such materials emit light (in an indirect way). Suited for large,
// dimly luminous surfaces like computer screens, the sky, etc.
func Flat(t texture.Texture) Material { return &flat{t} }

func (m *flat) Shade(_ *Ctx, _ *Scene, r *Ray, h HitCoords) Color {
	return m.texture.At(h.Local)
}
