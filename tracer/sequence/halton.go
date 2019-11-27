package sequence

import "math"

// Halton(b, i) returns the i'th element of the Halton series with base b.
// i starts from 0.
// The base b should be >= 2.
// See https://en.wikipedia.org/wiki/Halton_sequence
func Halton(b, i int) float64 {
	i++ // actual series starts from 1
	bf := float64(b)
	f := 1.0
	r := 0.0

	for i > 0 {
		f = f / bf
		r = r + f*(float64(i%b))
		i = int(math.Floor(float64(i) / bf))
	}
	return r
}
