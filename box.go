package bruteray

import (
	"fmt"
	"math"
)

type box struct {
	min, max Vec
	m        Material
}

func Box(center Vec, rx, ry, rz float64, m Material) Obj {
	return &box{
		min: center.Sub(Vec{rx, ry, rz}),
		max: center.Add(Vec{rx, ry, rz}),
		m:   m,
	}
}

func Cube(center Vec, r float64, m Material) Obj {
	return Box(center, r, r, r, m)
}

func (s *box) Hit(r *Ray, f *[]Shader) {
	min_ := s.min
	max_ := s.max

	tmin := min_.Sub(r.Start).Mul3(r.InvDir)
	tmax := max_.Sub(r.Start).Mul3(r.InvDir)

	txen := min(tmin[X], tmax[X])
	txex := max(tmin[X], tmax[X])

	tyen := min(tmin[Y], tmax[Y])
	tyex := max(tmin[Y], tmax[Y])

	tzen := min(tmin[Z], tmax[Z])
	tzex := max(tmin[Z], tmax[Z])

	ten := max3(txen, tyen, tzen)
	tex := min3(txex, tyex, tzex)

	if ten > tex {
		return
	}

	if math.IsNaN(ten) || math.IsNaN(tex) {
		return
	}

	*f = append(*f,
		Shader{T: ten, Norm: s.Normal(r.At(ten)), Material: s.m},
		Shader{T: tex, Norm: s.Normal(r.At(tex)), Material: s.m},
	)

}

func (s *box) Inside(v Vec) bool {
	return v[X] > s.min[X] && v[X] < s.max[X] &&
		v[Y] > s.min[Y] && v[Y] < s.max[Y] &&
		v[Z] > s.min[Z] && v[Z] < s.max[Z]
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
