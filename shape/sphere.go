package shape

import "math"
import . "github.com/barnex/bruteray/br"

func NewSphere(diam float64, m Material) *Sphere {
	return &Sphere{Vec{}, sqr(diam / 2), m}
}

type Sphere struct {
	c  Vec
	r2 float64
	m  Material
}

func (s *Sphere) Inside(p Vec) bool {
	v := p.Sub(s.c)
	return v.Len2() < s.r2
}

func (s *Sphere) Transl(d Vec) *Sphere {
	s.c.Transl(d)
	return s
}

func (s *Sphere) Hit1(r *Ray, f *[]Fragment) { s.HitAll(r, f) }

func (s *Sphere) HitAll(r *Ray, f *[]Fragment) {
	v := r.Start.Sub(s.c)
	d := r.Dir()
	vd := v.Dot(d)
	D := vd*vd - (v.Len2() - s.r2)
	if D < 0 {
		return
	}
	t1 := (-vd - math.Sqrt(D))
	t2 := (-vd + math.Sqrt(D))

	*f = append(*f,
		Fragment{T: t1, Norm: s.Normal(r.At(t1)), Material: s.m},
		Fragment{T: t2, Norm: s.Normal(r.At(t2)), Material: s.m},
	)
}

func (s *Sphere) Normal(pos Vec) Vec {
	n := pos.Sub(s.c)
	return n
}
