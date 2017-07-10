package r

type Material interface {
	Shade(e *Env, N int, pos Vec, norm Vec) Color
}

// -- flat

func Flat(c Color) Material {
	return &flat{c}
}

type flat struct {
	c Color
}

func (s *flat) Shade(e *Env, N int, pos, norm Vec) Color {
	return s.c
}

// -- normal

// debug shader
func ShadeNormal(c Color) Material {
	return shadeNormal{c}
}

type shadeNormal struct{ c Color }

func (s shadeNormal) Shade(e *Env, N int, pos, norm Vec) Color {
	v := Max(0, -norm[Z])
	return s.c.Mul(v)
}
