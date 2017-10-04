package bruteray

import (
	"sort"
)

// Intersection (boolean AND) of two objects.
func And(a, b Obj) Obj {
	return &and{a, b}
}

type and struct {
	a, b Obj
}

func (o *and) Hit(r *Ray, f *[]Surf) {

	o.a.Hit(r, f)
	if len(*f) == 0 {
		return
	}

	var fb []Surf
	o.b.Hit(r, &fb)

	var f3 []Surf

	for _, s := range *f {
		if o.b.Inside(r.At(s.T)) {
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

func (o *and) Inside(p Vec) bool {
	return o.a.Inside(p) && o.b.Inside(p)
}

// A or B
func Or(a, b Obj) Obj {
	return &or{a, b}
}

type or struct {
	a, b Obj
}

func (o *or) Hit(r *Ray, f *[]Surf) {

	o.a.Hit(r, f)

	var fb []Surf
	o.b.Hit(r, &fb)

	var f3 []Surf

	for _, s := range *f {
		if !o.b.Inside(r.At(s.T)) {
			f3 = append(f3, s)
		}
	}
	for _, s := range fb {
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

// A minus B
func Minus(a, b Obj) Obj {
	return &minus{a, b}
}

type minus struct {
	a, b Obj
}

func (o *minus) Hit(r *Ray, f *[]Surf) {

	o.a.Hit(r, f)
	if len(*f) == 0 {
		return
	}

	var fb []Surf
	o.b.Hit(r, &fb)

	var f3 []Surf

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

// "Hollow" AND: remove all parts of A that are not in B,
// treating a as a hollow surface (not volume).
// TODO: rename SurfaceAnd?
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

func Sort(f []Surf) {
	sort.Sort(byT(f))
}

type byT []Surf

func (s byT) Len() int           { return len(s) }
func (s byT) Less(i, j int) bool { return s[i].T < s[j].T }
func (s byT) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
