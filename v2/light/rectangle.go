package light

import (
	"github.com/barnex/bruteray/v2/builder"
	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	. "github.com/barnex/bruteray/v2/tracer"
)

type RectangleLight struct {
	*builder.Rectangle
	c Color
}

// TODO: if normal points the wrong way, show a big red cross or something
func NewRectangleLight(c Color, o, a, b Vec) *RectangleLight {
	v1 := a.Sub(o)
	v2 := b.Sub(o)
	area := v1.Cross(v2).Len()
	l := &RectangleLight{
		Rectangle: builder.NewRectangle(material.Flat(c), o, a, b),
		c:         c.Mul(area),
	}
	return l
}

var _ Light = (*RectangleLight)(nil)

func (l *RectangleLight) Sample(ctx *Ctx, target Vec) (Vec, Color) {
	u, v := ctx.Rng.Float64(), ctx.Rng.Float64()

	p := l.Origin().MAdd(u, l.CtrlVec(0)).MAdd(v, l.CtrlVec(1))
	n := l.CtrlVec(0).Cross(l.CtrlVec(1)).Normalized()
	p = p.MAdd(Tiny, n)

	delta := target.Sub(p)
	I := (1 / (2 * Pi)) / delta.Len2()
	I *= (n.Dot(delta.Normalized()))
	if I < 0 {
		I = 0
	}
	return p, l.c.Mul(I)
}
