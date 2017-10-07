package bruteray

import "math"

func Quad(center Vec, a Vec, b float64, m Material) Obj {
	return &quad{center, a, b, m}
}

// a0 x² + a1 y² + a2 z² = 1
type quad struct {
	c Vec // unused
	a Vec
	b float64
	m Material
}

func (s *quad) Hit(r *Ray, f *[]Shader) {
	a0 := s.a[0]
	a1 := s.a[1]
	a2 := s.a[2]

	s0 := r.Start[0]
	s1 := r.Start[1]
	s2 := r.Start[2]

	d0 := r.Dir[0]
	d1 := r.Dir[1]
	d2 := r.Dir[2]

	A := a0*d0*d0 + a1*d1*d1 + a2*d2*d2
	B := 2 * (a0*d0*s0 + a1*d1*s1 + a2*d2*s2)
	C := a0*s0*s0 + a1*s1*s1 + a2*s2*s2 - s.b

	V := math.Sqrt(B*B - 4*A*C)

	if math.IsNaN(V) {
		return
	}

	t1 := (-B - V) / (2 * A)
	t2 := (-B + V) / (2 * A)

	t1, t2 = sort2(t1, t2)
	*f = append(*f,
		Shader{T: t1, Norm: s.Normal(r.At(t1)), Material: s.m},
		Shader{T: t2, Norm: s.Normal(r.At(t2)), Material: s.m},
	)
}

func (s *quad) Inside(p Vec) bool {
	return s.a.Mul3(p).Dot(p) < s.b
}

func (s *quad) Normal(x Vec) Vec {
	return s.a.Mul3(x).Normalized()
}
