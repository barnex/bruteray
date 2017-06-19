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
	Encode(cam.Render(s), out, true)
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
