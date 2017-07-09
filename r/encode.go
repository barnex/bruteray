package r

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path"
)

const jpegQual = 90

func Encode(img Image, fname string, scale float64, overExp bool) error {
	//	img2 := MakeImage(len(img[0]), len(img))
	//	for i := range img {
	//		for j := range img[i] {
	//			v := img[i][j] * Color(scale)
	//			if !overExp {
	//				v = clip(v)
	//			}
	//			img2[i][j] = v
	//		}
	//	}
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

//// Normalizes colors into the [0,1] interval,
//// for maximum, non-overexposed contrast.
//func MaxContrast(img [][]Color) [][]Color {
//	min, max := Color(inf), Color(-inf)
//	for i := range img {
//		for j := range img[i] {
//			c := img[i][j]
//			if math.IsInf(float64(c), 0) {
//				continue
//			}
//			if c < min {
//				min = c
//			}
//			if c > max {
//				max = c
//			}
//		}
//	}
//
//	img2 := MakeImage(len(img[0]), len(img))
//	for i := range img {
//		for j := range img[i] {
//			c := img[i][j]
//			c = (c - min) / (max - min)
//			img2[i][j] = c
//		}
//	}
//	return img2
//}
//
//type Gray [][]Color
//
//func (g Gray) At(i, j int) color.Color {
//	c := g[j][i]
//	if c < 0 {
//		return color.NRGBA{R: 0, G: 0, B: 255, A: 255}
//	}
//	if c > 1 {
//		return color.NRGBA{R: 255, G: 0, B: 0, A: 255}
//	}
//	if *useSRGB {
//		c = srgb(c)
//	}
//	v := uint8(c * ((1 << 8) - 1))
//	return color.RGBA{v, v, v, 255}
//}
//
//func (g Gray) Bounds() image.Rectangle {
//	return image.Rect(0, 0, len(g[0]), len(g))
//}
//
//func (g Gray) ColorModel() color.Model {
//	return nil
//}

// clip color value between 0 and 1
func clip(v float64) float64 {
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	return v
}

// linear to sRGB gamma curve
// https://en.wikipedia.org/wiki/SRGB
func srgb(c float64) float64 {
	c = clip(c)
	if c <= 0.0031308 {
		return 12.92 * c
	}
	c = 1.055*math.Pow(float64(c), 1./2.4) - 0.05
	if c > 1 {
		return 1
	}
	return c
}
