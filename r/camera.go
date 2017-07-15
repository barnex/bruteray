package r

import "sync"

// Camera renders a scene into a raw intensity image.
type Cam struct {
	FocalLen float64
	pos      Vec
	transf   Matrix
	AA       bool
}

func Camera(focalLen float64) *Cam {
	return &Cam{
		FocalLen: focalLen,
		transf:   UnitMatrix(),
	}
}

func (c *Cam) Transl(r Vec) {
	c.pos = c.pos.Add(r)
}

func (c *Cam) Transf(t Matrix) {
	c.transf = *t.Mul(&c.transf)
}

func (c *Cam) Render(e *Env, maxRec int, img Image) {
	H := img.Bounds().Dy()
	var wg sync.WaitGroup
	const stride = 1
	for i := 0; i < H; i += stride {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.renderLine(e, maxRec, img, i, i+stride)
		}(i)
	}
	wg.Wait()
}

func (c *Cam) renderLine(e *Env, maxRec int, img Image, hMin, hMax int) {
	focalPoint := Vec{0, 0, -c.FocalLen}.Add(c.pos)
	W, H := img.Bounds().Dx(), img.Bounds().Dy()
	r := &Ray{}
	for i := hMin; i < hMax; i++ {
		for j := 0; j < W; j++ {
			// ray start point
			y0 := (-float64(i) + c.aa() + float64(H)/2) / float64(H)
			x0 := (float64(j) + c.aa() - float64(W)/2) / float64(H)
			start := Vec{x0, y0, 0}.Transf(&c.transf).Add(c.pos)

			// ray direction
			dir := Vec{0, 0, 1}
			if c.FocalLen != 0 {
				dir = start.Sub(focalPoint).Normalized().Transf(&c.transf)
			}
			dir = dir.Transf(&c.transf)

			// accumulate ray intensity
			r.Start = start
			r.Dir = dir
			v := e.Shade(r, maxRec)
			img[i][j] = v
		}
	}
}

// Anti-aliasing jitter
func (c *Cam) aa() float64 {
	if c.AA {
		return Rand()
	} else {
		return 0.5
	}
}
