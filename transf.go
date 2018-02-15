package bruteray

// Transf returns a transformed version of the object.
func Transf(o CSGObj, T *Matrix4) CSGObj {
	return &transformed{
		o,
		*T,
		*(T.Inv()),
	}
}

type transformed struct {
	orig CSGObj
	t    Matrix4
	tinv Matrix4
}

func (o *transformed) Hit(r *Ray, f *[]Fragment) {
	backup := *r

	r.Transf(&o.tinv)
	o.orig.Hit(r, f)
	for i := range *f {
		(*f)[i].Norm = (&o.t).TransfDir((*f)[i].Norm)
	}

	*r = backup
}

func (o *transformed) Inside(p Vec) bool {
	// TODO: currently untested
	return o.orig.Inside(o.tinv.TransfPoint(p))
}
