package material

import (
	"math/rand"

	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/texture"
	. "github.com/barnex/bruteray/v2/tracer"
)

type mate struct {
	texture texture.Texture
}

// Mate constructs a material with Lambertian ("diffuse") reflectance.
// E.g.: plaster, paper, rubber.
// See https://en.wikipedia.org/wiki/Lambertian_reflectance.
func Mate(t texture.Texture) Material {
	return &mate{t}
}

func (m *mate) Eval(ctx *Ctx, s *Scene, r *Ray, recDepth int, h HitCoords) Color {
	var acc Color

	normal := h.Normal.Towards(r.Dir)
	p := r.At(h.T).MAdd(Tiny, normal)
	secundary := ctx.Ray()
	secundary.Start = p

	for _, l := range s.Lights {
		lpos, intens := l.Sample(ctx, p)
		if intens == (Color{}) {
			continue
		}
		lDelta := lpos.Sub(p)
		lDir := lDelta.Normalized()
		cosTheta := lDir.Dot(normal)
		if cosTheta <= 0 {
			continue
		}
		lDist := lDelta.Len()
		secundary.Dir = lDir
		secundary.Len = lDist
		if s.IsOccluded(ctx, secundary) {
			continue
		}
		acc = acc.Add(intens.Mul(cosTheta))
	}

	sec := ctx.Ray()
	sec.Start = p.MAdd(Tiny, normal)
	sec.Dir = randVecCos(ctx.Rng, normal)
	// factor 1/2 due to non-uniform sampling: 1/2 = integral_[half sphere](cos(theta) d theta d phi)
	acc = acc.MAdd(0.5, s.EvalNonLuminous(ctx, sec, recDepth)) // does not include explicit lights
	ctx.PutRay(sec)

	refl := m.texture.At(h.Local)
	return acc.Mul3(refl)
}

// Random unit vector.
func randVec(rng *rand.Rand) Vec {
	return Vec{
		rng.NormFloat64(),
		rng.NormFloat64(),
		rng.NormFloat64(),
	}.Normalized()
}

// Random unit vector from the hemisphere around n
// (dot product with n >= 0).
func randVecDir(rng *rand.Rand, n Vec) Vec {
	v := randVec(rng)
	if v.Dot(n) < 0 {
		v = v.Mul(-1)
	}
	return v
}

// Random unit vector, sampled with probability cos(angle with dir).
// Used for diffuse inter-reflection importance sampling.
func randVecCos(rng *rand.Rand, dir Vec) Vec {
	v := randVecDir(rng, dir)
	for v.Dot(dir) < rng.Float64() {
		v = randVecDir(rng, dir)
	}
	return v
}
