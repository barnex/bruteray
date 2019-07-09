package light

import (
	"github.com/barnex/bruteray/v2/builder"
	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"

	. "github.com/barnex/bruteray/v2/color"
)

// Point light source (with fall-off).
func NewPointLight(intensity Color, pos Vec) *pointLight {
	return &pointLight{pos: pos, c: intensity}
}

var _ Light = (*pointLight)(nil)

type pointLight struct {
	pos Vec
	c   Color
}

func (l *pointLight) Init() {}

func (l *pointLight) Sample(ctx *Ctx, target Vec) (Vec, Color) {
	return l.pos, l.c.Mul((1 / (4 * Pi)) / target.Sub(l.pos).Len2())
}

func (l *pointLight) Intersect(_ *Ctx, _ *Ray) HitRecord {
	return HitRecord{}
}

func (l *pointLight) Bounds() builder.BoundingBox {
	return builder.BoundingBox{l.pos, l.pos}
}
