// Package test compares results to golden images.
package test

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/barnex/bruteray/v2/sampler"
	"github.com/barnex/bruteray/v2/tracer"
	//. "github.com/barnex/bruteray/v2/geom"
)

const (
	DefaultTolerance = 1e-6
	DefaultWidth     = 300
	DefaultHeight    = 200
)

// OnePass renders a scene in default size and with one monte-carlo pass,
// compares to golden testdata.
func OnePass(t *testing.T, s *tracer.Scene, tolerance float64) {
	t.Helper()
	testGolden(t, 1, renderNPass(s, 1, 1), tolerance)
}

func NPass(t *testing.T, s *tracer.Scene, recDepth, numPass int, tolerance float64) {
	t.Helper()
	testGolden(t, 1, renderNPass(s, recDepth, numPass), tolerance)
}

// Render renders to testdata/got without comparing to golden data.
//func Render(t *testing.T, s *tracer.Scene, recDepth, nPass, w, h int) {
//	t.Helper()
//	fname := path.Join("../testdata/got", Caller(1)+".jpg")
//	img := sampler.Uniform(s.ImageFunc(recDepth), nPass, w, h)
//	Save(t, img, fname)
//}

func renderNPass(s *tracer.Scene, recDepth, numPass int) image.Image {
	return sampler.Uniform(s.ImageFunc(recDepth), numPass, DefaultWidth, DefaultHeight)
}

func testGolden(t *testing.T, skip int, img image.Image, tolerance float64) {
	t.Helper()

	fname := Caller(skip+1) + ".png"
	got := path.Join("../testdata/got", fname)
	want := path.Join("../testdata", fname)
	Save(t, img, got)
	if deviation := DiffImg(t, got, want); deviation > tolerance {
		t.Errorf("difference between %v and %v = %v, want < %v", got, want, deviation, tolerance)
	}
}

func DiffImg(t *testing.T, a, b string) float64 {
	t.Helper()
	A := readImg(t, a)
	B := readImg(t, b)

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
	return deviation
}

func diff(a, b uint32) int {
	d := int(a) - int(b)
	if d < 0 {
		return -d
	}
	return d
}

func readImg(t *testing.T, fname string) image.Image {
	t.Helper()
	f, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(bufio.NewReader(f))
	if err != nil {
		t.Fatal(err)
	}
	return img
}

func Save(t *testing.T, img image.Image, fname string) {
	t.Helper()
	if err := save(img, fname); err != nil {
		t.Fatal(err)
	}
}

func save(img image.Image, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	defer b.Flush()

	switch path.Ext(fname) {
	case ".png":
		err = png.Encode(b, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(b, img, &jpeg.Options{Quality: 90})
	default:
		err = fmt.Errorf("save %q: unknown image format extension", fname)
	}
	return err
}

// Caller returns a name based on the calling function, skipping `skip` call frames. E.g.:
// 	github.com/barnex/bruteray/v2/object.TestSphere -> object.sphere
func Caller(skip int) string {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		panic("runtime.Caller failed")
	}
	name := strings.ToLower(runtime.FuncForPC(pc).Name())
	name = path.Base(name)
	name = strings.Replace(name, "_test", "", 1)
	name = strings.Replace(name, "test", "", 1)
	return name
}
