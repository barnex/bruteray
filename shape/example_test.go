package shape_test

import (
	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/doc"
	"github.com/barnex/bruteray/mat"
	. "github.com/barnex/bruteray/shape"
)

func ExampleNewSphere() {
	doc.Show(
		NewSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0}),
	)
	//Output:
	//ExampleNewSphere
}

func ExampleNewBox() {
	doc.Show(
		NewBox(0.5, 1, 0.5, mat.Diffuse(RED)).Transl(Vec{-1.5, 0.5, 0}),
		NewBox(1, 0.5, 1, mat.Diffuse(BLUE)).Transl(Vec{1.5, 0.5, 0}),
	)
	//Output:
	//ExampleNewBox
}
