package bruteray_test

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"path"
	"reflect"
	"testing"
	"time"

	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/sample"
)

func CompareNPass(t *testing.T, e *Env, number int, name string, nPass int, tolerance float64) {
	img := sample.MakeImage(testW, testH)

	start := time.Now()
	sample.MultiPass(e, img, nPass)
	duration := time.Since(start)

	CompareImg(t, e, img, number, name, tolerance)
	fmt.Println("t=", duration.Round(time.Millisecond/10))
}

// Compare renders the environment with standard resolution
// and compares the output against testdata/00number-name.png.
func Compare(t *testing.T, e *Env, number int, name string, tolerance float64) {
	CompareNPass(t, e, number, name, 1, tolerance)
}

func CompareImg(t *testing.T, e *Env, img sample.Image, number int, testName string, tol float64) {
	t.Helper()

	os.Mkdir("out", 0777)

	name := fmt.Sprintf("%03d-%v", number, testName)
	name = name + ".png"
	have := "out/" + name
	want := "testdata/" + name

	sample.Encode(img, have)
	deviation, err := imgComp(have, want)

	if err != nil {
		t.Fatal(err)
	}
	if deviation > tol {
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
	NPix := (A.Bounds().Dx() + 1) * (A.Bounds().Dx() + 1)
	deviation := float64(delta) / (3 * 255 * float64(NPix))
	fmt.Printf("%-25s: err=%1.3f ", path.Base(a), deviation)
	return deviation, nil
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
