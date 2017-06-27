package main

type Inter struct {
	Min, Max float64
}

var empty = Inter{inf, -inf}

func (a Inter) And(b Inter) Inter {
	if a.Max < b.Min || b.Max < a.Min {
		return empty
	}

	return Inter{Max(a.Min, b.Min), Min(a.Max, b.Max)}

	return empty
}

func (a Inter) Minus(b Inter) Inter {
	if a.Max < b.Min || b.Max < a.Min {
		return a
	}

	if b.Min < a.Min {
		return Inter{b.Max, a.Max}.Normalize()
	}
	return Inter{a.Min, b.Min}.Normalize()
}

func (a Inter) OK() bool {
	return a.Min <= a.Max // && a.Min > 0
}

//func (a Inter) Empty() bool {
//	return a.Min > a.Max
//}

func (a Inter) Normalize() Inter {
	if !a.OK() {
		return empty
	}
	return a
}
