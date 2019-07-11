package shape

import . "github.com/barnex/bruteray/v1/br"

func NewSheet(dir Vec, off float64, m Material) *Sheet {
	return &Sheet{dir: dir, off: off, m: m}
}

type Sheet struct {
	dir Vec
	off float64
	m   Material
}

func (s *Sheet) Hit1(r *Ray, f *[]Fragment) {
	t := s.hit(r)
	if t > 0 {
		*f = append(*f, Fragment{T: t, Norm: s.dir, Material: s.m})
	}
}

func (s *Sheet) hit(r *Ray) float64 {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir().Dot(s.dir)
	t := (s.off - rs) / rd
	return t
}
