package main

type Scene struct {
	objs []Obj
}

func (s *Scene) Intensity(r Ray) (float64, Color) {
	ival, shader := s.Intersect(r)
	if shader != nil {
		return ival.Min, shader.Intensity(r, ival.Min)
	} else {
		return inf, 0 //ambient(r.Dir)
	}
}

//func (s *Scene) ZMap(r Ray) float64 {
//	t, obj := s.Intersect(r)
//	if obj != nil {
//		return t
//	} else {
//		return inf
//	}
//}

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
