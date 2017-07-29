package bruteray

import (
	"image"
	"image/png"
	"math"
	"os"
	"reflect"
	"testing"
)

func CompareImg(t *testing.T, e *Env, img Image, name string) {
	t.Helper()

	os.Mkdir("out", 0777)

	name = name + ".png"
	have := "out/" + name
	want := "testdata/" + name

	Encode(img, have)
	deviation, err := imgComp(have, want)

	if err != nil {
		t.Fatal(err)
	}
	const tolerance = 10
	if deviation > tolerance {
		t.Errorf("%v: differs from reference by %v", name, deviation)
	}
}

func Compare(t *testing.T, e *Env, cam *Cam, name string) {
	t.Helper()

	img := MakeImage(testW, testH)
	cam.Render(e, testRec, img)
	CompareImg(t, e, img, name)
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

type helper struct {
	*testing.T
}

func Helper(tst *testing.T) helper {
	tst.Parallel()
	return helper{tst}
}

func (t helper) Eq(a, b interface{}) {
	t.Helper()

	if !reflect.DeepEqual(a, b) {
		t.Errorf("have: %v, want: %v", a, b)
	}
}

func (t helper) EqVec(have, want Vec) {
	const tol = 1e-6
	fail :=
		(math.Abs(have[X]-want[X]) > tol) ||
			(math.Abs(have[Y]-want[Y]) > tol) ||
			(math.Abs(have[Z]-want[Z]) > tol)
	if fail {
		t.Errorf("have %v, want %v", have, want)
	}
}
