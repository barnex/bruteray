package main

type slab struct {
	min, max float64
}

func Slab(min, max float64) *slab {
	return &slab{min, max}
}

func (s *slab) Intersect(r Ray) Inter {
	t0 := (s.min - r.Start.Y) / r.Dir.Y
	t1 := (s.max - r.Start.Y) / r.Dir.Y
	return Inter{t0, t1}
}
