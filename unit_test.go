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

	e.Add(Object(Sphere(Vec{0, 0, 1}, 0.25), Flat(WHITE)))
	c := Camera(0)

	Compare(t, e, c, "001-sphere")
}

// Test a sphere behind the camera
func TestBehindCam(t *testing.T) {
	e := NewEnv()

	e.Add(Object(Sphere(Vec{0, 0, -1}, 0.25), Flat(WHITE)))

	Compare(t, e, Camera(0), "002-behindcam")
}

// Test normal vectors
func TestNormal(t *testing.T) {
	e := NewEnv()

	e.Add(Object(Sphere(Vec{0, 0, 2}, 0.25), ShadeNormal(Ez)))
	e.Add(Object(Sphere(Vec{-0.5, 0, 2}, 0.25), ShadeNormal(Ex)))
	e.Add(Object(Sphere(Vec{0.5, 0, 2}, 0.25), ShadeNormal(Ey)))

	Compare(t, e, Camera(0), "003-normals")
}

// Test camera translation
func TestCamTransl(t *testing.T) {
	e := NewEnv()

	e.Add(Object(Sphere(Vec{0, 0, 2}, 0.25), ShadeNormal(Ez)))

	Compare(t, e, Camera(0).Transl(-0.5, -0.25, 0), "004-camtransl")
}

// Test camera rotation
func TestCamRot(t *testing.T) {
	e := NewEnv()

	r := 0.5
	nz := ShadeNormal(Ez)
	e.Add(Object(Sphere(Vec{0, 0, 0}, r), nz))
	e.Add(Object(Sphere(Vec{0, 0, 2}, r), nz))
	e.Add(Object(Sphere(Vec{0, 0, 4}, r), nz))
	e.Add(Object(Sphere(Vec{2, 0, 0}, r), nz))
	e.Add(Object(Sphere(Vec{2, 0, 2}, r), nz))
	e.Add(Object(Sphere(Vec{2, 0, 4}, r), nz))
	e.Add(Object(Sphere(Vec{-2, 0, 0}, r), nz))
	e.Add(Object(Sphere(Vec{-2, 0, 2}, r), nz))
	e.Add(Object(Sphere(Vec{-2, 0, 4}, r), nz))

	Compare(t, e, Camera(1).Transl(0, 4, -4).Transf(RotX4(pi/5)), "005-camrot")
}

// Test object transform
func TestObjTransf(t *testing.T) {
	e := NewEnv()

	r := 0.25
	sx := Object(Sphere(Vec{-0.5, 0, 2}, r), ShadeNormal(Ex))
	sy := Object(Sphere(Vec{0, 0, 2}, r), ShadeNormal(Ez))
	sz := Object(Sphere(Vec{0.5, 0, 2}, r), ShadeNormal(Ey))

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
	s1 := Object(Sphere(Vec{-r / 2, 0, 2}, r), ShadeNormal(Ez))
	s2 := Object(Sphere(Vec{r / 2, 0, 2}, r), ShadeNormal(Ey))
	s := ObjAnd(s1, s2)
	e.Add(s)

	Compare(t, e, Camera(0), "007-objand")
}

// Test two partially overlapping spheres
func TestOverlap(t *testing.T) {
	e := NewEnv()

	r := 0.5
	s1 := Object(Sphere(Vec{-r / 2, 0, 2}, r), ShadeNormal(Ez))
	s2 := Object(Sphere(Vec{r / 2, 0, 2}, r), ShadeNormal(Ey))
	e.Add(s1)
	e.Add(s2)

	Compare(t, e, Camera(0), "008-overlap")
}

// Make a cube out of 3 intersecting slabs
func TestSlabIntersect(t *testing.T) {
	e := NewEnv()

	r := 1.
	s1 := Object(Slab(Ex, -r, r), Flat(RED))
	s2 := Object(Slab(Ey, -r, r), Flat(GREEN))
	s3 := Object(Slab(Ez, -r, r), Flat(BLUE))
	cube := ObjAnd(ObjAnd(s1, s2), s3)
	cube = Transf(cube, RotY4(160*deg).Mul(RotX4(20*deg)))
	e.Add(cube)

	Compare(t, e, Camera(1).Transl(0, 0, -4), "009-slabintersect")
}

// Use sheets as green grass, blue sky and wall
func TestSheet(t *testing.T) {
	e := NewEnv()

	s1 := Object(Sheet(Ey, -1), Flat(GREEN))
	s2 := Object(Sheet(Ey, 4), Flat(BLUE))
	s3 := Object(Sheet(Ex, -10), Flat(WHITE))
	s4 := Object(Sphere(Vec{1.5, 0, 3}, 1), ShadeNormal(Ez))
	e.Add(s1, s2, s3, s4)

	Compare(t, e, Camera(1), "010-sheet")
}

// Test rectangles
func TestRect(t *testing.T) {
	e := NewEnv()

	const d = 0.5
	const z = 10
	nz := ShadeNormal(Ez)
	r1 := Object(Rect(Vec{-d, 0, z}, Ez, 0.2, 0.1, inf), nz)
	r2 := Transf(r1, RotZ4(-30*deg).Mul(Transl4(Vec{1, 0, 0})))
	r3 := ObjAnd(
		Object(Rect(Vec{0, 0, z}, Ez, 10, 10, 10), nz),
		Object(Sphere(Vec{0, 0, z}, 0.25), nz),
	)
	e.Add(r1, r2, r3)

	Compare(t, e, Camera(0), "011-rect")
}

// Test Axis Aligned Box
func TestBox(t *testing.T) {
	e := NewEnv()

	nz := ShadeNormal(Ez)
	b := Object(Box(Vec{0, 0, 0}, 2, 1, 1), nz)
	b = Transf(b, RotY4(150*deg))
	g := Object(Sheet(Ey, -1), Flat(GREEN.Mul(EV(-4))))
	e.Add(b, g)

	Compare(t, e, Camera(1).Transl(0, 0, -4), "012-box")
}

func TestDiffuse0(t *testing.T) {
	e := NewEnv()

	g := Object(Sheet(Ey, -1), Diffuse0(WHITE.Mul(EV(-1))))
	s := Object(Sphere(Vec{}, 1), Diffuse0(WHITE))
	e.Add(g, s)

	l := DirLight(Vec{1, 0.5, -4}, WHITE.Mul(EV(0)))
	e.AddLight(l)

	Compare(t, e, Camera(1).Transl(0, 0, -4), "013-diffuse0")
}

//func TestObjMinus(t *testing.T) {
//	e := NewEnv()
//
//	g := Object(Sheet(Ey, 0), Diffuse0(WHITE.Mul(EV(-1))))
//	r := 0.8
//	b := Object(Box(Vec{}, r, r, r), Diffuse0(WHITE.Mul(EV(-0))))
//	s := Object(Sphere(Vec{}, 1), Diffuse0(WHITE))
//	e.Add(g, s, b)
//
//	l := DirLight(Vec{2, 0.5, -4}, WHITE.Mul(EV(0)))
//	e.AddLight(l)
//
//	Compare(t, e, Camera(1).Transl(0, 1, -2).Transf(RotX4(10*deg)), "014-objminus")
//}
