package main

import "testing"

func TestIntervalAnd(t *testing.T) {
	cases := []struct {
		a, b, want Interval
	}{
		{empty, empty, empty},
		{Interval{}, Interval{}, Interval{}},
		{Interval{1, 2}, Interval{1, 2}, Interval{1, 2}},
		{Interval{1, 2}, Interval{3, 4}, empty},
		{Interval{1, 3}, Interval{2, 4}, Interval{2, 3}},
		{Interval{1, 4}, Interval{2, 3}, Interval{2, 3}},
	}

	for _, c := range cases {
		if have := c.a.And(c.b); have != c.want {
			t.Errorf("and %v, %v: have %v, want %v", c.a, c.b, have, c.want)
		}
		if have := c.b.And(c.a); have != c.want {
			t.Errorf("and %v, %v: have %v, want %v", c.b, c.a, have, c.want)
		}
	}
}

func TestIntervalMinus(t *testing.T) {
	cases := []struct {
		a, b, want Interval
	}{
		{empty, empty, empty},
		{Interval{0, 0}, Interval{0, 0}, Interval{0, 0}},
		{Interval{1, 2}, Interval{3, 4}, Interval{1, 2}},
		{Interval{1, 3}, Interval{2, 4}, Interval{1, 2}},
		{Interval{2, 4}, Interval{1, 3}, Interval{3, 4}},
		{Interval{1, 4}, Interval{2, 3}, Interval{1, 2}},
		{Interval{2, 3}, Interval{1, 4}, empty},
		{Interval{2, 3}, empty, Interval{2, 3}},
	}

	for _, c := range cases {
		if have := c.a.Minus(c.b); have != c.want {
			t.Errorf("minus %v, %v: have %v, want %v", c.a, c.b, have, c.want)
		}
	}
}
