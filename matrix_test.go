package bruteray

import (
	"fmt"
)

func ExampleMatrix4_Mul() {

	I := UnitMatrix4()
	I = I.Mul(I)
	fmt.Println(I)

	T := Transl4(Vec{2, 3, 4})
	fmt.Println(I.Mul(T))
	fmt.Println(T.Mul(I))

	// Output:
	//[1 0 0 0]
	//[0 1 0 0]
	//[0 0 1 0]
	//[0 0 0 1]
	//
	//[1 0 0 2]
	//[0 1 0 3]
	//[0 0 1 4]
	//[0 0 0 1]
	//
	//[1 0 0 2]
	//[0 1 0 3]
	//[0 0 1 4]
	//[0 0 0 1]
}

func ExampleMatrix4_Inv() {
	T := Transl4(Vec{2, 3, 4}).Mul(RotX4(Pi / 2).Mul(Transl4(Vec{1, 2, -3})).Mul(RotY4(Pi / 2)))
	//fmt.Println(T)
	fmt.Println(T.Mul(T.Inv()))
	fmt.Println(T.Inv().Mul(T))
	// Output:
	//[1 0 0 0]
	//[0 1 0 0]
	//[0 0 1 0]
	//[0 0 0 1]
	//
	//[1 0 0 0]
	//[0 1 0 0]
	//[0 0 1 0]
	//[0 0 0 1]
}

func ExampleRay_Transf() {
	T := Transl4(Vec{2, 3, -1})
	r := Ray{Vec{4, 3, 2}, Vec{-1, -2, -3}}
	r.Transf(T)
	fmt.Println(r)
	r.Transf(T.Inv())
	fmt.Println(r)

	// Output:
	//{[6 6 1] [-1 -2 -3]}
	//{[4 3 2] [-1 -2 -3]}
}
