package r

type Obj interface {
	Inters(r *Ray) BiSurf
}

func ObjShade(o Obj, e *Env, N int, r *Ray) Color {
	bi := o.Inters(r)
	return bi.S1.Shade(e, N, r)
}

// -- Primitive Object

func Object(s Shape, m Material) Obj {
	return &prim{s, m}
}

type prim struct {
	s Shape
	m Material
}

var _ Obj = (*prim)(nil)

func (o *prim) Inters(r *Ray) BiSurf {
	i := o.s.Inters(r)
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

//func(p*prim)

// -- Transforms

type transObj struct {
}

// -- CSG

func ObjAnd(a, b Obj) Obj {
	return &objAnd{a, b}
}

type objAnd struct{ a, b Obj }

func (o *objAnd) Inters(r *Ray) BiSurf {
	A := o.a.Inters(r)
	if !A.OK() {
		return A
	}

	B := o.b.Inters(r)
	if !B.OK() {
		return B
	}

	if A.Max() < B.Min() || B.Max() < A.Min() {
		return BiSurf{}
	}

	var bi BiSurf

	if A.Min() > B.Min() {
		bi.S1 = A.S1
	} else {
		bi.S1 = B.S1
	}

	if A.Max() < B.Max() {
		bi.S2 = A.S2
	} else {
		bi.S2 = B.S2
	}

	return bi
}