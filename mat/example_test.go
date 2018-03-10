package mat_test

import (
	"math"

	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/doc"
	. "github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/shape"
)

func ExampleUVCyl() {
	img := MustLoad("../assets/earth.jpg") // cylindrical projection
	r := 0.5
	globe := shape.NewSphere(2*r, nil)
	globe.Transl(Vec{0, r, 0})
	th := -30 * Deg
	uvmap := &UVCyl{
		P0: Vec{0, 0, 0},
		Pu: Vec{math.Sin(th), 0, -math.Cos(th)},
		Pv: Vec{0, 2 * r, 0},
	}
	globe.Mat = Diffuse(NewImgTex(img, uvmap))
	doc.Show(globe)
	//Output:
	//ExampleUVCyl
}

func ExampleUVAffine() {
	img := MustLoad("../assets/monalisa.jpg")
	cube := shape.NewBox(1/img.Aspect(), 1, 0.2, nil)
	cube.Transl(Vec{0, 0.5, 0})
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

func ExampleReflectFresnel() {
	mat := ReflectFresnel(1.5, BLACK)
	doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
	//Output:
	//ExampleReflectFresnel
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
