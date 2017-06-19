package main

// Shader provides a lazily evaluated color.
// It may be expensive and is only evaluated if really needed.
type Shader interface {
	Intensity(Ray, float64) Color
}

type flat struct {
	s Shape
	c Color
}

func Flat(s Shape, c Color) Obj {
	return &flat{s, c}
}

func (s *flat) Intensity(ray Ray, t float64) Color {
	return s.c
}

func (s *flat) Intersect(r Ray) (Inter, Shader) {
	return s.s.Intersect(r), s
}

type shadeNormal struct{ s Shape }

func ShadeNormal(s Shape) Obj {
	return shadeNormal{s}
}

func (s shadeNormal) Intersect(r Ray) (Inter, Shader) {
	return s.s.Intersect(r), s
}

func (s shadeNormal) Intensity(r Ray, t float64) Color {
	n := Normal(s.s, r, t)
	return Color(-n.Z)
}

type diffuse1 struct {
	scene *Scene
	shape Shape
	refl  float64
}

// Diffuse shading with shadows, but no interreflection
func Diffuse1(sc *Scene, sh Shape, refl float64) Obj {
	return &diffuse1{sc, sh, refl}
}

const off = 1 / (1024)

func (s *diffuse1) Intersect(r Ray) (Inter, Shader) {
	return s.shape.Intersect(r), s
}

func (s *diffuse1) Intensity(r Ray, t float64) Color {

	n := Normal(s.shape, r, t)

	p := r.At(t) //.MAdd(off, n)
	acc := 0.
	for _, light := range s.scene.sources {
		lightPos, flux := light.Sample()
		d := lightPos.Sub(p)
		if !s.scene.IntersectsAny(Ray{p.MAdd(1e-6, d.Normalized()), d.Normalized()}) {
			acc += s.refl * flux * Max(n.Dot(d.Normalized())/(d.Len2()), 0)
		}
	}
	return Color(acc)
}

//// Diffuse shading with shadows and interreflection
//func Diffuse2(reflect float64) Shader {
//	return func(r Ray, t float64, n Vec, N int) float64 {
//		acc := diffuse1(reflect, r, t, n, N)
//		p := r.At(t).MAdd(off, n)
//		d := RandVec(n)
//		sec := Ray{p, d}
//		acc += reflect * Intensity(sec, N+1, false) * n.Dot(d.Normalized())
//		return acc
//	}
//}
//
//
//func Reflective(reflect float64) Shader {
//	return func(r Ray, t float64, n Vec, N int) float64 {
//		p := r.At(t).MAdd(off, n)
//		dir2 := reflectVec(r.Dir, n)
//		return reflect * Intensity(Ray{p, dir2}, N+1, true)
//	}
//}
//
//func ReflectiveMate(reflect float64, jitter float64) Shader {
//	return func(r Ray, t float64, n Vec, N int) float64 {
//		p := r.At(t).MAdd(off, n)
//		dir2 := reflectVec(r.Dir, n).MAdd(jitter, randVec(n))
//		return reflect * Intensity(Ray{p, dir2}, N+1, true)
//	}
//}
//
//func ShaderAdd(a, b Shader) Shader {
//	return func(r Ray, t float64, n Vec, N int) float64 {
//		return a.Intensity(r, t, n, N) + b.Intensity(r, t, n, N)
//	}
//}
//
//// reflects v of the surface with normal n.
//func reflectVec(v, n Vec) Vec {
//	return v.MAdd(-2*v.Dot(n), n)
//}
