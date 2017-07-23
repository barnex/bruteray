package bruteray

import (
	"testing"
)

const (
	testW, testH = 300, 200 // test image size
	testRec      = 3        // test recursion depth
)

// Test a flat sphere
func TestSphere(t *testing.T) {
	e := NewEnv()

	e.Add(Sphere(Vec{0, 0, 1}, 0.25, Flat(WHITE)))
	c := Camera(0)

	Compare(t, e, c, "001-sphere")
}

// Test a sphere behind the camera
func TestBehindCam(t *testing.T) {
	e := NewEnv()

	e.Add(Sphere(Vec{0, 0, -1}, 0.25, Flat(WHITE)))

	Compare(t, e, Camera(0), "002-behindcam")
}

// Test normal vectors
func TestNormal(t *testing.T) {
	e := NewEnv()

	e.Add(Sphere(Vec{0, 0, 2}, 0.25, ShadeNormal(Ez)))
	e.Add(Sphere(Vec{-0.5, 0, 2}, 0.25, ShadeNormal(Ex)))
	e.Add(Sphere(Vec{0.5, 0, 2}, 0.25, ShadeNormal(Ey)))

	Compare(t, e, Camera(0), "003-normals")
}

// Test camera translation
func TestCamTransl(t *testing.T) {
	e := NewEnv()

	e.Add(Sphere(Vec{0, 0, 2}, 0.25, ShadeNormal(Ez)))

	Compare(t, e, Camera(0).Transl(-0.5, -0.25, 0), "004-camtransl")
}

// Test camera rotation
func TestCamRot(t *testing.T) {
	e := NewEnv()

	r := 0.5
	nz := ShadeNormal(Ez)
	e.Add(Sphere(Vec{0, 0, 0}, r, nz))
	e.Add(Sphere(Vec{0, 0, 2}, r, nz))
	e.Add(Sphere(Vec{0, 0, 4}, r, nz))
	e.Add(Sphere(Vec{2, 0, 0}, r, nz))
	e.Add(Sphere(Vec{2, 0, 2}, r, nz))
	e.Add(Sphere(Vec{2, 0, 4}, r, nz))
	e.Add(Sphere(Vec{-2, 0, 0}, r, nz))
	e.Add(Sphere(Vec{-2, 0, 2}, r, nz))
	e.Add(Sphere(Vec{-2, 0, 4}, r, nz))

	Compare(t, e, Camera(1).Transl(0, 4, -4).Transf(RotX4(pi/5)), "005-camrot")
}

// Test object transform
func TestObjTransf(t *testing.T) {
	e := NewEnv()

	r := 0.25
	sx := Sphere(Vec{-0.5, 0, 2}, r, ShadeNormal(Ex))
	sy := Sphere(Vec{0, 0, 2}, r, ShadeNormal(Ez))
	sz := Sphere(Vec{0.5, 0, 2}, r, ShadeNormal(Ey))

	rot := RotZ4(pi / 4)
	e.Add(Transf(sx, rot))
	e.Add(Transf(sy, rot))
	e.Add(Transf(sz, rot))

	Compare(t, e, Camera(0), "006-objtransf")
}

// Test intersection of two spheres
func TestObjAnd(t *testing.T) {
	e := NewEnv()

	r := 0.5
	s1 := Sphere(Vec{-r / 2, 0, 2}, r, ShadeNormal(Ez))
	s2 := Sphere(Vec{r / 2, 0, 2}, r, ShadeNormal(Ey))
	s := ObjAnd(s1, s2)
	e.Add(s)

	Compare(t, e, Camera(0), "007-objand")
}

// Test two partially overlapping spheres
func TestOverlap(t *testing.T) {
	e := NewEnv()

	r := 0.5
	s1 := Sphere(Vec{-r / 2, 0, 2}, r, ShadeNormal(Ez))
	s2 := Sphere(Vec{r / 2, 0, 2}, r, ShadeNormal(Ey))
	e.Add(s1)
	e.Add(s2)

	Compare(t, e, Camera(0), "008-overlap")
}

// Make a cube out of 3 intersecting slabs
func TestSlabIntersect(t *testing.T) {
	e := NewEnv()

	r := 1.
	s1 := Slab(Ex, -r, r, Flat(RED))
	s2 := Slab(Ey, -r, r, Flat(GREEN))
	s3 := Slab(Ez, -r, r, Flat(BLUE))
	cube := ObjAnd(ObjAnd(s1, s2), s3)
	cube = Transf(cube, RotY4(160*deg).Mul(RotX4(20*deg)))
	e.Add(cube)

	Compare(t, e, Camera(1).Transl(0, 0, -4), "009-slabintersect")
}

// Use sheets as green grass, blue sky and wall
func TestSheet(t *testing.T) {
	e := NewEnv()

	s1 := Sheet(Ey, -1, Flat(GREEN))
	s2 := Sheet(Ey, 4, Flat(BLUE))
	s3 := Sheet(Ex, -10, Flat(WHITE))
	s4 := Sphere(Vec{1.5, 0, 3}, 1, ShadeNormal(Ez))
	e.Add(s1, s2, s3, s4)

	Compare(t, e, Camera(1), "010-sheet")
}

// Test rectangles
func TestRect(t *testing.T) {
	e := NewEnv()

	const d = 0.5
	const z = 10
	nz := ShadeNormal(Ez)
	r1 := Rect(Vec{-d, 0, z}, Ez, 0.2, 0.1, inf, nz)
	r2 := Transf(r1, RotZ4(-30*deg).Mul(Transl4(Vec{1, 0, 0})))
	r3 := ObjAnd(
		Rect(Vec{0, 0, z}, Ez, 10, 10, 10, nz),
		Sphere(Vec{0, 0, z}, 0.25, nz),
	)
	e.Add(r1, r2, r3)

	Compare(t, e, Camera(0), "011-rect")
}

// Test Axis Aligned Box
func TestBox(t *testing.T) {
	e := NewEnv()

	nz := ShadeNormal(Ez)
	b := Box(Vec{0, 0, 0}, 2, 1, 1, nz)
	b = Transf(b, RotY4(150*deg))
	g := Sheet(Ey, -1, Flat(GREEN.Mul(EV(-4))))
	e.Add(b, g)

	Compare(t, e, Camera(1).Transl(0, 0, -4), "012-box")
}

func TestDiffuse0(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, -1, Diffuse0(WHITE.Mul(EV(-1))))
	s := Sphere(Vec{}, 1, Diffuse0(WHITE))
	e.Add(g, s)

	l := DirLight(Vec{1, 0.5, -4}, WHITE.Mul(EV(0)))
	e.AddLight(l)

	Compare(t, e, Camera(1).Transl(0, 0, -4), "013-diffuse0")
}

func TestObjOr1(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, -1, Diffuse0(WHITE.Mul(EV(-1))))
	white := Diffuse0(WHITE)
	box := Box(Vec{}, 1, 1, 1, white)
	bar1 := Box(Vec{}, 1.5, 0.5, 0.5, Diffuse0(RED))
	bar2 := Box(Vec{}, 0.5, 1.5, 0.5, Diffuse0(GREEN))
	bar3 := Box(Vec{}, 0.5, 0.5, 1.5, Diffuse0(BLUE))
	crux := ObjOr(ObjOr(box, bar1), ObjOr(bar2, bar3))
	crux = Transf(crux, RotY4(15*deg))
	e.Add(g, crux)

	l := DirLight(Vec{4, 2, -6}, WHITE.Mul(EV(0)))
	e.AddLight(l)

	Compare(t, e, Camera(1).Transl(0, 1.5, -5).Transf(RotX4(15*deg)), "014-objor1")
}

func TestObjOr2(t *testing.T) {
	e := NewEnv()

	g := Sheet(Ey, -1, Diffuse0(WHITE.Mul(EV(-1))))
	b1 := Box(Vec{-2, 0, -.1}, 0.5, 1, 1, Diffuse0(RED))
	b2 := Box(Vec{2, 0, -.1}, 0.5, 1, 1, Diffuse0(GREEN))
	or := ObjOr(b1, b2)
	or = Transf(or, RotY4(-15*deg))
	e.Add(g, or)

	l := DirLight(Vec{8, 2, 0}, WHITE.Mul(EV(0)))
	e.AddLight(l)

	Compare(t, e, Camera(1).Transl(0, 1.5, -5).Transf(RotX4(15*deg)), "015-objor2")
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

	e.AddLight(PointLight(Vec{0, 2.1, 0}, WHITE.Mul(EV(0))))

	Compare(t, e, Camera(1).Transl(0, 1, -2), "016-pointlight")
}

//func TestObjMinus(t *testing.T) {
//	e := NewEnv()
//
//	g := Object(Sheet(Ey, 0), Diffuse0(WHITE.Mul(EV(-1))))
//	r := 0.8
//
//	b := Object(Box(Vec{}, r, r, r), Diffuse0(WHITE.Mul(EV(-0))))
//	s := Object(Sphere(Vec{}, 1), Diffuse0(WHITE))
//
//	dome := ObjMinus(b, s)
//
//	e.Add(g, dome)
//
//	l := DirLight(Vec{2, 1.5, -4}, WHITE.Mul(EV(0)))
//	e.AddLight(l)
//
//	Compare(t, e, Camera(1).Transl(0, 1, -2).Transf(RotX4(10*deg)), "014-objminus")
//}

//todo: unit test cube intersect, unit test objminus
