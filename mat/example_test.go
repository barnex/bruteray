package mat_test

import (
	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/doc"
	. "github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/shape"
)

func ExampleUVAffine() {
	img := MustLoad("../assets/monalisa.jpg")
	cube := shape.NewBox(1, img.Aspect(), 0.2, nil)
	cube.Transl(Vec{0, img.Aspect() / 2, 0})
	uvmap := &UVAffine{
		P0: cube.Corner(-1, -1, 1),
		Pu: cube.Corner(1, -1, 1),
		Pv: cube.Corner(-1, 1, 1)}
	cube.Mat = Diffuse(NewImgTex(img, uvmap))
	doc.Show(cube)
	//Output:
	//ExampleUVAffine
}

func ExampleDiffuse() {
	mat := Diffuse(WHITE)
	doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
	//Output:
	//ExampleDiffuse
}

func ExampleBlend() {
	white := Diffuse(WHITE)
	refl := Reflective(WHITE)
	mat := Blend(0.95, white, 0.05, refl)
	doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
	//Output:
	//ExampleBlend
}

func ExampleReflective() {
	mat := Reflective(WHITE.EV(-1))
	doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
	//Output:
	//ExampleReflective
}

func ExampleRefractive() {
	mat := Refractive(1, 1.5)
	doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
	//Output:
	//ExampleRefractive
}

func ExampleCheckboard() {
	m1 := Diffuse(WHITE)
	m2 := Reflective(WHITE.EV(-3))
	mat := Checkboard(0.1, m1, m2)
	doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
	//Output:
	//ExampleCheckboard
}

func ExampleFlat() {
	mat := Flat(WHITE)
	doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
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
