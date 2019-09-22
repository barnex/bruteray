package image

import (
	"image"
	"image/color"

	colorf "github.com/barnex/bruteray/color"
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
	return image.Rect(0, 0, len(i[0]), len(i))
}

// At implements image.Image
func (img Image) At(i, j int) color.Color {
	return img[j][i]
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
