package main

// Camera renders a scene into a raw intensity image.
type Camera struct {
	W, H     int     // film size in pixels
	FocalLen float64 // focal length
}

func (c *Camera) Render(s *Scene) [][]float64 {
	img := MakeImage(c.W, c.H)
	c.iterate(s, img)
	return img
}

func (c *Camera) iterate(s *Scene, img [][]float64) {
	focalPoint := Vec{0, 0, -c.FocalLen}
	W := c.W
	H := c.H
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			// ray start point
			y0 := (-float64(i) + aa() + float64(H)/2) / float64(H)
			x0 := (float64(j) + aa() - float64(W)/2) / float64(H)
			start := Vec{x0, y0, 0}

			// ray direction
			dir := Vec{0, 0, 1}
			if c.FocalLen != 0 {
				dir = start.Sub(focalPoint).Normalized()
			}

			// accumulate ray intensity
			r := Ray{start, dir}
			v := s.Intensity(r)
			img[i][j] += v
		}
	}
}

func MakeImage(W, H int) [][]float64 {
	img := make([][]float64, H)
	for i := range img {
		img[i] = make([]float64, W)
	}
	return img
}

// Anti-aliasing jitter
func aa() float64 {
	return 0.5
	//return Rand()
}
