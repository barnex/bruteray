package r

type Env struct {
	objs    []Obj
	Ambient Surf
}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) Add(o Obj) {
	e.objs = append(e.objs, o)
}

func (e *Env) Shade(r *Ray, N int) Color {
	if N == 0 {
		return Color{}
	}

	surf := e.Ambient
	surf.T = inf

	for _, o := range e.objs {
		bi := o.Inters(r)
		if !bi.OK() {
			continue
		}
		if bi.S1.T < surf.T {
			surf = bi.S1
		}
	}

	return surf.Shade(e, N-1, r)
}
