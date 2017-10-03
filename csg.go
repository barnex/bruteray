package bruteray

// Intersection (boolean AND) of two objects.
//func And(a, b Obj) Obj {
//	return &objAnd{a, b}
//}
//
//type and struct {
//	a, b Obj
//}
//
//func (o *and) Hit(r *Ray, f *[]Surf) {
//
//	o.a.Hit(r, f)
//	if len(*f) == 0 {
//		return
//	}
//
//	var f2 []Surf
//	o.b.Hit(r, f2)
//
//	var f3 []Surf
//
//	for _, s := range *f{
//
//	}
//
//}

func HAnd(a, b Obj) Obj {
	return &hand{a: a, b: b}
}

type hand struct {
	a, b Obj
	noInside
}

func (o *hand) Hit(r *Ray, f *[]Surf) {

	o.a.Hit(r, f)
	if len(*f) == 0 {
		return
	}

	f2 := make([]Surf, 0, len(*f))
	for i, s := range *f {
		if o.b.Inside(r.At(s.T)) {
			f2 = append(f2, (*f)[i])
		}
	}
	*f = f2
}

//func toBi(x []BiSurf) BiSurf {
//	switch len(x) {
//	case 0:
//		return BiSurf{}
//	case 1:
//		return x[0]
//	default:
//		panic(fmt.Sprintf("TODO: CSG on > 1 interval"))
//	}
//}
//
//func (o *objAnd) Inters(r *Ray) []BiSurf { return o.Inters2(r).Slice() }
//
//func (o *objAnd) Hit(r *Ray) Surf {
//	return o.Inters2(r).Front()
//}
//
//func Or(a, b CSGObj) CSGObj {
//	return &objOr{a, b}
//}
//
//type objOr struct {
//	a, b CSGObj
//}
//
//func (o *objOr) Inters2(r *Ray) BiSurf {
//
//	A := toBi(o.a.Inters(r))
//	B := toBi(o.b.Inters(r))
//
//	// ray only hits one
//	if !A.OK() {
//		return B
//	}
//	if !B.OK() {
//		return A
//	}
//
//	// sort A in front of B
//	if A.S1.T > B.S1.T {
//		A, B = B, A
//	}
//
//	var bi BiSurf
//	bi.S1 = A.S1 // enter A
//
//	if A.S2.T < B.S1.T {
//		bi.S2 = A.S2 // non-overlapping: exit from A
//	} else {
//		bi.S2 = B.S2 // overlapping: exit from B
//	}
//
//	return bi
//}
//
//func (o *objOr) Inters(r *Ray) []BiSurf { return o.Inters2(r).Slice() }
//
//func (o *objOr) Hit(r *Ray) Surf {
//	return o.Inters2(r).Front()
//}
//
//// Carve away object b from a.
//func Minus(a, b CSGObj) CSGObj {
//	return &objMinus{a, b}
//}
//
//type objMinus struct {
//	a, b CSGObj
//}
//
//func (o *objMinus) Inters2(r *Ray) BiSurf {
//	// not intersecting A = not intersecting anything
//	A := toBi(o.a.Inters(r))
//	if !A.OK() {
//		return BiSurf{}
//	}
//
//	// not intersecting b = intersecting A fully
//	B := toBi(o.b.Inters(r))
//	if !B.OK() {
//		return A
//	}
//
//	// disjoint intervals =  intersecting A fully
//	if A.Max() < B.Min() || B.Max() < A.Min() {
//		return A
//	}
//
//	// non-trivial cases
//	var bi BiSurf
//	if B.Min() < A.Min() {
//		bi.S1 = B.S2
//		bi.S2 = A.S2
//	} else {
//		bi.S1 = B.S1
//		bi.S2 = A.S1
//	}
//	bi.Normalize()
//	return bi
//}
//
//func (o *objMinus) Hit(r *Ray) Surf {
//	return o.Inters2(r).Front()
//}
//
//func (o *objMinus) Inters(r *Ray) []BiSurf { return o.Inters2(r).Slice() }
