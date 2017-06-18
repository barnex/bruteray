package main

type Scene struct {
	objs []Obj
}

func (s *Scene) Intensity(r Ray) (float64, Color) {
	ival, obj := s.Intersect(r)
	if obj != nil {
		return ival.Min, obj.Intensity(r, ival.Min)
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

func (s *Scene) Intersect(r Ray) (Inter, Obj) {
	var (
		minT     = Inter{inf, inf}
		obj  Obj = nil
	)

	for _, o := range s.objs {
		ival, sub := o.Inters(r)
		if ival.Min < minT.Min && ival.Max > 0 {
			minT = ival
			obj = sub
		}
	}
	return minT, obj
}
