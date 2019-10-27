package ppm_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/barnex/bruteray/image"
	"github.com/barnex/bruteray/image/ppm"
	. "github.com/barnex/bruteray/tracer/types"
)

func TestEncodeAscii16(t *testing.T) {
	f, err := os.Create("testdata/ascii16.ppm")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	ppm.EncodeAscii16(w, testImg())
}

func TestEncode48BE(t *testing.T) {
	f, err := os.Create("testdata/binary48be.ppm")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	ppm.Encode48BE(w, testImg())
}

func testImg() image.Image {
	const W, H = 30, 20
	const max = 1<<16 - 1
	img := image.MakeImage(W, H)
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			img[j][i] = Color{R: float64(i) / W, G: float64(j) / H, B: 0}
		}
	}
	return img
}
