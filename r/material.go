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
func ShadeNormal() Material {
	return shadeNormal{}
}

type shadeNormal struct{}

func (s shadeNormal) Shade(e *Env, N int, pos, norm Vec) Color {
	v := norm[Z]
	if v < 0 {
		return RED.Mul(-v) // towards cam
	} else {
		return BLUE.Mul(v) // away from cam
	}
}
