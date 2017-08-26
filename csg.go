package bruteray

type CSGObj interface {
	Obj
	Inters(r *Ray) BiSurf
}

// -- CSG

// Intersection (boolean AND) of two objects.
func ObjAnd(a, b CSGObj) CSGObj {
	return &objAnd{a, b}
}

type objAnd struct {
	a, b CSGObj
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

func (o *objAnd) Hit(r *Ray) Surf {
	return o.Inters(r).Front()
}

func ObjOr(a, b CSGObj) CSGObj {
	return &objOr{a, b}
}

type objOr struct {
	a, b CSGObj
}

func (o *objOr) Inters(r *Ray) BiSurf {

	A := o.a.Inters(r)
	B := o.b.Inters(r)

	// ray only hits one
	if !A.OK() {
		return B
	}
	if !B.OK() {
		return A
	}

	// sort A in front of B
	if A.S1.T > B.S1.T {
		A, B = B, A
	}

	var bi BiSurf
	bi.S1 = A.S1 // enter A

	if A.S2.T < B.S1.T {
		bi.S2 = A.S2 // non-overlapping: exit from A
	} else {
		bi.S2 = B.S2 // overlapping: exit from B
	}

	return bi
}

func (o *objOr) Hit(r *Ray) Surf {
	return o.Inters(r).Front()
}

// Carve away object b from a.
func ObjMinus(a, b CSGObj) CSGObj {
	return &objMinus{a, b}
}

type objMinus struct {
	a, b CSGObj
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

func (o *objMinus) Hit(r *Ray) Surf {
	return o.Inters(r).Front()
}
