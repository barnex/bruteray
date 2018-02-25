package shape

import "github.com/barnex/bruteray/br"

func NCyl(dir int, diam float64, m br.Material) *cyl {
	r := diam / 2
	coeff := br.Vec{1, 1, 1}
	coeff[dir] = 0
	return &cyl{quad{br.Vec{}, coeff, r * r, m}}
}

type cyl struct {
	quad
}

func (s *cyl) Transl(d br.Vec) *cyl {
	s.quad.c.Transl(d)
	return s
}
