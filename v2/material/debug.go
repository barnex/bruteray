package material

import (
	"math"

	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
)

type shadeNormal struct {
	dir Vec
}

// Normal is a debug shader. Reveals the normal vector with respect to the camera position.
func Normal() Material {
	return shadeNormal{}
}

func (m shadeNormal) Eval(ctx *Ctx, s *Scene, r *Ray, recDepth int, h HitCoords) Color {
	toCam := s.Camera.Origin().Sub(r.At(h.T)).Normalized()
	v := h.Normal.Dot(toCam)
	if v > 0 {
		return White.Mul(v) // towards cam
	} else {
		return Red.Mul(-v) // away from cam
	}
}

// Normal2 is a debug shader. Reveals the normal vector with respect to the camera position.
func Normal2() Material {
	return shadeNormal2{}
}

type shadeNormal2 struct {
	dir Vec
}

func (m shadeNormal2) Eval(ctx *Ctx, s *Scene, r *Ray, recDepth int, h HitCoords) Color {
	toCam := s.Camera.Origin().Sub(r.At(h.T)).Normalized()
	v := h.Normal.Dot(toCam)
	v = math.Abs(v)
	return White.Mul(v) // towards cam
}

type grid struct {
	a Material
}

func Grid() Material {
	return &grid{
		a: Normal(),
	}
}

func (m *grid) Eval(ctx *Ctx, s *Scene, r *Ray, recDepth int, h HitCoords) Color {
	c := m.a.Eval(ctx, s, r, recDepth, h)
	p := r.At(h.T)
	for _, p := range p {
		if frac(math.Abs(p)) < 0.1 {
			c = c.Mul(0.8)
		}
	}
	return c
}

func frac(x float64) float64 {
	_, frac := math.Modf(x)
	return frac
}
