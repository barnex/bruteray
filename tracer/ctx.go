package tracer

import (
	"fmt"
	"math/rand"

	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/tracer/sequence"
	"github.com/barnex/bruteray/util"
)

// Ctx is a "execution context", scoped to the pixel currently being shaded.
// A Ctx lives though a single pixel shading process,
// similar to how Go's "context.Context" lives through a single RPC request.
//
// A Ctx carries:
//  - quasi-random sequences approprately seeded for the pixel
//  - a storage pool for obtaining and recycling Rays (without allocation)
// 	- statistics about the compute resources used
//  - the current recursion depth, for aborting recursion
//
// The reason for this is twofold. (1) The random generator, storage pool
// and statistics are not thread-safe, so each goroutine needs its own instance.
// This avoids locking, and significantly speeds up execution.
// (2) Quasi Monte Carlo requires quasi-random numbers that depend on the
// pixel and recursion depth. Therefore, the context needs access to the recursion
// depth, and we might as well make the context responsible for limiting recursion depth.
type Ctx struct {
	CurrentRecursionDepth int // counts recursion depth from 1.
	//currentRecursionBreadth int // counts recutions breath on first depth level, from 1.

	sequence1 sequence.Sequence
	sequence2 sequence.Sequence
	sequence3 sequence.Sequence
	sequenceL sequence.Sequence
	AA        sequence.Sequence
	//Rng       *rand.Rand // TODO: rm

	rays  pool
	Stats Stats
}

const pseudo = false

func NewCtx(numPix int) *Ctx {
	sh := randomShifts(123, numPix)
	//sh2 := randomShifts(456, numPix)
	c := &Ctx{
		sequence1: sequence.NewHalton(2, 3, 1, sh),
		sequence2: sequence.NewHalton(5, 7, 1, sh),
		sequence3: sequence.PseudoRandom(),
		sequenceL: sequence.NewHalton(5, 7, 11, sh),
		AA:        sequence.NewHalton(2, 3, 1, sh),
		rays:      pool{new: func() interface{} { return new(Ray) }},
	}

	if pseudo {
		c.sequence1 = sequence.PseudoRandom()
		c.sequence2 = sequence.PseudoRandom()
		c.sequence3 = sequence.PseudoRandom()
		c.sequenceL = sequence.PseudoRandom()
		c.AA = sequence.PseudoRandom()
	}
	return c
}

func (c *Ctx) Init(pixel, pass int) {
	c.sequence1.Init(pixel, pass)
	c.sequence2.Init(pixel, pass)
	c.sequence3.Init(pixel, pass)
	c.sequenceL.Init(pixel, pass)
}

func (c *Ctx) Generate2() (u, v float64) {
	if c.CurrentRecursionDepth == 0 {
		panic(fmt.Sprintf("Ctx.Generate2: bad recursion depth: %v", c.CurrentRecursionDepth))
	}
	switch c.CurrentRecursionDepth {
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

// IsInitial returns whether this is the context of the initial ray cast by the camera.
// I.e., returns true when we are at the root of recursion.
// At the root of recursion, we may apply some expensive Media like fog,
// or chose to use Quasi Monte Carlo rather than the regular method.
func (c *Ctx) IsInitial() bool {
	util.Assert(c.CurrentRecursionDepth != 0)
	return c.CurrentRecursionDepth == 1
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

func randomShifts(seed int64, n int) []geom.Vec2 {
	rng := rand.New(rand.NewSource(seed))
	shifts := make([]geom.Vec2, n)
	for i := range shifts {
		shifts[i] = geom.Vec2{rng.Float64(), rng.Float64()}
	}
	return shifts
}
