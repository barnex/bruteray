package imagef

import (
	"image"
	"image/color"

	colorf "github.com/barnex/bruteray/imagef/colorf"
)

type ImageGray [][]float64

// TODO: remove
func MakeImageGray(w, h int) ImageGray {
	list := make([]float64, w*h)
	img := make(ImageGray, h)
	for i := range img {
		img[i] = list[i*w : (i+1)*w]
	}
	return img
}

// Bounds implements image.Image
func (i ImageGray) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(i[0]), len(i))
}

// At implements image.Image
func (img ImageGray) At(i, j int) colorf.Color {
	c := img[j][i]
	return colorf.Color{c, c, c}
}

// ColorModel implements image.Image
func (i ImageGray) ColorModel() color.Model {
	return nil
}
