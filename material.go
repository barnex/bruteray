package bruteray

type Material interface {
	Shade(e *Env, r *Ray, N int, pos Vec, norm Vec) Color
}

// -- flat

// Flat shader always returns the same color.
func Flat(c Color) Material {
	return &flat{c}
}

type flat struct {
	c Color
}

func (s *flat) Shade(e *Env, r *Ray, N int, pos, norm Vec) Color {
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
	refl Color
}

func (s *diffuse0) Shade(e *Env, r *Ray, N int, pos, norm Vec) Color {
	var acc Color
	for _, l := range e.lights {
		acc = acc.Add(s.lightIntensity(e, pos, norm, l))
	}
	return acc
}

const off = 1e-6

func (s *diffuse0) lightIntensity(e *Env, pos, norm Vec, l Light) Color {
	lpos, intens := l.Sample(e, pos)

	pos = pos.MAdd(off, norm)
	secundary := Ray{Start: pos, Dir: lpos.Sub(pos).Normalized()}

	t := e.IntersectAny(&secundary)

	lightT := lpos.Sub(pos).Len()
	if (t > 0) && t < lightT { // intersection between start and light position
		return Color{} // shadow
	} else {
		return s.refl.Mul(Re(norm.Dot(secundary.Dir))).Mul3(intens)
	}
}

//    -- diffuse 1

func Diffuse1(c Color) Material {
	return &diffuse1{diffuse0{c}}
}

type diffuse1 struct {
	diffuse0
}

func (s *diffuse1) Shade(e *Env, r *Ray, N int, pos, norm Vec) Color {
	var acc Color
	for _, l := range e.lights {
		acc = acc.Add(s.lightIntensity(e, pos, norm, l))
	}

	// random ray

	sec := &Ray{pos.MAdd(off, norm), RandVecCos(e, norm)}
	acc = acc.Add(s.refl.Mul3(e.Shade(sec, N-1)))

	return acc
}

// -- ambient

type ShadeDir func(Vec) Color

func (s ShadeDir) Shade(e *Env, N int, pos, norm Vec) Color {
	return s(pos)
}

// -- reflective

func Reflective(c Color) Material {
	return &reflective{c}
}

type reflective struct {
	c Color
}

func (s *reflective) Shade(e *Env, r *Ray, N int, pos, norm Vec) Color {
	r2 := &Ray{pos, r.Dir.Reflect(norm)}
	return e.Shade(r2, N-1).Mul3(s.c)
}

// -- normal

// Debug shader: colors according to the normal vector projected on dir.
func ShadeNormal(dir Vec) Material {
	return &shadeNormal{dir}
}

type shadeNormal struct {
	dir Vec
}

func (s *shadeNormal) Shade(e *Env, r *Ray, N int, pos, norm Vec) Color {
	v := norm.Dot(s.dir)
	if v < 0 {
		return RED.Mul(-v) // towards cam
	} else {
		return BLUE.Mul(v) // away from cam
	}
}
