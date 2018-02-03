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

func Waves(n int, K Vec, col func(float64) Material) Material {
	rng := rand.New(rand.NewSource(int64(1)))
	terms := make([]term, n)
	for i := range terms {
		r := randVec(rng)
		r = r.Mul3(K)
		terms[i].k = r.Mul(1 - 0.5*rng.Float64())
	}
	return &series{terms, col}
}

func (m *series) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	v := 0.0

	pos := r.At(frag.T)
	for _, t := range m.terms {
		v += sin(t.k.Dot(pos))
	}

	v /= sqrt(float64(len(m.terms)))

	return m.col(v).Shade(e, N, r, frag)
}

type series struct {
	terms []term
	col   func(float64) Material
}

type term struct {
	k Vec
}
