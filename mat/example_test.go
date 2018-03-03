package mat_test

import (
	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/doc"
	. "github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/shape"
)

func ExampleDiffuse() {
	doc.Show(
		shape.NewSphere(1, Diffuse(WHITE)).Transl(Vec{0, 0.5, 0}),
	)
	//Output:
	//ExampleDiffuse
}

func ExampleBlend() {
	white := Diffuse(WHITE)
	refl := Reflective(WHITE)
	doc.Show(
		shape.NewSphere(1, Blend(0.95, white, 0.05, refl)).Transl(Vec{0, 0.5, 0}),
	)
	//Output:
	//ExampleBlend
}

func ExampleReflective() {
	doc.Show(
		shape.NewSphere(1, Reflective(WHITE.EV(-1))).Transl(Vec{0, 0.5, 0}),
	)
	//Output:
	//ExampleReflective
}

func ExampleRefractive() {
	doc.Show(
		shape.NewSphere(1, Refractive(1, 1.5)).Transl(Vec{0, 0.5, 0}),
	)
	//Output:
	//ExampleRefractive
}

func ExampleFlat() {
	doc.Show(
		shape.NewSphere(1, Flat(WHITE)).Transl(Vec{0, 0.5, 0}),
	)
	//Output:
	//ExampleFlat
}

func ExampleDebugShape() {
	e := NewEnv()
	e.Add(shape.NewSheet(Ey, 0, DebugShape(WHITE)))
	e.Add(shape.NewSphere(1, DebugShape(WHITE)).Transl(Vec{0, 0.5, 0}))
	// Note: no light source added
	doc.Example(e)
	//Output:
	//ExampleDebugShape
}
