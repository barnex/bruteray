package r

type BiSurf struct {
	S1, S2 Surf
}

func (b *BiSurf) Min() float64 { return b.S1.T }
func (b *BiSurf) Max() float64 { return b.S2.T }
func (b *BiSurf) OK() bool     { return b.Min() != 0 && b.Max() != 0 }

type Surf struct {
	T    float64
	Norm Vec
	Material
}

func (s *Surf) Shade(e *Env, N int, r *Ray) Color {
	pos := r.At(s.T)
	norm := s.Norm.Towards(r.Dir)
	return s.Material.Shade(e, N, pos, norm)
}
