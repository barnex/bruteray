package tracer

import (
	"fmt"

	"github.com/barnex/bruteray/random"
	"github.com/barnex/bruteray/util"
)

// Ctx is a thread-local context
// for passing mutable state like
// random number generators and allocation pools.
type Ctx struct {
	currentRecursionDepth   int // counts recursion depth from 1.
	currentRecursionBreadth int // counts recutions breath on first depth level, from 1.

	sequence1 random.Sequence
	sequence2 random.Sequence
	sequence3 random.Sequence
	sequenceL random.Sequence
	AA        random.Sequence
	//RngTODORemove *rand.Rand
	rays pool

	Stats Stats
}

//var RandomSequence = random.PseudoRandom
var RandomSequence = random.NewHalton23()

func NewCtx() *Ctx {
	return &Ctx{
		sequence1: random.NewHalton(2, 3, 1),
		sequence2: random.NewHalton(5, 7, 1),
		sequence3: random.PseudoRandom(),
		sequenceL: random.NewHalton(5, 7, 11),
		AA:        random.PseudoRandom(),
		//RngTODORemove: rand.New(rand.NewSource(int64(seed + 2000000))),
		rays: pool{new: func() interface{} { return new(Ray) }},
	}
}

func (c *Ctx) Init(pixel, pass int) {
	c.sequence1.Init(pixel, pass)
	c.sequence2.Init(pixel, pass)
	c.sequence3.Init(pixel, pass)
	c.sequenceL.Init(pixel, pass)
}

func (c *Ctx) Generate2() (u, v float64) {
	if c.currentRecursionDepth == 0 {
		panic(fmt.Sprintf("Ctx.Generate2: bad recursion depth: %v", c.currentRecursionDepth))
	}
	switch c.currentRecursionDepth {
	case 1:
		return c.sequence1.Generate2()
	case 2:
		return c.sequence2.Generate2()
	default:
		return c.sequence3.Generate2()
	}
}

func (c *Ctx) GenerateLens() (u, v float64) {
	return c.sequence3.Generate2()
}

//func (c *Ctx) Init(seed, index int) {
//	c.currentRecursionDepth = 0
//	c.currentRecursionBreadth = 0
//	//c.Sequence.Init(seed, index)
//}

//func (c *Ctx) NextRecursionBreadth() int {
//	if c.currentRecursionDepth > 1 {
//		return 0
//	}
//	c.currentRecursionBreadth++
//	return c.currentRecursionBreadth
//}
//
//func (c *Ctx) Init(seed, index int) {
//	c.currentRecursionDepth = 0
//	c.currentRecursionBreadth = 0
//	c.Sequence.Init(seed, index)
//}

// IsInitial returns whether this is the context of the initial ray cast by the camera.
// I.e., returns true when we are at the root of recursion.
// At the root of recursion, we may apply some expensive Media like fog,
// or chose to use Quasi Monte Carlo rather than the regular method.
func (c *Ctx) IsInitial() bool {
	util.Assert(c.currentRecursionDepth != 0)
	return c.currentRecursionDepth == 1
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
