package mat

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/raster"
)

type Texture3D interface {
	At(Vec) Color
}

type Image interface {
	AtUV(u, v float64) Color
}

func NewImgTex(img raster.Image, mapper UVMapper) *ImgTex {
	return &ImgTex{img, mapper}
}

type ImgTex struct {
	img   raster.Image
	uvmap UVMapper
}

// TODO: remove?
func (c *ImgTex) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	return c.At(r.At(frag.T))
}

func (c *ImgTex) At(pos Vec) Color {
	u, v := c.uvmap.Map(pos)

	// pixel mapping
	w := c.img.Bounds().Dx()
	h := c.img.Bounds().Dy()
	i := clamp(int(u*float64(w)), w)
	j := clamp(int(v*float64(h)), h)
	return c.img[j][i]
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

func MustLoad(name string) raster.Image {
	img, err := Load(name)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func Load(name string) (raster.Image, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	srgb, _, err := image.Decode(f)

	w := srgb.Bounds().Dx()
	h := srgb.Bounds().Dy()
	img := raster.MakeImage(w, h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, _ := srgb.At(x, y).RGBA()
			Y := h - 1 - y
			img[Y][x] = Color{
				linear(float64(r) / 0xffff),
				linear(float64(g) / 0xffff),
				linear(float64(b) / 0xffff),
			}
		}
	}
	return img, err
}

// sRGB to linear conversion
// https://en.wikipedia.org/wiki/SRGB
func linear(s float64) float64 {
	if s <= 0.04045 {
		return s / 12.92
	}
	const a = 0.055
	return math.Pow((s+a)/(1+a), 2.4)
}
