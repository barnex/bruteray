package main

type Inter struct {
	Min, Max float64
}

var empty = Inter{-inf, -inf}

func (a Inter) And(b Inter) Inter {
	if a.Max < b.Min || b.Max < a.Min {
		return empty
	}

	return Inter{Max(a.Min, b.Min), Min(a.Max, b.Max)}

	return empty
}

func (a Inter) OK() bool {
	return a.Min < a.Max
}

func (a Inter) Empty() bool {
	return a.Min >= a.Max
}
