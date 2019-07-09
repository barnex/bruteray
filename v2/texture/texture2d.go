package texture

import (
	. "github.com/barnex/bruteray/v2/color"
)

type Texture2D interface {
	AtUV(u, v float64) Color
}

// TODO: reuse package geom transforms!
func Pan(t Texture2D, deltaU, deltaV float64) Texture2D {
	return &panned{t, 1, 1, deltaU, deltaV}
}

func ScaleUV(t Texture2D, scaleU, scaleV float64) Texture2D {
	return &panned{t, 1 / scaleU, 1 / scaleV, 0, 0}
}

type panned struct {
	orig             Texture2D
	iScaleU, iScaleV float64
	deltaU, deltaV   float64
}

func (p *panned) AtUV(u, v float64) Color {
	return p.orig.AtUV(u*p.iScaleU-p.deltaU, v*p.iScaleV-p.deltaV)
}
