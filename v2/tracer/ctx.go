package tracer

import (
	"math/rand"
	"time"

	. "github.com/barnex/bruteray/v2/geom"
)

// Ctx is a thread-local context
// for passing mutable state like
// random number generators and allocation pools.
type Ctx struct {
	Rng  *rand.Rand // Random-number generator for use by one thread
	rays pool

	PixelCount int // Tracks total number of pixels evaluated
	RayCount   int // Tracks total number of rays evaluated
}

func NewCtx(seed int64) *Ctx {
	return &Ctx{
		Rng:  rand.New(rand.NewSource(int64(seed) + time.Now().UnixNano())),
		rays: pool{new: func() interface{} { return new(Ray) }},
	}
}

// Ray returns a new Ray, allocated from a pool.
// PutRay should be called to recycle the Ray.
func (c *Ctx) Ray() *Ray {
	r := c.rays.get().(*Ray)
	*r = Ray{Len: Inf}
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
