package objects

import (
	"fmt"
	"math"

	. "github.com/barnex/bruteray/tracer/types"
)

// Box returns an (axis-aligned) box with given size and center.
func Box(m Material, width, height, depth float64, center Vec) Interface {
	return &box{
		bounds: BoundingBox{
			Min: Vec{-width / 2, -height / 2, -depth / 2},
			Max: Vec{width / 2, height / 2, depth / 2},
		}.translated(center),
		origin: center,
		mat:    m,
	}
}

type box struct {
	bounds BoundingBox
	origin Vec
	mat    Material
}

func (s *box) Bounds() BoundingBox {
	return s.bounds
}

func (b *box) Inside(p Vec) bool {
	return b.bounds.inside(p)
}

func (b *box) Intersect(r *Ray) HitRecord {
	t := b.bounds.intersect(r)
	if !(t > 0) { // handles NaN gracefully
		return HitRecord{}
	}
	p := r.At(t)
	return HitRecord{T: t, Normal: b.normal(p), Material: b.mat, Local: p.Sub(b.origin)}
}

func (b *box) normal(p Vec) Vec {
	s := &b.bounds
	for i := range p {
		if approx(p[i], s.Min[i]) {
			return unit[i].Mul(-1)
		}
		if approx(p[i], s.Max[i]) {
			return unit[i]
		}
	}
	panic(fmt.Sprint("box.normal", p, s.Min, s.Max))
}

func approx(a, b float64) bool {
	return math.Abs(a-b) < 1e-8
}

var unit = [3]Vec{Ex, Ey, Ez}
