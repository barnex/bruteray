package main

import (
	"image"
	"image/png"
	"os"
	"testing"
)

const (
	testW, testH = 300, 200
)

// Two flat-shaded spheres, partially overlapping.
func TestOverlap(tst *testing.T) {
	t := Helper(tst)

	const r = 0.25
	s := &Scene{
		objs: []Obj{
			Flat(Sphere(Vec{-r / 2, 0, 3}, r), 1.0),
			Flat(Sphere(Vec{r / 2, 0, 3}, r), 0.5),
		},
	}

	t.Compare(s, "001-overlap")
}

// A sphere behind the camera, should not be visible
func TestBehindCam(tst *testing.T) {
	t := Helper(tst)

	const r = 0.25
	objects := []Obj{
		Flat(Sphere(Vec{0, 0, -3}, r), 1),
	}
	s := &Scene{
		objs: objects,
	}

	t.Compare(s, "002-behindcam")
}

// Intersection of flat-shaded spheres
func TestIntersect(tst *testing.T) {
	t := Helper(tst)

	const r = 0.25
	s1 := Flat(Sphere(Vec{-r / 2, 0, 3}, r), 1)
	s2 := Flat(Sphere(Vec{r / 2, 0, 3}, r), 0.5)
	s := &Scene{
		objs: []Obj{
			&ObjAnd{s1, s2},
		},
	}

	t.Compare(s, "003-intersect")
}

// Intersection of spheres, as shapes (not objects)
func TestIntersectShape(tst *testing.T) {
	t := Helper(tst)

	const r = 0.25
	s1 := Sphere(Vec{-r / 2, 0, 3}, r)
	s2 := Sphere(Vec{r / 2, 0, 3}, r)
	sh := ShapeAnd{s1, s2}

	s := &Scene{
		objs: []Obj{
			Flat(sh, 1),
		},
	}

	t.Compare(s, "004-intersectshape")
}

// Minus of spheres, as shapes (not objects)
func TestMinusShape(tst *testing.T) {
	t := Helper(tst)

	const r = 0.5
	s1 := Sphere(Vec{-r / 2, 0, 3}, r)
	s2 := Sphere(Vec{r / 2, 0, 3}, r)
	sh := ShapeMinus{s1, s2}
	s := &Scene{
		objs: []Obj{
			Flat(sh, 1),
		},
	}

	t.Compare(s, "005-minusshape")
}

// Intersection of normal.z-shaded spheres
func TestSphereNormals(tst *testing.T) {
	t := Helper(tst)

	const r = 0.25
	s3 := ShadeNormal(Sphere(Vec{0, -0.5, 3}, 2*r))
	s1 := ShadeNormal(Sphere(Vec{-r / 2, 0, 3}, r))
	s2 := ShadeNormal(Sphere(Vec{r / 2, 0, 3}, r))
	s := &Scene{
		objs: []Obj{
			ObjAnd{s3, s1},
			ObjAnd{s3, s2},
		},
	}

	t.Compare(s, "006-spherenormals")
}

func TestBox(tst *testing.T) {
	t := Helper(tst)

	b1 := ABox(Vec{-2, -2, 4}, Vec{-1, -1, 3})
	b2 := ABox(Vec{2, -2, 4}, Vec{1, -1, 3})
	b3 := ABox(Vec{-2, 2, 4}, Vec{-1, 1, 3})
	b4 := ABox(Vec{2, 2, 4}, Vec{1, 1, 3})
	s := &Scene{}
	s.objs = []Obj{
		Diffuse1(s, b1, 1),
		Diffuse1(s, b2, 1),
		Diffuse1(s, b3, 1),
		Diffuse1(s, b4, 1),
	}
	s.sources = []Source{
		&PointSource{Vec{0.5, 0.3, -5}, 60},
	}

	t.CompareCam(s, "007-box", Camera(testW, testH, 0.7))
}

func TestReflection(tst *testing.T) {
	t := Helper(tst)

	s := &Scene{}

	const h = 2
	ground := Diffuse1(s, Slab(-h, -h-100), 0.5)
	sp := Sphere(Vec{-0.5, -1, 8}, 2)
	die := &ShapeAnd{sp, Slab(-h+.2, -.2)}
	dice := Diffuse1(s, die, 0.95)
	s.objs = []Obj{
		ground,
		dice,
		Reflective(s, Sphere(Vec{3, -1, 10}, 1), 0.9),
	}
	s.sources = []Source{
		&PointSource{Vec{6, 10, 2}, 180},
	}
	s.amb = func(Vec) Color { return 1 }

	cam := Camera(testW, testH, 1)
	cam.Transf = RotX(-10 * deg)

	t.CompareCam(s, "008-reflections", cam)
}

type helper struct {
	*testing.T
}

func Helper(tst *testing.T) helper {
	return helper{tst}
}

func (t helper) Compare(s *Scene, name string) {
	//t.Helper()
	cam := Camera(testW, testH, 0)
	t.CompareCam(s, name, cam)
}

func (t helper) CompareCam(s *Scene, name string, cam *Cam) {
	//t.Helper()
	out := name + ".png"
	Encode(cam.Render(s), out, 1/(float64(cam.N)), true)
	ref := "testdata/" + out
	deviation, err := imgComp(out, ref)

	if err != nil {
		t.Fatal(err)
	}
	if deviation > 0 {
		t.Errorf("%v: differs from reference by %v", name, deviation)
	}
}

func imgComp(a, b string) (float64, error) {
	A, err := imgRead(a)
	if err != nil {
		return 0, err
	}
	B, err := imgRead(b)
	if err != nil {
		return 0, err
	}

	delta := 0
	for y := 0; y < A.Bounds().Max.Y; y++ {
		for x := 0; x < A.Bounds().Max.X; x++ {
			r1, g1, b1, _ := A.At(x, y).RGBA()
			r2, g2, b2, _ := B.At(x, y).RGBA()
			delta += diff(r1, r2) + diff(g1, g2) + diff(b1, b2)
		}
	}
	return float64(delta) / (3 * 255), nil
}

func diff(a, b uint32) int {
	d := int(a) - int(b)
	if d < 0 {
		return -d
	}
	return d
}

func imgRead(fname string) (image.Image, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}
