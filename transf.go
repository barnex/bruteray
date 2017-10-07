package bruteray

// Transf returns a transformed version of the object.
func Transf(o Obj, T *Matrix4) Obj {
	return &transformed{
		o,
		*T,
		*(T.Inv()),
	}
}

type transformed struct {
	orig Obj
	t    Matrix4
	tinv Matrix4
}

func (o *transformed) Hit(r *Ray, f *[]Shader) {
	r2 := *r
	r2.Transf(&o.tinv)
	o.orig.Hit(&r2, f)
	for i := range *f {
		(*f)[i].Norm = (&o.t).TransfDir((*f)[i].Norm)
	}
}

func (o *transformed) Inside(p Vec) bool {
	// TODO: currently untested
	return o.orig.Inside(o.tinv.TransfPoint(p))
}
