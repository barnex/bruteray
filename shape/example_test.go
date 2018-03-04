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

func ExampleNewCylinder() {
	cyl := NewCylinder(Y, Vec{0, 0.5, 0}, 1, 1, mat.Diffuse(RED))
	doc.Show(cyl)
	//Output:
	//ExampleNewCylinder
}

func ExampleNewSheet() {
	sheet := NewSheet(Ey, 0.1, mat.Diffuse(RED))
	doc.Show(sheet)
	//Output:
	//ExampleNewSheet
}

func ExampleNewSphere() {
	sphere := NewSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
	doc.Show(sphere)
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
