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
	DefaultTolerance = 1e-6
	DefaultWidth     = 300
	DefaultHeight    = 200
)

func Compare(t testing.TB, tolerance float64, img image.Image) {
	t.Helper()
	fname := testName() + ".png"
	got := path.Join(testdataDir(), "got", fname)
	want := path.Join(testdataDir(), fname)
	Save(t, img, got)
	if deviation := DiffImg(t, got, want); deviation > tolerance {
		t.Errorf("difference between %v and %v = %v, want < %v", got, want, deviation, tolerance)
	}
}

// OnePass renders a scene with default size and one monte-carlo pass,
// and compares the result to testdata.
func OnePass(t *testing.T, s *Scene, c Camera, tolerance float64) {
	t.Helper()
	t.Parallel()
	Compare(t, tolerance, renderNPass(s, c, 1, DefaultWidth, DefaultHeight))
}

// NPass is like OnePass but allows to set the number of monte-carlo passes.
func NPass(t *testing.T, s *Scene, c Camera, numPass int, tolerance float64) {
	t.Helper()
	t.Parallel()
	Compare(t, tolerance, renderNPass(s, c, numPass, DefaultWidth, DefaultHeight))
}

// NPassSize is like NPass, but allows to set the image size.
// Intended to run costly tests at lower resolution.
func NPassSize(t *testing.T, s *Scene, c Camera, numPass, width, height int, tolerance float64) {
	t.Helper()
	t.Parallel()
	Compare(t, tolerance, renderNPass(s, c, numPass, width, height))
}

// QuadView renders the scene from 4 points of view
// 	- the camera
//	- the X, Y and Z direction, orthgraphically
// and compares the result to testdata.
func QuadView(t *testing.T, s *Scene, c Camera, fov float64, tolerance float64) {
	t.Helper()
	Compare(t, tolerance, renderQuadView(s, c, fov, 1, DefaultWidth, DefaultHeight))
}

// QuadViewN is like QuadView but allows to set the number of passes.
func QuadViewN(t *testing.T, s *Scene, c Camera, fov float64, nPass int, tolerance float64) {
	t.Helper()
	t.Parallel()
	Compare(t, tolerance, renderQuadView(s, c, fov, nPass, DefaultWidth, DefaultHeight))
}

func Benchmark(b *testing.B, s *Scene, c Camera, tolerance float64) {
	b.Helper()
	width := DefaultWidth
	height := DefaultHeight
	numPix := width * height
	b.ReportAllocs()
	//b.ReportMetric(float64(numPix), "pix/op")
	b.SetBytes(int64(numPix)) // "Bytes" means pixels
	b.ResetTimer()
	var img image.Image
	for i := 0; i < b.N; i++ {
		img = renderNPass(s, c, 1, width, height)
	}
	b.StopTimer()
	//_ = img
	Compare(b, tolerance, img)
}

func renderNPass(s *Scene, c Camera, numPass, width, height int) image.Image {
	antiAlias := (numPass > 1)
	return sampler.Uniform(s.ImageFunc(c), numPass, width, height, antiAlias)
}

func renderQuadView(s *Scene, c Camera, fov float64, numPass, w, h int) image.Image {
	comp := image.NewNRGBA(image.Rect(0, 0, 2*w, 2*h))

	drawAt(comp, 0, 0,
		renderNPass(s, c, numPass, w, h),
	)
	drawAt(comp, w, h,
		renderNPass(s, cameras.NewIsometric(Y, fov), numPass, w, h),
	)
	drawAt(comp, 0, h,
		renderNPass(s, cameras.NewIsometric(X, fov), numPass, w, h),
	)
	drawAt(comp, w, 0,
		renderNPass(s, cameras.NewIsometric(Z, fov), numPass, w, h),
	)

	return comp
}

func drawAt(dst draw.Image, x, y int, src image.Image) {
	w := src.Bounds().Dx()
	h := src.Bounds().Dy()
	draw.Draw(dst, image.Rect(x, y, x+w, y+h), src, image.Pt(0, 0), draw.Src)
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

// testName returns a name based on the calling Test function. E.g.:
// 	github.com/barnex/bruteray/object.TestSphere -> object.sphere
func testName() string {
	for skip := 0; ; skip++ {
		pc, _, _, ok := runtime.Caller(skip)
		if !ok {
			panic("runtime.Caller failed")
		}
		caller := path.Base(runtime.FuncForPC(pc).Name()) // without package directory
		noPkg := path.Ext(caller)[1:]                     // without package
		if !(strings.HasPrefix(noPkg, "Test") || strings.HasPrefix(noPkg, "Benchmark")) {
			continue
		}
		name := strings.Replace(caller, "_test", "", 1)
		name = strings.Replace(name, "Test", "", 1)
		name = strings.ToLower(name)
		return name
	}
}

// testdataDir walks up the filesystem starting from the working directory
// in search of an existing subdirectory called "testdata".
func testdataDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for dir := wd; dir != "/"; dir = path.Dir(dir) {
		testDir := path.Join(dir, "testdata")
		if _, err := os.Stat(testDir); err == nil {
			return testDir
		}
	}
	panic("no directory named testdata found in " + wd)
}

func checkRay(r *Ray) {
	const tol = 1 / 1024.
	if math.Abs(1-r.Dir.Len()) > tol {
		panic(fmt.Sprintf("BUG: unnormalized ray dir: %v (len=%v)", r.Dir, r.Dir.Len()))
	}
}
