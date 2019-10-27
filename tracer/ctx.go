package tracer

import (
	"math/rand"
	"time"
)

// Ctx is a thread-local context
// for passing mutable state like
// random number generators and allocation pools.
type Ctx struct {
	Rng  *rand.Rand // Random-number generator for use by one thread. TODO: hide
	rays pool

	// TODO: struct Stats
	PixelCount int // Tracks total number of pixels evaluated
	RayCount   int // Tracks total number of rays evaluated
	// TODO: NaN count!
}

func NewCtx(seed int64) *Ctx {
	return &Ctx{
		Rng:  rand.New(rand.NewSource(int64(seed) + time.Now().UnixNano())),
		rays: pool{new: func() interface{} { return new(Ray) }},
	}
}

// TODO: LHS
func (c *Ctx) Sample2() (u, v float64) {
	return c.Rng.Float64(), c.Rng.Float64()
}

func (c *Ctx) SampleDisk() (u, v float64) {
	u, v = c.Sample2()
	for u*u+v*v > 1 {
		u, v = c.Sample2()
	}
	return u, v
}

// Ray returns a new Ray, allocated from a pool.
// PutRay should be called to recycle the Ray.
// TODO: rename NewRay
func (c *Ctx) Ray() *Ray {
	r := c.rays.get().(*Ray)
	*r = Ray{}
	return r
}

// Put recycles Rays returned by GetRay.
func (c *Ctx) PutRay(r *Ray) {
	c.rays.put(r)
}

type pool struct {
	new func() interface{}
	p   []interface{}
}

func (p *pool) get() interface{} {
	if len(p.p) == 0 {
		return p.new()
	}
	fb := p.p[len(p.p)-1]
	p.p = p.p[:len(p.p)-1]
	return fb
}

func (p *pool) put(v interface{}) {
	p.p = append(p.p, v)
}
