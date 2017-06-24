package main

type Shader interface {
	Shade(e *Env, r *Ray, t float64, n Vec, N int) Color
}

type shadeFn func(e *Env, r *Ray, t float64, n Vec, N int) Color

func (f shadeFn) Shade(e *Env, r *Ray, t float64, n Vec, N int) Color {
	return f(e, r, t, n, N)
}

func Flat(v Color) Shader {
	return shadeFn(func(*Env, *Ray, float64, Vec, int) Color {
		return v
	})
}

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

func Diffuse1(refl float64) Shader {
	return shadeFn(func(e *Env, r *Ray, t float64, n Vec, N int) Color {
		return Color(refl) * directDiffuse(e, r, t, n)
	})
}

const off = 1e-3

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

func ShadeNormal() Shader {
	return shadeFn(func(e *Env, r *Ray, t float64, n Vec, N int) Color {
		return Color(-n.Z)
	})
}

func Diffuse2(refl float64) Shader {
	return shadeFn(func(e *Env, r *Ray, t float64, n Vec, N int) Color {
		acc := Color(refl) * directDiffuse(e, r, t, n)
		p := r.At(t).MAdd(off, n)
		d := RandVecDir(n)
		sec := &Ray{p, d}
		I := e.Shade(sec, N-1)
		acc += I * Color(refl*Max(n.Dot(d.Normalized()), 0))
		return acc
	})
}

//// Diffuse shading with shadows and interreflection
//func (s *diffuse2) Intensity(r Ray, t float64, N int) Color {
//	n := Normal(s.shape, r, t)
//	acc := Color(s.refl) * directDiffuse(s.scene, r, t, n)
//	p := r.At(t).MAdd(1e-6, n)
//	d := RandVecDir(n)
//	sec := Ray{p, d}
//	_, I := s.scene.Intensity(sec, N-1)
//	acc += I * Color(s.refl*Max(n.Dot(d.Normalized()), 0))
//	return acc
//}
//
//func Reflective(sc *Env, sh Shape, refl float64) Obj {
//	return &reflective{sc, sh, refl}
//}
//
//func (s *reflective) Intersect(r Ray) (Inter, Shader) {
//	return s.shape.Intersect(r), s
//}
//
//// Diffuse shading with shadows and interreflection
//func (s *reflective) Intensity(r Ray, t float64, N int) Color {
//}
//
////func Reflective(reflect float64) Shader {
////	return func(r Ray, t float64, n Vec, N int) float64 {
////		p := r.At(t).MAdd(off, n)
////		dir2 := reflectVec(r.Dir, n)
////		return reflect * Intensity(Ray{p, dir2}, N+1, true)
////	}
////}
////
////func ReflectiveMate(reflect float64, jitter float64) Shader {
////	return func(r Ray, t float64, n Vec, N int) float64 {
////		p := r.At(t).MAdd(off, n)
////		dir2 := reflectVec(r.Dir, n).MAdd(jitter, randVec(n))
////		return reflect * Intensity(Ray{p, dir2}, N+1, true)
////	}
////}
////
////func ShaderAdd(a, b Shader) Shader {
////	return func(r Ray, t float64, n Vec, N int) float64 {
////		return a.Intensity(r, t, n, N) + b.Intensity(r, t, n, N)
////	}
////}
////
