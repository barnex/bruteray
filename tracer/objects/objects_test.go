package objects

import (
	"math"
	"math/rand"
	"testing"

	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/imagef/colorf"
	"github.com/barnex/bruteray/tracer/cameras"
	"github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/types"
	"github.com/barnex/bruteray/util"
)

var fov = 90 * Deg // legacy field-of-view

func TestAnd(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{1, 1.0, 1}),
			},
			And(
				Sphere(test.White, 1, Vec{-.1, 0, 0}),
				Sphere(test.Cyan, 1, Vec{+.1, 0, 0}),
			),
			test.Sheet(test.Checkers1, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 1.5}),
		4,
		test.DefaultTolerance,
	)
}

func TestAnd_Bounds(t *testing.T) {
	test.QuadView(t,
		NewScene(
			3,
			[]Light{},
			WithBounds(
				And(
					Sphere(test.Normal, 1, Vec{-.1, 0, 0}),
					Sphere(test.Normal, 1, Vec{+.1, 0, 0}),
				),
			),
			test.Sheet(test.Normal, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 1.5}),
		4,
		test.DefaultTolerance,
	)
}

func TestAndNot(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{1.0, 1.2, 1.0}),
			},
			And(
				Sphere(test.White, 1, Vec{0, 0, 0}),
				Not(Sphere(test.Cyan, 0.6, Vec{0, 0, 0.4})),
			),
			test.Sheet(test.Checkers1, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 1.5}),
		4,
		test.DefaultTolerance,
	)
}

func TestAndNot_Box(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{0.0, 1.2, 1}),
			},
			And(
				Box(test.Cyan, 1, 1, 0.3, Vec{0, 0, 0}),
				Not(
					Transformed(
						Cylinder(test.White, 0.5, 2, Vec{0, 0, 0}),
						geom.Rotate(O, Ex, fov),
					),
				),
			),
			test.Sheet(test.Checkers1, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 1.5}),
		4,
		test.DefaultTolerance,
	)
}

func BenchmarkAndNot(b *testing.B) {
	test.Benchmark(b,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{1.0, 1.2, 1.0}),
			},
			Bounded(
				And(
					Sphere(test.White, 1, Vec{0, 0, 0}),
					Not(Sphere(test.Cyan, 0.6, Vec{0, 0, 0.4})),
				),
			),
			test.Sheet(test.Checkers1, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 1.5}),
		test.DefaultTolerance,
	)
}

func TestAndNot_Bounds(t *testing.T) {
	test.QuadView(t,
		NewScene(
			3,
			[]Light{},
			WithBounds(
				And(
					Sphere(test.Normal, 1, Vec{0, 0, 0}),
					Not(Sphere(test.Normal, 0.6, Vec{0, 0, 0.4})),
				),
			),
			test.Sheet(test.Normal, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 1.5}),
		4,
		test.DefaultTolerance,
	)
}

func TestAndNot_Hollow(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{1.0, 1.5, 3.0}),
			},
			And(
				Hollow(Sphere(test.White, 1, Vec{0, 0, 0})),
				Not(Sphere(test.Cyan, 0.6, Vec{0, 0, 0.4})),
			),
			test.Sheet(test.Checkers1, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 1.5}),
		4,
		test.DefaultTolerance,
	)
}

func TestBackdrop(t *testing.T) {
	test.OnePass(t,
		NewScene(
			1,
			[]Light{},
			Backdrop(test.Checkers(test.Flat(colorf.White), test.Flat(colorf.Blue))),
		),
		cameras.Projective(fov).Translate(Vec{0, 1, 0}),
		test.DefaultTolerance,
	)
}

// TODO: normal looks flaky.
func TestBox(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{2, 2, 2}),
			},
			Box(test.Checkers1, 1, 2, 1, Vec{-2, 0.5, 0}),
			Box(test.Checkers1, 1, 1, 3, Vec{1.5, 0.5, 0}),
			test.Sheet(test.Checkers2, 0),
		),
		cameras.Projective(fov).Translate(Vec{0, 2, 5}),
		8,
		test.DefaultTolerance,
	)
}

func TestCylinder(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{0, 1.5, 1}),
			},
			Cylinder(test.Checkers1, 1, 0.5, Vec{0, 0, 0}),
			Cylinder(test.Checkers3, 1.5, 2, Vec{1, 0.25, -1}),

			test.Sheet(test.Checkers2, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0.7, 2.5}),
		8,
		test.DefaultTolerance,
	)
}

func TestCylinder_Bounds(t *testing.T) {
	test.QuadView(t,
		NewScene(
			5,
			[]Light{},
			WithBounds(Cylinder(test.Normal, 1, 0.5, Vec{0, 0, 0})),
			WithBounds(Cylinder(test.Normal, 1.5, 2, Vec{1, 0.25, -1})),
			test.Sheet(test.Normal, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0.7, 2.5}),
		8,
		test.DefaultTolerance,
	)
}

func TestDisk_Bounds(t *testing.T) {
	test.QuadView(t,
		NewScene(
			3,
			[]Light{},
			WithBounds(Disk(test.Normal, 1, Vec{-1, 0, 0})),
			WithBounds(Disk(test.Normal, 2, Vec{2, 2, -1})),
			test.Sheet(test.Normal, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0.7, 2.5}),
		8,
		test.DefaultTolerance,
	)
}

func TestDisk_Shadow(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{1, 2, -1}),
			},
			Disk(test.Checkers1, 1, Vec{-1, 0.5, 0}),
			Disk(test.Checkers1, 2, Vec{2, 3, -1}),
			test.Sheet(test.Checkers4, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0.7, 2.5}),
		8,
		test.DefaultTolerance,
	)
}

//func TestIsoSurface_Discontinuous(t *testing.T) {
//	cam := cameras.NewProjective(Vec{1, 1.2, 3.1})
//	cam.Pitch(-20 * Deg)
//	test.QuadView(t,
//		NewScene(
//			IsoSurface(test.Normal, 2, 0.4, 2, func(u, v float64) float64 {
//	if (int(u*2+10000)+int(v*2+10000))%2 == 0 {
//		return m.a.Eval(ctx, s, r, recDepth, h)
//	} else {
//		return m.b.Eval(ctx, s, r, recDepth, h)
//	}
//			}),
//			test.PointLight(Vec{3, 2, 2.}),
//			test.Sheet(test.Checkers2, 0.001), // offset so shadow rays start in the box
//		),
//		cam,
//		5,
//		test.DefaultTolerance,
//	)
//}

// Test IsoSurface for the constant functions y=0 and y=1,
// The mininum and maximum allowed values.
// These are exactly at the bottom and top of the bounding box, respectively.
// At b8c4432a320b4b1d179ca11 this was bleeding.
func TestIsoSurface_Const(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{3, 2, 2.}),
			},
			Transformed(
				IsoSurface(test.Normal, 1, 0.4, 1, func(u, v float64) float64 {
					return 0
				}),
				geom.Translate(Vec{1, 0, 0}),
			),
			Transformed(
				IsoSurface(test.Normal, 1, 0.4, 1, func(u, v float64) float64 {
					return 1
				}),
				geom.Translate(Vec{0, 0, 0}),
			),
			test.Sheet(test.Checkers2, -0.001), // offset so it does not bleed with surface at y=0
		),
		cameras.Projective(fov).Translate(Vec{1, 1.2, 2.5}).YawPitchRoll(0, -20*Deg, 0),
		5,
		test.DefaultTolerance,
	)
}

func TestIsoSurface_Discontinuous(t *testing.T) {
	t.Skip("TODO")
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{3, 2, 2.}),
			},
			IsoSurface(test.Normal, 2, 0.2, 2, func(u, v float64) float64 {
				if (int(u*2+10000)+int(v*2+10000))%2 == 0 {
					return 0.001
				} else {
					return 0.999
				}
			}),
			test.Sheet(test.Blue, -0.001),
		),
		cameras.Projective(fov).Translate(Vec{1.4, 1.2, 3.1}).YawPitchRoll(0, -20*Deg, 0),
		7,
		test.DefaultTolerance,
	)
}

func TestIsoSurface_Sinc(t *testing.T) {
	t.Skip("TODO: incorrect when ray starts in box. E.g. shadow ray")
	//defer disableChecks()()
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{3, 2, 2.}),
			},
			IsoSurface(test.Normal, 2, 0.4, 2, func(u, v float64) float64 {
				u -= 0.5 // shift to center
				v -= 0.5
				r := math.Sqrt(u*u+v*v) * 20
				return 0.5*util.Sinc(r) + 0.5 // shift to [0..1]
			}),
			test.Sheet(test.Checkers2, 0.001), // offset so shadow rays start in the box
		),
		cameras.Projective(fov).Translate(Vec{1, 1.2, 3.1}).YawPitchRoll(0, -20*Deg, 0),
		5,
		test.DefaultTolerance,
	)
}

func TestMesh_Bunny(t *testing.T) {
	t.Skip("TODO: mesh has holes, NaNs")
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{2, 3, 4}),
			},
			PlyFile(test.Checkers1, "../../assets/bunny_res4.ply",
				geom.Scale(O, 20), geom.Translate(Vec{0, -0.7, 0})),
			test.Sheet(test.Checkers2, 0),
		),
		cameras.Projective(fov).Translate(Vec{0, 1.5, 3.5}),
		8,
		test.DefaultTolerance,
	)
}

func TestMesh_Teapot(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{2, 3, 4}),
			},
			PlyFile(test.Checkers1, "../../assets/teapot.ply",
				geom.Scale(O, 0.25), geom.Rotate(O, Ex, -90*Deg)),
			test.Sheet(test.Checkers2, 0),
		),
		cameras.Projective(fov).Translate(Vec{0, 1.5, 3.5}),
		8,
		test.DefaultTolerance,
	)
}

func BenchmarkMesh_Teapot(b *testing.B) {
	test.Benchmark(b,
		NewScene(
			1,
			[]Light{},
			PlyFile(test.Normal, "../../assets/teapot.ply",
				geom.Scale(O, 0.25), geom.Rotate(O, Ex, -90*Deg), geom.Rotate(O, Ey, -20*Deg)),
		),
		cameras.Projective(fov).Translate(Vec{0, 1.2, 3.5}),
		test.DefaultTolerance,
	)
}

func TestSphere(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{1, 2, 0}),
			},
			Sphere(test.Checkers1, 1, Vec{0, 0, 0}),
			Sphere(test.Checkers3, 1.5, Vec{1, 0.25, -1}),

			test.Sheet(test.Checkers2, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0.7, 2.5}),
		8,
		test.DefaultTolerance,
	)
}

func TestSphere_Bounds(t *testing.T) {
	test.QuadView(t,
		NewScene(
			5,
			[]Light{},
			WithBounds(Sphere(test.Normal, 1, Vec{0, 0, 0})),
			WithBounds(Sphere(test.Normal, 1.5, Vec{1, 0.25, -1})),
			test.Sheet(test.Normal, -0.50001),
		),
		cameras.Projective(fov).Translate(Vec{0, 0.7, 2.5}),
		8,
		test.DefaultTolerance,
	)
}

func TestSphere_Inside(t *testing.T) {
	s := Sphere(nil, 1, Vec{0, 0, 1})
	cases := map[Vec]bool{
		{0, 0, 1}:    true,
		{0, 0, 1.49}: true,
		{0, 0, 1.51}: false,
		{0.49, 0, 1}: true,
		{0.51, 0, 1}: false,
		{0, 0.49, 1}: true,
		{0, 0.51, 1}: false,
	}
	for point, want := range cases {
		got := s.Inside(point)
		if got != want {
			t.Errorf("Inside(%v): got: %v, want: %v", point, got, want)
		}
	}
}

func TestTransformed_Cylinder(t *testing.T) {
	cyl := Cylinder(test.Checkers1, 0.5, 1, Vec{0, 0, 0})
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{1, 2, 0}),
			},
			Transformed(cyl, geom.Scale(O, 0.5)),
			Transformed(cyl, geom.Translate(Vec{2, 0, -2})),
			Transformed(cyl, geom.ComposeLR(
				geom.Rotate(O, Ex, 90*Deg),
				geom.Translate(Vec{2, 0, 0}),
			)),
			Transformed(cyl, geom.ComposeLR(
				geom.Rotate(O, Ez, 45*Deg),
				geom.Rotate(O, Ex, 90*Deg),
				geom.Translate(Vec{0, 0, -2}),
			)),
			test.Sheet(test.Checkers2, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0.5, 0.7, 3}),
		8,
		test.DefaultTolerance,
	)
}

func TestTransformed_Rotate(t *testing.T) {
	cyl := Cylinder(test.Checkers1, 0.5, 0.5, Vec{0, 0, 0})
	tcyl := Transformed(cyl, geom.Translate(Vec{-1, 0, 0}))
	test.QuadView(t,
		NewScene(
			1,
			[]Light{},
			test.Sphere(test.Yellow, 0.2, Vec{1, 0, 0}), // center of rotation
			cyl,
			tcyl,
			Transformed(cyl, geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg)),
			Transformed(tcyl, geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg)),
			Transformed(cyl, geom.ComposeLR(
				geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg),
				geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg),
			)),
			Transformed(cyl, geom.ComposeLR(
				geom.Translate(Vec{-1, 0, 0}),
				geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg),
				geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg),
			)),
			Transformed(
				Transformed(
					Transformed(cyl,
						geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg),
					),
					geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg),
				),
				geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg),
			),
			Transformed(
				Transformed(
					Transformed(tcyl,
						geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg),
					),
					geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg),
				),
				geom.Rotate(Vec{1, 0, 0}, Ez, -45*Deg),
			),
			test.Sheet(test.Checkers2, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 1, 3}),
		8,
		test.DefaultTolerance,
	)
}

func TestTransformed_Bounds_Cylinder(t *testing.T) {
	cyl := Cylinder(test.Normal, 0.5, 1, Vec{0, 0, 0})
	test.QuadView(t,
		NewScene(
			5,
			[]Light{},
			WithBounds(Transformed(cyl, geom.Scale(O, 0.5))),
			WithBounds(Transformed(cyl, geom.Translate(Vec{2, 0, -2}))),
			WithBounds(Transformed(cyl, geom.ComposeLR(
				geom.Rotate(O, Ex, 90*Deg),
				geom.Translate(Vec{2, 0, 0}),
			))),
			WithBounds(Transformed(cyl, geom.ComposeLR(
				geom.Rotate(O, Ez, 45*Deg),
				geom.Rotate(O, Ex, 90*Deg),
				geom.Translate(Vec{0, 0, -2}),
			))),
			test.Sheet(test.Normal, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0.5, 0.7, 3}),
		8,
		test.DefaultTolerance,
	)
}

func BenchmarkTransformed_Cylinder(b *testing.B) {
	cyl := Cylinder(test.Checkers1, 0.5, 1, Vec{0, 0, 0})
	test.Benchmark(b,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{1, 2, 0}),
			},
			Transformed(cyl, geom.Scale(O, 0.5)),
			Transformed(cyl, geom.Translate(Vec{2, 0, -2})),
			Transformed(cyl, geom.ComposeLR(
				geom.Rotate(O, Ex, 90*Deg),
				geom.Translate(Vec{2, 0, 0}),
			)),
			Transformed(cyl, geom.ComposeLR(
				geom.Rotate(O, Ex, 90*Deg),
				geom.Rotate(O, Ez, 45*Deg),
				geom.Translate(Vec{0, 0, -2}),
			)),
			test.Sheet(test.Checkers2, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0.8, 0.7, 2.2}),
		test.DefaultTolerance,
	)
}

// Mutiple chained invocations of Transformed are supposed to get optimized into just one tranform.
// This test
//  - checks that the optimized version yields the same result
//    (golden data was recoreded before the optimization)
//  - benchmarks that there is indeed a speed-up
//    (checked once manually: 3MPix/s -> 8MPix/s)
//  - verifies, in a whitebox manner, that the optimized object indeed
//    has one layer of transformes wrapped around it.
func BenchmarkTransformed_Optimize(b *testing.B) {
	orig := Cylinder(test.Checkers1, 0.5, 1, Vec{0, 0, 0})
	t1 := Transformed(orig, geom.Scale(O, 0.9))
	t2 := Transformed(t1, geom.Rotate(O, Ex, 5*Deg))
	t3 := Transformed(t2, geom.Translate(Vec{0.3, 0, 0}))
	t4 := Transformed(t3, geom.Rotate(O, Ez, 30*Deg))
	t5 := Transformed(t4, geom.Scale(Vec{-0.5, -0.4, 0.2}, 1.5))

	if _, ok := t5.(*transformed).orig.(*transformed); ok {
		b.Errorf("Chained transforms were not optimized, got: %#v", t5)
	}

	test.Benchmark(b,
		NewScene(
			1,
			[]Light{},
			WithBounds(t5),
		),
		cameras.Projective(fov).Translate(Vec{0.8, 0.7, 2.2}),
		test.DefaultTolerance,
	)
}

// Test Tree against golden image obtained without tree
// (slow but certain).
func TestTree(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{},
			Tree(randomSpheres(300, 0.3)...),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 2.5}),
		5,
		2e-5,
	)
}

func TestTree_Bounds(t *testing.T) {
	test.QuadView(t,
		NewScene(
			5,
			[]Light{},
			WithBounds(Tree(randomSpheres(300, 0.3)...)),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 2.5}),
		5,
		test.DefaultTolerance,
	)
}

func BenchmarkTree10(b *testing.B) {
	test.Benchmark(b,
		NewScene(
			1,
			[]Light{},
			Tree(randomSpheres(10, 0.1)...),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 2.5}),
		test.DefaultTolerance,
	)
}

func BenchmarkTree100(b *testing.B) {
	test.Benchmark(b,
		NewScene(1, []Light{},
			Tree(randomSpheres(100, 0.1)...),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 2.5}),
		test.DefaultTolerance,
	)
}

func BenchmarkTree1000(b *testing.B) {
	test.Benchmark(b,
		NewScene(
			1,
			[]Light{},
			Tree(randomSpheres(1000, 0.1)...),
		),
		cameras.Projective(fov).Translate(Vec{0, 0, 2.5}),
		test.DefaultTolerance,
	)
}

func randomSpheres(N int, r float64) []Interface {
	rng := rand.New(rand.NewSource(123))
	rnd := func() float64 {
		return 2*rng.Float64() - 1
	}
	var spheres []Interface
	for i := 0; i < N; i++ {
		spheres = append(spheres, Sphere(
			test.Normal,
			((r/3)*rng.Float64()+(2*r/3)),
			Vec{rnd(), rnd(), rnd()},
		))
	}
	return spheres
}

//// At 8595458d9d77d757 this paniced with bad HitRecord: {9e+99 [0 0 0] [0 0 0] <nil>}
//func TestTree_Regression1(t *testing.T) {
//	cyl := Cylinder(test.Checkers1, 0.5, 1, Vec{0, 0, 0})
//	test.OnePass(t,
//		NewScene(
//			Tree(
//				Transformed(cyl, geom.Scale(O, 0.5)),
//				Transformed(cyl, geom.Translate(Vec{2, 0, -2})),
//				Transformed(cyl, geom.Rotate(O, Ex, 90*Deg).Compose(geom.Translate(Vec{2, 0, 0}))),
//				Transformed(cyl, geom.Rotate(O, Ex, 90*Deg).Compose(geom.Rotate(O, Ez, 45*Deg)).Compose(geom.Translate(Vec{0, 0, -2}))),
//			),
//			test.Sheet(test.Checkers2, -0.5),
//			test.PointLight(Vec{1, 2, 0}),
//		),
//		cameras.NewProjective(Vec{0.8, 0.7, 2.2}),
//		test.DefaultTolerance,
//	)
//}

func TestTriangle(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{1, 2, 3}),
			},
			Triangle(test.Checkers1, Vec{2, 1, 0}, Vec{1, 0, 0}, Vec{2, 0, 0}),
			Triangle(test.Checkers2, Vec{0, 0, 0}, Vec{0, 0, -2}, Vec{-1, 0, 1}),
			Triangle(test.Checkers3, Vec{0, 0, 0}, Vec{0, 1, 0}, Vec{1, 0, 1}),
			test.Sheet(test.Checkers4, -0.0001),
		),
		cameras.Projective(fov).Translate(Vec{0, 1, 2.5}),
		8,
		0.001,
	)
}

func TestParametric_Cylinder(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{0, 1, 0}),
			},
			Parametric(test.Checkers1, 8, 10, func(u, v float64) Vec {
				v *= 2 * Pi
				r := 0.2
				return Vec{
					u,
					r * math.Sin(v),
					r * math.Cos(v),
				}
			}),
			test.Sheet(test.Checkers2, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 0.5, 1.5}),
		3,
		test.DefaultTolerance,
	)
}

func TestParametric_Torus(t *testing.T) {
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{1, 2, 0}),
			},
			Parametric(test.Checkers1, 20, 10, func(u, v float64) Vec {
				u *= 2 * Pi
				v *= 2 * Pi
				r := 0.2
				R := 1.0
				return Vec{
					r * math.Sin(v),
					(R + r*math.Cos(v)) * math.Cos(u),
					(R + r*math.Cos(v)) * math.Sin(u),
				}
			}),
			test.Sheet(test.Checkers2, -0.5),
		),
		cameras.Projective(fov).Translate(Vec{0, 1, 0}),
		8,
		test.DefaultTolerance,
	)
}

func TestQuadrilateral(t *testing.T) {
	// Quadrilateral UV mapping is not perfect yet (see comment on func Quadrilateral).
	// Therefore, we do not test UV mapping until it has been improved.
	test.QuadView(t,
		NewScene(
			1,
			[]Light{
				test.PointLight(Vec{-1, 2, 1}),
			},
			Quadrilateral(test.Yellow, Vec{0, 0, 0}, Vec{2, 0, 0}, Vec{1, 0, -1}, Vec{0, 0, -1}),
			Quadrilateral(test.Blue, Vec{0, 1, 0}, Vec{-1, 2, 0}, Vec{-2, 0, 0}, Vec{0, 0, 0}),
			test.Sheet(test.Checkers4, -0.0001),
		),
		cameras.Projective(fov).Translate(Vec{0, 1, 2.5}),
		8,
		test.DefaultTolerance,
	)
}

//func disableChecks() func() {
//	backup := tracer.Check
//	tracer.Check = false
//	return func() { tracer.Check = backup }
//}
