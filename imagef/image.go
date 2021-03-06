// Package imagef provides a floating-point image implementation.
package imagef

import (
	"image"
	"image/color"

	"github.com/barnex/bruteray/imagef/colorf"
)

type Image [][]colorf.Color

func MakeImage(w, h int) Image {
	list := make([]colorf.Color, w*h)
	img := make(Image, h)
	for i := range img {
		img[i] = list[i*w : (i+1)*w]
	}
	return img
}

// Bounds implements image.Image
func (i Image) Bounds() image.Rectangle {
	w, h := i.Size()
	return image.Rect(0, 0, w, h)
}

func (i Image) Size() (width, height int) {
	if len(i) == 0 {
		return 0, 0
	}
	return len(i[0]), len(i)
}

func (i Image) NumPixels() int {
	w, h := i.Size()
	return w * h
}

// At implements image.Image
func (img Image) At(i, j int) color.Color {
	return img[j][i]
}

func (img Image) RGBAAt(i, j int) (r, g, b, a uint32) {
	panic("yes")
	return img[j][i].RGBA()
}

// ColorModel implements image.Image
func (i Image) ColorModel() color.Model {
	return nil
}

func (i Image) Copy() Image {
	w, h := i.Bounds().Dx(), i.Bounds().Dy()
	cpy := MakeImage(w, h)
	for iy := 0; iy < h; iy++ {
		for ix := 0; ix < w; ix++ {
			cpy[iy][ix] = i[iy][ix]
		}
	}
	return cpy
}
