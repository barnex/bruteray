package bruteray

import "math/rand"

func Checkboard(stride float64, a, b Material) Material {
	return &checkboard{1 / stride, a, b}
}

type checkboard struct {
	invs float64
	a, b Material
}

func (c *checkboard) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
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

func Distort(seed int, n int, K Vec, ampli float64, m Material) Material {
	return &distort{
	//f: [
	}
}

type distort struct {
	orig Material
	f    [3]series
}

func (m distort) Shade() Color {

}

func Waves(seed int, n int, K Vec, col func(float64) Material) Material {
	return &waves{makeWaveSeries(seed, n, K), col}
}

type waves struct {
	series
	col func(float64) Material
}

func makeWaveSeries(seed int, n int, K Vec) {
	rng := rand.New(rand.NewSource(int64(seed)))
	terms := make([]term, n)
	for i := range terms {
		r := randVec(rng)
		r = r.Mul3(K)
		terms[i].k = r.Mul(1 - 0.5*rng.Float64())
	}
	return terms
}

func (m *series) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	v := m.Eval(r.At(frag.T))
	return m.col(v).Shade(e, N, r, frag)
}

// series is a sum of a number of sines.
type series struct {
	terms []term
}

func (s *series) Eval(pos Vec) float64 {
	v := 0.0
	for _, t := range s.terms {
		v += sin(t.k.Dot(pos))
	}
	v /= sqrt(float64(len(s.terms)))
	return v
}

type term struct {
	k Vec
}
