package sampler

import (
	"sync"
	"time"

	colorf "github.com/barnex/bruteray/color"
	imagef "github.com/barnex/bruteray/image"
	"github.com/barnex/bruteray/tracer"
)

type Sampler struct {
	f         tracer.ImageFunc
	sum       Image
	sumSq     Image
	n         [][]int
	antiAlias bool
	//	placement func(*Sampler)int??

	Stats tracer.Stats
}

type ImageFunc func(x, y int) colorf.Color

func Uniform(f tracer.ImageFunc, numPass, w, h int, antiAlias bool) imagef.Image {
	//antiAlias := (numPass > 1)
	s := New(f, w, h, antiAlias)
	s.Sample(numPass)
	return s.Image()
}

func New(f tracer.ImageFunc, w, h int, antiAlias bool) *Sampler {
	return &Sampler{
		f:         f,
		sum:       imagef.MakeImage(w, h),
		sumSq:     imagef.MakeImage(w, h),
		n:         makeInt2D(w, h),
		antiAlias: antiAlias,
	}
}

func (s *Sampler) Sample(nPass int) {
	s.SampleWithCancel(nPass, make(chan struct{}))
}

func (s *Sampler) SampleWithCancel(nPass int, cancel chan struct{}) {
	start := time.Now()

	w, h := s.imageSize()

	//numPix := w * h
	//budget := numPix * nPass
	//s.totalPasses += nPass
	//total := s.TotalVariance()
	//fmt.Println("pass", s.totalPasses, "var/pix:", total/float64(numPix))
	//fmt.Println(s.Stats())

	var wg sync.WaitGroup
	for iy := 0; iy < h; iy++ {
		wg.Add(1)
		go func(iy int) {
			defer wg.Done()
			ctx := tracer.NewCtx() // ???????????????????????????????????????????
			if isCanceled(cancel) {
				return
			}
			for ix := 0; ix < w; ix++ {
				s.samplePixel(ctx, ix, iy, w, h, nPass)
			}
		}(iy)
	}
	wg.Wait()

	s.Stats.WallTime += time.Since(start)
	//s.Stats.Add(&ctx.Stats)
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
	w = s.sum.Bounds().Dx()
	h = s.sum.Bounds().Dy()
	return w, h
}

func (s *Sampler) samplePixel(ctx *tracer.Ctx, ix, iy, w, h int, n int) {
	ctx.Stats.NumPixels++
	xi, yi := IndexToCam(w, h, float64(ix), float64(iy))
	pixs := imagef.PixelSize(w, h)
	for pass := 0; pass < n; pass++ {
		pixNum := (w*iy + ix)
		ctx.Init(pixNum, s.n[iy][ix])
		s.n[iy][ix]++

		x := xi
		y := yi
		if s.antiAlias { // need deeper dimension sampler eg.g halton5,7
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

func (s *Sampler) Image() imagef.Image {
	w, h := s.imageSize()
	img := imagef.MakeImage(w, h)
	for iy := 0; iy < h; iy++ {
		for ix := 0; ix < w; ix++ {
			img[iy][ix] = s.sum[iy][ix].Mul(1 / float64(s.n[iy][ix]))
		}
	}
	return img
}

func makeInt2D(w, h int) [][]int {
	list := make([]int, w*h)
	img := make([][]int, h)
	for i := range img {
		img[i] = list[i*w : (i+1)*w]
	}
	return img
}
