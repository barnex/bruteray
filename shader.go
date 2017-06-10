package main

type Shader func(r Ray, t float64, normal Vec, N int) float64

func (s Shader) Intensity(r Ray, t float64, n Vec, N int) float64 {
	return s(r, t, n, N)
}

func Flat(v float64) Shader {
	return func(r Ray, t float64, normal Vec, N int) float64 {
		return v
	}
}

func Diffuse1(reflect float64) Shader {
	return func(r Ray, t float64, n Vec, N int) float64 {
		p := r.At(t).MAdd(off, n)

		acc := 0.
		for _, light := range sources {
			lightPos, flux := light.Sample()
			d := lightPos.Sub(p)
			if !intersectsAny(Ray{p, d.Normalized()}) {
				acc += reflect * flux * n.Dot(d) / (d.Len2())
			}
		}
		return acc
	}
}

func intersectsAny(r Ray) bool {
	_, _, obj := FirstIntersect(r)
	return obj != nil
}

func Reflective(reflect float64) Shader {
	return func(r Ray, t float64, n Vec, N int) float64 {
		p := r.At(t).MAdd(off, n)
		dir2 := reflectVec(r.Dir, n)
		return reflect * Intensity(Ray{p, dir2}, N+1)
	}
}

// reflects v of the surface with normal n.
func reflectVec(v, n Vec) Vec {
	return v.MAdd(-2*v.Dot(n), n)
}

//func ShadeDiffuse() ShaderFunc {
//	return func(t float64, n Vec, r Ray, rec int) float64 {
//		p := r.At(t)
//		d := scene.light.Sub(p)
//		return 100*0.8*n.Dot(d.Normalized())/(d.Len2()) + scene.amb
//	}
//}
//
//const off = 1e-3 // tiny offset to avoid bleeding
//
//func WithShadow(sf ShaderFunc) ShaderFunc {
//	return func(t float64, n Vec, r Ray, rec int) float64 {
//
//		p := r.At(t)
//		d := scene.light.Sub(p).Normalized()
//
//		secondary := Ray{p.MAdd(off, d), d} // todo: rm
//		if !intersAny(secondary, scene.objs) {
//			return sf(t, n, r, rec-1) // not occluded, original shader
//		}
//		return scene.amb // occluded: ambient light
//	}
//}
//
//func ShadeGlobal() ShaderFunc {
//	return func(t float64, n Vec, r Ray, rec int) float64 {
//		p := r.At(t).MAdd(off, n)
//
//		a := 0.0
//		const N = 300
//
//		for i := 0; i < N; i++ {
//			secondary := Ray{p, randVec(n)}
//			v2, p2, ok := PixelShade(scene, secondary, rec-1)
//			if !ok {
//				continue
//			}
//			d := p2.Sub(p).Normalized()
//			v := v2 * n.Dot(d)
//			a += v
//		}
//		//assert(v >= 0)
//		return a / N
//	}
//}
//
//
