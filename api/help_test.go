package api

import (
	"testing"

	"github.com/barnex/bruteray/tracer/test"
)

func benchmark(b *testing.B, s Spec) {
	s.InitDefaults()
	test.Benchmark(b,
		s.Scene(),
		s.Camera,
		999999999,
	)
}
