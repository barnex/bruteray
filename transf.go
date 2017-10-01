package bruteray

type transObj struct {
	orig Obj
	t    Matrix4
	tinv Matrix4
}

// TODO: non-csg version
func Transf(o Obj, T *Matrix4) Obj {
	return &transObj{
		o,
		*T,
		*(T.Inv()),
	}
}

//func (o *transObj) Inters(r *Ray) []BiSurf {
//	r2 := *r
//	r2.Transf(&o.tinv)
//
//	bi := o.orig.Inters(&r2)
//
//	for i := range bi {
//		bi[i].S1.Norm = (&o.t).TransfDir(bi[i].S1.Norm)
//		bi[i].S2.Norm = (&o.t).TransfDir(bi[i].S2.Norm)
//	}
//	return bi
//}

func (o *transObj) Hit(r *Ray) Surf {
	r2 := *r
	r2.Transf(&o.tinv)
	return o.orig.Hit(&r2)
}
