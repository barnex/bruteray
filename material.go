package bruteray

type Material interface {
	Shade(e *Env, N int, r *Ray, frag *Fragment) Color
}

// A Flat material always returns the same color.
// Useful for debugging, or for rare cases like
// a computer screen or other extended, dimly luminous surface.
func Flat(c Color) Material {
	return &flat{c}
}

type flat struct {
	c Color
}

func (s *flat) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	return s.c
}

// A Diffuse material appears perfectly mate,
// like paper or plaster.
// See https://en.wikipedia.org/wiki/Lambertian_reflectance.
func Diffuse(c Color) Material {
	return &diffuse{diffuse0{c}}
}

type diffuse struct {
	diffuse0
}

func (s *diffuse) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	// This is the core of bi-directonial path tracing.

	pos, norm := r.At(frag.T-offset), frag.Norm

	// accumulates the result
	var acc Color

	// first sum over all explicit sources
	// (with fall-off)
	for _, l := range e.lights {
		acc = acc.Add(s.lightIntensity(e, pos, norm, l))
	}

	// add one random ray to sample (only!) the indirect sources.
	// (no fall-off, the chance of hitting an object
	// automatically falls off correctly with distance).
	// Choose the random ray via importance sampling.
	sec := NewRay(pos.MAdd(offset, norm), randVecCos(e, norm))
	acc = acc.Add(s.refl.Mul3(e.ShadeNonLum(sec, N-1))) // does not include explicit lights

	return acc
}

// Diffuse material with direct illumination only (no interreflection).
// Intended for debugging or rapid previews. Diffuse is much more realistic.
func Diffuse0(c Color) Material {
	return &diffuse0{c}
}

type diffuse0 struct {
	refl Color
}

func (s *diffuse0) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	pos, norm := r.At(frag.T-offset), frag.Norm
	var acc Color
	for _, l := range e.lights {
		acc = acc.Add(s.lightIntensity(e, pos, norm, l))
	}
	return acc
}

func (s *diffuse0) lightIntensity(e *Env, pos, norm Vec, l Light) Color {
	lpos, intens := l.Sample(e, pos)

	secundary := NewRay(pos, lpos.Sub(pos).Normalized())

	t := e.IntersectAny(secundary)

	lightT := lpos.Sub(pos).Len()
	if (t > 0) && t < lightT { // intersection between start and light position
		return Color{} // shadow
	} else {
		return s.refl.Mul(re(norm.Dot(secundary.Dir()))).Mul3(intens)
	}
}

// Diffuse material with direct illumination only and no shadows.
// Intended for the tutorial.
func Diffuse00(c Color) Material {
	return &diffuse00{c}
}

type diffuse00 struct {
	refl Color
}

func (s *diffuse00) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	pos, norm := r.At(frag.T-offset), frag.Norm
	var acc Color
	for _, l := range e.lights {
		lpos, intens := l.Sample(e, pos)
		secundary := NewRay(pos, lpos.Sub(pos).Normalized())
		i := s.refl.Mul(re(norm.Dot(secundary.Dir()))).Mul3(intens)
		acc = acc.Add(i)
	}
	return acc
}

func (s *diffuse00) lightIntensity(e *Env, pos, norm Vec, l Light) Color {
	lpos, intens := l.Sample(e, pos)

	secundary := NewRay(pos, lpos.Sub(pos).Normalized())

	t := e.IntersectAny(secundary)

	lightT := lpos.Sub(pos).Len()
	if (t > 0) && t < lightT { // intersection between start and light position
		return Color{} // shadow
	} else {
		return s.refl.Mul(re(norm.Dot(secundary.Dir()))).Mul3(intens)
	}
}

// ShadeDir returns a color based on the direction of a ray.
// Used for shading the ambient background, E.g., the sky.
type ShadeDir func(dir Vec) Color

func (s ShadeDir) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	pos := r.At(frag.T - offset)
	return s(pos)
}

// A Reflective surface. E.g.:
// 	Reflective(WHITE)        // perfectly reflective, looks like shiny metal
// 	Reflective(WHITE.EV(-1)) // 50% reflective, looks like darker metal
// 	Reflective(RED)          // Reflects only red, looks like metal in transparent red candy-wrap.
func Reflective(c Color) Material {
	return &reflective{c}
}

type reflective struct {
	c Color
}

func (s *reflective) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	pos, norm := r.At(frag.T-offset), frag.Norm
	r2 := NewRay(pos, r.Dir().Reflect(norm))
	return e.ShadeAll(r2, N-1).Mul3(s.c)
}

func Refractive(n float64, obj Insider) Material {
	return &refractive{n, obj}
}

type refractive struct {
	n float64 // index of refraction
	i Insider
}

func (s *refractive) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	return Color{}
}

// Blend mixes two materials with certain weights. E.g.:
// 	Blend(0.9, Diffuse(WHITE), 0.1, Reflective(WHITE))  // 90% mate + 10% reflective, like a shiny billiard ball.
func Blend(a float64, matA Material, b float64, matB Material) Material {
	return &blend{a, matA, b, matB}
}

// Shiny is shorthand for Blend-ing diffuse + reflection, e.g.:
// Shiny(WHITE, 0.1) // a white billiard ball, 10% specular reflection
func Shiny(c Color, reflectivity float64) Material {
	return Blend(1-reflectivity, Diffuse(c), reflectivity, Reflective(WHITE))
}

type blend struct {
	a    float64
	matA Material
	b    float64
	matB Material
}

func (s *blend) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	ca := s.matA.Shade(e, N, r, frag)
	cb := s.matB.Shade(e, N, r, frag)

	return ca.Mul(s.a).MAdd(s.b, cb)
}

// Debug shader: colors according to the normal vector projected on dir.
func ShadeNormal(dir Vec) Material {
	return &shadeNormal{dir}
}

type shadeNormal struct {
	dir Vec
}

func (s *shadeNormal) Shade(e *Env, N int, r *Ray, frag *Fragment) Color {
	v := frag.Norm.Dot(s.dir)
	if v < 0 {
		return RED.Mul(-v) // towards cam
	} else {
		return BLUE.Mul(v) // away from cam
	}
}
