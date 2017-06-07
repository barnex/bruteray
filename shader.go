package main

type ShaderFunc func(pos, normal Vec, r Ray, rec int) float64

func ShadeFlat(v float64) ShaderFunc {
	return func(p, n Vec, r Ray, rec int) float64 {
		return v
	}
}

func ShadeDiffuse() ShaderFunc {
	return func(p, n Vec, r Ray, rec int) float64 {
		d := scene.light.Sub(p).Normalized()
		return 0.8*n.Dot(d) + scene.amb
	}
}

func WithShadow(sf ShaderFunc) ShaderFunc {
	return func(p, n Vec, r Ray, rec int) float64 {

		d := scene.light.Sub(p).Normalized()

		secondary := Ray{p.MAdd(0.01, d), d}
		if !intersAny(secondary, scene.objs) {
			return sf(p, n, r, rec-1) // not occluded, original shader
		}
		return scene.amb // occluded: ambient light
	}
}

func ShadeReflect() ShaderFunc {
	return func(p, n Vec, r Ray, rec int) float64 {
		p = p.MAdd(0.01, n) // make sure we're outside
		dir2 := reflect(r.Dir, n)
		secondary := Ray{p, dir2}
		return 0.5*PixelShade(scene, secondary, rec-1) + scene.amb
	}
}

// reflects v of the surface with normal n.
func reflect(v, n Vec) Vec {
	return v.MAdd(-2*v.Dot(n), n)
}
