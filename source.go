package main

type Source interface {
	Sample() (Vec, float64)
}

func PointLight(pos Vec, flux float64) Source {
	return &pointSource{pos, flux}
}

type pointSource struct {
	Pos  Vec
	Flux float64
}

func (s *pointSource) Sample() (Vec, float64) {
	return s.Pos, s.Flux
}

func SmoothLight(pos Vec, flux float64, size float64) Source {
	return &bulbSource{pos, flux, size}
}

type bulbSource struct {
	Pos  Vec
	Flux float64
	R    float64
}

func (s *bulbSource) Sample() (Vec, float64) {
	p := Vec{RandNorm(), RandNorm(), RandNorm()}
	for p.Len2() > 9 {
		p = Vec{RandNorm(), RandNorm(), RandNorm()}
	}
	p = p.Mul(s.R).Add(s.Pos)
	return p, s.Flux
}
