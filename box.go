package bruteray

import (
	"fmt"
	"math"
)

// -- box (axis aligned)

type box struct {
	min, max Vec
}

func Box(center Vec, rx, ry, rz float64, m Material) CSGObj {
	return &prim{&box{
		min: center.Sub(Vec{rx, ry, rz}),
		max: center.Add(Vec{rx, ry, rz}),
	}, m}
}

func Cube(center Vec, r float64, m Material) CSGObj {
	return Box(center, r, r, r, m)
}

func (s *box) Inters2(r *Ray) Interval {
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

func (s *box) Inters(r *Ray) []Interval {
	return s.Inters2(r).Slice()
}

func (s *box) Hit(r *Ray) float64 {
	return s.Inters2(r).Front()
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

// -- slab

func Slab(dir Vec, off1, off2 float64, m Material) CSGObj {
	return &prim{&slab{dir, off1, off2}, m}
}

type slab struct {
	dir        Vec
	off1, off2 float64
}

func (s *slab) Normal(pos Vec) Vec {
	return s.dir
}

func (s *slab) Inters2(r *Ray) Interval {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t1 := (s.off1 - rs) / rd
	t2 := (s.off2 - rs) / rd
	t1, t2 = sort(t1, t2)
	return Interval{t1, t2}.Fix().check()
}

func (s *slab) Inters(r *Ray) []Interval {
	return s.Inters2(r).Slice()
}

func (s *slab) Hit(r *Ray) float64 {
	return s.Inters2(r).Front()
}
