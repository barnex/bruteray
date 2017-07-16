package main

import (
	"fmt"
	"testing"
)

func TestMatrix(tst *testing.T) {
	//t := Helper(tst)

	A := Matrix{
		{1, 0, 0},
		{0, 0, 1},
		{0, -1, 0},
	}

	b := Vec{2, 0, 0}
	fmt.Println(b.Transf(&A))

	b = Vec{0, 2, 0}
	fmt.Println(b.Transf(&A))

	b = Vec{0, 0, 2}
	fmt.Println(b.Transf(&A))
	fmt.Println()

	B := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(A.Mul(&B))
	fmt.Println()
	fmt.Println(B.Mul(&A))
}

func TestMatrix4(tst *testing.T) {

	//T := Transl4(Vec{1, 2, 3})

	//a := Vec{-2, 3, -1}

	//fmt.Println(T.TransfPoint(a))
	//fmt.Println(T.TransfDir(a))

	//T = RotX4(pi / 2)

	//fmt.Println(T.TransfPoint(a))
	//fmt.Println(T.TransfDir(a))

	//a = Vec{1, 2, 3}
	//tr := Transl4(Vec{10, 10, 10})
	//rot := RotX4(pi / 2)

	//_ = tr
	//_ = rot

	//T = *(tr.Mul(&rot))
	////T = *(rot.Mul(&tr))
	////T = tr
	////T = rot

	//fmt.Println(&T)

	//fmt.Println()
	//fmt.Println(T.TransfPoint(a))
	////fmt.Println(T.TransfDir(a))
}
