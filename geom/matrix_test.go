package geom

import (
	"fmt"
	"math"
)

func ExampleMatrix_Mul() {
	theta := 45 * math.Pi / 180
	c := math.Cos(theta)
	s := math.Sin(theta)
	a := Matrix{{c, s, 0}, {-s, c, 0}, {0, 0, 1}}
	fmt.Printf("% 4.1f", a.Mul(&a))

	//Output:
	// [[ 0.0  1.0  0.0] [-1.0  0.0  0.0] [ 0.0  0.0  1.0]]
}

func ExampleMatrix_MulVec() {
	theta := 30 * math.Pi / 180
	c := math.Cos(theta)
	s := math.Sin(theta)

	m := Matrix{{c, s, 0}, {-s, c, 0}, {0, 0, 1}}
	fmt.Printf("% 3f\n", m.MulVec(Vec{1, 0, 0}))
	fmt.Printf("% 3f\n", m.MulVec(Vec{0, 1, 0}))
	fmt.Printf("% 3f\n", m.MulVec(Vec{0, 0, 1}))

	//Output:
	// [ 0.866025  0.500000  0.000000]
	// [-0.500000  0.866025  0.000000]
	// [ 0.000000  0.000000  1.000000]
}

func ExampleMatrix_Inverse() {
	m := Matrix{{1, 2, 3}, {3, -1, 2}, {2, 3, -1}}
	inv := m.Inverse()
	check := inv.Mul(&m)
	fmt.Printf("% 4.3f", check)

	//Output:
	// [[ 1.000  0.000 -0.000] [-0.000  1.000  0.000] [-0.000 -0.000  1.000]]
}
