// Package test provides test helpers for BruteRay.
// 	- comparing rendered images to golden testdata
// 	- minimal implementations of a few shapes and materials
//
// Rendered images and golden testdata are saved under directories testdata/got and testdata/,
// respectively.
//
// The minimal implementations of Sphere and Flat material
// allow tests without complex dependencies between packages like objects, materials, textures.
package test

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/barnex/bruteray/sampler"
	"github.com/barnex/bruteray/tracer/cameras"
	. "github.com/barnex/bruteray/tracer/types"
)

const (
	TestdataDir      = "../testdata"
	DefaultTolerance = 1e-6
	DefaultWidth     = 300
	DefaultHeight    = 200
)

func QuadView(t *testing.T, s *Scene, c Camera, fov float64, tolerance float64) {
	t.Helper()
	testGolden(t, 1, renderQuadView(s, c, fov, 1, 1, DefaultWidth, DefaultHeight), tolerance)
}

func QuadViewN(t *testing.T, s *Scene, c Camera, fov float64, recDepth, nPass int, tolerance float64) {
	t.Helper()
	t.Parallel()
	testGolden(t, 1, renderQuadView(s, c, fov, recDepth, nPass, DefaultWidth, DefaultHeight), tolerance)
}

// OnePass renders a scene in default size and with one monte-carlo pass,
// compares to golden testdata.
func OnePass(t *testing.T, s *Scene, c Camera, tolerance float64) {
	t.Helper()
	t.Parallel()
	testGolden(t, 1, renderNPass(s, c, 1, 1, DefaultWidth, DefaultHeight), tolerance)
}

func Benchmark(b *testing.B, s *Scene, c Camera, tolerance float64) {
	b.Helper()
	recDepth := 3
	width := DefaultWidth
	height := DefaultHeight
	numPix := width * height
	b.ReportAllocs()
	//b.ReportMetric(float64(numPix), "pix/op")
	b.SetBytes(int64(numPix)) // "Bytes" means pixels
	b.ResetTimer()
	var img image.Image
	for i := 0; i < b.N; i++ {
		img = renderNPass(s, c, recDepth, 1, width, height)
	}
	b.StopTimer()
	//_ = img
	testGolden(b, 1, img, tolerance)
}

func NPass(t *testing.T, s *Scene, c Camera, recDepth, numPass int, tolerance float64) {
	t.Helper()
	t.Parallel()
	testGolden(t, 1, renderNPass(s, c, recDepth, numPass, DefaultWidth, DefaultHeight), tolerance)
}

func NPassSize(t *testing.T, s *Scene, c Camera, recDepth, numPass, width, height int, tolerance float64) {
	t.Helper()
	t.Parallel()
	testGolden(t, 1, renderNPass(s, c, recDepth, numPass, width, height), tolerance)
}

func renderNPass(s *Scene, c Camera, recDepth, numPass, width, height int) image.Image {
	return sampler.Uniform(s.ImageFunc(c, recDepth), numPass, width, height)
}

func renderQuadView(s *Scene, c Camera, fov float64, recDepth, numPass, w, h int) image.Image {
	comp := image.NewNRGBA(image.Rect(0, 0, 2*w, 2*h))

	drawAt(comp, 0, 0,
		renderNPass(s, c, recDepth, numPass, w, h),
	)
	drawAt(comp, w, h,
		renderNPass(s, cameras.NewIsometric(Y, fov), recDepth, numPass, w, h),
	)
	drawAt(comp, 0, h,
		renderNPass(s, cameras.NewIsometric(X, fov), recDepth, numPass, w, h),
	)
	drawAt(comp, w, 0,
		renderNPass(s, cameras.NewIsometric(Z, fov), recDepth, numPass, w, h),
	)

	return comp
}

func drawAt(dst draw.Image, x, y int, src image.Image) {
	w := src.Bounds().Dx()
	h := src.Bounds().Dy()
	draw.Draw(dst, image.Rect(x, y, x+w, y+h), src, image.Pt(0, 0), draw.Src)
}

func testGolden(t testing.TB, skip int, img image.Image, tolerance float64) {
	t.Helper()
	fname := Caller(skip+1) + ".png"
	got := path.Join(TestdataDir, "got", fname)
	want := path.Join(TestdataDir, fname)
	Save(t, img, got)
	if deviation := DiffImg(t, got, want); deviation > tolerance {
		t.Errorf("difference between %v and %v = %v, want < %v", got, want, deviation, tolerance)
	}
}

func DiffImg(t testing.TB, a, b string) float64 {
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

func readImg(t testing.TB, fname string) image.Image {
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

func Save(t testing.TB, img image.Image, fname string) {
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
// 	github.com/barnex/bruteray/object.TestSphere -> object.sphere
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

func checkRay(r *Ray) {
	const tol = 1 / 1024.
	if math.Abs(1-r.Dir.Len()) > tol {
		panic(fmt.Sprintf("BUG: unnormalized ray dir: %v (len=%v)", r.Dir, r.Dir.Len()))
	}
}
