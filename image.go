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
	return img[j][i]
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

//func (img Image) Mul(x float64) {
//	for i := range img {
//		for j := range img[i] {
//			img[i][j].Mul(x)
//		}
//	}
//}

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
