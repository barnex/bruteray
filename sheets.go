package main

type sheet struct {
	pos float64
	dir Vec
}

// An infinite sheet perpendicular to dir,
// at distance pos from the origin.
// E.g.: pos=1, dir=Ey for a horizontal sheet at Y=1.
func Sheet(pos float64, dir Vec) Shape {
	return &sheet{pos, dir}
}

func (s *sheet) Inters(r *Ray) Inter {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t := (s.pos - rs) / rd
	return Inter{t, t}
}

func (s *sheet) Normal(r *Ray, t float64) Vec {
	return s.dir.Towards(r.Dir)
}

type rect struct {
	pos, dir   Vec
	rx, ry, rz float64
}

// A rectangle (i.e. finite sheet) at given position,
// with normal vector dir and half-axes rx, ry, rz.
func Rect(pos, dir Vec, rx, ry, rz float64) Shape {
	return &rect{pos, dir, rx, ry, rz}
}

func (s *rect) Inters(r *Ray) Inter {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t := (s.pos.Dot(s.dir) - rs) / rd
	if t < 0 {
		return empty
	}
	p := r.At(t).Sub(s.pos)
	if p.X < -s.rx || p.X > s.rx ||
		p.Y < -s.ry || p.Y > s.ry ||
		p.Z < -s.rz || p.Z > s.rz {
		return empty
	}
	return Inter{t, t}
}

func (s *rect) Normal(r *Ray, t float64) Vec {
	return s.dir.Towards(r.Dir)
}
