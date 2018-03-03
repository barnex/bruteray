package shape_test

import (
	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/doc"
	"github.com/barnex/bruteray/mat"
	. "github.com/barnex/bruteray/shape"
)

func ExampleSphere() {
	e := NewEnv()
	sphere := NSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
	e.Add(sphere)
	doc.Example(e)
	//Output:
	//ExampleSphere
}
