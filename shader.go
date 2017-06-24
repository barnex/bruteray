package main

type Shader interface {
	Shade(e *Env, r *Ray, t float64, n Vec, N int) Color
}

type shadeFn func(e *Env, r *Ray, t float64, n Vec, N int) Color

func (f shadeFn) Shade(e *Env, r *Ray, t float64, n Vec, N int) Color {
	return f(e, r, t, n, N)
}

// Flat shading always returns the same color.
func Flat(v Color) Shader {
	return shadeFn(func(*Env, *Ray, float64, Vec, int) Color {
		return v
	})
}

// Diffuse shading without shadows.
func Diffuse0(refl float64) Shader {
	return shadeFn(func(e *Env, r *Ray, t float64, n Vec, N int) Color {
		acc := 0.0
		for _, light := range e.sources {
			lightPos, flux := light.Sample()
			d := lightPos.Sub(r.At(t))
			acc += flux * Max(n.Dot(d.Normalized())/(d.Len2()), 0)
		}
		return Color(refl * acc)
	})
}

// Diffuse shading with shadows, but no inter-reflection.
func Diffuse1(refl float64) Shader {
	return shadeFn(func(e *Env, r *Ray, t float64, n Vec, N int) Color {
		return Color(refl) * directDiffuse(e, r, t, n)
	})
}

// Diffuse shading with shadows and inter-reflection.
func Diffuse2(refl float64) Shader {
	return shadeFn(func(e *Env, r *Ray, t float64, n Vec, N int) Color {
		acc := Color(refl) * directDiffuse(e, r, t, n)
		p := r.At(t).MAdd(off, n)
		d := RandVecDir(n)
		sec := &Ray{p, d}
		I := e.Shade(sec, N-1)
		acc += I * Color(refl*Max((2)*n.Dot(d.Normalized()), 0))
		return acc
	})
}

const off = 1e-6

func directDiffuse(e *Env, r *Ray, t float64, n Vec) Color {
	p := r.At(t)
	acc := 0.
	for _, light := range e.sources {
		lightPos, flux := light.Sample()
		d := lightPos.Sub(p)
		p2 := p.MAdd(off, n) // anti-bleeding offset
		sec := Ray{p2, d.Normalized()}
		t, _ := e.Hit(&sec) // TODO: could hit any except this
		own := d.Len()
		//own := inf
		if !(t > 0 && t < own) {
			acc += flux * Max(n.Dot(d.Normalized())/(d.Len2()), 0)
		}
	}
	return Color(acc)
}

// Reflective shading.
func Reflective(refl float64) Shader {
	return shadeFn(func(e *Env, r *Ray, t float64, n Vec, N int) Color {
		p := r.At(t)
		dir2 := reflectVec(r.Dir, n)
		sec := &Ray{p.MAdd(off, n), dir2}
		I := e.Shade(sec, N-1) // TODO: pass N
		return Color(refl) * I
	})
}

// reflects v of the surface with normal n.
func reflectVec(v, n Vec) Vec {
	return v.MAdd(-2*v.Dot(n), n)
}

func CheckBoard(a, b Shader) Shader {
	return shadeFn(func(e *Env, r *Ray, t float64, n Vec, N int) Color {
		p := r.At(t)
		x := int(p.X + (1 << 10))
		z := int(p.Z + (1 << 10))
		if mod(x+z, 2) == 0 {
			return a.Shade(e, r, t, n, N)
		} else {
			return b.Shade(e, r, t, n, N)
		}
	})
}

func mod(x, y int) int {
	m := x % y
	if x < 0 {
		m = (m + y) % y
	}
	return m
}

// Shade with Z component of normal vector (for debugging)
func ShadeNormal() Shader {
	return shadeFn(func(e *Env, r *Ray, t float64, n Vec, N int) Color {
		return Color(-n.Z)
	})
}
