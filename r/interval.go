package r

import "fmt"

type Interval struct {
	Min, Max float64
}

func Interv(min, max float64) Interval {
	i := Interval{min, max}
	i.check()
	return i
}

func (i Interval) OK() bool {
	i.check()
	return (i != Interval{})
}

func (i Interval) check() {
	if i.Min > i.Max {
		panic(fmt.Sprintf("bad interval: %v", i))
	}
}

//const inf = 1e99
//
//var empty = Interval{inf, -inf}
//
//func (a Interval) And(b Interval) Interval {
//	check(a)
//	check(b)
//
//	if a.Max < b.Min || b.Max < a.Min {
//		return empty
//	}
//
//	return Interval{Max(a.Min, b.Min), Min(a.Max, b.Max)}
//
//	return empty
//}
//
//func (a Interval) Minus(b Interval) Interval {
//	check(a)
//	check(b)
//
//	if a.Max < b.Min || b.Max < a.Min {
//		return a
//	}
//
//	if b.Min < a.Min {
//		return Interval{b.Max, a.Max}.normalize()
//	}
//	return Interval{a.Min, b.Min}.normalize()
//}
//
//func (i Interval) normalize() Interval {
//	if !i.OK() {
//		return empty
//	}
//	return i
//}
//
//func (a Interval) OK() bool {
//	return a.Min <= a.Max
//}
//
//func check(i Interval) {
//	if i != empty && i.Min > i.Max {
//		panic(fmt.Sprintf("bad interval:", i))
//	}
//}
