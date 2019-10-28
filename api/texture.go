package api

import "github.com/barnex/bruteray/texture"

type Texture struct {
	texture.Texture
}

func LoadTexture(file string) Texture {
	return Texture{texture.MustLoad(file)}
}

func LoadHeightMap(file string) Texture {
	return Texture{texture.HeightMap(file)}
}

func (t Texture) Scale(scaleU, scaleV float64) Texture {
	return Texture{texture.ScaleUV(t.Texture, scaleU, scaleV)}
}

func (t Texture) HeightMap() func(u, v float64) float64 {
	return func(u, v float64) float64 {
		return t.Texture.At(Vec{u, v, 0}).R
	}
}
