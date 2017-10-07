package bruteray

// A rectangle (i.e. finite sheet) at given position,
// with normal vector dir and half-axes rx, ry, rz.
func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj {
	return &rect{pos, dir, rx, ry, rz, m, noInside{}}
}

type rect struct {
	pos, dir   Vec
	rx, ry, rz float64
	m          Material
	noInside
}

func (s *rect) Hit(r *Ray, f *[]Shader) {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t := (s.pos.Dot(s.dir) - rs) / rd
	p := r.At(t).Sub(s.pos)
	if p[X] < -s.rx || p[X] > s.rx ||
		p[Y] < -s.ry || p[Y] > s.ry ||
		p[Z] < -s.rz || p[Z] > s.rz {
		return
	}
	*f = append(*f,
		Shader{T: t, Norm: s.dir, Material: s.m},
	)
}
