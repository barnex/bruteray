package bruteray

import "math/rand"

var rng = rand.New(rand.NewSource(1))

// Uniform random number.
func Rand() float64 {
	return rng.Float64()
}

// Normal random number.
func RandNorm() float64 {
	return rng.NormFloat64()
}

// Random unit vector.
func RandVec() Vec {
	return Vec{RandNorm(), RandNorm(), RandNorm()}.Normalized()

}

// Random unit vector, dot product with n >= 0.
func RandVecDir(n Vec) Vec {
	v := RandVec()
	if v.Dot(n) < 0 {
		v = v.Mul(-1)
	}
	return v
}

// Random unit vector, sampled with probability cos(angle with dir).
// Used for diffuse inter-reflection importance sampling.
func RandVecCos(dir Vec) Vec {
	v := RandVecDir(dir)
	for v.Dot(dir) < Rand() {
		v = RandVecDir(dir)
	}
	return v
}
