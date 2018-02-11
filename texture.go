package bruteray

import (
	"math"
	"math/rand"
)

func Checkboard(stride float64, a, b Material) Material {
	return &checkboard{1 / stride, a, b}
}

type checkboard struct {
	invs float64
	a, b Material
}

func (c *checkboard) Shade(e *Env, N int, r *Ray, frag Fragment) Color {
	pos := r.At(frag.T)
	x := int(pos[X]*c.invs + 10000)
	y := int(pos[Z]*c.invs + 10000)
	i := (x + y) & 0x1
	if i == 0 {
		return c.a.Shade(e, N, r, frag)
	} else {
		return c.b.Shade(e, N, r, frag)
	}
}

func Bricks(stride, width float64, a, b Material) Material {
	return &bricks{1 / stride, width, a, b}
}

type bricks struct {
	invs  float64
	width float64
	a, b  Material
}

func (c *bricks) Shade(e *Env, N int, r *Ray, frag Fragment) Color {
	pos := r.At(frag.T)

	x, y, z := pos[X]*c.invs, pos[Y]*c.invs, pos[Z]*c.invs

	rx := x - math.Floor(x)
	ry := y - math.Floor(y)
	rz := z - math.Floor(z)

	ix := int(math.Floor(x)) + 100000
	iy := int(math.Floor(y)) + 100000
	iz := int(math.Floor(z)) + 100000

	w := c.width

	hit := ry < w

	hit = hit || (rx < w) && (((ix+iy)%2) == 1)
	hit = hit || (rz < w) && (((iz+iy)%2) == 1)

	//(math.Floor(y) < w) ||
	//(math.Ceil(y) > 1-w) ||
	//(math.Floor(z) < w) ||
	//(math.Ceil(z) > 1-w)
	if hit {
		return c.b.Shade(e, N, r, frag)
	} else {
		return c.a.Shade(e, N, r, frag)
	}
}

func Distort(seed int, n int, K Vec, ampli float64, orig Material) Material {
	m := &distort{
		orig:  orig,
		ampli: ampli,
	}
	for i := range m.f {
		m.f[i] = makeWaveSeries(seed*i*100, n, K)
	}
	return m
}

type distort struct {
	orig  Material
	ampli float64
	f     [3]series
}

func (m distort) Shade(e *Env, N int, r *Ray, frag Fragment) Color {
	pos := r.At(frag.T)
	var delta Vec
	for i := range delta {
		delta[i] = m.f[i].Eval(pos)
	}
	frag.Norm = frag.Norm.MAdd(m.ampli, delta).Normalized()
	return m.orig.Shade(e, N, r, frag)
}

func Waves(seed int, n int, K Vec, col func(float64) Material) Material {
	return &waves{makeWaveSeries(seed, n, K), col}
}

type waves struct {
	series
	col func(float64) Material
}

func (m *waves) Shade(e *Env, N int, r *Ray, frag Fragment) Color {
	v := m.Eval(r.At(frag.T))
	return m.col(v).Shade(e, N, r, frag)
}

// series is a sum of a number of sines.
type series struct {
	terms []term
}

type term struct {
	k Vec
}

func makeWaveSeries(seed int, n int, K Vec) series {
	rng := rand.New(rand.NewSource(int64(seed)))
	terms := make([]term, n)
	for i := range terms {
		r := randVec(rng)
		r = r.Mul3(K)
		terms[i].k = r.Mul(1 - 0.5*rng.Float64())
	}
	return series{terms}
}

func (s *series) Eval(pos Vec) float64 {
	v := 0.0
	for _, t := range s.terms {
		v += sin(t.k.Dot(pos))
	}
	v /= sqrt(float64(len(s.terms)))
	return v
}
