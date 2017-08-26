package bruteray

type BiSurf struct {
	S1, S2 Surf
}

func (b *BiSurf) Min() float64 {
	return b.S1.T
}

func (b *BiSurf) Max() float64 {
	return b.S2.T
}

func (b BiSurf) Front() Surf {
	if b.S1.T > 0 {
		return b.S1
	}
	return b.S2
}

func (b *BiSurf) OK() bool {
	return b.Min() != 0 && b.Max() != 0
}

// swap surfaces if not in sorted order
func (b *BiSurf) Normalize() {
	if b.S1.T > b.S2.T {
		b.S1, b.S2 = b.S2, b.S1
	}
}

type Surf struct {
	T    float64
	Norm Vec
	Material
}

func (s *Surf) Shade(e *Env, N int, r *Ray) Color {
	pos := r.At(s.T)
	norm := s.Norm.Towards(r.Dir)
	return s.Material.Shade(e, r, N, pos, norm)
}
