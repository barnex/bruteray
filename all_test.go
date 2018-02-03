package bruteray

import (
	"fmt"
	"testing"
)

const (
	testW, testH = 300, 200 // test image size
	testRec      = 3        // test recursion depth
)

// Test a flat sphere
func TestSphere(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, 1}, 0.25, Flat(WHITE)),
	)

	Compare(t, e, "001-sphere")
}

// Test a sphere behind the camera
func TestBehindCam(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, -1}, 0.25, Flat(WHITE)),
	)

	Compare(t, e, "002-behindcam")
}

// Test normal vectors
func TestNormal(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, 2}, 0.25, ShadeNormal(Ez)),
		Sphere(Vec{-0.5, 0, 2}, 0.25, ShadeNormal(Ex)),
		Sphere(Vec{0.5, 0, 2}, 0.25, ShadeNormal(Ey)),
	)

	Compare(t, e, "003-normals")
}

// Test camera translation
func TestCamTransl(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, 2}, 0.25, ShadeNormal(Ez)),
	)
	e.Camera = Camera(0).Transl(-0.5, -0.25, 0)

	Compare(t, e, "004-camtransl")
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
	e.Camera = Camera(1).Transl(0, 4, -4).Transf(RotX4(Pi / 5))

	Compare(t, e, "005-camrot")
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

	Compare(t, e, "006-objtransf")
}

// Test intersection of two spheres
func TestObjAnd(t *testing.T) {
	e := NewEnv()

	r := 0.5
	s1 := Sphere(Vec{-r / 2, 0, 2}, r, ShadeNormal(Ez))
	s2 := Sphere(Vec{r / 2, 0, 2}, r, ShadeNormal(Ey))
	s := And(s1, s2)
	e.Add(s)

	Compare(t, e, "007-objand")
}

// Test two partially overlapping spheres
func TestOverlap(t *testing.T) {
	e := NewEnv()

	r := 0.5
	s1 := Sphere(Vec{-r / 2, 0, 2}, r, ShadeNormal(Ez))
	s2 := Sphere(Vec{r / 2, 0, 2}, r, ShadeNormal(Ey))
	e.Add(s1)
	e.Add(s2)

	Compare(t, e, "008-overlap")
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
	e.Camera = Camera(1).Transl(0, 0, -5)

	Compare(t, e, "009-slabintersect")
}

// Use sheets as green grass, blue sky and wall
func TestSheet(t *testing.T) {
	e := NewEnv()

	s1 := Sheet(Ey, -1, Flat(GREEN))
	s2 := Sheet(Ey, 4, Flat(BLUE))
	s3 := Sheet(Ex, -10, Flat(WHITE))
	s4 := Sphere(Vec{1.5, 0, 3}, 1, ShadeNormal(Ez))
	e.Add(s1, s2, s3, s4)
	e.Camera = Camera(1).Transl(0, 0, -1)

	Compare(t, e, "010-sheet")
}

// Test rectangles
func TestRect(t *testing.T) {
	e := NewEnv()

	const d = 0.5
	const z = 10
	nz := ShadeNormal(Ez)
	r1 := Rect(Vec{-d, 0, z}, Ez, 0.2, 0.1, inf, nz)
	r2 := Transf(r1, RotZ4(-30*Deg).Mul(Transl4(Vec{1, 0, 0})))
	r3 := And(
		Rect(Vec{0, 0, z}, Ez, 10, 10, 10, nz),
		Sphere(Vec{0, 0, z}, 0.25, nz),
	)
	e.Add(r1, r2, r3)

	Compare(t, e, "011-rect")
}

// Test Axis Aligned Box
func TestBox(t *testing.T) {
	e := NewEnv()

	nz := ShadeNormal(Ez)
	b := Box(Vec{0, 0, 0}, 2, 1, 1, nz)
	b = Transf(b, RotY4(150*Deg))
	g := Sheet(Ey, -1, Flat(GREEN.Mul(EV(-4))))
	e.Add(b, g)
	e.Camera = Camera(1).Transl(0, 0, -5)

	Compare(t, e, "012-box")
}

func TestDiffuse0(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, -1, Diffuse0(WHITE.Mul(EV(-1))))
	s := Sphere(Vec{}, 1, Diffuse0(WHITE))
	e.Add(g, s)

	l := DirLight(Vec{1, 0.5, -4}, WHITE.Mul(EV(0)))
	e.AddLight(l)
	e.Camera = Camera(1).Transl(0, 0, -5)

	Compare(t, e, "013-diffuse0")
}

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

	l := DirLight(Vec{4, 2, -6}, WHITE.Mul(EV(0)))
	e.AddLight(l)
	e.Camera = Camera(1).Transl(0, 1.5, -6).Transf(RotX4(15 * Deg))

	Compare(t, e, "014-objor1")
}

func TestObjOr2(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, -1, Diffuse0(WHITE.Mul(EV(-1))))
	b1 := Box(Vec{-2, 0, -.1}, 0.5, 1, 1, Diffuse0(RED))
	b2 := Box(Vec{2, 0, -.1}, 0.5, 1, 1, Diffuse0(GREEN))
	or := Or(b1, b2)
	or = Transf(or, RotY4(-15*Deg))
	e.Add(g, or)

	l := DirLight(Vec{8, 2, 0}, WHITE.Mul(EV(0)))
	e.AddLight(l)

	e.Camera = Camera(1).Transl(0, 1.5, -6).Transf(RotX4(15 * Deg))

	Compare(t, e, "015-objor2")
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

	e.AddLight(PointLight(Vec{0, 2.1, 0}, WHITE.Mul(4*Pi)))
	e.Camera = Camera(1).Transl(0, 1, -3)

	Compare(t, e, "016-pointlight")
}

// Test convergence of diffuse interreflection:
//
// 1) We are inside a highly reflective white box containing a point light.
// If the pre-factor for interreflection is slightly too high,
// deeper recursion will diverge to an infinitely bright image instead of converge.
//
// 2) We are inside a 100% reflective white box containing a point light.
// This is not a physical situation and the intensity should diverge to infinity.
// If the pre-factor for interreflection is slightly too low, divergence will not happen.
func TestDiffuse1(t *testing.T) {
	for _, refl := range []float64{0.8, 1} {
		refl := refl
		for _, r := range []int{1, 16, 128} {
			e := whitebox(refl)
			e.Recursion = r
			t.Run(fmt.Sprintf("refl=%v,rec=%v", refl, e.Recursion), func(t *testing.T) {
				t.Parallel()
				img := MakeImage(testW/4, testH/4)
				nPass := 2
				MultiPass(e, img, nPass)
				name := fmt.Sprintf("017-diffuse1-refl%v-rec%v", refl, e.Recursion)
				CompareImg(t, e, img, name, 10)
			})
		}
	}
}

func whitebox(refl float64) *Env {
	e := NewEnv()
	white := Diffuse(WHITE.Mul(refl))
	e.Add(
		Sheet(Ey, -1, white),
		Sheet(Ey, 1, white),
		Sheet(Ex, -1, white),
		Sheet(Ex, 1, white),
		Sheet(Ez, -1, white),
		Sheet(Ez, 1, white),
	)
	e.AddLight(PointLight(Vec{}, WHITE.Mul(EV(-3)).Mul(4*Pi)))
	e.Camera = Camera(0.75).Transl(0, 0, -1)
	e.Camera.AA = true
	return e
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
	e.AddLight(PointLight(Vec{1, 4, -4}, WHITE.Mul(EV(5)).Mul(4*Pi)))
	e.Camera = Camera(1).Transl(0, 1, -3)
	Compare(t, e, "019-shadowbehind")
}

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
		SphereLight(Vec{3, 3, 1}, 0.1, WHITE.Mul(EV(8))),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-5))))

	e.Camera = Camera(1).Transl(0, 2.5, -7).Transf(RotX4(22 * Deg))
	e.Camera.AA = true

	img := MakeImage(testW, testH)
	nPass := 8
	e.Recursion = 3
	MultiPass(e, img, nPass)
	name := "020-luminous-object"
	CompareImg(t, e, img, name, 10)
}

func TestQuad(t *testing.T) {
	e := NewEnv()
	e.Add(
		Sheet(Ey, -1.0, Diffuse0(WHITE.Mul(EV(-0.3)))),
		Quad(Vec{0, 0, 0}, Vec{1, -1, 1}, 1, Diffuse0(RED)),
	)
	e.AddLight(
		PointLight(Vec{3, 3, -7}, WHITE.Mul(EV(9))),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-5))))

	e.Camera = Camera(1).Transl(0, 0, -7)
	e.Camera.AA = false
	Compare(t, e, "021-quad-hyper")
}

func TestObjMinus(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, 0, Diffuse0(WHITE.Mul(EV(-1))))
	r := 0.8

	b := Box(Vec{}, r, r, r, Diffuse0(WHITE.Mul(EV(-0))))
	s := Sphere(Vec{}, 1, Diffuse0(WHITE))

	dome := Minus(b, s)

	e.Add(g, dome)

	l := DirLight(Vec{2, 1.5, -4}, WHITE.Mul(EV(0)))
	e.AddLight(l)
	e.Camera = Camera(1).Transl(0, 1, -3).Transf(RotX4(10 * Deg))

	Compare(t, e, "014-objminus")
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
		SphereLight(Vec{3, 3, 1}, 0.1, WHITE.Mul(EV(8))),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-5))))

	e.Camera = Camera(1).Transl(0, 2.5, -7).Transf(RotX4(22 * Deg))
	e.Camera.AA = true

	img := MakeImage(testW, testH)
	nPass := 8
	e.Recursion = 3
	MultiPass(e, img, nPass)
	name := "022-hollowand"
	CompareImg(t, e, img, name, 10)
}

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
			RectLight(Vec{0, 0, 0}, 1, 0, 1, intens), // remove for reference image
		)
	}

	e.Camera = Camera(1).Transl(0, -.5, -6)

	img := MakeImage(testW, testH)
	nPass := 50
	e.Recursion = 2
	MultiPass(e, img, nPass)
	name := "023-rectlight"
	CompareImg(t, e, img, name, 10)

}

func TestCornellBox(t *testing.T) {
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
		RectLight(Vec{0, h - 1e-4, 0}, 120/2, 0, 120/2, Color{1.0, 1.0, 0.6}.EV(18)),
	)

	e.SetAmbient(Flat(WHITE.EV(-6)))

	focalLen := 0.035 / 0.025
	e.Camera = Camera(focalLen).Transl(0, h/2, -1000)
	e.Camera.AA = true
	e.Recursion = 10
	e.Cutoff = EV(3)

	nPass := 20
	img := MakeImage(testW, testH)
	MultiPass(e, img, nPass)
	name := "024-cornellbox"
	CompareImg(t, e, img, name, 10)
}

func TestDOF(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sheet(Ey, 0, Checkboard(1, Flat(WHITE), Flat(BLACK))),
	)

	e.SetAmbient(Flat(WHITE.EV(-2)))

	e.Camera = Camera(1).Transl(0, 10, 0).Transf(RotX4(30 * Deg)) //.RotScene(30 * Deg).Transf(RotX4(30 * Deg))
	e.Camera.AA = true
	e.Camera.Aperture = 1
	e.Camera.Focus = 20
	e.Recursion = 1

	nPass := 30
	img := MakeImage(testW, testH)
	MultiPass(e, img, nPass)
	CompareImg(t, e, img, "025-dof", 10)
}

func TestDiafragmDisk(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, 20}, 1, Flat(WHITE.EV(8))),
	)

	e.Camera = Camera(1)
	e.Camera.AA = true
	e.Camera.Aperture = 0.7
	e.Camera.Focus = 2
	e.Recursion = 1

	nPass := 100
	img := MakeImage(testW, testH)
	MultiPass(e, img, nPass)
	CompareImg(t, e, img, "026-diaphragm-disk", 20)
}

func TestDiaphragmHex(t *testing.T) {
	e := NewEnv()

	e.Add(
		Sphere(Vec{0, 0, 20}, 1, Flat(WHITE.EV(8))),
	)

	e.Camera = Camera(1)
	e.Camera.AA = true
	e.Camera.Aperture = 0.7
	e.Camera.Focus = 2
	e.Camera.Diaphragm = DiaHex
	e.Recursion = 1

	nPass := 100
	img := MakeImage(testW, testH)
	MultiPass(e, img, nPass)
	CompareImg(t, e, img, "027-diaphragm-hex", 20)
}

//func TestWaves(t *testing.T) {
//	e := NewEnv()
//
//	a := Diffuse(WHITE.EV(-.3))
//	b := Diffuse(WHITE.EV(-3))
//	f := func(v float64) Material {
//		if v > 0 {
//			return a
//		} else {
//			return b
//		}
//	}
//	m := Waves(20, Vec{1, 1, 1}, f)
//	e.Add(
//
//		Sheet(Ey, 0, m),
//	)
//
//	e.Camera = Camera(1).Transl(0, 100, 0).Transf(RotX4(90 * Deg))
//	e.Camera.AA = true
//
//	Compare(t, e, "028-waves")
//}
