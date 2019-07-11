package raster

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path"

	"github.com/barnex/bruteray/v1/br"
)

// Image with float64 precission.
type Image [][]br.Color

func MakeImage(W, H int) Image {
	img := make(Image, H)
	for i := range img {
		img[i] = make([]br.Color, W)
	}
	return img
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(i[0]), len(i))
}

func (i Image) Aspect() float64 {
	w := float64(i.Bounds().Dx())
	h := float64(i.Bounds().Dy())
	return h / w
}

func (img Image) At(i, j int) color.Color {
	return img[j][i]
}

func (img Image) AtUV(u, v float64) br.Color {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	i := clamp(int(u*float64(w)), w)
	j := clamp(int(v*float64(h)), h)
	return img[j][i]
}

func clamp(v, max int) int {
	if v < 0 {
		return 0
	}
	if v >= max {
		return max - 1
	}
	return v
}

func (i Image) ColorModel() color.Model {
	return nil
}

// Adds img2 to img, overwriting img.
func (img Image) Add(img2 Image) {
	for i := range img {
		for j := range img[i] {
			img[i][j] = img[i][j].Add(img2[i][j])
		}
	}
}

// Mul multiplies colors in the image by x, in place (!).
// Returns img for easy chaining.
func (img Image) Mul(x float64) Image {
	for i := range img {
		for j := range img[i] {
			img[i][j] = img[i][j].Mul(x)
		}
	}
	return img
}

const jpegQual = 95

func MustEncode(img Image, fname string) {
	if err := Encode(img, fname); err != nil {
		panic(err)
	}
}

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
