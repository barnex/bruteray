package bruteray

import (
	"math"
	"math/rand"
)

// Env stores the entire environment
// (all objects, light sources, ... in the scene)
// as well as a random-number generator needed for iterative rendering.
type Env struct {
	objs      []Obj     // non-source objects
	lights    []Light   // light sources
	all       []Obj     // objs + lights
	Ambient   Shader    // Shades the background at infinity, when no object is hit
	rng       rand.Rand // Random-number generator for use by one thread
	Camera    *Cam      // Camera determines the point of view
	Recursion int       // Maximum allowed recursion depth.
	Cutoff    float64   // Maximum allowed brightness. Used to suppres spurious caustics.
}

// NewEnv creates an empty environment
// to which objects can be added later.
func NewEnv() *Env {
	return &Env{
		Ambient:   Shader{T: inf, Material: Flat(BLACK)},
		rng:       *(newRng()),
		Camera:    Camera(0),
		Recursion: DefaultRec,
		Cutoff:    math.Inf(1),
	}
}

// Default recursion depth for NewEnv
const DefaultRec = 6

// Returns a copy with its own random number generator,
// so it can be used from a different thread.
func (e *Env) Copy() *Env {
	e2 := *e
	e2.rng = *(newRng())
	return &e2
}

// Adds an object to the scene.
func (e *Env) Add(o ...Obj) {
	e.objs = append(e.objs, o...)
	e.all = append(e.all, o...)
}

// Adds a light source to the scene.
func (e *Env) AddLight(l ...Light) {
	e.lights = append(e.lights, l...)
	for _, l := range l {
		e.all = append(e.all, l)
	}
}

// Sets the background color.
func (e *Env) SetAmbient(m Material) {
	e.Ambient = Shader{T: inf, Material: m}
}

// Calculate intensity seen by ray,
// caused by all objects including lights.
// Used by specular surfaces
// who make no distinction between light sources and regular objects.
func (e *Env) ShadeAll(r *Ray, N int) Color {
	return e.shade(r, N, e.all)
}

// Calculate intensity seen by ray,
// caused by objects but excluding lights.
// Used for diffuse inter reflection
// where contributions of light sources are added separately.
func (e *Env) ShadeNonLum(r *Ray, N int) Color {
	return e.shade(r, N, e.objs)
}

// Calculate intensity seen by ray, with maximum recursion depth N.
// who = objs, lights, or all.
func (e *Env) shade(r *Ray, N int, who []Obj) Color {
	if N <= 0 {
		return Color{}
	}

	surf := e.Ambient
	surf.T = inf

	hit := make([]Shader, 0, 2)

	for _, o := range who {
		hit = hit[:0]
		o.Hit(r, &hit)

		for i := range hit {
			t := hit[i].T
			if t < surf.T && t > 0 {
				surf = hit[i]
			}
		}
	}
	c := surf.Shade(e, N-1, r)

	return c
}

// Returns t > 0 if r intersects any object.
// Used to determine shadows.
func (e *Env) IntersectAny(r *Ray) float64 {

	T := inf
	hit := make([]Shader, 0, 2)

	for _, o := range e.objs {
		hit = hit[:0]
		o.Hit(r, &hit)

		for i := range hit {
			t := hit[i].T
			if t < T && t > 0 {
				T = t
			}
		}
	}

	if T == inf {
		T = 0
	}
	return T
}
