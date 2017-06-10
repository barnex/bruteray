package main

//type ShaderFunc func(t float64, normal Vec, r Ray, rec int) float64
//
//func ShadeFlat(v float64) ShaderFunc {
//	return func(t float64, n Vec, r Ray, rec int) float64 {
//		return v
//	}
//}
//
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
//func ShadeReflect() ShaderFunc {
//	return func(t float64, n Vec, r Ray, rec int) float64 {
//		p := r.At(t)
//		//p = p.MAdd(0.01, n) // make sure we're outside
//		dir2 := reflect(r.Dir, n)
//		secondary := Ray{p, dir2}
//		v, p, ok := PixelShade(scene, secondary, rec-1)
//		if !ok {
//			return scene.amb
//		}
//		return v
//
//		//d := scene.light.Sub(p).Normalized()
//		//return v * n.Dot(d)
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
//// reflects v of the surface with normal n.
//func reflect(v, n Vec) Vec {
//	return v.MAdd(-2*v.Dot(n), n)
//}
