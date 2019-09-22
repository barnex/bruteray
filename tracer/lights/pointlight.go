package lights

import (
	. "github.com/barnex/bruteray/tracer/types"
)

func PointLight(power Color, pos Vec) Light {
	return &point{pos, power}
}

type point struct {
	pos   Vec
	power Color
}

func (l *point) Sample(ctx *Ctx, target Vec) (Vec, Color) {
	return l.pos, l.power.Mul((1 / (4 * Pi)) / target.Sub(l.pos).Len2())
}

func (l *point) Object() Object {
	return l
}

func (*point) Intersect(*Ray) HitRecord {
	return HitRecord{}
}
