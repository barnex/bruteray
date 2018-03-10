package shape

import (
	. "github.com/barnex/bruteray/br"
)

func NewFunction(min, max Vec, m Material, f func() float64) *function {
	return &function{
		box: &Box{Min: min, Max: max, Mat: m},
	}
}

type function struct {
	box *Box
	f   func() float64
}

func (s *function) Hit1(r *Ray, f *[]Fragment) {
	s.box.HitAll(r, f)
	if len(*f) != 2 {
		panic(len(*f))
	}

	t1, t2 := (*f)[0].T, (*f)[1].T

}
