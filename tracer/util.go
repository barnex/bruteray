package tracer

import (
	"fmt"
	"math"
)

//var Check = true

//func log(x ...interface{}) {
//	fmt.Fprintln(os.Stderr, x...)
//}

func Frontmost(a, b *HitRecord) HitRecord {
	if b.T < a.T {
		a, b = b, a
	}
	if a.T > 0 {
		return *a
	}
	if b.T > 0 {
		return *b
	}
	return HitRecord{}
}

func FrontT(t1, t2 float64) float64 {
	if t2 < t1 {
		t1, t2 = t2, t1
	}
	if t1 > 0 {
		return t1
	}
	if t2 > 0 {
		return t2
	}
	return 0
}

func CheckHit(o Object, r *Ray, h *HitRecord) {
	if h.T == 0 {
		return
	}
	if math.IsNaN(h.T +
		h.Normal[0] + h.Normal[1] + h.Normal[2] +
		h.Local[0] + h.Local[1] + h.Local[2]) { //|| h.Local == (Vec{}) {
		panic(fmt.Sprintf("Scene: Intersect: bad HitRecord\nObject: %#v\nRay: %v\nHitRecord: %v", o, r, *h))
		// Note: *h, otherwise escape analysis thinks h escpaes, causing an alloc per ray, 3x overall performance penalty.
	}
}

func CheckRay(r *Ray) {
	const tol = 1e-9
	if math.Abs(1-r.Dir.Len2()) > tol {
		panic(fmt.Sprintf("unnormalized Ray dir: %v (len=%v)", r.Dir, r.Dir.Len()))
	}
}
