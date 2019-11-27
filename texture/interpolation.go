package texture

/*
	Interpolation turns an in-memory image (integer indices)
	into a texture (float indices)
*/

import (
	"math"

	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/imagef"
	"github.com/barnex/bruteray/imagef/colorf"
	"github.com/barnex/bruteray/util"
)

func Nearest(img imagef.Image) Texture {
	return &nearest{img}
}

type nearest struct{ img imagef.Image }

func (n *nearest) At(p geom.Vec) colorf.Color {
	return n.AtUV(p[0], p[1])
}

func (n *nearest) AtUV(u, v float64) colorf.Color {
	img := n.img
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	i := int(u*float64(w-1) + 0.5)
	j := int(v*float64(h-1) + 0.5)
	return atIndex(img, i, j)
}

type bilinear struct{ img imagef.Image }

func (n *bilinear) At(p geom.Vec) colorf.Color {
	return n.AtUV(p[0], p[1])
}

// https://en.wikipedia.org/wiki/Bilinear_interpolation
func (f *bilinear) AtUV(u, v float64) colorf.Color {
	img := f.img
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	if u != u || v != v { //NaN
		return colorf.Color{}
	}

	X := warp(u) * float64(w-1)
	Y := warp(v) * float64(h-1)
	x0 := int(X)
	y0 := int(Y)
	x1 := x0 + 1
	if x1 >= w {
		x1 = 0
	}
	y1 := y0 + 1
	if y1 >= h {
		y1 = 0
	}
	x := util.Frac(X)
	y := util.Frac(Y)

	c00 := atIndex(img, x0, y0)
	c01 := atIndex(img, x0, y1)
	c10 := atIndex(img, x1, y0)
	c11 := atIndex(img, x1, y1)

	return colorf.Color{
		R: bilin(c00.R, c01.R, c10.R, c11.R, x, y),
		G: bilin(c00.G, c01.G, c10.G, c11.G, x, y),
		B: bilin(c00.B, c01.B, c10.B, c11.B, x, y),
	}
}

func warp(x float64) float64 {
	_, x = math.Modf(x)
	if x < 0 {
		x = 1 + x
	}
	if x == 1 {
		x = 0
	}
	return x
}

func bilin(f00, f01, f10, f11, x, y float64) float64 {
	return f00*(1-x)*(1-y) + f10*x*(1-y) + f01*(1-x)*y + f11*x*y
}

func atIndex(img imagef.Image, i, j int) colorf.Color {
	return img[j][i]
}
