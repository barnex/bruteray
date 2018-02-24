// package csg provides constructive solid geometry operations on shapes.
package csg

import (
	"math"
	"sync"

	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/shape"
)

var (
	fbBuf = sync.Pool{
		New: func() interface{} {
			fb := (make([]Fragment, 0, 4))
			return &fb
		},
	}
)

func getFrags() *[]Fragment {
	fb := fbBuf.Get().(*[]Fragment)
	*fb = (*fb)[:0]
	return fb
}

func putFrags(fb *[]Fragment) {
	fbBuf.Put(fb)
}

func init() {
	shape.CsgAnd_ = And // TODO: remove
}

// Intersection (boolean AND) of two objects.
func And(a, b CSGObj) CSGObj {
	return &and{a, b}
}

type and struct {
	a, b CSGObj
}

func (o *and) Hit1(r *Ray, f *[]Fragment) { o.HitAll(r, f) }

func (o *and) HitAll(r *Ray, f *[]Fragment) {

	fa := getFrags()
	defer putFrags(fa)
	o.a.HitAll(r, fa)
	if len(*fa) == 0 {
		return
	}
	for _, s := range *fa {
		if o.b.Inside(r.At(s.T)) {
			*f = append(*f, s)
		}
	}

	fb := getFrags()
	defer putFrags(fb)
	o.b.HitAll(r, fb)
	for _, s := range *fb {
		if o.a.Inside(r.At(s.T)) {
			*f = append(*f, s)
		}
	}
}

func (o *and) Inside(p Vec) bool {
	return o.a.Inside(p) && o.b.Inside(p)
}

// Union (logical OR) of two objects.
// TODO: remove in favor of MultiOr
func Or(a, b CSGObj) CSGObj {
	return &or{a, b}
}

type or struct {
	a, b CSGObj
}

func (o *or) Hit1(r *Ray, f *[]Fragment) { o.HitAll(r, f) }

func (o *or) HitAll(r *Ray, f *[]Fragment) {

	fa := getFrags()
	defer putFrags(fa)

	o.a.HitAll(r, fa)
	if len(*fa) == 0 {
		o.b.HitAll(r, f)
		return
	}

	fb := getFrags()
	defer putFrags(fb)
	o.b.HitAll(r, fb)

	for _, s := range *fa {
		if !o.b.Inside(r.At(s.T)) {
			*f = append(*f, s)
		}
	}
	for _, s := range *fb {
		if !o.a.Inside(r.At(s.T)) {
			*f = append(*f, s)
		}
	}
}

func (o *or) Inside(p Vec) bool {
	return o.a.Inside(p) || o.b.Inside(p)
}

func MultiOr(o ...CSGObj) CSGObj {
	return &multiOr{o}
}

type multiOr struct {
	o []CSGObj
}

func (o *multiOr) Hit1(r *Ray, f *[]Fragment) { o.HitAll(r, f) }

func (o *multiOr) HitAll(r *Ray, f *[]Fragment) {
	fa := getFrags()
	defer putFrags(fa)

	for i, a := range o.o {
		*fa = (*fa)[:0]
		a.HitAll(r, fa)

		for _, s := range *fa {

			pos := r.At(s.T)
			inside := false

			for j, b := range o.o {
				if i == j {
					continue
				}
				if b.Inside(pos) {
					inside = true
					break
				}
			}
			if !inside {
				*f = append(*f, s)
			}

		}
	}
}

func (o *multiOr) Inside(pos Vec) bool {
	for _, o := range o.o {
		if o.Inside(pos) {
			return true
		}
	}
	return false
}

// Union (logical OR) of two objects, without optimizing result.
// Best suited for a small number of simple objects.
//func Or0(a, b CSGObj) CSGObj {
//	return &or0{a, b}
//}
//
//type or0 struct {
//	a, b CSGObj
//}
//
//func (o *or0) Hit1(r *Ray, f *[]Fragment) {}
//
//func (o *or0) Hit(r *Ray, f *[]Fragment) {
//	o.a.Hit(r, f)
//	fa := *f
//
//	fb := (*f)[len(fa):]
//	o.b.Hit(r, &fb)
//
//	*f = append(fa, fb...)
//}
//
//func (o *or0) Inside(p Vec) bool {
//	return o.a.Inside(p) || o.b.Inside(p)
//}

// Subtraction (logical AND NOT) of two objects
func Minus(a, b CSGObj) CSGObj {
	return &minus{a, b}
}

type minus struct {
	a, b CSGObj
}

func (o *minus) Hit1(r *Ray, f *[]Fragment) { o.HitAll(r, f) }

func (o *minus) HitAll(r *Ray, f *[]Fragment) {

	o.a.HitAll(r, f)
	if len(*f) == 0 {
		return
	}

	var fb []Fragment
	o.b.HitAll(r, &fb)

	var f3 []Fragment

	for _, s := range *f {
		if !o.b.Inside(r.At(s.T)) {
			f3 = append(f3, s)
		}
	}
	for _, s := range fb {
		if o.a.Inside(r.At(s.T)) {
			f3 = append(f3, s)
		}
	}
	//Sort(f3)
	*f = f3
}

func (o *minus) Inside(p Vec) bool {
	return o.a.Inside(p) && !o.b.Inside(p)
}

func Cutout(a, b CSGObj) CSGObj {
	return &cutout{a, b}
}

type cutout struct {
	a CSGObj
	b CSGObj // TODO -> Insider
}

func (o *cutout) Hit1(r *Ray, f *[]Fragment) { o.HitAll(r, f) }

func (o *cutout) HitAll(r *Ray, f *[]Fragment) {
	o.a.HitAll(r, f)
	for i, s := range *f {
		if o.b.Inside(r.At(s.T)) {
			(*f)[i].T = math.NaN()
		}
	}
}

func (o *cutout) Inside(pos Vec) bool {
	return o.a.Inside(pos)
}

// Intersection, treating A as a hollow object.
// Equivalent to, but more efficient than And(Hollow(a), b)
func SurfaceAnd(a Obj, b CSGObj) Obj {
	return &hand{a: a, b: b}
}

type hand struct {
	a Obj
	b CSGObj
}

func (o *hand) Hit1(r *Ray, f *[]Fragment) {

	o.a.Hit1(r, f)
	if len(*f) == 0 {
		return
	}

	// TODO: optimize
	f2 := make([]Fragment, 0, len(*f))
	for i, s := range *f {
		if o.b.Inside(r.At(s.T)) {
			f2 = append(f2, (*f)[i])
		}
	}
	*f = f2
}

// Hollow turns a into a hollow surface.
// E.g.: a filled cylinder into a hollow tube.
func Hollow(o CSGObj) CSGObj {
	return hollow{o}
}

type hollow struct {
	CSGObj
}

func (hollow) Inside(Vec) bool {
	return false
}

func Inverse(o CSGObj) CSGObj {
	return inverse{o}
}

type inverse struct {
	CSGObj
}

func (o inverse) Inside(p Vec) bool {
	return !o.CSGObj.Inside(p)
}

//func Sort(f []Fragment) {
//	sort.Sort(byT(f))
//}
//
//type byT []Fragment
//
//func (s byT) Len() int           { return len(s) }
//func (s byT) Less(i, j int) bool { return s[i].T < s[j].T }
//func (s byT) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
