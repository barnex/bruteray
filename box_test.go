package bruteray

import "testing"

func TestBoxIntersect(tst *testing.T) {
	t := Helper(tst)

	b := Box(Vec{}, 1, 1, 1, Flat(BLACK)).(*prim).s

	t.Eq(b.Inters(&Ray{Vec{2, 0, 0}, Vec{1, 0, 0}}), Interval{})
	t.Eq(b.Inters(&Ray{Vec{2, 0, 0}, Vec{-1, 0, 0}}), Interval{1, 3})
	t.Eq(b.Inters(&Ray{Vec{0, 0, 0}, Vec{-1, 0, 0}}), Interval{-1, 1})

}
