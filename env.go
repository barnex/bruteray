package bruteray

import (
	"math/rand"
)

// Env stores the entire environment
// (all objects, light sources, ... in the scene)
type Env struct {
	objs    []Obj   // non-source objects
	lights  []Light // light sources
	all     []Obj   // objs + lights
	Ambient Surf
	rng     rand.Rand
	Camera  *Cam
}

func NewEnv() *Env {
	return &Env{
		Ambient: Surf{T: inf, Material: Flat(BLACK)},
		rng:     *(newRng()),
		Camera:  Camera(0),
	}
}

func (e *Env) SetAmbient(m Material) {
	e.Ambient = Surf{T: inf, Material: m}
}

// Returns a copy with its own random number generator,
// so it can be used from a different thread.
func (e *Env) Copy() *Env {
	e2 := *e
	e2.rng = *(newRng())
	return &e2
}

func (e *Env) Add(o ...Obj) {
	e.objs = append(e.objs, o...)
	e.all = append(e.all, o...)
}

func (e *Env) AddLight(l ...Light) {
	e.lights = append(e.lights, l...)
	for _, l := range l {
		e.all = append(e.all, l)
	}
}

// Calculate intensity seen by ray,
// caused by objects including lights.
// Used by specular surfaces
// who make no distinction between light sources and regular objects.
func (e *Env) ShadeAll(r *Ray, N int) Color {
	//fmt.Println(e.objs)
	//fmt.Println(e.lights)
	//fmt.Println(e.all)
	return e.shade(r, N, e.all)
}

// Calculate intensity seen by ray,
// caused by objects but excluding lights.
// Used for diffuse inter reflection
// where contributions of light sources are added separately.
// TODO: once a ray has hit a diffuse surface, luminous objects should be excluded at subsequent specular reflections.
// Otherwise we get caustics, which are not rendered nicely.
func (e *Env) ShadeNonLum(r *Ray, N int) Color {
	return e.shade(r, N, e.objs)
}

// Calculate intensity seen by ray,
// with maximum recursion depth N.
func (e *Env) shade(r *Ray, N int, who []Obj) Color {
	if N <= 0 {
		return Color{}
	}

	surf := e.Ambient
	surf.T = inf

	for _, o := range who {
		bi := o.Inters(r)
		Interval{bi.S1.T, bi.S2.T}.check()
		if !bi.OK() {
			continue
		}
		assert(bi.S1.T <= bi.S2.T)
		if t := bi.S1.T; t < surf.T && t > 0 {
			surf = bi.S1
		}
	}
	return surf.Shade(e, N-1, r)
}

// Returns t > 0 if r intersects any object
func (e *Env) IntersectAny(r *Ray) float64 {
	t := inf
	I := -1
	for i, o := range e.objs {
		bi := o.Inters(r)
		if !bi.OK() {
			continue
		}
		if bi.S1.T < 0 {
			continue
		}
		if bi.S1.T < t {
			t = bi.S1.T
			I = i
		}
	}
	if I == -1 {
		t = 0
	}
	return t
}
