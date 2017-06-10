package main

import (
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

var (
	quality = flag.Int("q", 80, "JPEG quality")
)

func Encode(img [][]float64, fname string, div float64) error {
	img2 := MakeImage(*width, *height)
	fac := 1 / div
	for i := range img {
		for j := range img[i] {
			v := img[i][j] * fac
			if !*overExp {
				v = clip(v, 0, 1)
			}
			img2[i][j] = v
		}
	}
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	return jpeg.Encode(f, Gray(img2), &jpeg.Options{Quality: *quality})
}

type Gray [][]float64

func (g Gray) At(i, j int) color.Color {
	c := g[j][i]
	if c < 0 {
		return color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	}
	if c > 1 {
		return color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	}
	//c = srgb(c)
	v := uint8(c * ((1 << 8) - 1))
	return color.RGBA{v, v, v, 255}
}

func (g Gray) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(g[0]), len(g))
}

func (g Gray) ColorModel() color.Model {
	return nil
}

// linear to sRGB gamma curve
// https://en.wikipedia.org/wiki/SRGB
func srgb(c float64) float64 {
	if c <= 0.0031308 {
		return 12.92 * c
	}
	c = 1.055*math.Pow(c, 1./2.4) - 0.05
	if c > 1 {
		return 1
	}
	return c
}
