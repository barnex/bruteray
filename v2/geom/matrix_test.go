package geom

import (
	"fmt"
)

func ExampleMatrix_String() {
	fmt.Println(Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	//Output:
	//[[1 2 3] [4 5 6] [7 8 9]]
}

//func ExampleMatrix_Inv(){
//	m := Matrix{{1,2,3},{3,-1,2},{2,3,-1}}
//	inv := m.Inv()
//	check := inv.Mul(&m)
//	fmt.Println(check.Sprintf("%7f"))
//
//	//Output:
//	//[[1 0 0] [0 1 0] [0 0 1]]
//}
