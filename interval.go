package main

type Interval struct {
	Min, Max float64
}

var empty = Interval{inf, -inf}

func (a Interval) And(b Interval) Interval {
	if a.Max < b.Min || b.Max < a.Min {
		return empty
	}

	return Interval{Max(a.Min, b.Min), Min(a.Max, b.Max)}

	return empty
}

func (a Interval) Minus(b Interval) Interval {
	if a.Max < b.Min || b.Max < a.Min {
		return a
	}

	if b.Min < a.Min {
		return Interval{b.Max, a.Max}.Normalize()
	}
	return Interval{a.Min, b.Min}.Normalize()
}

func (a Interval) OK() bool {
	return a.Min <= a.Max
}

func (a Interval) Normalize() Interval {
	if !a.OK() {
		return empty
	}
	return a
}
