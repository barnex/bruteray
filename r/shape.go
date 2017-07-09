package r

type Shape interface {
	Inters(r *Ray) Interval
}
