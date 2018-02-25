package shape

import "math"
import . "github.com/barnex/bruteray/br"

// TODO: remove
var CsgAnd_ func(a, b CSGObj) CSGObj

// Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj {
	r := diam / 2
	coeff := Vec{1, 1, 1}
	coeff[dir] = 0
	infCyl := &quad{center, coeff, r * r, m}
	capped := CsgAnd_(infCyl, Slab(Unit[dir], center[dir]-h/2, center[dir]+h/2, m))
	return capped
}

func Quad(center Vec, a Vec, b float64, m Material) CSGObj {
	return &quad{center, a, b, m}
}

// a0 x² + a1 y² + a2 z² = 1
type quad struct {
	c Vec
	a Vec
	b float64
	m Material
}

func (s *quad) Hit1(r *Ray, f *[]Fragment) { s.HitAll(r, f) }

func (s *quad) HitAll(r *Ray, f *[]Fragment) {
	a0 := s.a[0]
	a1 := s.a[1]
	a2 := s.a[2]

	s0 := r.Start[0] - s.c[0]
	s1 := r.Start[1] - s.c[1]
	s2 := r.Start[2] - s.c[2]

	d := r.Dir()
	d0 := d[0]
	d1 := d[1]
	d2 := d[2]

	A := a0*d0*d0 + a1*d1*d1 + a2*d2*d2
	B := 2 * (a0*d0*s0 + a1*d1*s1 + a2*d2*s2)
	C := a0*s0*s0 + a1*s1*s1 + a2*s2*s2 - s.b

	V := math.Sqrt(B*B - 4*A*C)

	if math.IsNaN(V) {
		return
	}

	t1 := (-B - V) / (2 * A)
	t2 := (-B + V) / (2 * A)

	//t1, t2 = sort2(t1, t2)
	*f = append(*f,
		Fragment{T: t1, Norm: s.Normal(r.At(t1)), Material: s.m},
		Fragment{T: t2, Norm: s.Normal(r.At(t2)), Material: s.m},
	)
}

func (s *quad) Inside(p Vec) bool {
	p = p.Sub(s.c)
	return s.a.Mul3(p).Dot(p) < s.b
}

func (s *quad) Normal(x Vec) Vec {
	x = x.Sub(s.c)
	return s.a.Mul3(x)
}
