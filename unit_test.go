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

type helper struct {
	*testing.T
}

func Helper(tst *testing.T) helper {
	tst.Parallel()
	return helper{tst}
}

func (t helper) Compare(s *Scene, name string) {
	//t.Helper()
	cam := Camera(testW, testH, 0)
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
