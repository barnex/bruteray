package r

import (
	"fmt"
)

func ExampleMatrix4() {

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

func ExampleRay_Transf() {
	T := Transl4(Vec{2, 3, -1})
	r := Ray{Vec{4, 3, 2}, Vec{-1, -2, -3}}
	r.Transf(T)
	fmt.Println(r)

	// Output:
	//{[6 6 1] [-1 -2 -3]}
}
