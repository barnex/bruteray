package main

type slab struct {
	min, max float64
}

// A slab along Y (XZ plane, aka horizontal)
func Slab(min, max float64) *slab {
	return &slab{min, max}
}

func (s *slab) Intersect(r Ray) Inter {
	t0 := (s.min - r.Start.Y) / r.Dir.Y
	t1 := (s.max - r.Start.Y) / r.Dir.Y
	t0, t1 = sort(t0, t1)
	return Inter{t0, t1}
}

func (s *slab) Normal(r Ray, t float64) Vec {
	return Vec{0, 1, 0}.Towards(r.Dir)
}

func sort(t0, t1 float64) (float64, float64) {
	if t0 < t1 {
		return t0, t1
	}
	return t1, t0
}
