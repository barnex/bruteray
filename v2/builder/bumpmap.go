package builder

import (
	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/texture"
	. "github.com/barnex/bruteray/v2/tracer"
)

type bumpMapped struct {
	orig Builder
	bump texture.Texturef
}

func BumpMapped(orig Builder, bumpMap texture.Texturef) Builder {
	return &bumpMapped{orig, bumpMap}
}

const bumpDiff = 1. / (1024 * 1024)

func (m *bumpMapped) Bounds() BoundingBox {
	return m.orig.Bounds()
}

func (m *bumpMapped) Intersect(ctx *Ctx, r *Ray) HitRecord {
	original := m.orig.Intersect(ctx, r)
	if original.T <= 0 {
		return original
	}

	r2 := ctx.Ray()
	defer ctx.PutRay(r2)
	r2.Dir = r.Dir

	r2.Start = r.Start.Add(Ez.Mul(bumpDiff))
	orig := m.orig.Intersect(ctx, r2)
	origP := r2.At(orig.T)
	c := r2.At(orig.T).MAdd(m.bump.At(origP), orig.Normal)

	r2.Start = r.Start.Add(Ey.Mul(bumpDiff))
	orig = m.orig.Intersect(ctx, r2)
	origP = r2.At(orig.T)
	a := r2.At(orig.T).MAdd(m.bump.At(origP), orig.Normal)

	r2.Start = r.Start.Add(Ex.Mul(bumpDiff))
	orig = m.orig.Intersect(ctx, r2)
	origP = r2.At(orig.T)
	b := r2.At(orig.T).MAdd(m.bump.At(origP), orig.Normal)

	original.Normal = TriangleNormal(c, a, b)
	return original
}

func (m *bumpMapped) Init() {
	m.orig.Init()
}
