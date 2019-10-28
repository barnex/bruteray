package objects

import (
	"math"

	. "github.com/barnex/bruteray/tracer/types"
)

func Sphere(m Material, diam float64, center Vec) Interface {
	r := diam / 2
	return &quadric{
		mat:    m,
		origin: center,
		a:      Vec{1, 1, 1},
		b:      r * r,
		bounds: BoundingBox{
			Min: Vec{-r, -r, -r},
			Max: Vec{+r, +r, +r},
		}.translated(center).withMargin(Tiny),
	}
	// NOTE:
	// We add a tiny offset to the boudning box, for in case the quadric surface
	// touches the box. When that happens, Bounds().Inside() is noisy and causes bleeding.
	// This is most clearly visible in, e.g., tests with isometric projections.
}

func Cylinder(m Material, diam, height float64, center Vec) Interface {
	r := diam / 2
	return &quadric{
		mat:    m,
		origin: center,
		a:      Vec{1, 0, 1},
		b:      r * r,
		bounds: BoundingBox{
			Min: Vec{-r, -height / 2, -r},
			Max: Vec{+r, +height / 2, +r},
		}.translated(center).withMargin(Tiny),
	}
}

func CylinderDir(m Material, dir int, diam, height float64, center Vec) Interface {
	r := diam / 2
	a := Vec{1, 1, 1}
	a[dir] = 0
	bMin := Vec{-r, -r, -r}
	bMin[dir] = -height / 2
	return &quadric{
		mat:    m,
		origin: center,
		a:      a,
		b:      r * r,
		bounds: BoundingBox{
			Min: bMin,
			Max: bMin.Mul(-1),
		}.translated(center).withMargin(Tiny),
	}
}

func CylinderWithCaps(m Material, diam, height float64, center Vec) Interface {
	return And(Cylinder(m, diam, height, center), Box(m, diam, height, diam, center))
}

//var tinyVec = Vec{Tiny, Tiny, Tiny}

// a0 x² + a1 y² + a2 z² = b
type quadric struct {
	origin Vec
	a      Vec
	b      float64
	bounds BoundingBox
	mat    Material
}

func (q *quadric) Bounds() BoundingBox {
	return q.bounds
}

func (s *quadric) Intersect(r *Ray) HitRecord {
	s0 := r.Start[0] - s.origin[0]
	s1 := r.Start[1] - s.origin[1]
	s2 := r.Start[2] - s.origin[2]

	d0 := r.Dir[0]
	d1 := r.Dir[1]
	d2 := r.Dir[2]

	a0 := s.a[0]
	a1 := s.a[1]
	a2 := s.a[2]

	A := a0*d0*d0 + a1*d1*d1 + a2*d2*d2
	B := 2 * (a0*d0*s0 + a1*d1*s1 + a2*d2*s2)
	C := a0*s0*s0 + a1*s1*s1 + a2*s2*s2 - s.b

	D := B*B - 4*A*C
	if D < 0 {
		return HitRecord{}
	}
	V := math.Sqrt(D)

	t1 := (-B - V) / (2 * A)
	t2 := (-B + V) / (2 * A)
	//t1, t2 = util.Sort2(t1, t2)
	if t2 < t1 {
		t1, t2 = t2, t1
	}

	if t1 > 0 {
		t := t1
		p := r.At(t)
		if s.bounds.inside(p) {
			return HitRecord{T: t, Normal: s.normal(p), Material: s.mat, Local: p.Sub(s.origin)}
		}
	}

	if t2 > 0 {
		t := t2
		p := r.At(t)
		if s.bounds.inside(p) {
			return HitRecord{T: t, Normal: s.normal(p), Material: s.mat, Local: p.Sub(s.origin)}
		}
	}

	return HitRecord{}
}

func (s *quadric) Inside(p Vec) bool {
	if !s.bounds.inside(p) {
		return false
	}
	p_o := p.Sub(s.origin)
	return mul3(s.a, p_o).Dot(p_o) < s.b
}

func (s *quadric) normal(x Vec) Vec {
	x_o := x.Sub(s.origin)
	return mul3(s.a, x_o)
}

func mul3(a, b Vec) Vec {
	return Vec{
		a[X] * b[X],
		a[Y] * b[Y],
		a[Z] * b[Z],
	}
}
