package util

import (
	"fmt"
)

func ExampleFrac() {
	for _, x := range []float64{-2.1, -2, -1.9, -1, -0.1, 0, 0.1, 0.9, 1, 1.1, 2, 2.1} {
		fmt.Printf("Frac(% .1f) = % .1f\n", x, Frac(x))
	}
	//Output:
	//Frac(-2.1) =  0.9
	//Frac(-2.0) =  0.0
	//Frac(-1.9) =  0.1
	//Frac(-1.0) =  0.0
	//Frac(-0.1) =  0.9
	//Frac( 0.0) =  0.0
	//Frac( 0.1) =  0.1
	//Frac( 0.9) =  0.9
	//Frac( 1.0) =  0.0
	//Frac( 1.1) =  0.1
	//Frac( 2.0) =  0.0
	//Frac( 2.1) =  0.1
}
