package texture

/*
	Interpolation turns an in-memory image (integer indices)
	into a texture (float indices)
*/

import (
	"github.com/barnex/bruteray/v2/color"
	"github.com/barnex/bruteray/v2/image"
	"github.com/barnex/bruteray/v2/util"
)

func Nearest(img image.Image) Texture2D {
	return &nearest{img}
}

type nearest struct {
	img image.Image
}

func (n *nearest) AtUV(u, v float64) color.Color {
	u = util.Frac(u)
	v = util.Frac(v)
	img := n.img
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	i := int(u*float64(w-1) + 0.5)
	j := int(v*float64(h-1) + 0.5)
	return img[j][i]
}
