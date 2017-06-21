package main

import (
	"log"
	"math"
)

// Camera renders a scene into a raw intensity image.
type Cam struct {
	Img      [][]Color
	ZMap     [][]Color
	FocalLen float64
	N        int
	Pitch    float64
}

func Camera(w, h int, focalLen float64) *Cam {
	return &Cam{
		Img:      MakeImage(w, h),
		ZMap:     MakeImage(w, h),
		FocalLen: focalLen,
	}
}

func (c *Cam) Size() (int, int) {
	return len(c.Img[0]), len(c.Img)
}

func (c *Cam) Render(s *Scene) [][]Color {
	c.iterate(s)
	return c.Img
}

func (c *Cam) Iterate(s *Scene, N int) [][]Color {
	for i := 0; i < N; i++ {
		c.iterate(s)
	}
	return c.Img
}

func (c *Cam) iterate(s *Scene) {
	focalPoint := Vec{0, 0, -c.FocalLen}
	W, H := c.Size()
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			// ray start point
			y0 := (-float64(i) + aa() + float64(H)/2) / float64(H)
			x0 := (float64(j) + aa() - float64(W)/2) / float64(H)
			start := Vec{x0, y0, 0}.RotX(c.Pitch)

			// ray direction
			dir := Vec{0, 0, 1}
			if c.FocalLen != 0 {
				dir = start.Sub(focalPoint).Normalized().RotX(c.Pitch)
			}
			dir = dir.RotX(c.Pitch)

			// accumulate ray intensity
			r := Ray{start, dir}
			t, v := s.Intensity(r, *maxRec)
			if math.IsNaN(float64(v)) {
				log.Println("ERROR: got NaN")
				continue
			}
			c.Img[i][j] += v
			c.ZMap[i][j] = Color(-t)
		}
	}
	c.N++
}

func MakeImage(W, H int) [][]Color {
	img := make([][]Color, H)
	for i := range img {
		img[i] = make([]Color, W)
	}
	return img
}

// Anti-aliasing jitter
func aa() float64 {
	//	return 0.5
	return Rand()
}
