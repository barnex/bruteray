package mat

import (
	"math"

	. "github.com/barnex/bruteray/br"
)

// ShadeDir returns a color based on the direction of a ray.
// Used for shading the ambient background, E.g., the sky.
type ShadeDir func(dir Vec) Color

func (s ShadeDir) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	return s(r.Dir())
}

// SkyDome maps a fisheye image on the sky.
func SkyDome(tex Image, rot float64) ShadeDir {
	return ShadeDir(
		func(dir Vec) Color {
			dir = dir.Normalized()
			r := math.Sqrt(dir[Z]*dir[Z] + dir[X]*dir[X])
			r = math.Asin(r) / (math.Pi / 2)
			//dir = dir.Mul(r)
			th := math.Atan2(dir[Z], dir[X]) + rot
			x := r * math.Cos(th)
			y := r * math.Sin(th)
			u := 0.5 + x*0.5
			v := 0.5 + y*0.5
			return tex.AtUV(u, v)
		})
}
