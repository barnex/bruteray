package sampler

import (
	"runtime"
	"sync"
	"time"

	imagef "github.com/barnex/bruteray/image"
	"github.com/barnex/bruteray/tracer"
)

// Uniform returns a uniformly sampled image,
// averaging nPass samples per pixel.
// Multi-pass images are anti-aliased.
func Uniform(f tracer.ImageFunc, nPass, w, h int) Image {
	img := imagef.MakeImage(w, h)

	// channel with work items: line numbers to render
	ch := make(chan int, (w+1)*(h+1)*nPass)
	for iy := 0; iy < h; iy++ {
		for pass := 0; pass < nPass; pass++ {
			ch <- iy
		}
	}
	close(ch)

	// render lines and passes in parallel
	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		ctx := tracer.NewCtx(time.Now().UnixNano() + int64(i))
		go func() {
			defer wg.Done()
			for iy := range ch {
				addLine(ctx, img, f, iy)
			}
		}()
	}
	wg.Wait()

	// normalize final image
	normalize := 1 / float64(nPass)
	for iy := 0; iy < h; iy++ {
		for ix := 0; ix < w; ix++ {
			img[iy][ix] = img[iy][ix].Mul(normalize)
		}
	}
	return img
}

func addLine(ctx *tracer.Ctx, img Image, f tracer.ImageFunc, iy int) {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	for ix := 0; ix < w; ix++ {
		x, y := IndexToCam(w, h, float64(ix), float64(iy))
		c := f(ctx, x, y)
		if !c.IsNaN() {
			img[iy][ix] = img[iy][ix].Add(c)
		}
	}
}
