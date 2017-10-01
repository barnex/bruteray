package bruteray

// -- sheet (infinite)

func Sheet(dir Vec, off float64) Obj {
	return &sheet{dir: dir, off: off}
}

type sheet struct {
	dir Vec
	off float64
	noInside
}

//func (s *sheet) Normal(pos Vec) Vec {
//	return s.dir
//}

func (s *sheet) Hit(r *Ray, f *[]Surf) {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t := (s.off - rs) / rd

	*f = append(*f, Surf{T: t, Norm: s.dir})
}

type noInside struct{}

func (noInside) Inside(Vec) bool {
	return false
}

// --rectangle (finite sheet)

//// A rectangle (i.e. finite sheet) at given position,
//// with normal vector dir and half-axes rx, ry, rz.
//func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj {
//	return &rect{pos, dir, rx, ry, rz}
//}
//
//type rect struct {
//	pos, dir   Vec
//	rx, ry, rz float64
//}
//
//// TODO: rm
//func (s *rect) Inters2(r *Ray) Interval {
//	rs := r.Start.Dot(s.dir)
//	rd := r.Dir.Dot(s.dir)
//	t := (s.pos.Dot(s.dir) - rs) / rd
//	p := r.At(t).Sub(s.pos)
//	if p[X] < -s.rx || p[X] > s.rx ||
//		p[Y] < -s.ry || p[Y] > s.ry ||
//		p[Z] < -s.rz || p[Z] > s.rz {
//		return Interval{}
//	}
//	return Interval{t, t}.Fix().check()
//}
//
//func (s *rect) Inters(r *Ray) []Interval {
//	return s.Inters2(r).Slice()
//}
//
//func (s *rect) Hit(r *Ray) float64 {
//	return s.Inters2(r).Front()
//}
//
//func (s *rect) Normal(p Vec) Vec {
//	return s.dir
//}
