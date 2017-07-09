package r

import (
	"image"
	"image/color"
)

type Image [][]Color

func MakeImage(W, H int) Image {
	img := make(Image, H)
	for i := range img {
		img[i] = make([]Color, W)
	}
	return img
}

func (i Image) Size() (int, int) {
	return len(i[0]), len(i)
}

func (img Image) At(i, j int) color.Color {
	c := img[j][i]

	c.R = srgb(c.R)
	c.G = srgb(c.G)
	c.B = srgb(c.B)

	r := uint8(c.R * ((1 << 8) - 1))
	g := uint8(c.G * ((1 << 8) - 1))
	b := uint8(c.B * ((1 << 8) - 1))

	return color.RGBA{r, g, b, 255}
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(i[0]), len(i))
}

func (i Image) ColorModel() color.Model {
	return nil
}
