package bruteray

import (
	"fmt"
	"math"
)

type Shape interface {
	// Hit returns the t-value where the ray first hits the surface.
	// I.e. r.At(t) lies on the surface.
	// If the ray intersects the surface multiple times,
	// the smallest positive t must be returned.
	// t<=0 is interpreted as the ray not intersecting the surface.
	Hit(r *Ray) float64

	// Ray-shape intersection.
	// Special cases:
	// 	Shape lies entirely behind ray start point: return Interval{}
	// 	Ray start point lies inside shape: Interval.Min < 0
	// 	Shape does not intersect ray at all: return Interval{}
	// 	Shape lies entirely in front of ray start: Min & Max > 0
	Inters(r *Ray) Interval

	// Normal vector at position.
	// Does not necessarily need to point outwards.
	Normal(pos Vec) Vec
}

// -- sphere

func Sphere(center Vec, radius float64, m Material) Obj {
	return &prim{&sphere{center, sqr(radius)}, m}
}

type sphere struct {
	c  Vec
	r2 float64
}

func (s *sphere) Normal(pos Vec) Vec {
	n := pos.Sub(s.c).Normalized().check()
	return n
}

func (s *sphere) Inters(r *Ray) Interval {
	v := r.Start.Sub(s.c)
	d := r.Dir
	vd := v.Dot(d)
	D := sqr(vd) - (v.Len2() - s.r2)
	if D < 0 {
		return Interval{}
	}
	t1 := (-vd - math.Sqrt(D))
	t2 := (-vd + math.Sqrt(D))
	return Interval{t1, t2}.Fix().check()
}

func (s *sphere) Hit(r *Ray) float64 {
	return s.Inters(r).Front()
}

// -- box (axis aligned)

type box struct {
	min, max Vec
}

func Box(center Vec, rx, ry, rz float64, m Material) Obj {
	return &prim{&box{
		min: center.Sub(Vec{rx, ry, rz}),
		max: center.Add(Vec{rx, ry, rz}),
	}, m}
}

func Cube(center Vec, r float64, m Material) Obj {
	return Box(center, r, r, r, m)
}

func (s *box) Inters(r *Ray) Interval {
	min_ := s.min
	max_ := s.max

	tmin := min_.Sub(r.Start).Div3(r.Dir)
	tmax := max_.Sub(r.Start).Div3(r.Dir)

	txen := min(tmin[X], tmax[X])
	txex := max(tmin[X], tmax[X])

	tyen := min(tmin[Y], tmax[Y])
	tyex := max(tmin[Y], tmax[Y])

	tzen := min(tmin[Z], tmax[Z])
	tzex := max(tmin[Z], tmax[Z])

	ten := max3(txen, tyen, tzen)
	tex := min3(txex, tyex, tzex)

	if ten > tex {
		return Interval{}
	}

	if math.IsNaN(ten) || math.IsNaN(tex) {
		return Interval{}
	}

	return Interval{ten, tex}.Fix().check()
}

func (s *box) Hit(r *Ray) float64 {
	return s.Inters(r).Front()
}

func (s *box) Normal(p Vec) Vec {
	p.check()
	for i := range p {
		if approx(p[i], s.min[i]) || approx(p[i], s.max[i]) {
			return Unit[i]
		}
	}

	panic(fmt.Sprint("box.normal", p, s.min, s.max))
}

func approx(a, b float64) bool {
	return math.Abs(a-b) < 1e-4
}

// -- sheet (infinite)

func Sheet(dir Vec, off float64, m Material) Obj {
	return &prim{&sheet{dir, off}, m}
}

type sheet struct {
	dir Vec
	off float64
}

func (s *sheet) Normal(pos Vec) Vec {
	return s.dir
}

func (s *sheet) Inters(r *Ray) Interval {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t := (s.off - rs) / rd
	return Interval{t, t}.Fix().check()
}

func (s *sheet) Hit(r *Ray) float64 {
	return s.Inters(r).Front()
}

// --rectangle (finite sheet)

// A rectangle (i.e. finite sheet) at given position,
// with normal vector dir and half-axes rx, ry, rz.
func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj {
	return &prim{&rect{pos, dir, rx, ry, rz}, m}
}

type rect struct {
	pos, dir   Vec
	rx, ry, rz float64
}

func (s *rect) Inters(r *Ray) Interval {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t := (s.pos.Dot(s.dir) - rs) / rd
	p := r.At(t).Sub(s.pos)
	if p[X] < -s.rx || p[X] > s.rx ||
		p[Y] < -s.ry || p[Y] > s.ry ||
		p[Z] < -s.rz || p[Z] > s.rz {
		return Interval{}
	}
	return Interval{t, t}.Fix().check()
}

func (s *rect) Hit(r *Ray) float64 {
	return s.Inters(r).Front()
}

func (s *rect) Normal(p Vec) Vec {
	return s.dir
}

// -- slab

func Slab(dir Vec, off1, off2 float64, m Material) Obj {
	return &prim{&slab{dir, off1, off2}, m}
}

type slab struct {
	dir        Vec
	off1, off2 float64
}

func (s *slab) Normal(pos Vec) Vec {
	return s.dir
}

func (s *slab) Inters(r *Ray) Interval {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t1 := (s.off1 - rs) / rd
	t2 := (s.off2 - rs) / rd
	t1, t2 = sort(t1, t2)
	return Interval{t1, t2}.Fix().check()
}

func (s *slab) Hit(r *Ray) float64 {
	return s.Inters(r).Front()
}
