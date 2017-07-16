package main

type slab struct {
	min, max float64
	dir      Vec
}

// A slab along Y (XZ plane, aka horizontal)
func Slab(min, max float64) *slab {
	return &slab{min, max, Vec{0, 1, 0}}
}

// A slab along a certain (normal) direction,
// e.g. Ey. for horizontal.
func SlabD(min, max float64, dir Vec) *slab {
	return &slab{min, max, dir}
}

func (s *slab) Inters(r *Ray) Interval {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t0 := (s.min - rs) / rd
	t1 := (s.max - rs) / rd
	t0, t1 = Sort(t0, t1)
	return Interval{t0, t1}
}

func (s *slab) Normal(r *Ray, t float64) Vec {
	return s.dir.Towards(r.Dir)
}
