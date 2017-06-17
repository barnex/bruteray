package main

type Scene struct {
	objs []Obj
}

func (s *Scene) Intensity(r Ray) float64 {
	t, obj := s.FirstIntersect(r)
	if obj != nil {
		return obj.Intensity(r, t)
	} else {
		return 0 //ambient(r.Dir)
	}
}

func (s *Scene) FirstIntersect(r Ray) (float64, Obj) {
	var (
		minT     = inf
		obj  Obj = nil
	)

	for _, o := range s.objs {
		ival := o.Inters(r)
		if ival.Min < minT && ival.Max > 0 {
			minT = ival.Min
			obj = o
		}
	}
	return minT, obj
}
