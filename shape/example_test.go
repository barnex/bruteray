package shape_test

import (
	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/doc"
	"github.com/barnex/bruteray/mat"
	. "github.com/barnex/bruteray/shape"
)

func ExampleNewBox() {
	doc.Show(
		NewBox(1, 1, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0}),
	)
	//Output:
	//ExampleNewBox
}

//func ExampleNewCylinder() {
//	doc.Show(
//		NewCylinder(Y, Vec{0, 0.5, 0}, 1, 0.5, mat.Diffuse(RED)),
//	)
//	//Output:
//	//ExampleNewCylinder
//}

func ExampleNewSheet() {
	doc.Show(
		NewSheet(Ey, 0.1, mat.Diffuse(RED)),
	)
	//Output:
	//ExampleNewSheet
}

func ExampleNewSphere() {
	doc.Show(
		NewSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0}),
	)
	//Output:
	//ExampleNewSphere
}

func ExampleAnd() {
	cube := NewBox(1, 1, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
	sphere := NewSphere(1.5, mat.Diffuse(BLUE)).Transl(cube.Corner(1, 1, -1))
	doc.Show(And(cube, sphere))
	//Output:
	//ExampleAnd
}

func ExampleOr() {
	cube := NewBox(1, 1, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
	sphere := NewSphere(1.0, mat.Diffuse(BLUE)).Transl(cube.Corner(1, 1, -1))
	doc.Show(Or(cube, sphere))
	//Output:
	//ExampleOr
}

func ExampleMinus() {
	cube := NewBox(1, 1, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
	sphere := NewSphere(1.5, mat.Diffuse(BLUE)).Transl(cube.Corner(1, 1, -1))
	doc.Show(Minus(cube, sphere))
	//Output:
	//ExampleMinus
}

func ExampleCutout() {
	cube := NewBox(1, 1, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
	sphere := NewSphere(1.5, mat.Diffuse(BLUE)).Transl(cube.Corner(1, 1, -1))
	doc.Show(Cutout(cube, sphere))
	//Output:
	//ExampleCutout
}
