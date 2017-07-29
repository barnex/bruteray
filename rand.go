package bruteray

import (
	"math/rand"
	"sync/atomic"
)

func newRng() *rand.Rand {
	return rand.New(rand.NewSource(seed()))
}

var _seed int64 = 1

func seed() int64 {
	return atomic.AddInt64(&_seed, 1)
}

// Uniform random number.
func Rand(e *Env) float64 {
	return e.rng.Float64()
}

// Normal random number.
func RandNorm(e *Env) float64 {
	return e.rng.NormFloat64()
}

// Random unit vector.
func RandVec(e *Env) Vec {
	return Vec{
		e.rng.NormFloat64(),
		e.rng.NormFloat64(),
		e.rng.NormFloat64(),
	}.Normalized()

}

// Random unit vector, dot product with n >= 0.
func RandVecDir(e *Env, n Vec) Vec {
	v := RandVec(e)
	if v.Dot(n) < 0 {
		v = v.Mul(-1)
	}
	return v
}

// Random unit vector, sampled with probability cos(angle with dir).
// Used for diffuse inter-reflection importance sampling.
func RandVecCos(e *Env, dir Vec) Vec {
	v := RandVecDir(e, dir)
	for v.Dot(dir) < e.rng.Float64() {
		v = RandVecDir(e, dir)
	}
	return v
}
