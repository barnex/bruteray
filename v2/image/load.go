package image

import (
	"bufio"
	"image"
	"log"
	"os"

	"github.com/barnex/bruteray/v2/color"

	_ "image/jpeg"
	_ "image/png"
)

func MustLoad(fname string) Image {
	img, err := Load(fname)
	if err != nil {
		log.Fatalf("Error loading %q: %v", fname, err)
	}
	return img
}

func Load(fname string) (Image, error) {
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
			img[Y][x] = color.Color{
				color.SRGBToLinear(float64(r) / 0xffff),
				color.SRGBToLinear(float64(g) / 0xffff),
				color.SRGBToLinear(float64(b) / 0xffff),
			}
		}
	}
	return img, nil
}
