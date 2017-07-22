package bruteray

type Obj interface {
	Inters(r *Ray) BiSurf
}

// -- Primitive Object (shape + material)
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
	orig Obj
	t    Matrix4
	tinv Matrix4
}

func Transf(o Obj, T *Matrix4) Obj {
	return &transObj{
		o,
		*T,
		*(T.Inv()),
	}
}

func (o *transObj) Inters(r *Ray) BiSurf {
	r2 := *r
	r2.Transf(&o.tinv)
	bi := o.orig.Inters(&r2)
	bi.S1.Norm = (&o.t).TransfDir(bi.S1.Norm)
	bi.S2.Norm = (&o.t).TransfDir(bi.S2.Norm)
	return bi
}

// -- CSG

// Intersection (boolean AND) of two objects.
func ObjAnd(a, b Obj) Obj {
	return &objAnd{a, b}
}

type objAnd struct {
	a, b Obj
}

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

// Carve away object b from a.
func ObjMinus(a, b Obj) Obj {
	return &objMinus{a, b}
}

type objMinus struct {
	a, b Obj
}

func (o *objMinus) Inters(r *Ray) BiSurf {
	// not intersecting A = not intersecting anything
	A := o.a.Inters(r)
	if !A.OK() {
		return BiSurf{}
	}

	// not intersecting b = intersecting A fully
	B := o.b.Inters(r)
	if !B.OK() {
		return A
	}

	// disjoint intervals =  intersecting A fully
	if A.Max() < B.Min() || B.Max() < A.Min() {
		return A
	}

	// non-trivial cases
	var bi BiSurf
	if B.Min() < A.Min() {
		bi.S1 = B.S2
		bi.S2 = A.S2
	} else {
		bi.S1 = B.S1
		bi.S2 = A.S1
	}
	bi.Normalize()
	return bi
}
