package sequence

import (
	"math/rand"
	"time"

	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/util"
)

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
	shifts         []geom.Vec2
}

func NewHalton(baseU, baseV, stride int, shifts []geom.Vec2) Sequence {
	return &halton{
		baseU:  baseU,
		baseV:  baseV,
		offset: 0,
		stride: stride,
		index:  0,
		shifts: shifts,
	}
}

func (s *halton) Init(pixel, pass int) {
	sh := s.shifts[pixel]
	s.shiftU, s.shiftV = sh[0], sh[1]
	s.index = pass
}

func (s *halton) Generate2() (u, v float64) {
	u = util.Frac(Halton(s.baseU, s.index) + s.shiftU)
	v = util.Frac(Halton(s.baseV, s.index) + s.shiftV)
	return u, v
}
