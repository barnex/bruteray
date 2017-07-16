package main

import "testing"

func TestVec(t *testing.T) {
	testVec(t, Vec{1, 2, 3}.Add(Vec{4, 5, 6}), Vec{5, 7, 9}, 0)
	testVec(t, Vec{1, 2, 3}.MAdd(2, Vec{4, 5, 6}), Vec{9, 12, 15}, 0)
	testVec(t, Vec{1, 2, 3}.Sub(Vec{1, 3, 2}), Vec{0, -1, 1}, 0)
	testVec(t, Vec{1, 2, 3}.Mul(2), Vec{2, 4, 6}, 0)
	testVec(t, Vec{2, 4, 6}.Div(2), Vec{1, 2, 3}, 0)
	testVec(t, Vec{2, 6, 12}.Div3(Vec{1, 2, 3}), Vec{2, 3, 4}, 0)
	testVec(t, Vec{0, 3, 4}.Normalized(), Vec{0, 3. / 5., 4. / 5.}, 1e-15)
	testVec(t, Vec{0, 0, 0}.Normalized(), Vec{0, 0, 0}, 0)
	testVec(t, Vec{1, 0, 0}.Cross(Vec{0, 1, 0}), Vec{0, 0, 1}, 0)
	testVec(t, Vec{0, 1, 0}.Cross(Vec{1, 0, 0}), Vec{0, 0, -1}, 0)
	testVec(t, Vec{-1, 7, 4}.Cross(Vec{-5, 8, 4}), Vec{-4, -16, 27}, 0)
	testFloat(t, Vec{0, 3, 4}.Len(), 5, 1e-15)
	testFloat(t, Vec{0, 3, 4}.Len2(), 25, 1e-15)
}
