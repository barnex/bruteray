package main

type Source interface {
	Sample() (Vec, float64)
}

type PointSource struct {
	Pos  Vec
	Flux float64
}

func (s *PointSource) Sample() (Vec, float64) {
	return s.Pos, s.Flux
}

type BulbSource struct {
	Pos  Vec
	Flux float64
	R    float64
}

func (s *BulbSource) Sample() (Vec, float64) {
	return Vec{RandNorm(), RandNorm(), RandNorm()}.Mul(s.R).Add(s.Pos), s.Flux
}
