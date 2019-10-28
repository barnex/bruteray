package random

import (
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/util"
)

//TODO:
//	seed = pixel number (limited to 16M)
//	subSequence: stratifies for each diffuce, light, lens, ...
//	index: incremented internaly upon Generate
//	Ctx returns Stratified or Pseudo for initial or recursive

type Sequence interface {
	Init(pixel, pass int)
	Generate2() (u, v float64)
}

func PseudoRandom() Sequence {
	return (*pseudoRandom)(rand.New(rand.NewSource(time.Now().UnixNano())))
}

type pseudoRandom rand.Rand

func (r *pseudoRandom) Init(pixel, pass int) {
}

func (r *pseudoRandom) Generate2() (u, v float64) {
	rng := (*rand.Rand)(r)
	return rng.Float64(), rng.Float64()
}

type halton struct {
	baseU, baseV   int
	offset, stride int
	index          int
	shiftU, shiftV float64
}

func NewHalton23() Sequence {
	return NewHalton(2, 3, 1)
}

func NewHalton(baseU, baseV, stride int) Sequence {
	return &halton{
		baseU:  baseU,
		baseV:  baseV,
		offset: 0,
		stride: stride,
		index:  0,
	}
}

func (s *halton) Init(pixel, pass int) {
	s.shiftU, s.shiftV = randomShift(pixel)
	s.index = pass
}

var (
	shifts []geom.Vec2
	shlock sync.Mutex
)

func randomShift(i int) (u, v float64) {
	shlock.Lock()
	for len(shifts) <= i {
		shifts = append(shifts, geom.Vec2{})
	}
	if shifts[i] == (geom.Vec2{}) {
		shifts[i] = geom.Vec2{rand.Float64(), rand.Float64()}
	}
	s := shifts[i]
	shlock.Unlock()
	return s[0], s[1]
}

func (s *halton) Generate2() (u, v float64) {
	//if subsequence >= s.stride {
	//	panic(fmt.Sprintf("halton.generate: subsequence %v > stride %v", subsequence, s.stride))
	//}
	//i := s.offset + s.index*s.stride // TODO
	//s.index++
	u = util.Frac(Halton(s.baseU, s.index) + s.shiftU)
	v = util.Frac(Halton(s.baseV, s.index) + s.shiftV)
	return u, v
}

// https://en.wikipedia.org/wiki/Halton_sequence
func Halton(b, i int) float64 {
	i++ // actual series starts from 1
	f := 1.0
	r := 0.0

	for i > 0 {
		f = f / float64(b)
		r = r + f*(float64(i%b))
		i = int(math.Floor(float64(i) / float64(b)))
	}
	return r
}
