package shape

import (
	"math"

	. "github.com/barnex/bruteray/br"
)

func NewFunction(min, max float64, m Material, f Func) *function {
	return &function{
		top: NewSheet(Ey, max, nil),
		bot: NewSheet(Ey, min, nil),
		Mat: m,
		f:   f,
	}
}

type function struct {
	top, bot *Sheet
	Mat      Material
	f        Func
}

type Func func(Vec) float64

func (s *function) Hit1(r *Ray, f *[]Fragment) {
	min, max := s.bot.hit(r), s.top.hit(r)
	if min < 0 || max < 0 {
		return
	}

	t, n := s.normal(r, min, max)
	*f = append(*f, Fragment{T: t, Norm: n, Material: s.Mat})
}

func (s *function) Inside(p Vec) bool {
	return s.f(p) < 0
}

func (s *function) bisect(r *Ray, min, max float64) float64 {

	const tol = 1e-9

	in, out := max, min
	if s.f(r.At(in)) < 0 {
		in, out = out, in
	}

	for math.Abs(in-out)/(in+out) > tol {
		mid := (in + out) / 2
		if s.f(r.At(mid)) > 0 {
			in = mid
		} else {
			out = mid
		}
	}

	return out
}

func (s *function) normal(r *Ray, min, max float64) (float64, Vec) {
	t := s.bisect(r, min, max)
	c := r.At(t)

	ra := *r
	ra.SetDir(ra.Dir().Add(Vec{1e-5, 0, 0}))
	ta := s.bisect(&ra, min, max)
	a := ra.At(ta)

	rb := *r
	rb.SetDir(rb.Dir().Add(Vec{0, 1e-5, 0}))
	tb := s.bisect(&rb, min, max)
	b := rb.At(tb)

	a = a.Sub(c)
	b = b.Sub(c)

	n := b.Cross(a).Normalized()
	return t, n
}
