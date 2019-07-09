package builder

import (
	"math"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/texture"
	. "github.com/barnex/bruteray/v2/tracer"
)

const marchOffset = (1. / (1024 * 1024))

type isoSurface struct {
	bounds   Builder //geom.Frame?
	function texture.Texturef
	material Material
}

func NewIsoSurface() Builder {
	return &isoSurface{}
}

func (s *isoSurface) Bounds() BoundingBox {
	return s.bounds.Bounds()
}

func (s *isoSurface) Init() {
	s.bounds.Init()
}

func (s *isoSurface) Intersect(ctx *Ctx, r *Ray) HitRecord {
	t1, t2 := intersect2(s.bounds, ctx, r)
	if t2 <= 0 {
		return HitRecord{}
	}

	t := s.bisect(r, t1, t2)
	n := s.normal(r, t1, t2)
	return HitRecord{T: t, Normal: n, Material: s.material}
}

// t2 == 0 means no intersection
func intersect2(b Object, ctx *Ctx, r *Ray) (t1, t2 float64) {
	// (0) No intersection: t1 <= 0
	//             -----------
	//             |         |
	//             |         1  *-------->
	//             |         |
	//             -----------
	// (1) Ray crosses through: t1 > 0, t2 > 0
	//             -----------
	//             |         |
	//        *----1---------2----->
	//             |         |
	//             -----------
	//
	// (2) Ray starts inside: t1 < 0, t2 > 0
	//             -----------
	//             |         |
	//             1----*----2----->
	//             |         |
	//             -----------

	t1 = b.Intersect(ctx, r).T

	// case (0): No intersection
	if t1 <= 0 {
		return 0, 0
	}

	// There is at least one intersection point.
	// Start a second ray just beyond that point
	// to look for the next one.
	r2 := ctx.Ray()
	defer ctx.PutRay(r2)
	*r2 = *r
	r2.Start = r.At(t1 + marchOffset)
	t2_ := b.Intersect(ctx, r2).T

	// Ray crosses
	if t2_ > 0 {
		t2 = t2_ + t1 + marchOffset
		return t1, t2
	}

	*r2 = *r
	r2.Dir = r2.Dir.Mul(-1)
	t2_ = b.Intersect(ctx, r2).T
	t2 = -t2_

	return t2, t1
}

func (f *isoSurface) bisect(r *Ray, min, max float64) float64 {
	s := f.function
	const tol = 1e-9
	in, out := max, min
	if s.At(r.At(in)) < 0 {
		in, out = out, in
	}

	for math.Abs(in-out)/(in+out) > tol {
		mid := (in + out) / 2
		if s.At(r.At(mid)) > 0 {
			in = mid
		} else {
			out = mid
		}
	}
	return out
}

func (s *isoSurface) normal(r *Ray, min, max float64) Vec {
	// TODO: use d/dx, d/dy, d/dz
	t := s.bisect(r, min, max)
	c := r.At(t)

	ra := *r
	ra.Start = ra.Start.Add(Ex.Mul(bumpDiff))
	ta := s.bisect(&ra, min, max)
	a := ra.At(ta)

	rb := *r
	ra.Start = rb.Start.Add(Ey.Mul(bumpDiff))
	tb := s.bisect(&rb, min, max)
	b := rb.At(tb)

	return TriangleNormal(c, a, b)
}
