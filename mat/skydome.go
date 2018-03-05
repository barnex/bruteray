package mat

import . "github.com/barnex/bruteray/br"

// ShadeDir returns a color based on the direction of a ray.
// Used for shading the ambient background, E.g., the sky.
type ShadeDir func(dir Vec) Color

func (s ShadeDir) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color {
	return s(r.Dir())
}

func Skybox(tex Image) ShadeDir {
	return ShadeDir(
		func(dir Vec) Color {
			dir = dir.Normalized()
			u := 0.5 + dir[X]*0.5
			v := 0.5 + dir[Z]*0.5
			return tex.AtUV(u, v)
		})
}
