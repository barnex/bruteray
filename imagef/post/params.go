// Package post implements image post-processing effects, like bloom.
package post

import "github.com/barnex/bruteray/imagef"

type Params struct {
	Gaussian BloomParams
	Airy     BloomParams
	Star     BloomParams
}

type BloomParams struct {
	Radius    float64
	Amplitude float64
	Threshold float64
}

func (p *Params) ApplyTo(img imagef.Image, pixelSize float64) imagef.Image {
	if b := p.Gaussian; b.Radius != 0 {
		img = ApplyGaussianBloom(img, pixelSize, b.Radius, b.Amplitude, b.Threshold)
	}
	if b := p.Airy; b.Radius != 0 {
		img = ApplyAiryBloom(img, pixelSize, b.Radius, b.Amplitude, b.Threshold)
	}
	if b := p.Star; b.Radius != 0 {
		img = ApplyStarBloom(img, pixelSize, b.Radius, b.Amplitude, b.Threshold)
	}
	return img
}

func ApplyGaussianBloom(img imagef.Image, pixelSize, radius, amplitude, threshold float64) imagef.Image {
	widthPix := radius / pixelSize
	numPix := int(5*widthPix) + 1
	K := Gaussian(numPix, widthPix)
	img2 := img.Copy()
	AddConvolution(img2, img, K, amplitude, threshold)
	return img2
}

func ApplyAiryBloom(img imagef.Image, pixelSize, radius, amplitude, threshold float64) imagef.Image {
	widthPix := radius / pixelSize
	numPix := int(8*widthPix) + 1
	K := Airy(numPix, widthPix)
	img2 := img.Copy()
	AddConvolution(img2, img, K, amplitude, threshold)
	return img2
}

func ApplyStarBloom(img imagef.Image, pixelSize, radius, amplitude, threshold float64) imagef.Image {
	widthPix := radius / pixelSize
	numPix := int(widthPix)
	K := starKernel(numPix)
	img2 := img.Copy()
	AddConvolution(img2, img, K, amplitude, threshold)
	return img2
}
