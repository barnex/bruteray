package bruteray

import "math"

func Sphere(center Vec, radius float64) Obj {
	return &sphere{center, sqr(radius)}
}

type sphere struct {
	c  Vec
	r2 float64
}

func (s *sphere) Inside(p Vec) bool {
	v := p.Sub(s.c)
	return v.Len2() < s.r2
}

func (s *sphere) Hit(r *Ray, f *[]Surf) {
	v := r.Start.Sub(s.c)
	d := r.Dir
	vd := v.Dot(d)
	D := sqr(vd) - (v.Len2() - s.r2)
	if D < 0 {
		return
	}
	t1 := (-vd - math.Sqrt(D))
	t2 := (-vd + math.Sqrt(D))

	*f = append(*f,
		Surf{T: t1, Norm: s.Normal(r.At(t1))},
		Surf{T: t2, Norm: s.Normal(r.At(t2))},
	)
}

func (s *sphere) Normal(pos Vec) Vec {
	n := pos.Sub(s.c).Normalized().check()
	return n
}
