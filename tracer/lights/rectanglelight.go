package lights

import (
	"github.com/barnex/bruteray/tracer/objects"
	. "github.com/barnex/bruteray/tracer/types"
)

type rectangle struct {
	w, h       float64
	center     Vec
	totalPower Color // W/m2 (pi?)
	object     objects.Interface
}

// TODO: if normal points the wrong way, show a big red cross or something

func RectangleLight(brightness Color, w, h float64, center Vec) Light {
	mat := &twoSided{brightness, Color{1, 0, 0}} // TODO: black
	return &rectangle{
		w:          w,
		h:          h,
		center:     center,
		totalPower: brightness.Mul(w * h),
		object:     objects.Rectangle(mat, w, h, center),
	}
}

type twoSided struct {
	front, back Color
}

func (m *twoSided) Eval(ctx *Ctx, s *Scene, r *Ray, recDepth int, h HitCoords) Color {
	if r.Dir[Y] > 0 {
		return m.front
	} else {
		return m.back
	}
}

func (l *rectangle) Sample(ctx *Ctx, target Vec) (Vec, Color) {
	u, v := ctx.Sample2()
	p := Vec{(u - 0.5) * l.w, Tiny, (v - 0.5) * l.h,}.Add(l.center)
	n := Vec{0, -1, 0}

	delta := target.Sub(p)
	I := (1 / Pi) / delta.Len2()
	I *= (n.Dot(delta.Normalized()))
	if I < 0 {
		I = 0
	}
	return p, l.totalPower.Mul(I)
}

func (l *rectangle) Object() Object {
	return l.object
}
