package shape

import . "github.com/barnex/bruteray/br"

// A rectangle (i.e. finite sheet) at given position,
// with normal vector dir and half-axes rx, ry, rz.
//
// TODO: pass Vec normal, U, V
func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj {
	return &rect{pos, dir, rx, ry, rz, m}
}

type rect struct {
	pos, dir   Vec
	rx, ry, rz float64
	m          Material
}

func (s *rect) Hit1(r *Ray, f *[]Fragment) { s.HitAll(r, f) }

func (s *rect) HitAll(r *Ray, f *[]Fragment) {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir().Dot(s.dir)
	t := (s.pos.Dot(s.dir) - rs) / rd
	p := r.At(t).Sub(s.pos)
	if p[X] < -s.rx || p[X] > s.rx ||
		p[Y] < -s.ry || p[Y] > s.ry ||
		p[Z] < -s.rz || p[Z] > s.rz {
		return
	}
	*f = append(*f,
		Fragment{T: t, Norm: s.dir, Material: s.m},
	)
}
