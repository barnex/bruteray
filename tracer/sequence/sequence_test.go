package sequence

import (
	"fmt"
)

func ExampleHalton_2() {
	for i := 0; i < 10; i++ {
		fmt.Println(Halton(2, i))
	}

	//Output:
	// 0.5
	// 0.25
	// 0.75
	// 0.125
	// 0.625
	// 0.375
	// 0.875
	// 0.0625
	// 0.5625
	// 0.3125
}

func ExampleHalton_3() {
	for i := 0; i < 10; i++ {
		fmt.Println(Halton(3, i))
	}

	//Output:
	// 0.3333333333333333
	// 0.6666666666666666
	// 0.1111111111111111
	// 0.4444444444444444
	// 0.7777777777777777
	// 0.2222222222222222
	// 0.5555555555555556
	// 0.8888888888888888
	// 0.037037037037037035
	// 0.37037037037037035
}
