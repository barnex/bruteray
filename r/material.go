package r

type Material interface {
	Shade(e *Env, N int, pos Vec, norm Vec) Color
}

func Flat(c Color) Material {
	return &flat{c}
}

type flat struct {
	c Color
}

func (f *flat) Shade(e *Env, N int, pos, norm Vec) Color {
	return f.c
}
