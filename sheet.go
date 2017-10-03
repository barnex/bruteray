package bruteray

func Sheet(dir Vec, off float64, m Material) Obj {
	return &sheet{dir: dir, off: off, m: m}
}

type sheet struct {
	dir Vec
	off float64
	m   Material
	noInside
}

func (s *sheet) Hit(r *Ray, f *[]Surf) {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t := (s.off - rs) / rd

	*f = append(*f, Surf{T: t, Norm: s.dir, Material: s.m})
}
