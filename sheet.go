package bruteray

// -- sheet (infinite)

func Sheet(dir Vec, off float64, m Material) Obj {
	return &prim{&sheet{dir, off}, m}
}

type sheet struct {
	dir Vec
	off float64
}

func (s *sheet) Normal(pos Vec) Vec {
	return s.dir
}

// TODO: rm, is not a CSGObj
func (s *sheet) Inters2(r *Ray) Interval {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t := (s.off - rs) / rd
	return Interval{t, t}.Fix().check()
}

func (s *sheet) Inters(r *Ray) []Interval {
	return s.Inters2(r).Slice()
}

func (s *sheet) Hit(r *Ray) float64 {
	return s.Inters2(r).Front()
}

// --rectangle (finite sheet)

// A rectangle (i.e. finite sheet) at given position,
// with normal vector dir and half-axes rx, ry, rz.
func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj {
	return &prim{&rect{pos, dir, rx, ry, rz}, m}
}

type rect struct {
	pos, dir   Vec
	rx, ry, rz float64
}

// TODO: rm
func (s *rect) Inters2(r *Ray) Interval {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t := (s.pos.Dot(s.dir) - rs) / rd
	p := r.At(t).Sub(s.pos)
	if p[X] < -s.rx || p[X] > s.rx ||
		p[Y] < -s.ry || p[Y] > s.ry ||
		p[Z] < -s.rz || p[Z] > s.rz {
		return Interval{}
	}
	return Interval{t, t}.Fix().check()
}

func (s *rect) Inters(r *Ray) []Interval {
	return s.Inters2(r).Slice()
}

func (s *rect) Hit(r *Ray) float64 {
	return s.Inters2(r).Front()
}

func (s *rect) Normal(p Vec) Vec {
	return s.dir
}
