package main

import (
	"fmt"
	"log"
	"math"
)

// Camera renders a scene into a raw intensity image.
type Cam struct {
	FocalLen float64
	N        int
	Pos      Vec
	Transf   Matrix
	AA       bool
}

func Camera(focalLen float64) *Cam {
	return &Cam{
		FocalLen: focalLen,
		Transf:   UnitMatrix(),
	}
}

func (c *Cam) Render(s *Env, img Image) {
	focalPoint := Vec{0, 0, -c.FocalLen}.Add(c.Pos)
	W, H := img.Size()
	r := &Ray{}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			// ray start point
			y0 := (-float64(i) + c.aa() + float64(H)/2) / float64(H)
			x0 := (float64(j) + c.aa() - float64(W)/2) / float64(H)
			start := Vec{x0, y0, 0}.Transf(&c.Transf).Add(c.Pos)

			// ray direction
			dir := Vec{0, 0, 1}
			if c.FocalLen != 0 {
				dir = start.Sub(focalPoint).Normalized().Transf(&c.Transf)
			}
			dir = dir.Transf(&c.Transf)

			// accumulate ray intensity
			r.Start = start
			r.Dir = dir
			v := s.Shade(r, *maxRec)
			if math.IsNaN(float64(v)) {
				log.Println("ERROR: got NaN")
				continue
			}
			if v > 100 {
				fmt.Println("too big", v)
				v = 100
			}
			img[i][j] += v
			//c.ZMap[i][j] = Color(-t)
		}
	}
	c.N++
}

// Anti-aliasing jitter
func (c *Cam) aa() float64 {
	if c.AA {
		return Rand()
	} else {
		return 0.5
	}
}
