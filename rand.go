package bruteray

// Utilities for generating random numbers and vectors.

import (
	"math"
	"math/rand"
	"sync/atomic"
	"time"
)

func newRng() *rand.Rand {
	return rand.New(rand.NewSource(seed()))
}

var _seed int64 = 1

func seed() int64 {
	return atomic.AddInt64(&_seed, 12345) + time.Now().UnixNano()
}

// Uniform random number.
func random(e *Env) float64 {
	return e.rng.Float64()
}

// Normal random number.
func randNorm(e *Env) float64 {
	return e.rng.NormFloat64()
}

// Random unit vector.
func randVec(e *Env) Vec {
	return Vec{
		e.rng.NormFloat64(),
		e.rng.NormFloat64(),
		e.rng.NormFloat64(),
	}.Normalized()

}

// Random unit vector from the hemisphere around n
// (dot product with n >= 0).
func randVecDir(e *Env, n Vec) Vec {
	v := randVec(e)
	if v.Dot(n) < 0 {
		v = v.Mul(-1)
	}
	return v
}

// Random unit vector, sampled with probability cos(angle with dir).
// Used for diffuse inter-reflection importance sampling.
func randVecCos(e *Env, dir Vec) Vec {
	v := randVecDir(e, dir)
	for v.Dot(dir) < e.rng.Float64() {
		v = randVecDir(e, dir)
	}
	return v
}

// randCircle draws a point from the unit disk.
func randCircle(e *Env) (x, y float64) {
	x = 2*random(e) - 1
	y = 2*random(e) - 1
	for sqrt(x*x+y*y) > 1 {
		x = 2*random(e) - 1
		y = 2*random(e) - 1
	}
	return x, y
}

// randHex draws a point from the unit hexagon.
func randHex(e *Env) (x, y float64) {
	x = 2*random(e) - 1
	y = 2*random(e) - 1
	for abs(y) > sqrt3/2 || abs(x+y/sqrt3) > 1 || abs(x-y/sqrt3) > 1 {
		x = 2*random(e) - 1
		y = 2*random(e) - 1
	}
	return x, y
}

func abs(x float64) float64 { return math.Abs(x) }

const sqrt3 = 1.7320508075688772935
