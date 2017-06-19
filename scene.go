package main

type Scene struct {
	objs    []Obj
	sources []Source
	amb     func(Vec) Color
}

func (s *Scene) Ambient(r Ray) Color {
	if s.amb == nil {
		return 0
	}
	return s.amb(r.Dir)
}

func (s *Scene) Intensity(r Ray, N int) (float64, Color) {
	if N == 0 {
		return inf, s.Ambient(r)
	}
	ival, shader := s.Intersect(r)
	if shader != nil {
		return ival.Min, shader.Intensity(r, ival.Min, N)
	} else {
		return inf, s.Ambient(r)
	}
}

func (s *Scene) Intersect(r Ray) (Inter, Shader) {
	var (
		minT          = Inter{inf, inf}
		shader Shader = nil
	)

	for _, o := range s.objs {
		ival, s := o.Intersect(r)
		if ival.Min < minT.Min && ival.Max > 0 {
			minT = ival
			shader = s
		}
	}
	return minT, shader
}

func (s *Scene) IntersectsAny(r Ray) bool {
	_, shader := s.Intersect(r)
	return shader != nil
}
