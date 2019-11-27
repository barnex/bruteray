package imagef

import (
	"bufio"
	"image"
	"log"
	"os"

	"github.com/barnex/bruteray/imagef/colorf"

	_ "image/jpeg"
	_ "image/png"
)

func MustLoad(fname string) Image {
	img, err := Load(fname, colorf.SRGBToLinear)
	if err != nil {
		log.Fatalf("Error loading %q: %v", fname, err)
	}
	return img
}

func Load(fname string, colorspace func(float64) float64) (Image, error) {
	if colorspace == nil {
		colorspace = colorf.SRGBToLinear
	}
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	srgb, _, err := image.Decode(bufio.NewReader(f))
	if err != nil {
		return nil, err
	}

	w := srgb.Bounds().Dx()
	h := srgb.Bounds().Dy()
	img := MakeImage(w, h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, _ := srgb.At(x, y).RGBA()
			Y := h - 1 - y
			img[Y][x] = colorf.Color{
				colorspace(float64(r) / 0xffff),
				colorspace(float64(g) / 0xffff),
				colorspace(float64(b) / 0xffff),
			}
		}
	}
	return img, nil
}

func Linear(s float64) float64 { return s }
