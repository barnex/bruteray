package main

type ShaderFunc func(t float64, normal Vec, r Ray, rec int) float64

func ShadeFlat(v float64) ShaderFunc {
	return func(t float64, n Vec, r Ray, rec int) float64 {
		return v
	}
}

func ShadeDiffuse() ShaderFunc {
	return func(t float64, n Vec, r Ray, rec int) float64 {
		p := r.At(t)
		d := scene.light.Sub(p)
		return 100*0.8*n.Dot(d.Normalized())/(d.Len2()) + scene.amb
	}
}

func WithShadow(sf ShaderFunc) ShaderFunc {
	return func(t float64, n Vec, r Ray, rec int) float64 {

		p := r.At(t)
		d := scene.light.Sub(p).Normalized()

		off := 1e-3                         // tiny offset to avoid bleeding
		secondary := Ray{p.MAdd(off, d), d} // todo: rm
		if !intersAny(secondary, scene.objs) {
			return sf(t, n, r, rec-1) // not occluded, original shader
		}
		return scene.amb // occluded: ambient light
	}
}

func ShadeReflect() ShaderFunc {
	return func(t float64, n Vec, r Ray, rec int) float64 {
		p := r.At(t)
		//p = p.MAdd(0.01, n) // make sure we're outside
		dir2 := reflect(r.Dir, n)
		secondary := Ray{p, dir2}
		return 0.5*PixelShade(scene, secondary, rec-1) + scene.amb
	}
}

// reflects v of the surface with normal n.
func reflect(v, n Vec) Vec {
	return v.MAdd(-2*v.Dot(n), n)
}
