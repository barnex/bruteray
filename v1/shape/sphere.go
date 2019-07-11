package shape

import "math"
import . "github.com/barnex/bruteray/v1/br"

func NewSphere(diam float64, m Material) *Sphere {
	return &Sphere{Vec{}, sqr(diam / 2), m}
}

type Sphere struct {
	Center Vec
	r2     float64
	Mat    Material
}

func (s *Sphere) Radius() float64 {
	return math.Sqrt(s.r2)
}

func (s *Sphere) Inside(p Vec) bool {
	v := p.Sub(s.Center)
	return v.Len2() < s.r2
}

func (s *Sphere) Transl(d Vec) *Sphere {
	s.Center.Transl(d)
	return s
}

func (s *Sphere) Hit1(r *Ray, f *[]Fragment) { s.HitAll(r, f) }

func (s *Sphere) HitAll(r *Ray, f *[]Fragment) {
	v := r.Start.Sub(s.Center)
	d := r.Dir()
	vd := v.Dot(d)
	D := vd*vd - (v.Len2() - s.r2)
	if D < 0 {
		return
	}
	t1 := (-vd - math.Sqrt(D))
	t2 := (-vd + math.Sqrt(D))

	*f = append(*f,
		Fragment{T: t1, Norm: s.Normal(r.At(t1)), Material: s.Mat},
		Fragment{T: t2, Norm: s.Normal(r.At(t2)), Material: s.Mat},
	)
}

func (s *Sphere) Normal(pos Vec) Vec {
	n := pos.Sub(s.Center)
	return n
}
