package br

import (
	"math"
	"math/rand"
)

// Env stores the entire environment
// (all objects, light sources, ... in the scene)
// as well as a random-number generator needed for iterative rendering.
type Env struct {
	objs        []Obj     // non-source objects
	Lights      []Light   // light sources
	all         []Obj     // objs + lights
	Ambient     Fragment  // Shades the background at infinity, when no object is hit
	Rng         rand.Rand // Random-number generator for use by one thread // TODO: rm
	Recursion   int       // Maximum allowed recursion depth.
	Fog         float64   // Fog distance
	IndirectFog bool      // Include fog interreflection

	fragPool Pool
	rayPool  Pool

	Cutoff float64 // Maximum allowed brightness. Used to suppress spurious caustics. TODO rm
	//Camera      *Cam      // Camera determines the point of view
}

// NewEnv creates an empty environment
// to which objects can be added later.
func NewEnv() *Env {
	return &Env{
		Ambient: Fragment{T: inf, Material: BLACK},
		Rng:     *(newRng()),
		//Camera:    Camera(0),
		Recursion: DefaultRec,
		Cutoff:    math.Inf(1),
		fragPool:  Pool{New: func() interface{} { v := make([]Fragment, 0, 8); return &v }},
		rayPool:   Pool{New: func() interface{} { return new(Ray) }},
	}
}

// Default recursion depth for NewEnv
const DefaultRec = 6

// TODO: rm
func (e *Env) Copy() *Env {
	e2 := *e
	e2.Rng = *(newRng())
	return &e2
}

// Adds an object to the scene.
func (e *Env) Add(o ...Obj) {
	e.objs = append(e.objs, o...)
	e.all = append(e.all, o...)
}

// Adds a light source to the scene.
func (e *Env) AddLight(l ...Light) {
	e.Lights = append(e.Lights, l...)
	for _, l := range l {
		e.all = append(e.all, l)
	}
}

// Adds a light source to the scene.
// The source itself is not visible, only its light.
func (e *Env) AddInvisibleLight(l ...Light) {
	e.Lights = append(e.Lights, l...)
}

// Sets the background color.
func (e *Env) SetAmbient(m Material) {
	e.Ambient = Fragment{T: inf, Material: m}
}

// Calculate intensity seen by ray,
// caused by all objects including lights.
// Used by specular surfaces
// who make no distinction between light sources and regular objects.
func (e *Env) ShadeAll(ctx *Ctx, r *Ray, N int) Color {
	return e.Shade(ctx, r, N, e.all)
}

// Calculate intensity seen by ray,
// caused by objects but excluding lights.
// Used for diffuse inter reflection
// where contributions of light sources are added separately.
func (e *Env) ShadeNonLum(ctx *Ctx, r *Ray, N int) Color {
	return e.Shade(ctx, r, N, e.objs)
}

func (e *Env) fb() *[]Fragment {
	fb := e.fragPool.Get().(*[]Fragment)
	*fb = (*fb)[:0]
	return fb
}

func (e *Env) rfb(fb *[]Fragment) {
	e.fragPool.Put(fb)
}

// Calculate intensity seen by ray, with maximum recursion depth N.
// who = objs, lights, or all.
func (e *Env) Shade(ctx *Ctx, r *Ray, N int, who []Obj) Color {
	if N <= 0 {
		return Color{}
	}

	var surf Fragment

	surf.T = inf
	surf.Material = e.Ambient.Material

	hit := e.fb()
	defer e.rfb(hit)

	for _, o := range who {
		o.Hit1(r, hit)

		for i := range *hit {
			t := (*hit)[i].T
			if t < surf.T && t > 0 {
				surf = (*hit)[i]
				surf.Object = o
			}
		}
	}

	// Shapes are not obliged to normalized and orient their surface normal.
	// Let's do it here, only for the normal that's going to be used.
	surf.Norm = surf.Norm.Normalized().Towards(r.Dir())

	if e.Fog != 0 && N == e.Recursion && e.Recursion > 1 {
		// add fog only to primary ray,
		// it's very expensive and the indirect effect is hardly visible.
		return e.withFog(ctx, surf, N-1, r)
	} else {
		return surf.Shade(ctx, e, N-1, r)
	}
}

func (e *Env) withFog(ctx *Ctx, surf Fragment, N int, r *Ray) Color {
	tObject := surf.T
	tScatter := e.Rng.ExpFloat64() * e.Fog
	if tScatter > tObject {
		return surf.Shade(ctx, e, N, r) // hit object without scattering
	}
	// else: ray scattered on fog

	//c := Color{}
	c := surf.Shade(ctx, e, N, r)
	pos := r.At(tScatter)
	for _, l := range e.Lights {
		lpos, intens := l.Sample(e, pos)
		secundary := e.NewRay(pos, lpos.Sub(pos).Normalized())
		defer e.RRay(secundary) //TODO: out of loop
		lightT := lpos.Sub(pos).Len()
		if e.Occludes(secundary, lightT) { // intersection between start and light position
			// shadow
		} else {
			c = c.MAdd(1/e.Fog, intens)
		}
	}

	if e.IndirectFog {
		r2 := e.NewRay(pos, randVec(&e.Rng))
		defer e.RRay(r2)
		fogc := e.Shade(ctx, r2, 1, e.objs)
		c = c.Add(fogc)
	}

	return c
}

// Returns t > 0 if r intersects any object.
// Used to determine shadows.
//func (e *Env) IntersectAny(r *Ray) float64 {
//
//	T := inf
//	hit := make([]Fragment, 0, 2)
//
//	for _, o := range e.objs {
//		hit = hit[:0]
//		o.Hit(r, &hit)
//
//		for i := range hit {
//			t := hit[i].T
//			if t < T && t > 0 {
//				T = t
//			}
//		}
//	}
//
//	if T == inf {
//		T = 0
//	}
//	return T
//}

// Occludes returns true when an object intersects r
// between t=0 and t=endpoint.
// This means a light source at endpoint casts a shadow at the ray start point.
func (e *Env) Occludes(r *Ray, endpoint float64) bool {

	hit := e.fb()
	defer e.rfb(hit)

	for _, o := range e.objs {

		o.Hit1(r, hit)

		for i := range *hit {
			t := (*hit)[i].T
			if t > 0 && t < endpoint {
				return true
			}
		}
	}
	return false
}
