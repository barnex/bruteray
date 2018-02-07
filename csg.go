package bruteray

import (
	"sort"
	"sync"
)

var (
	fbBuf = sync.Pool{
		New: func() interface{} {
			fb := (make([]Fragment, 0, 4))
			return &fb
		},
	}
)

func fb() *[]Fragment {
	fb := fbBuf.Get().(*[]Fragment)
	*fb = (*fb)[:0]
	return fb
}

func rfb(fb *[]Fragment) {
	fbBuf.Put(fb)
}

// Intersection (boolean AND) of two objects.
func And(a, b Obj) Obj {
	return &and{a, b}
}

type and struct {
	a, b Obj
}

func (o *and) Hit(r *Ray, f *[]Fragment) {

	o.a.Hit(r, f)
	if len(*f) == 0 {
		return
	}

	fb := fb()
	defer rfb(fb)
	o.b.Hit(r, fb)

	var f3 []Fragment

	for _, s := range *f {
		if o.b.Inside(r.At(s.T)) {
			f3 = append(f3, s)
		}
	}
	for _, s := range *fb {
		if o.a.Inside(r.At(s.T)) {
			f3 = append(f3, s)
		}
	}
	Sort(f3)
	*f = f3
}

func (o *and) Inside(p Vec) bool {
	return o.a.Inside(p) && o.b.Inside(p)
}

// Union (logical OR) of two objects.
func Or(a, b Obj) Obj {
	return &or{a, b}
}

type or struct {
	a, b Obj
}

func (o *or) Hit(r *Ray, f *[]Fragment) {

	o.a.Hit(r, f)

	fb := fb()
	defer rfb(fb)
	o.b.Hit(r, fb)

	var f3 []Fragment

	for _, s := range *f {
		if !o.b.Inside(r.At(s.T)) {
			f3 = append(f3, s)
		}
	}
	for _, s := range *fb {
		if !o.a.Inside(r.At(s.T)) {
			f3 = append(f3, s)
		}
	}
	Sort(f3)
	*f = f3
}

func (o *or) Inside(p Vec) bool {
	return o.a.Inside(p) || o.b.Inside(p)
}

// Subtraction (logical AND NOT) of two objects
func Minus(a, b Obj) Obj {
	return &minus{a, b}
}

type minus struct {
	a, b Obj
}

func (o *minus) Hit(r *Ray, f *[]Fragment) {

	o.a.Hit(r, f)
	if len(*f) == 0 {
		return
	}

	var fb []Fragment
	o.b.Hit(r, &fb)

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
	Sort(f3)
	*f = f3
}

func (o *minus) Inside(p Vec) bool {
	return o.a.Inside(p) && !o.b.Inside(p)
}

// Intersection, treating A as a hollow object.
// Equivalent to, but more efficient than And(Hollow(a), b)
func SurfaceAnd(a, b Obj) Obj {
	return &hand{a: a, b: b}
}

type hand struct {
	a, b Obj
	noInside
}

func (o *hand) Hit(r *Ray, f *[]Fragment) {

	o.a.Hit(r, f)
	if len(*f) == 0 {
		return
	}

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
func Hollow(o Obj) Obj {
	return hollow{o}
}

type hollow struct {
	Obj
}

func (hollow) Inside(Vec) bool {
	return false
}

func Inverse(o Obj) Obj {
	return inverse{o}
}

type inverse struct {
	Obj
}

func (o inverse) Inside(p Vec) bool {
	return !o.Obj.Inside(p)
}

func Sort(f []Fragment) {
	sort.Sort(byT(f))
}

type byT []Fragment

func (s byT) Len() int           { return len(s) }
func (s byT) Less(i, j int) bool { return s[i].T < s[j].T }
func (s byT) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
