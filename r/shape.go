package r

import "math"

type Shape interface {
	Inters(r *Ray) Interval
}

type sphere struct {
	c  Vec
	r2 float64
}

func Sphere(center Vec, radius float64) *sphere {
	return &sphere{center, Sqr(radius)}
}

//func (s *sphere) Normal(r *Ray, t float64) Vec {
//        return r.At(t).Sub(s.c).Normalized()
//}

func (s *sphere) Inters(r *Ray) Interval {
	v := r.Start.Sub(s.c)
	d := r.Dir
	vd := v.Dot(d)
	D := Sqr(vd) - (v.Len2() - s.r2)
	if D < 0 {
		return Interval{}
	}
	t1 := (-vd - math.Sqrt(D))
	t2 := (-vd + math.Sqrt(D))
	//t1, t2 = Sort(t1, t2)
	return Interv(t1, t2)
}
