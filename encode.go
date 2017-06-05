package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func Encode(img [][]float64, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, Gray(img))
}

type Gray [][]float64

func (g Gray) At(i, j int) color.Color {
	return color.Gray16{uint16(g[j][i] * ((1 << 16) - 1))}
}

func (g Gray) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(g[0]), len(g))
}

func (g Gray) ColorModel() color.Model {
	return nil
}
