package main

type slab struct {
	min, max float64
	dir      Vec
}

// A slab along Y (XZ plane, aka horizontal)
func Slab(min, max float64) *slab {
	return &slab{min, max, Vec{0, 1, 0}}
}

func SlabD(min, max float64, dir Vec) *slab {
	return &slab{min, max, dir}
}

func (s *slab) Inters(r *Ray) Inter {
	rs := r.Start.Dot(s.dir)
	rd := r.Dir.Dot(s.dir)
	t0 := (s.min - rs) / rd
	t1 := (s.max - rs) / rd
	t0, t1 = Sort(t0, t1)
	return Inter{t0, t1}
}

func (s *slab) Normal(r Ray, t float64) Vec {
	return s.dir.Towards(r.Dir)
}

func (s *slab) Hit(r *Ray) float64 {
	return s.Inters(r).Min
}
