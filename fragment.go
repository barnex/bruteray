package bruteray

type Surf struct {
	T    float64
	Norm Vec
	Material
}

func (s *Surf) Shade(e *Env, N int, r *Ray) Color {
	pos := r.At(s.T)
	norm := s.Norm.Towards(r.Dir)
	pos = pos.MAdd(offset, norm)
	return s.Material.Shade(e, r, N, pos, norm)
}
