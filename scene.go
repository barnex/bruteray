package main

type Env struct {
	objs    []Obj
	sources []Source
	amb     func(Vec) Color
}

func (s *Env) Shade(r Ray, _ float64, N int) Color {
	if N == 0 {
		return s.Ambient(r)
	}
	t, shader := s.Intersect(r)
	if shader != nil {
		return shader.Intensity(r, t, N)
	} else {
		return s.Ambient(r)
	}
}

func (s *Env) Intersect(r Ray) (float64, Shader) {
	var (
		minT   float64 = 0
		shader Shader  = nil
	)

	for i, o := range s.objs {
		t := o.Hit(r)
		if t < 0 {
			panic(fmt.Sprintf("object %v: %#v: t=%v", i, o, t))
		}
		if t < minT && t > 0 {
			minT = t
			shader = o
		}
	}
	return minT, shader
}

func (s *Env) IntersectsAny(r Ray) bool {
	_, shader := s.Intersect(r)
	return shader != nil
}

func (s *Env) Ambient(r Ray) Color {
	if s.amb == nil {
		return 0
	}
	return s.amb(r.Dir)
}
