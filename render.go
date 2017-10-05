package bruteray

import (
	"runtime"
	"sync"
)

func Render1Thread(e *Env, img Image) {
	H := img.Bounds().Dy()
	for i := 0; i < H; i++ {
		renderLine(e, img, i)
	}
}

func Render(e *Env, img Image) {
	render(e, img, runtime.NumCPU())
}

func render(e *Env, img Image, numCPU int) {
	H := img.Bounds().Dy()

	// numCPU goroutines will each render
	// one line at a time taken from ch.
	ch := make(chan int, H+1)
	for i := 0; i < H; i++ {
		ch <- i
	}
	close(ch)

	var wg sync.WaitGroup
	for i := 0; i < numCPU; i++ {
		eCopy := e.Copy()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range ch {
				renderLine(eCopy, img, i)
			}
		}()
	}
	wg.Wait()
}

func MultiPass(e *Env, img Image, N int) {
	multiPass(e, img, N, runtime.NumCPU())
}

func multiPass(e *Env, img Image, N int, numCPU int) {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	for i := 0; i < N; i++ {
		acc := MakeImage(w, h)
		render(e, acc, numCPU)
		img.Add(acc)
	}
	img.Mul(1 / float64(N))
}

func renderLine(e *Env, img Image, i int) {
	c := e.Camera
	focalPoint := Vec{0, 0, -c.FocalLen}
	W, H := img.Bounds().Dx(), img.Bounds().Dy()
	r := &Ray{}
	for j := 0; j < W; j++ {
		// ray start point
		y0 := (-float64(i) + c.aa(e) + float64(H)/2) / float64(H)
		x0 := (float64(j) + c.aa(e) - float64(W)/2) / float64(H)
		r.Start = Vec{x0, y0, 0}

		// ray direction
		r.Dir = Vec{0, 0, 1}
		if c.FocalLen != 0 {
			r.Dir = r.Start.Sub(focalPoint).Normalized()
		}

		// camera transform
		r.Transf(&(c.transf))

		// accumulate ray intensity
		c := e.ShadeAll(r, e.Recursion)

		// clip to avoid caustic noise
		if c.R > e.Cutoff || c.G > e.Cutoff || c.B > e.Cutoff {
			c = c.Mul(1 / c.Max())
		}

		img[i][j] = c
	}
}
