package bruteray

func Checkboard(stride float64, a, b Material) Material {
	return &checkboard{1 / stride, a, b}
}

type checkboard struct {
	invs float64
	a, b Material
}

func (c *checkboard) Shade(e *Env, r *Ray, N int, pos, norm Vec) Color {
	x := int(pos[X]*c.invs + 10000)
	y := int(pos[Z]*c.invs + 10000)
	i := (x + y) & 0x1
	if i == 0 {
		return c.a.Shade(e, r, N, pos, norm)
	} else {
		return c.b.Shade(e, r, N, pos, norm)
	}
}
