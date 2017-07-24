package bruteray

type Material interface {
	Shade(e *Env, N int, pos Vec, norm Vec) Color
}

// -- flat

// Flat shader always returns the same color.
func Flat(c Color) Material {
	return &flat{c}
}

type flat struct {
	c Color
}

func (s *flat) Shade(e *Env, N int, pos, norm Vec) Color {
	return s.c
}

// -- diffuse

//    -- diffuse 0

// Diffuse material with direct illumination only (no interreflection).
// Intended for rapid previews.
func Diffuse0(c Color) Material {
	return &diffuse0{c}
}

type diffuse0 struct {
	c Color
}

func (s *diffuse0) Shade(e *Env, N int, pos, norm Vec) Color {
	var acc Color
	for _, l := range e.lights {
		acc = acc.Add(s.shade(e, pos, norm, l))
	}
	return acc
}

const off = 1e-6

func (s *diffuse0) shade(e *Env, pos, norm Vec, l Light) Color {
	lpos, intens := l.Sample(pos)

	pos = pos.MAdd(off, norm)
	secundary := Ray{Start: pos, Dir: lpos.Sub(pos).Normalized()}

	t := e.IntersectAny(&secundary)

	lightT := lpos.Sub(pos).Len()
	if t > 0 && t < lightT { // intersection between start and light position
		return Color{} // shadow
	} else {
		return s.c.Mul(Re(norm.Dot(secundary.Dir))).Mul3(intens)
	}
}

//    -- diffuse 1

func Diffuse1(c Color) Material {
	return &diffuse1{diffuse0{c}}
}

type diffuse1 struct {
	diffuse0
}

func (s *diffuse1) Shade(e *Env, N int, pos, norm Vec) Color {
	var acc Color
	for _, l := range e.lights {
		acc = acc.Add(s.shade(e, pos, norm, l))
	}

	// random ray

	return acc
}

// -- normal

// Debug shader: colors according to the normal vector projected on dir.
func ShadeNormal(dir Vec) Material {
	return &shadeNormal{dir}
}

type shadeNormal struct {
	dir Vec
}

func (s *shadeNormal) Shade(e *Env, N int, pos, norm Vec) Color {
	v := norm.Dot(s.dir)
	if v < 0 {
		return RED.Mul(-v) // towards cam
	} else {
		return BLUE.Mul(v) // away from cam
	}
}
