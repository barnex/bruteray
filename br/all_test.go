package br_test

import (
	"math"
	"testing"

	. "github.com/barnex/bruteray/br"
	. "github.com/barnex/bruteray/csg"
	"github.com/barnex/bruteray/light"
	. "github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
	. "github.com/barnex/bruteray/shape"
	. "github.com/barnex/bruteray/transf"
)

const (
	testW, testH = 300, 200 // test image size
	testRec      = 3        // test recursion depth
)

const defaultTol = 1.0

// Test a flat sphere
func TestSphere(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, 1}, 0.25, Flat(WHITE)),
	)

	Compare(t, raster.Camera(0), e, 1, "sphere", defaultTol)
}

// Test a sphere behind the camera,
// it should be invisible.
func TestBehindCam(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, -1}, 0.25, Flat(WHITE)),
	)

	Compare(t, raster.Camera(0), e, 2, "behindcam", defaultTol)
}

// Test normal vectors.
func TestNormal(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, 2}, 0.25, ShadeNormal(Ez)),
		Sphere(Vec{-0.5, 0, 2}, 0.25, ShadeNormal(Ex)),
		Sphere(Vec{0.5, 0, 2}, 0.25, ShadeNormal(Ey)),
	)

	Compare(t, raster.Camera(0), e, 3, "normals", defaultTol)
}

// Test camera translation.
func TestCamTransl(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, 2}, 0.25, ShadeNormal(Ez)),
	)
	cam := raster.Camera(0).Transl(-0.5, -0.25, 0)

	Compare(t, cam, e, 4, "camtransl", defaultTol)
}

// Test camera rotation
func TestCamRot(t *testing.T) {
	e := NewEnv()

	r := 0.5
	nz := ShadeNormal(Ez)
	e.Add(
		Sphere(Vec{0, 0, 0}, r, nz),
		Sphere(Vec{0, 0, 2}, r, nz),
		Sphere(Vec{0, 0, 4}, r, nz),
		Sphere(Vec{2, 0, 0}, r, nz),
		Sphere(Vec{2, 0, 2}, r, nz),
		Sphere(Vec{2, 0, 4}, r, nz),
		Sphere(Vec{-2, 0, 0}, r, nz),
		Sphere(Vec{-2, 0, 2}, r, nz),
		Sphere(Vec{-2, 0, 4}, r, nz),
	)
	cam := raster.Camera(1).Transl(0, 4, -4).Transf(RotX4(Pi / 5))

	Compare(t, cam, e, 5, "camrot", defaultTol)
}

// Test object transform
func TestObjTransf(t *testing.T) {
	e := NewEnv()

	r := 0.25
	sx := Sphere(Vec{-0.5, 0, 2}, r, ShadeNormal(Ex))
	sy := Sphere(Vec{0, 0, 2}, r, ShadeNormal(Ez))
	sz := Sphere(Vec{0.5, 0, 2}, r, ShadeNormal(Ey))

	rot := RotZ4(Pi / 4)
	e.Add(Transf(sx, rot))
	e.Add(Transf(sy, rot))
	e.Add(Transf(sz, rot))

	Compare(t, raster.Camera(0), e, 6, "objtransf", defaultTol)
}

// Test intersection of two spheres
func TestObjAnd(t *testing.T) {
	e := NewEnv()

	r := 0.5
	s1 := Sphere(Vec{-r / 2, 0, 2}, r, ShadeNormal(Ez))
	s2 := Sphere(Vec{r / 2, 0, 2}, r, ShadeNormal(Ey))
	s := And(s1, s2)
	e.Add(s)

	Compare(t, raster.Camera(0), e, 7, "objand", defaultTol)
}

// Test two partially overlapping spheres
func TestOverlap(t *testing.T) {
	e := NewEnv()

	r := 0.5
	s1 := Sphere(Vec{-r / 2, 0, 2}, r, ShadeNormal(Ez))
	s2 := Sphere(Vec{r / 2, 0, 2}, r, ShadeNormal(Ey))
	e.Add(s1)
	e.Add(s2)

	Compare(t, raster.Camera(0), e, 8, "overlap", defaultTol)
}

// Make a cube out of 3 intersecting slabs
func TestSlabIntersect(t *testing.T) {
	e := NewEnv()

	r := 1.
	s1 := Slab(Ex, -r, r, Flat(RED))
	s2 := Slab(Ey, -r, r, Flat(GREEN))
	s3 := Slab(Ez, -r, r, Flat(BLUE))
	cube := And(And(s1, s2), s3)
	cube = Transf(cube, RotY4(160*Deg).Mul(RotX4(20*Deg)))
	e.Add(cube)
	cam := raster.Camera(1).Transl(0, 0, -5)

	Compare(t, cam, e, 9, "slabintersect", defaultTol)
}

// Use sheets as green grass, blue sky and wall
func TestSheet(t *testing.T) {
	e := NewEnv()

	s1 := Sheet(Ey, -1, Flat(GREEN))
	s2 := Sheet(Ey, 4, Flat(BLUE))
	s3 := Sheet(Ex, -10, Flat(WHITE))
	s4 := Sphere(Vec{1.5, 0, 3}, 1, ShadeNormal(Ez))
	e.Add(s1, s2, s3, s4)
	cam := raster.Camera(1).Transl(0, 0, -1)

	Compare(t, cam, e, 10, "sheet", defaultTol)
}

// Test rectangles
func TestRect(t *testing.T) {
	e := NewEnv()

	const d = 0.5
	const z = 10
	nz := ShadeNormal(Ez)
	r1 := Rect(Vec{-d, 0, z}, Ez, 0.2, 0.1, math.Inf(1), nz)
	r2 := TransfNonCSG(r1, RotZ4(-30*Deg).Mul(Transl4(Vec{1, 0, 0})))
	r3 := SurfaceAnd(
		Rect(Vec{0, 0, z}, Ez, 10, 10, 10, nz),
		Sphere(Vec{0, 0, z}, 0.25, nz),
	)
	e.Add(r1, r2, r3)

	Compare(t, raster.Camera(0), e, 11, "rect", defaultTol)
}

// Test Axis Aligned Box
func TestBox(t *testing.T) {
	e := NewEnv()

	nz := ShadeNormal(Ez)
	b := Box(Vec{0, 0, 0}, 2, 1, 1, nz)
	b = Transf(b, RotY4(150*Deg))
	g := Sheet(Ey, -1, Flat(GREEN.Mul(EV(-4))))
	e.Add(b, g)

	cam := raster.Camera(1).Transl(0, 0, -5)
	Compare(t, cam, e, 12, "box", defaultTol)
}

// Test Diffuse material without interreflection.
func TestDiffuse0(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, -1, Diffuse0(WHITE.Mul(EV(-1))))
	s := Sphere(Vec{}, 1, Diffuse0(WHITE))
	e.Add(g, s)

	l := light.DirLight(Vec{1, 0.5, -4}, WHITE.Mul(EV(0)))
	e.AddLight(l)

	cam := raster.Camera(1).Transl(0, 0, -5)
	Compare(t, cam, e, 13, "diffuse0", defaultTol)
}

// Test CSG OR of objects.
func TestObjOr1(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, -1, Diffuse0(WHITE.Mul(EV(-1))))
	white := Diffuse0(WHITE)
	box := Box(Vec{}, 1, 1, 1, white)
	bar1 := Box(Vec{}, 1.5, 0.5, 0.5, Diffuse0(RED))
	bar2 := Box(Vec{}, 0.5, 1.5, 0.5, Diffuse0(GREEN))
	bar3 := Box(Vec{}, 0.5, 0.5, 1.5, Diffuse0(BLUE))
	crux := Or(Or(box, bar1), Or(bar2, bar3))
	crux = Transf(crux, RotY4(15*Deg))
	e.Add(g, crux)

	l := light.DirLight(Vec{4, 2, -6}, WHITE.Mul(EV(0)))
	e.AddLight(l)

	cam := raster.Camera(1).Transl(0, 1.5, -6).Transf(RotX4(15 * Deg))
	Compare(t, cam, e, 14, "objor1", defaultTol)
}

// Test CSG OR of objects.
func TestObjOr2(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, -1, Diffuse0(WHITE.Mul(EV(-1))))
	b1 := Box(Vec{-2, 0, -.1}, 0.5, 1, 1, Diffuse0(RED))
	b2 := Box(Vec{2, 0, -.1}, 0.5, 1, 1, Diffuse0(GREEN))
	or := Or(b1, b2)
	or = Transf(or, RotY4(-15*Deg))
	e.Add(g, or)

	l := light.DirLight(Vec{8, 2, 0}, WHITE.Mul(EV(0)))
	e.AddLight(l)

	cam := raster.Camera(1).Transl(0, 1.5, -6).Transf(RotX4(15 * Deg))
	Compare(t, cam, e, 15, "objor2", defaultTol)
}

func TestPointLight(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, 0, Diffuse0(WHITE))
	l := Sheet(Ex, -1, Diffuse0(RED))
	r := Sheet(Ex, 1, Diffuse0(GREEN))
	c := Sheet(Ey, 2.2, Diffuse0(WHITE))
	b := Sheet(Ez, 1.1, Diffuse0(WHITE))
	e.Add(g, l, r, c, b)

	R := 0.4
	s := Sphere(Vec{-0.3, R, 0}, R, Diffuse0(WHITE))
	e.Add(s)

	e.AddLight(light.PointLight(Vec{0, 2.1, 0}, WHITE.Mul(4*Pi)))

	cam := raster.Camera(1).Transl(0, 1, -3)
	Compare(t, cam, e, 16, "pointlight", defaultTol)
}

// Test CSG MINUS of objects.
func TestObjMinus(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, 0, Diffuse0(WHITE.Mul(EV(-1))))
	r := 0.8

	b := Box(Vec{}, r, r, r, Diffuse0(WHITE.Mul(EV(-0))))
	s := Sphere(Vec{}, 1, Diffuse0(WHITE))

	dome := Minus(b, s)

	e.Add(g, dome)

	l := light.DirLight(Vec{2, 1.5, -4}, WHITE.Mul(EV(0)))
	e.AddLight(l)

	cam := raster.Camera(1).Transl(0, 1, -3).Transf(RotX4(10 * Deg))
	Compare(t, cam, e, 17, "objminus", defaultTol)
}

// There is a box buried underneath the floor,
// it should not cast a shadow.
func TestShadowBehind(t *testing.T) {
	e := NewEnv()
	const r = 0.8
	e.Add(
		Sheet(Ey, 0, Diffuse0(WHITE)),
		Box(Vec{0, -1, 0}, r, r, r, Diffuse0(WHITE)),
	)
	e.AddLight(light.PointLight(Vec{1, 4, -4}, WHITE.Mul(EV(5)).Mul(4*Pi)))

	cam := raster.Camera(1).Transl(0, 1, -3)
	Compare(t, cam, e, 19, "shadowbehind", defaultTol)
}

// Test that we can see the reflection of a light.
func TestLuminousObject(t *testing.T) {
	e := NewEnv()
	e.Add(
		Sheet(Ey, -1.0, Diffuse(WHITE.Mul(EV(-0.3)))),
		Sphere(Vec{0, 0.5, 3}, 1.5, Shiny(RED, EV(-3))),
		Sphere(Vec{-2, 0.1, 0}, 1.1, Shiny(BLUE.EV(-0.3), EV(-3))),
		Sphere(Vec{2, 0, -1}, 1, Shiny(GREEN.EV(-1), EV(-3))),
		Sphere(Vec{0, -0.2, -2}, 0.8, Shiny(WHITE, EV(-2))),
		Sphere(Vec{4, 4, 2}, 1, Diffuse(WHITE.EV(-8))),
		//Sphere(Vec{7, 7, -5}, 1, Flat(WHITE.EV(-2))),
	)
	e.AddLight(
		light.Sphere(Vec{3, 3, 1}, 0.1, WHITE.Mul(EV(8))),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-5))))
	e.Recursion = 3

	cam := raster.Camera(1).Transl(0, 2.5, -7).Transf(RotX4(22 * Deg))
	cam.AA = true
	img := raster.MakeImage(testW, testH)
	nPass := 8
	raster.MultiPass(cam, e, img, nPass)
	CompareImg(t, cam, e, img, 20, "luminous-object", 10)
}

// Test a quad surface
func TestQuad(t *testing.T) {
	e := NewEnv()
	e.Add(
		Sheet(Ey, -1.0, Diffuse0(WHITE.Mul(EV(-0.3)))),
		Quad(Vec{0, 0, 0}, Vec{1, -1, 1}, 1, Diffuse0(RED)),
	)
	e.AddLight(
		light.PointLight(Vec{3, 3, -7}, WHITE.Mul(EV(9))),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-5))))

	cam := raster.Camera(1).Transl(0, 0, -7)
	cam.AA = false
	Compare(t, cam, e, 21, "quad-hyper", defaultTol)
}

func TestHollowAnd(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sheet(Ey, -1.0, Diffuse(WHITE.Mul(EV(-0.3)))),
		SurfaceAnd(
			Sphere(Vec{}, 1, Shiny(RED, EV(-3))),
			Slab(Ey, -0.3, 0.3, Flat(RED)),
		),
	)
	e.AddLight(
		light.Sphere(Vec{3, 3, 1}, 0.1, WHITE.Mul(EV(8))),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-5))))
	e.Recursion = 3

	cam := raster.Camera(1).Transl(0, 2.5, -7).Transf(RotX4(22 * Deg))
	cam.AA = true
	nPass := 8
	tol := 10.0
	CompareNPass(t, cam, e, 22, "hollowand", nPass, tol, testW, testH)
}

// Test a rectangular light source.
// Light fall-off should follow cosine.
func TestRectLight(t *testing.T) {
	e := NewEnv()

	r := 0.3

	e.Add(
		Sphere(Vec{0, 2, 0}, r, Diffuse(WHITE)),
		Sphere(Vec{1.41, 1.41, 0}, r, Diffuse(WHITE)),
		Sphere(Vec{2, 0, 0}, r, Diffuse(WHITE)),
		Sphere(Vec{1.41, -1.41, 0}, r, Diffuse(WHITE)),
		Sphere(Vec{0, -2, 0}, r, Diffuse(WHITE)),
		Sphere(Vec{-2, 2, 0}, r, Diffuse(WHITE)),
		Sphere(Vec{-2, -2, 0}, r, Diffuse(WHITE)),
	)

	reference := false
	intens := WHITE.EV(-1)
	if reference {
		e.Add(
			Rect(Vec{0, 0, 0}, Ey, 1, 1, 1, Flat(intens)), // add for the reference image
		)
	} else {
		e.AddLight(
			light.RectLight(Vec{0, 0, 0}, 1, 0, 1, intens), // remove for reference image
		)
	}
	e.Recursion = 2

	cam := raster.Camera(1).Transl(0, -.5, -6)
	nPass := 50
	CompareNPass(t, cam, e, 23, "rectlight", nPass, 10, testW, testH)

}

func TestCornellBox(t *testing.T) {
	t.Skip("TODO: fix rectlight")
	e := NewEnv()

	white := Diffuse(WHITE.EV(-.6))
	green := Diffuse(GREEN.EV(-1.6))
	red := Diffuse(RED.EV(-.6))

	const (
		w = 550
		h = 550
		U = 66666
	)

	e.Add(
		Rect(Vec{0, 0, 0}, Ey, w/2, U, w/2, white),
		Rect(Vec{0, h, 0}, Ey, w/2, U, w/2, white),
		Rect(Vec{0, h / 2, w / 2}, Ez, w/2, h/2, U, white),
		Rect(Vec{w / 2, h / 2, 0}, Ex, U, h/2, w/2, green),
		Rect(Vec{-w / 2, h / 2, 0}, Ex, U, h/2, w/2, red),
		Transf(Box(Vec{120, 80, -80}, 80, 80, 80, white), RotY4(-18*Deg)),
		Transf(Box(Vec{-50, 165, 100}, 85, 180, 70, white), RotY4(20*Deg)),
	)

	e.AddLight(
		light.RectLight(Vec{0, h - 1e-4, 0}, 120/2, 0, 120/2, Color{1.0, 1.0, 0.6}.EV(18)),
	)

	e.SetAmbient(Flat(WHITE.EV(-6)))
	e.Recursion = 10
	e.Cutoff = EV(3)

	focalLen := 0.035 / 0.025
	cam := raster.Camera(focalLen).Transl(0, h/2, -1000)
	cam.AA = true
	nPass := 20
	CompareNPass(t, cam, e, 24, "cornellbox", nPass, 10, testW, testH)
}

func TestDOF(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sheet(Ey, 0, Checkboard(1, Flat(WHITE), Flat(BLACK))),
	)

	e.SetAmbient(Flat(WHITE.EV(-2)))

	cam := raster.Camera(1).Transl(0, 10, 0).Transf(RotX4(30 * Deg)) //.RotScene(30 * Deg).Transf(RotX4(30 * Deg))
	cam.AA = true
	cam.Aperture = 1
	cam.Focus = 20
	e.Recursion = 1
	nPass := 30
	CompareNPass(t, cam, e, 25, "dof", nPass, 10, testW, testH)
}

func TestDiafragmDisk(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, 20}, 1, Flat(WHITE.EV(8))),
	)

	cam := raster.Camera(1)
	cam.AA = true
	cam.Aperture = 0.8
	cam.Focus = 2
	e.Recursion = 1

	nPass := 300
	CompareNPass(t, cam, e, 26, "diaphragm-disk", nPass, 8, testW/4, testH/4)
}

func TestDiaphragmHex(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, 20}, 1, Flat(WHITE.EV(8))),
	)
	e.Recursion = 1

	cam := raster.Camera(1)
	cam.AA = true
	cam.Aperture = 0.8
	cam.Focus = 2
	cam.Diaphragm = DiaHex
	nPass := 300
	CompareNPass(t, cam, e, 27, "diaphragm-hex", nPass, 8, testW/4, testH/4)
}

func TestCyl(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sheet(Ey, 0.0, Checkboard(1, Diffuse0(WHITE), Diffuse0(GREY))),
		Cyl(Y, Vec{1, 0.5, 0}, 2, 0.5, Diffuse0(RED)),
		Cyl(X, Vec{-0.8, 0.5, 0}, 1, 0.5, Diffuse0(GREEN)),
		Cyl(Z, Vec{-2, 1, 2}, 2, 1, Diffuse0(BLUE)),
	)
	e.AddLight(
		light.PointLight(Vec{-3, 5, -2}, WHITE.Mul(EV(8))),
		light.PointLight(Vec{3, 8, -6}, WHITE.Mul(EV(8))),
	)

	cam := raster.Camera(1).Transl(0, 3, -3).Transf(RotX4(30 * Deg))
	cam.AA = false
	Compare(t, cam, e, 28, "cyl", defaultTol)
}

func TestMultiOr(t *testing.T) {
	e := NewEnv()

	const (
		w      = 3.0            // central width
		h1     = 3.0            // pillar height
		pill   = 0.5            // pillar width
		d      = 3 * (w + pill) // central depth
		pointy = 0.8            //
		b      = 0.5
	)

	chestz := func(pos Vec, w, h1, d, pointy float64, m Material) CSGObj {
		const off = 1e-4
		c1 := Cyl(Z, Vec{pointy / 2, h1, 0}.Add(pos), w+pointy, d-off, m)
		c2 := Cyl(Z, Vec{-pointy / 2, h1, 0}.Add(pos), w+pointy, d+off, m)
		ceil := And(c1, c2)
		box := NBox(w, h1, d+2*off, m).Transl(pos.Add(Vec{Y: h1 / 2}))
		return Or(box, ceil)
	}

	chestx := func(pos Vec, w, h1, d, pointy float64, m Material) CSGObj {
		const off = 1e-4
		c1 := Cyl(X, Vec{0, h1, pointy / 2}.Add(pos), w+pointy, d-off, m)
		c2 := Cyl(X, Vec{0, h1, -pointy / 2}.Add(pos), w+pointy, d+off, m)
		ceil := And(c1, c2)
		box := NBox(d+2*off, h1, w, m).Transl(pos.Add(Vec{Y: h1 / 2}))
		return Or(box, ceil)
	}

	white := ShadeShape(WHITE)
	tileB := ShadeShape(WHITE.EV(-5))
	tileW := ShadeShape(WHITE.EV(-1))

	e.Add(
		Sheet(Ey, 0.01, Checkboard(1, tileW, tileB)),
		MultiOr(
			chestz(Vec{}, w, h1, d, pointy, white),
			chestz(Vec{Y: b}, w-b, h1-b, d+4, pointy, Flat(WHITE)),

			chestz(Vec{X: -w - pill}, w, h1, d, pointy, white),
			chestz(Vec{X: +w + pill}, w, h1, d, pointy, white),

			chestx(Vec{}, w, h1, d, pointy, white),
			chestx(Vec{Z: -w - pill}, w, h1, d, pointy, white),
			chestx(Vec{Z: +w + pill}, w, h1, d, pointy, white),
		),
	)

	cam := raster.Camera(0.65).Transl(0, 2.2, -4.9).RotScene(0 * Deg).Transf(RotX4(-0 * Deg))
	Compare(t, cam, e, 29, "multi-or", defaultTol)
}

func TestCSG2(t *testing.T) {
	e := NewEnv()

	white := Diffuse0(WHITE.EV(-.3))
	red := Diffuse0(RED.EV(-.6))
	blue := Diffuse0(BLUE.EV(-.6))

	floor := Sheet(Ey, 0, white)

	cube := NBox(1, 1, 1, red).Transl(Vec{0, 0.5001, 0})
	sphere := NSphere(1, blue).Transl(cube.Corner(1, 1, -1))

	e.Add(
		floor,
		Minus(cube, sphere),
	)

	e.AddLight(
		light.PointLight(Vec{4, 5, -2.5}, WHITE.EV(8)),
	)

	e.SetAmbient(Flat(WHITE.EV(-2)))
	e.Recursion = 4

	cam := raster.Camera(1).Transl(0, 2.1, -2.5).RotScene(30 * Deg).Transf(RotX4(30 * Deg))
	Compare(t, cam, e, 30, "csg2", defaultTol)
}

//func TestDistort(t *testing.T) {
//	e := NewEnv()
//
//	m := Shiny(RED, EV(-2))
//	m = Distort(1, 10, Vec{50, 50, 50}, 0.03, m)
//
//	e.Add(
//		Sheet(Ey, 0, Checkboard(1, Diffuse0(BLUE.EV(-3)), Diffuse0(WHITE))),
//		Sphere(Vec{0, 0.5, 0}, 1, m),
//	)
//
//	e.AddLight(
//		SphereLight(Vec{4, 5, -3}, 1, WHITE.EV(10)),
//	)
//
//	e.Camera = Camera(1).Transl(0, 2, -3).Transf(RotX4(25 * Deg))
//
//	Compare(t, e, 25, "028-waves")
//}

//func TestTexture(t *testing.T) {
//	e := NewEnv()
//
//	m := Texture(MustLoad("assets/monalisa.jpg"), 100)
//
//	e.Add(
//		Sheet(Ey, 0, Checkboard(1, Diffuse0(BLUE.EV(-3)), Diffuse0(WHITE))),
//		Sphere(Vec{0, 0.5, 0}, 1, m),
//	)
//
//	e.AddLight(
//		light.Sphere(Vec{4, 5, -3}, 1, WHITE.EV(10)),
//	)
//
//	cam := raster.Camera(1).Transl(0, 2, -3).Transf(RotX4(25 * Deg))
//
//	Compare(t, cam, e, 30, "texture", defaultTol)
//}
