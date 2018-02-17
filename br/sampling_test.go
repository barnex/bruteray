package br

import (
	"fmt"
	"math"
	"testing"
)

func TestImportanceSampling(t *testing.T) {

	testFn := []func(Vec) float64{
		func(v Vec) float64 { return 1 },
		func(v Vec) float64 { return re(1 - v[Y]) },
		func(v Vec) float64 { return re(1 - v[Y]*v[Y] + v[Z]*v[X]) },
		func(v Vec) float64 { return re(v[Y]) },
		//func(v Vec) float64 { return Sqr(Sqr(1 - v[Y])) }, // converges slowly (but correct with large N)
		func(v Vec) float64 { return sqrt(1 - v[Y]) },
	}

	const (
		N   = 3e4
		tol = 5e-2 // a few times 1/sqrt(N)
	)

	for i, f := range testFn {
		i := i
		t.Run(fmt.Sprintf("f%v", i), func(t *testing.T) {
			//t.Parallel() // TODO: enable if rng is multi-threaded
			for _, dir := range []Vec{Ex, Ey, Ez, Vec{1, 2, 3}.Normalized()} {
				want := uniformInt(f, N, dir)
				have := importanceInt(f, N, dir)
				if math.Abs(have-want) > tol || math.IsNaN(have) {
					t.Errorf("importance sampling: dir %v: have %.2f, want %.2f", dir, have, want)
				}
			}
		})
	}
}

func re(x float64) float64 { return math.Max(x, 0) }

func uniformInt(f func(Vec) float64, N int, dir Vec) float64 {
	e := NewEnv()
	acc := 0.0
	for i := 0; i < N; i++ {
		V := randVecDir(&e.Rng, dir)
		acc += f(V) * dir.Dot(V) * 2 * Pi
	}
	return acc / float64(N)
}

func importanceInt(f func(Vec) float64, N int, dir Vec) float64 {
	e := NewEnv()
	acc := 0.0
	for i := 0; i < N; i++ {
		V := RandVecCos(e, dir)
		acc += f(V) * Pi
	}
	return acc / float64(N)
}