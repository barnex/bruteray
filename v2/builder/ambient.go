package builder

import . "github.com/barnex/bruteray/v2/tracer"
import . "github.com/barnex/bruteray/v2/color"

// TODO: UVMapping?
type Ambient struct {
	//Texture UVTexture
	C Color
}

func (a *Ambient) Init() {}

func (a *Ambient) Intersect(c *Ctx, r *Ray) HitRecord {
	return HitRecord{T: 1e99, Material: a}
}

func (a *Ambient) Eval(ctx *Ctx, s *Scene, r *Ray, recDepth int, h HitCoords) Color {
	return a.C
}

func (a *Ambient) Bounds() BoundingBox {
	return infBox()
}
