// Package raster provides functionality to render a bruteray.Env into an image.
package raster

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/barnex/bruteray/br"
)

// SinglePass renders a scene by evaluating each pixel once.
// Suited if no Monte-Carlo methods (e.g. Diffuse Material) are involved.
func SinglePass(cam *Cam, e *br.Env, img Image) {
	render(cam, e, img, runtime.NumCPU())
}

func render(cam *Cam, e *br.Env, img Image, numCPU int) {
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
		eCopy := e.Copy() // TODO: rm
		ctx := &br.Ctx{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range ch {
				renderLine(ctx, cam, eCopy, img, i)
			}
		}()
	}
	wg.Wait()
}

// MutliPass renders a scene by evaluating each pixel a fixed number of times and averaging the results.
// Suited for Monte-Carlo methods.
func MultiPass(cam *Cam, e *br.Env, img Image, passes int) {
	multiPass(cam, e, img, passes, runtime.NumCPU())
}

func multiPass(cam *Cam, e *br.Env, img Image, passes int, numCPU int) {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	for i := 0; i < passes; i++ {
		acc := MakeImage(w, h)
		render(cam, e, acc, numCPU)
		img.Add(acc)
	}
	img.Mul(1 / float64(passes))
}

// RenderLoop starts an infinite loop, continuously improving the image quality.
// Intermediate images can be queried from the peek channel.
// Used for live previews.
func RenderLoop(cam *Cam, e *br.Env, w, h int, peek chan chan Image) {
	img := MakeImage(w, h)
	passes := 0

	onePass := func() {
		start := time.Now()
		acc := MakeImage(w, h)
		render(cam, e, acc, runtime.NumCPU())
		passes++
		rate := float64(w*h) / time.Since(start).Seconds()
		log.Printf("pass: %v, %.2f Mpixel/s", passes, rate/1e6)
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

func renderLine(ctx *br.Ctx, cam *Cam, e *br.Env, img Image, i int) {
	W, H := img.Bounds().Dx(), img.Bounds().Dy()

	r := e.NewRay(br.Vec{}, br.Vec{})
	defer e.RRay(r)
	for j := 0; j < W; j++ {

		cam.RayFrom(e, i, j, W, H, r)

		// accumulate ray intensity
		c := e.ShadeAll(ctx, r, e.Recursion)

		// clip to avoid caustic noise
		if c.R > e.Cutoff || c.G > e.Cutoff || c.B > e.Cutoff {
			c = c.Mul(e.Cutoff / c.Max())
		}

		img[i][j] = c
	}
}
