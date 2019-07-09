package post

import (
	"path"
	"testing"

	"github.com/barnex/bruteray/v2/color"
	"github.com/barnex/bruteray/v2/image"
	"github.com/barnex/bruteray/v2/test"
)

func TestConvolution_Gaussian(t *testing.T) {
	w, h := 300, 200
	src := image.MakeImage(w, h)
	src[100][150] = color.Color{1, 1, 1} // below threshold
	src[150][100] = color.Color{10, 10, 5}
	src[0][0] = color.Color{10, 0, 0}
	src[0][299] = color.Color{0, 10, 0}
	src[199][299] = color.Color{0, 0, 10}
	k := Gaussian(20, 5)

	dst := image.MakeImage(w, h)
	AddConvolution(dst, src, k, 5, 1)
	fname := "post.convolution_gaussian.png"
	got := path.Join("../testdata/got", fname)
	want := path.Join("../testdata", fname)
	test.Save(t, dst, got)
	tolerance := 1e-6
	if deviation := test.DiffImg(t, got, want); deviation > tolerance {
		t.Errorf("difference between %v and %v = %v, want < %v", got, want, deviation, tolerance)
	}
}

func TestConvolution_Airy(t *testing.T) {
	w, h := 300, 200
	src := image.MakeImage(w, h)
	src[100][150] = color.Color{1, 1, 1}
	src[150][100] = color.Color{1, 1, 1}
	src[0][0] = color.Color{1, 0, 0}
	src[0][299] = color.Color{0, 1, 0}
	src[199][299] = color.Color{0, 0, 1}
	k := Airy(20, 1)

	dst := image.MakeImage(w, h)
	AddConvolution(dst, src, k, 50, 0.1)
	fname := "post.convolution_airy.png"
	got := path.Join("../testdata/got", fname)
	want := path.Join("../testdata", fname)
	test.Save(t, dst, got)
	tolerance := 1e-6
	if deviation := test.DiffImg(t, got, want); deviation > tolerance {
		t.Errorf("difference between %v and %v = %v, want < %v", got, want, deviation, tolerance)
	}
}

func TestConvolution_Star(t *testing.T) {
	w, h := 300, 200
	src := image.MakeImage(w, h)
	src[100][150] = color.Color{1, 1, 1} // below threshold
	src[150][100] = color.Color{10, 10, 5}
	src[0][0] = color.Color{10, 0, 0}
	src[0][299] = color.Color{0, 10, 0}
	src[199][299] = color.Color{0, 0, 10}
	k := starKernel(20)

	dst := image.MakeImage(w, h)
	AddConvolution(dst, src, k, 5, 1)
	fname := "post.convolution_star.png"
	got := path.Join("../testdata/got", fname)
	want := path.Join("../testdata", fname)
	test.Save(t, dst, got)
	tolerance := 1e-6
	if deviation := test.DiffImg(t, got, want); deviation > tolerance {
		t.Errorf("difference between %v and %v = %v, want < %v", got, want, deviation, tolerance)
	}
}
