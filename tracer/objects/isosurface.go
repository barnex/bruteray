package objects

import (
	"fmt"

	. "github.com/barnex/bruteray/tracer/types"
)

// IsoSurface constructs a surface defined by the height function:
//	y = dz * f(u, v)
// where (u, v) are varied in the interval (0, 1), and f is expected to return a value in (0,1).
//
// The surface is meshless. A numerical method is used to find ray-object intersection.
// This works best when the object is not viewed under a grazing angle,
// and when the surface gradient is small compared to 1 (i.e. has low "frequency").
//
// A texture can be used as height map, yielding "relief mapping"
// (See https://en.wikipedia.org/wiki/Relief_mapping_(computer_graphics) ).
//
// TODO: provide a way to tune the intersection so that it works with high frequencies.
func IsoSurface(mat Material, dx, dy, dz float64, f func(u, v float64) float64) Interface {
	return &isoSurface{
		bounds: BoundingBox{
			Min: Vec{0, 0, 0},
			Max: Vec{dx, dy, dz},
		},
		Tol:  1. / (1024 * 1024),
		mat:  mat,
		hmap: f,
	}
}

type isoSurface struct {
	bounds BoundingBox
	Tol    float64
	mat    Material
	hmap   func(u, v float64) float64
}

func (s *isoSurface) Bounds() BoundingBox {
	return s.bounds
}

func (s *isoSurface) Inside(p Vec) bool {
	return s.bounds.inside(p) && p[Y] < s.heightAt(p)
}

func (s *isoSurface) Intersect(r *Ray) HitRecord {
	t1, t2 := intersectAABB2(&s.bounds, r)
	// not intersecting bounding box: early return
	if !(t2 > 0) {
		return HitRecord{}
	}

	// brute-force intersection with height map
	t := s.bisect(r, t1-Tiny, t2+Tiny) // Tiny offset to avoid bleeding when surface exactly touches bounding box
	if t == 0 {
		return HitRecord{}
	}
	p := r.At(t)
	uv := s.localUV(p)
	n := s.normalAtUV(uv)

	return HitRecord{T: t, Normal: n, Material: s.mat, Local: Vec{uv[0], uv[1]}}
}

func (f *isoSurface) isoValueAt(p Vec) float64 {
	// TODO: rescale so it always lies in [-1,1]
	// That makes an absoulte tolerance on Y meaningful
	return p[Y] - (f.heightAt(p))
}

func (f *isoSurface) heightAt(p Vec) float64 {
	return f.heightAtUV(f.localUV(p))
}

func (o *isoSurface) heightAtUV(uv Vec2) float64 {
	const tol = 1. / 1024.
	u, v := uv[0], uv[1]
	f := o.hmap(u, v)
	if f < -tol || f > 1+tol {
		panic(fmt.Sprintf("IsoSurface: User function: bad value at u=%v, v=%v: value=%v. Want [0..1]", u, v, f))
	}
	return o.height() * f
}

func (o *isoSurface) height() float64 {
	return o.bounds.Max[Y]
}

func (f *isoSurface) partialAt(p Vec) Vec2 {
	return f.partialAtUV(f.localUV(p))
}

func (f *isoSurface) partialAtUV(uv Vec2) Vec2 {
	u, v := uv[0], uv[1]
	const delta = 1. / 4096. // TODO: use pixel pitch
	partialU := (1 / delta) * (f.heightAtUV(Vec2{u + delta/2, v}) - f.heightAtUV(Vec2{u - delta/2, v}))
	partialV := (1 / delta) * (f.heightAtUV(Vec2{u, v + delta/2}) - f.heightAtUV(Vec2{u, v - delta/2}))
	return Vec2{partialU, partialV}
}

// TODO: precompute and store
func (f *isoSurface) normalAtUV(uv Vec2) Vec {
	diff := f.partialAtUV(uv)
	return Vec{-diff[0], 1, -diff[1]}
}

func (o *isoSurface) localUV(p Vec) Vec2 {
	return Vec2{p[X] / o.bounds.Max[X], p[Z] / o.bounds.Max[Z]}
}

func (f *isoSurface) bisect(r *Ray, min, max float64) float64 {

	f1 := f.isoValueAt(r.At(min))
	f2 := f.isoValueAt(r.At(max))

	// no zero crossing
	if f1*f2 > 0 {
		return 0
	}

	in, out := max, min
	if f.isoValueAt(r.At(out)) > 0 {
		in, out = out, in
	}

	tol := f.Tol
	for (out - in) > tol {
		mid := (in + out) / 2
		if f.isoValueAt(r.At(mid)) > 0 {
			in = mid
		} else {
			out = mid
		}
	}

	return out
	//t := (in + out) / 2
	//if math.Abs(f.isoValueAt(r.At(t))) < (tol / 2) {
	//	return t
	//}
	//return 0
}

//// find both intersection points with an arbitrary convex shape
//// t2 == 0 means no intersection
//func intersect2(b Object, ctx *Ctx, r *Ray) (t1, t2 float64) {
//	// (0) No intersection: t1 <= 0
//	//             -----------
//	//             |         |
//	//             |         1  *-------->
//	//             |         |
//	//             -----------
//	// (1) Ray crosses through: t1 > 0, t2 > 0
//	//             -----------
//	//             |         |
//	//        *----1---------2----->
//	//             |         |
//	//             -----------
//	//
//	// (2) Ray starts inside: t1 < 0, t2 > 0
//	//             -----------
//	//             |         |
//	//             1----*----2----->
//	//             |         |
//	//             -----------
//
//	t1 = b.Intersect(ctx, r).T
//
//	// case (0): No intersection
//	if t1 <= 0 {
//		return 0, 0
//	}
//
//	// There is at least one intersection point.
//	// Start a second ray just beyond that point
//	// to look for the next one.
//	r2 := ctx.Ray()
//	defer ctx.PutRay(r2)
//	*r2 = *r
//	r2.Start = r.At(t1 + marchOffset)
//	t2_ := b.Intersect(ctx, r2).T
//
//	// Ray crosses
//	if t2_ > 0 {
//		t2 = t2_ + t1 + marchOffset
//		return t1, t2
//	}
//
//	*r2 = *r
//	r2.Dir = r2.Dir.Mul(-1)
//	t2_ = b.Intersect(ctx, r2).T
//	t2 = -t2_
//
//	return t2, t1
//}
