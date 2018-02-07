package bruteray

import (
	"log"
	"runtime"
	"sync"
	"time"
)

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

func MultiPass(e *Env, img Image, passes int) {
	multiPass(e, img, passes, runtime.NumCPU())
}

func multiPass(e *Env, img Image, passes int, numCPU int) {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	for i := 0; i < passes; i++ {
		acc := MakeImage(w, h)
		render(e, acc, numCPU)
		img.Add(acc)
	}
	img.Mul(1 / float64(passes))
}

func RenderLoop(e *Env, w, h int, peek chan chan Image) {
	img := MakeImage(w, h)
	passes := 0

	onePass := func() {
		start := time.Now()
		acc := MakeImage(w, h)
		render(e, acc, runtime.NumCPU())
		passes++
		rate := float64(w*h) / time.Since(start).Seconds()
		log.Printf("pass: %v, %.1f Mpixel/s", passes, rate/1e6)
		img.Add(acc)
	}

	for {
		select {
		default:
			onePass()
		case resp := <-peek:
			log.Println("peeking...")
			cpy := MakeImage(w, h)
			scale := 1 / float64(passes)
			for i := range cpy {
				for j := range cpy[i] {
					cpy[i][j] = img[i][j].Mul(scale)
				}
			}
			resp <- cpy
			onePass() // after peeking, make sure we render at least one pass
		}
	}
}

func renderLine(e *Env, img Image, i int) {
	W, H := img.Bounds().Dx(), img.Bounds().Dy()

	r := e.NewRay(Vec{}, Vec{})
	defer e.RRay(r)
	for j := 0; j < W; j++ {

		e.Camera.RayFrom(e, i, j, W, H, r)

		// accumulate ray intensity
		c := e.ShadeAll(r, e.Recursion)

		// clip to avoid caustic noise
		if c.R > e.Cutoff || c.G > e.Cutoff || c.B > e.Cutoff {
			c = c.Mul(e.Cutoff / c.Max())
		}

		img[i][j] = c
	}
}
