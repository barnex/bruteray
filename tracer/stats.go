package tracer

import "time"

type Stats struct {
	NumPixels int // number of pixels evaluated
	NumRays   int // number of rays evaluated
	NumNaN    int // number of NaN colors encountered
	WallTime  time.Duration
}

func (s *Stats) Add(b *Stats) {
	if b.WallTime != 0 {
		panic("tracer.Stats.Add: non-zero walltime")
	}
	s.NumPixels += b.NumPixels
	s.NumRays += b.NumRays
	s.NumNaN += b.NumNaN
}
