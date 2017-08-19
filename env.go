package bruteray

import (
	"math/rand"
)

// Env stores the entire environment
// (all objects, light sources, ... in the scene)
type Env struct {
	objs    []Obj
	lights  []Light
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

// Returns a copy with its own random number generator,
// so it can be used from a different thread.
func (e *Env) Copy() *Env {
	e2 := *e
	e2.rng = *(newRng())
	return &e2
}

// Returns a copy where expensive elements are replaced by fast ones,
// E.g. Smooth lights replaced by point lights, etc.
func (e *Env) Preview() *Env {
	e2 := *e
	e2.lights = make([]Light, len(e.lights))
	for i, l := range e.lights {
		e2.lights[i] = toPreview(l)
	}
	return &e2
}

func (e *Env) Add(o ...Obj) {
	e.objs = append(e.objs, o...)
}

func (e *Env) AddLight(l ...Light) {
	e.lights = append(e.lights, l...)
}

// Calculate intensity seen by ray,
// with maximum recursion depth N.
func (e *Env) Shade(r *Ray, N int) Color {
	if N <= 0 {
		return Color{}
	}

	surf := e.Ambient
	surf.T = inf

	for _, o := range e.objs {
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
