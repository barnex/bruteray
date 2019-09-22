package post

import (
	"math"

	"github.com/barnex/bruteray/color"
	"github.com/barnex/bruteray/image"
	"github.com/barnex/bruteray/util"
)

func AddConvolution(dst, src, kern image.Image, amplitude, offset float64) image.Image {
	w, h := src.Bounds().Dx(), src.Bounds().Dy()

	kw, kh := kern.Bounds().Dx(), kern.Bounds().Dy()
	kdx, kdy := kw/2, kh/2

	// normalize the kernel to a distribution, then multiply by amplitude
	// normalization is needed for the bloom brightness to be independent of image resolution.
	var sum float64
	for iy := range kern {
		for ix := range kern[iy] {
			c := kern[iy][ix]
			sum += (c.R + c.G + c.B) / 3
		}
	}
	k := kern.Copy()
	factor := amplitude / sum
	for iy := range k {
		for ix := range k[iy] {
			k[iy][ix] = k[iy][ix].Mul(factor)
		}
	}

	for iy := 0; iy < h; iy++ {
		for ix := 0; ix < w; ix++ {
			c := src[iy][ix]
			if c.Max() < offset {
				continue
			}
			c.R = util.Re(c.R - offset)
			c.G = util.Re(c.G - offset)
			c.B = util.Re(c.B - offset)

			for iky := 0; iky < kh; iky++ {
				idy := iy + iky - kdy
				if idy < 0 || idy >= h {
					continue
				}
				for ikx := 0; ikx < kw; ikx++ {
					idx := ix + ikx - kdx
					if idx < 0 || idx >= w {
						continue
					}
					dst[idy][idx] = dst[idy][idx].Add(c.Mul3(k[iky][ikx]))
				}
			}
		}
	}
	return dst
}

func Gaussian(pixels int, width float64) image.Image {
	return kernel(pixels, func(u, v float64) color.Color {
		r2 := (u*u + v*v) / (width * width)
		c := math.Exp(-r2)
		return color.Color{c, c, c}
	})
}

func Airy(pixels int, width float64) image.Image {
	return kernel(pixels, func(u, v float64) color.Color {

		r := math.Sqrt((u*u + v*v)) / width
		if r == 0 {
			return color.Color{1, 1, 1}
		}
		R := airy((500. / 700.) * r)
		G := airy((500. / 500.) * r)
		B := airy((500. / 450.) * r)
		return color.Color{R, G, B}
	})
}
func airy(r float64) float64 {
	return util.Sqr(2 * math.J1(r) / r)
}

func kernel(pixels int, f func(u, v float64) color.Color) image.Image {
	n := 2*pixels + 1
	center := n / 2
	k := image.MakeImage(n, n)
	for iky := 0; iky < n; iky++ {
		v := float64(iky - center)
		for ikx := 0; ikx < n; ikx++ {
			u := float64(ikx - center)
			k[iky][ikx] = f(u, v)
		}
	}
	return k
}

func starKernel(pixels int) image.Image {
	n := 2*pixels + 1
	k := image.MakeImage(n, n)
	for i := 2; i < pixels-1; i++ {

		v := util.Sqr(float64(pixels-i) / float64(pixels))

		col := color.Color{v, v, v}

		c := pixels
		k[c][c+i] = col
		k[c][c-i] = col
		k[c+i][c] = col
		k[c-i][c] = col

		// diagonal
		col = col.Mul(1 / math.Sqrt2)
		k[c+i][c+i] = col
		k[c-i][c-i] = col
		k[c+i][c-i] = col
		k[c-i][c+i] = col
	}
	return k
}
