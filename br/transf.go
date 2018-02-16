package br

// Transf returns a transformed version of the object.
// TODO: also for non-csg?
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

func (o *transformed) Hit1(r *Ray, f *[]Fragment) { o.HitAll(r, f) }

func (o *transformed) HitAll(r *Ray, f *[]Fragment) {
	backup := *r

	r.Transf(&o.tinv)
	o.orig.HitAll(r, f)
	for i := range *f {
		(*f)[i].Norm = (&o.t).TransfDir((*f)[i].Norm)
	}

	*r = backup
}

func (o *transformed) Inside(p Vec) bool {
	// TODO: currently untested
	return o.orig.Inside(o.tinv.TransfPoint(p))
}

//TODO: rename
func TransfNonCSG(o Obj, T *Matrix4) Obj {
	return &transformedNonCSG{
		o,
		*T,
		*(T.Inv()),
	}
}

type transformedNonCSG struct {
	orig Obj
	t    Matrix4
	tinv Matrix4
}

func (o *transformedNonCSG) Hit1(r *Ray, f *[]Fragment) { o.HitAll(r, f) }

func (o *transformedNonCSG) HitAll(r *Ray, f *[]Fragment) {
	backup := *r

	r.Transf(&o.tinv)
	o.orig.Hit1(r, f)
	for i := range *f {
		(*f)[i].Norm = (&o.t).TransfDir((*f)[i].Norm)
	}

	*r = backup
}
