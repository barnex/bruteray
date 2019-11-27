package texture

import (
	"log"

	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/imagef"
	. "github.com/barnex/bruteray/imagef/colorf"
)

func MustLoad(file string) Texture {
	img := imagef.MustLoad(file)
	return &bilinear{img}
}

func HeightMap(file string) Texture {
	img, err := imagef.Load(file, imagef.Linear)
	if err != nil {
		log.Fatalln("heightmap: read", file, ":", err)
	}
	return &bilinear{img}
}

// TODO: reuse package geom transforms!
func Pan(t Texture, deltaU, deltaV float64) Texture {
	return &panned{t, 1, 1, deltaU, deltaV}
}

func ScaleUV(t Texture, scaleU, scaleV float64) Texture {
	return &panned{t, 1 / scaleU, 1 / scaleV, 0, 0}
}

type panned struct {
	orig             Texture
	iScaleU, iScaleV float64
	deltaU, deltaV   float64
}

func (p *panned) AtUV(u, v float64) Color {
	return p.orig.At(geom.Vec{u*p.iScaleU - p.deltaU, v*p.iScaleV - p.deltaV, 0})
}

func (p *panned) At(P geom.Vec) Color {
	return p.AtUV(P[0], P[1])
}
