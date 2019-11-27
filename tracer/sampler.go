package tracer

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"

	. "github.com/barnex/bruteray/imagef"
	. "github.com/barnex/bruteray/imagef/colorf"
	"github.com/barnex/bruteray/util"
)

func Uniform(f ImageFunc, numPass, w, h int, antiAlias bool) Image {
	s := NewSampler(f, w, h, antiAlias)
	s.Sample(numPass)
	return s.Image()
}

// A Sampler renders a ray-traced image.
type Sampler struct {
	f         ImageFunc
	sum       Image
	sumSq     Image
	n         [][]int
	antiAlias bool
	//	placement func(*Sampler)int??

	Stats Stats
	//Convergence []struct{samples int, error float64}
}

func NewSampler(f ImageFunc, w, h int, antiAlias bool) *Sampler {
	return &Sampler{
		f:         f,
		sum:       MakeImage(w, h),
		sumSq:     MakeImage(w, h),
		n:         makeInt2D(w, h),
		antiAlias: antiAlias,
	}
}

func (s *Sampler) Sample(nPass int) {
	s.SampleWithCancel(nPass, make(chan struct{}))
}

// TODO: cancellation
func (s *Sampler) SampleWithCancel(nPass int, _ chan struct{}) {
	start := time.Now()

	w, h := s.imageSize()
	const tileSize = 16
	work := tessellate(w, h, tileSize)

	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := NewCtx(w * h)
			for t := range work {
				s.sampleTile(ctx, t, nPass)
			}
		}()
	}
	wg.Wait()

	s.Stats.WallTime += time.Since(start)
	//s.Stats.Add(&ctx.Stats)
}

func (s *Sampler) sampleTile(ctx *Ctx, t tile, nPass int) {
	for iy := t.y0; iy < t.y1; iy++ {
		for ix := t.x0; ix < t.x1; ix++ {
			s.samplePixel(ctx, ix, iy, nPass)
		}
	}
}

func (s *Sampler) samplePixel(ctx *Ctx, ix, iy, n int) {
	w, h := s.imageSize()
	ctx.Stats.NumPixels++
	xi, yi := IndexToCam(w, h, float64(ix), float64(iy))
	pixs := PixelSize(w, h)
	for pass := 0; pass < n; pass++ {
		pixNum := (w*iy + ix)
		ctx.Init(pixNum, s.n[iy][ix])
		s.n[iy][ix]++

		x := xi
		y := yi
		if s.antiAlias {
			u, v := ctx.AA.Generate2()
			x += pixs * (u - 0.5)
			y += pixs * (v - 0.5)
		}

		c := s.f(ctx, x, y)
		if !c.IsNaN() {
			s.sum[iy][ix] = s.sum[iy][ix].Add(c)
			s.sumSq[iy][ix].R += c.R * c.R
			s.sumSq[iy][ix].G += c.G * c.G
			s.sumSq[iy][ix].B += c.B * c.B
		}
	}
}

// tessellate divides an image of size (w,h) into tiles
// no larger than (tileSize,tileSize).
// The tiles are returned in buffered and closed channel,
// ready for worker goroutines to range over.
func tessellate(w, h, tileSize int) chan tile {
	tilesW := divUp(w, tileSize)
	tilesH := divUp(h, tileSize)
	tilesN := tilesW * tilesH
	work := make(chan tile, tilesN)
	for x0 := 0; x0 < w; x0 += tileSize {
		x1 := min(x0+tileSize, w)
		for y0 := 0; y0 < h; y0 += tileSize {
			y1 := min(y0+tileSize, h)
			work <- tile{x0, y0, x1, y1}
		}
	}
	close(work)
	return work
}

// divUp returns x/y, but rounded up rather than down.
func divUp(x, y int) int {
	return ((x - 1) / y) + 1
}

type tile struct {
	x0, y0 int
	x1, y1 int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isCanceled(c chan struct{}) bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}

func (s *Sampler) imageSize() (w, h int) {
	return s.sum.Size()
}

// ??
func PixelSize(w, h int) float64 {
	wf, hf := float64(w), float64(h)
	minf := math.Min(wf, hf)
	return 1 / minf
}

func makeInt2D(w, h int) [][]int {
	list := make([]int, w*h)
	img := make([][]int, h)
	for i := range img {
		img[i] = list[i*w : (i+1)*w]
	}
	return img
}

// Image returns the rendered image currently accumulated by the Sampler.
// Successive calls to Sample progressively improve this image's quality.
func (s *Sampler) Image() Image {
	return s.memoize(s.imagef)
}

// StdDev returns an image whose pixel values are the standard deviation
// of all samples currently accumulated.
func (s *Sampler) StdDev() Image {
	return s.memoize(s.stddevf)
}

// memoize turns an image function into stored image.
func (s *Sampler) memoize(f func(ix, iy int) Color) Image {
	w, h := s.imageSize()
	img := MakeImage(w, h)
	for iy := 0; iy < h; iy++ {
		for ix := 0; ix < w; ix++ {
			img[iy][ix] = f(ix, iy)
		}
	}
	return img
}

func (s *Sampler) imagef(ix, iy int) Color {
	return s.sum[iy][ix].Mul(1 / float64(s.n[iy][ix]))
}

func (s *Sampler) stddevf(ix, iy int) Color {
	v := s.variancef(ix, iy)
	v.R = math.Sqrt(v.R)
	v.G = math.Sqrt(v.G)
	v.B = math.Sqrt(v.B)
	return v
}

func (s *Sampler) variancef(ix, iy int) Color {
	sum := s.sum[iy][ix]
	sumSq := s.sumSq[iy][ix]
	n := float64(s.n[iy][ix])
	var v Color
	d := 1 / (n * (n - 1))
	v.R = (n*sumSq.R - util.Sqr(sum.R)) * d
	v.G = (n*sumSq.G - util.Sqr(sum.G)) * d
	v.B = (n*sumSq.B - util.Sqr(sum.B)) * d
	return v
}

// IndexToCam maps a pixel index {ix, iy} inside an image with given width and height
// onto a u,v coordinate strictly inside the interval [0,1].
// If the image's aspect ratio width:height is not square,
// then either u or v will not span the entire [0,1] interval.
//
// Half-pixel offsets are applied so that the borders in u,v correspond
// exactly to pixel borders (not centers). This transformation is sketched below:
//
// 	             +----------------+ (u,v=1,1)
// 	             |                |
//  (x,y=-.5,-.5)+----------------+
// 	             |                |
// 	             |                |
// 	             +----------------+ (x,y=w-.5,h-.5)
// 	             |                |
// 	    (u,v=0,0)+----------------+
//
// Note that the v axis points up, while the y axis points down.
func IndexToCam(w, h int, ix, iy float64) (u, v float64) {
	W := float64(w)
	H := float64(h)

	if ix < -0.5 || iy < -0.5 || ix > W-0.5 || iy > H-0.5 {
		panic(fmt.Sprintf("IndexToCam: pixel index out of range: w=%v, h=%v, x=%v, y=%v",
			w, h, ix, iy))
	}

	u = linterp(-0.5, 0, W-0.5, 1, ix)
	v = linterp(-0.5, 0.5+0.5*(H/W), H-0.5, 0.5-0.5*(H/W), iy)
	return u, v
}

// linear interpolation
// 	x1 -> y1
// 	x2 -> y2
// 	x  -> y
func linterp(x1, y1, x2, y2, x float64) (y float64) {
	return y1 + (y2-y1)*(x-x1)/(x2-x1)
}
