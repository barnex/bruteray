package bruteray

func Slab(dir Vec, off1, off2 float64, m Material) Obj {
	return &slab{dir.Normalized(), off1, off2, m}
}

type slab struct {
	dir        Vec
	off1, off2 float64
	m          Material
}

func (s *slab) Hit(r *Ray, f *[]Surf) {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	if rd == 0 {
		return
	}
	t1 := (s.off1 - rs) / rd
	t2 := (s.off2 - rs) / rd
	t1, t2 = sort(t1, t2)

	*f = append(*f,
		Surf{T: t1, Norm: s.dir, Material: s.m},
		Surf{T: t2, Norm: s.dir, Material: s.m},
	)
}

func (s *slab) Inside(v Vec) bool {
	proj := v.Dot(s.dir)
	return proj > s.off1 && proj < s.off2
}

func (s *slab) Normal(pos Vec) Vec {
	return s.dir
}
