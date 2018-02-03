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

func Waves() Material {
	rng := rand.New(rand.NewSource(1))
	terms := make([]term, 30)
	for i := range terms {
		r := randVec(rng)
		r[Y] = 0
		r = r.Normalized()
		terms[i].k = r.Mul(1 - 0.5*rng.Float64())
	}
	return &series{terms}
}

func (m *series) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	v := 0.0

	pos := r.At(frag.T)
	for _, t := range m.terms {
		v += sin(t.k.Dot(pos))
	}

	v /= sqrt(float64(len(m.terms)))

	v /= 2
	//v = v * v
	if v > 0 {
		return Color{v, v, v}
	} else {
		return Color{0, 0, 0}
	}
}

type series struct {
	terms []term
}

type term struct {
	k Vec
}
