package bruteray

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path"
)

// Image with float64 precission.
type Image [][]Color

func MakeImage(W, H int) Image {
	img := make(Image, H)
	for i := range img {
		img[i] = make([]Color, W)
	}
	return img
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(i[0]), len(i))
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

func (i Image) ColorModel() color.Model {
	return nil
}

const jpegQual = 95

func Encode(img Image, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	switch path.Ext(fname) {
	default:
		return fmt.Errorf("unknown format: %q", fname)
	case ".jpg", ".jpeg":
		return jpeg.Encode(f, img, &jpeg.Options{Quality: jpegQual})
	case ".png":
		return png.Encode(f, img)
	}
}
