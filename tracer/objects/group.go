package objects

//import (
//	. "github.com/barnex/bruteray/tracer/types"
//)
//
//// TODO: remove in favor of tree
//func Group(obj ...Object) Object {
//	return group(obj)
//}
//
//type group []Object
//
//func (g group) Intersect(r *Ray) HitRecord {
//	if len(g) == 0 {
//		return HitRecord{}
//	}
//	front := HitRecord{T: inf}
//	ok := false
//	for _, o := range g {
//		h := o.Intersect(r)
//		if h.T > 0 && h.T < front.T {
//			front = h
//			ok = true
//		}
//	}
//	if !ok {
//		return HitRecord{}
//	}
//	return front
//}
//
