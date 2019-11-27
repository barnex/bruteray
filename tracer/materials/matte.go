package materials

import (
	"math/rand"

	"github.com/barnex/bruteray/texture"
	"github.com/barnex/bruteray/tracer/sequence"
	. "github.com/barnex/bruteray/tracer/types"
)

type matte struct {
	texture texture.Texture
}

// Matte constructs a material with Lambertian ("diffuse") reflectance.
// E.g.: plaster, paper, rubber.
//
// The texture determines the reflectivity in each point ("diffuse map").
// It should be between 0 and 1 (in each color channel), for physicaly possible materials.
//
// See https://en.wikipedia.org/wiki/Lambertian_reflectance.
//
// The ray tracing algorithm implemented here is a flavour of bidirectional path tracing:
// A ray is shot forward from the camera onto the scene.
// When it hits a matte surface, we gather the light from all light sources
// to give the direct illumination. To that we add the (appropriately weighted)
// contribution of one random ray. This gives the indirect illumination.
// The random ray's color is determined recurively, thus again
// taking into account all light sources, etc. (up to a maximum depth).
//
// E.g. in the sketch below, Ray a goes from the camera to a matte surface.
// At the intersection point we take into account the intensity of the light
// (properly checking for shadows, of course). To this we add the brightness seen by
// a properly chosen random ray b. This brightness again contains a contribution
// of the light source at the point where ray b intersects a matter surface,
// and so on.
//
//         light
//           | \
//           |  \  #
//           |   v #
//     cam   |   / #
//        \  |  /b #
//        a\ v /   #
//   ###############
//
// This separation of direct and indirect illumination causes significantly
// faster convergence for the common case of relatively small light sources.
func Matte(t texture.Texture) Material {
	return &matte{t}
}

// Eval implements tracer.Material.
func (m *matte) Shade(ctx *Ctx, s *Scene, r *Ray, h HitCoords) Color {
	var acc Color

	normal := flipTowards(h.Normal, r.Dir)
	sec := ctx.Ray()
	p := r.At(h.T).MAdd(Tiny, normal)
	sec.Start = p

	for _, l := range s.Lights() {

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
		sec.Dir = lDir

		//if s.IsOccluded(sec, lDist) {
		//	continue
		//}

		intens = s.Occlude(sec, lDist, intens)

		acc = acc.Add(intens.Mul(cosTheta))
	}

	sec.Start = p.MAdd(Tiny, normal)
	//sec.Dir = randVecCos(ctx.Rng, normal)
	u, v := ctx.Generate2()
	sec.Dir = sequence.CosineSphere(u, v, normal)
	acc = acc.Add(s.LightFieldIndirect(ctx, sec)) // does not include explicit lights
	ctx.PutRay(sec)

	refl := m.texture.At(h.Local)
	return acc.Mul3(refl)
}

// flipTowards flips normal vector n to point towards (i.e., against) direction d,
// if not already so. Used to ensure that normal vectors point towards the ray direction.
func flipTowards(n, d Vec) Vec {
	if n.Dot(d) > 0 {
		return n.Mul(-1)
	}
	return n
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
