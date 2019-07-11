package br

import (
	"math/rand"
	"time"
)

// Ctx is a thread-local context
// for passing mutable state like
// random number generators and allocation pools.
type Ctx struct {
	Rng      *rand.Rand // Random-number generator for use by one thread
	fragPool Pool
	rayPool  Pool
}

func NewCtx(seed int) *Ctx {
	return &Ctx{
		Rng:      rand.New(rand.NewSource(int64(seed) + time.Now().UnixNano())),
		fragPool: Pool{New: func() interface{} { v := make([]Fragment, 0, 8); return &v }},
		rayPool:  Pool{New: func() interface{} { return new(Ray) }},
	}
}

// GetRay returns a new Ray, allocated from a pool.
// PutRay should be called to recycle the Ray.
func (c *Ctx) GetRay(start, dir Vec) *Ray {
	r := c.rayPool.Get().(*Ray)
	r.Start = start
	r.SetDir(dir)
	return r
}

// PutRay recycles Rays returned by GetRay.
func (c *Ctx) PutRay(r *Ray) {
	c.rayPool.Put(r)
}

// GetFrags returns a new []Fragment, allocated from a pool.
// PutFrags should be called to recycle.
func (c *Ctx) GetFrags() *[]Fragment {
	fb := c.fragPool.Get().(*[]Fragment)
	*fb = (*fb)[:0]
	return fb
}

// PutFrags recycles values returned by GetFrags.
func (c *Ctx) PutFrags(fb *[]Fragment) {
	c.fragPool.Put(fb)
}
