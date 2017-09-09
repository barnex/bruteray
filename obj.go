package bruteray

type Obj interface {
	Hit(r *Ray) Surf
}

// -- Primitive Object (shape + material)
type prim struct {
	s Shape
	m Material
}

var _ Obj = (*prim)(nil)

func (o *prim) Inters2(r *Ray) BiSurf {
	i := o.s.Inters2(r)
	if !i.OK() {
		return BiSurf{}
	}
	i.check()
	n1 := o.s.Normal(r.At(i.Min))
	n2 := o.s.Normal(r.At(i.Max))
	return BiSurf{
		S1: Surf{T: i.Min, Norm: n1, Material: o.m},
		S2: Surf{T: i.Max, Norm: n2, Material: o.m},
	}
}

func (o *prim) Hit(r *Ray) Surf {
	return o.Inters2(r).Front()
}

//func(p*prim)

// -- Transforms

type transObj struct {
	orig CSGObj
	t    Matrix4
	tinv Matrix4
}

// TODO: non-csg version
func Transf(o CSGObj, T *Matrix4) CSGObj {
	return &transObj{
		o,
		*T,
		*(T.Inv()),
	}
}

func (o *transObj) Inters2(r *Ray) BiSurf {
	r2 := *r
	r2.Transf(&o.tinv)
	bi := o.orig.Inters2(&r2)
	bi.S1.Norm = (&o.t).TransfDir(bi.S1.Norm)
	bi.S2.Norm = (&o.t).TransfDir(bi.S2.Norm)
	return bi
}

func (o *transObj) Hit(r *Ray) Surf {
	return o.Inters2(r).Front()
}
