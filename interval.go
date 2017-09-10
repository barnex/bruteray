package bruteray

import (
	"fmt"
	"math"
)

// An Interval along a ray.
// 	Max >= Min
// 	Max >= 0
// The empty interval is {0, 0}.
type Interval struct {
	Min, Max float64
}

// Returns the smallest, positive value of (min, max).
// Used to implement Hit in terms of Intersect.
func (i Interval) Front() float64 {
	if i.Min > 0 {
		return i.Min
	}
	return i.Max
}

// For backwards compatibility.
// TODO: rm
func (i Interval) Slice() []Interval {
	if !i.OK() {
		return nil
	}
	return []Interval{i}
}

//
func (i Interval) Fix() Interval {
	if i.Max <= 0 {
		return Interval{}
	}
	return i.check()
}

func (i Interval) OK() bool {
	i.check()
	return (i != Interval{})
}

func (i Interval) check() Interval {
	if math.IsNaN(i.Min) || math.IsNaN(i.Max) ||
		i.Min > i.Max || i.Max < 0 {
		panic(fmt.Sprintf("bad interval: %v", i))
	}
	return i
}
