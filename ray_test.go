package main

import "testing"

func TestRay(t *testing.T) {
	r := Ray{Vec{1, 2, 3}, Vec{1, -1, 2}}
	testVec(t, r.At(2), Vec{3, 0, 7}, 0)
}
