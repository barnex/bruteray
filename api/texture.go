package api

import "github.com/barnex/bruteray/texture"

type Texture struct {
	texture.Texture
}

func LoadTexture(file string) Texture {
	return Texture{texture.MustLoad(file)}
}

func (t Texture) Scale(scaleU, scaleV float64) Texture {
	return Texture{texture.ScaleUV(t.Texture, scaleU, scaleV)}
}
