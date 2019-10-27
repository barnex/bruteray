package sampler

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"
	"time"

	colorf "github.com/barnex/bruteray/color"
	imagef "github.com/barnex/bruteray/image"
	"github.com/barnex/bruteray/tracer"
	"github.com/barnex/bruteray/util"
)

type Image = imagef.Image
type ImageGray = imagef.ImageGray

type Adaptive struct {
	f           tracer.ImageFunc
	sum         Image
	sumSq       Image
	n           ImageGray // TODO: int
	totalPasses int
	AntiAlias   bool

	PixelCount, RayCount int
	WallTime             time.Duration
}

func NewAdaptive(f tracer.ImageFunc, w, h int, antiAlias bool) *Adaptive {
	return &Adaptive{
		f:         f,
		sum:       imagef.MakeImage(w, h),
		sumSq:     imagef.MakeImage(w, h),
		n:         imagef.MakeImageGray(w, h),
		AntiAlias: antiAlias,
	}
}

func (s *Adaptive) Sample(nPass int) {
	s.SampleNumCPU(runtime.NumCPU(), nPass)
}

func (s *Adaptive) SampleNumCPU(nCPU, nPass int) {
	s.SampleNumCPUWithCancel(nCPU, nPass, make(chan struct{}))
}

func (s *Adaptive) SampleNumCPUWithCancel(nCPU, nPass int, cancel chan struct{}) {
	w, h := s.Bounds().Dx(), s.Bounds().Dy()
	numPix := w * h
	budget := numPix * nPass
	s.totalPasses += nPass
	total := s.TotalVariance()
	fmt.Println("pass", s.totalPasses, "var/pix:", total/float64(numPix))

	// channel with work items: line numbers to render
	ch := make(chan int, numPix)
	for iy := 0; iy < h; iy++ {
		ch <- iy
	}
	close(ch)

	ctx := make([]*tracer.Ctx, nCPU)
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < nCPU; i++ {
		wg.Add(1)
		ctx[i] = tracer.NewCtx(time.Now().UnixNano() + int64(i))
		go func(i int) {
			defer wg.Done()
			for iy := range ch {
				select {
				case <-cancel:
					return
				default:
				}
				s.SampleLine(ctx[i], iy, w, h, budget, total)
			}
		}(i)
	}
	wg.Wait()

	s.WallTime += time.Since(start)
	for _, c := range ctx {
		s.PixelCount += c.PixelCount
		s.RayCount += c.RayCount
	}
}

func (s *Adaptive) Image() image.Image {
	return stratifiedImage{s}
}

func (s *Adaptive) StoredImage() imagef.Image {
	w, h := s.Bounds().Dx(), s.Bounds().Dy()
	cpy := imagef.MakeImage(w, h)
	for iy := 0; iy < h; iy++ {
		for ix := 0; ix < w; ix++ {
			cpy[iy][ix] = s.Color64At(ix, iy)
		}
	}
	return cpy
}

func (s *Adaptive) Stats() string {
	t := s.WallTime.Seconds()
	return fmt.Sprintf("%.1f rays/s, %.1f pix/s", float64(s.RayCount)/t, float64(s.PixelCount)/t)
}

func (s *Adaptive) SampleLine(ctx *tracer.Ctx, iy, w, h, budget int, totalVariance float64) {
	for ix := 0; ix < w; ix++ {
		n := 1
		if s.totalPasses > 30 {
			N := (float64(budget) * s.filteredVariance(ix, iy) / totalVariance)
			if N <= 0 || N > 1000 || math.IsNaN(N) {
				n = 1
			}
			n = util.Dither(ctx.Rng, N)
		}
		if n > 100 {
			n = 100
		}
		s.samplePixel(ctx, ix, iy, w, h, n)
	}
}

func (s *Adaptive) samplePixel(ctx *tracer.Ctx, ix, iy, w, h int, n int) {
	ctx.PixelCount++
	xi, yi := IndexToCam(w, h, float64(ix), float64(iy))
	pixs := imagef.PixelSize(w, h)
	for pass := 0; pass < n; pass++ {
		x := xi
		y := yi
		if s.AntiAlias {
			x += pixs * (ctx.Rng.Float64() - 0.5)
			y += pixs * (ctx.Rng.Float64() - 0.5)
		}

		c := s.f(ctx, x, y)
		if !c.IsNaN() {
			s.sum[iy][ix] = s.sum[iy][ix].Add(c)
			s.sumSq[iy][ix].R += c.R * c.R
			s.sumSq[iy][ix].G += c.G * c.G
			s.sumSq[iy][ix].B += c.B * c.B
			s.n[iy][ix]++
		}
	}
}

func (s *Adaptive) filteredVariance(i, j int) float64 {
	return util.Max3(s.varianceGamma(i-1, j), s.varianceGamma(i, j), s.varianceGamma(i+1, j))
	//return s.varianceGamma(i, j)
}

// Returns the gamma-corrected (i.e. perceived) noise level at pixel i,j.

func (s *Adaptive) varianceGamma(i, j int) float64 {
	if i < 0 || i >= len(s.sum[0]) {
		return 0
	}
	c := s.sum[j][i]
	v := s.variance3(i, j)

	return varGamma(c.R, v.R) + varGamma(c.G, v.G) + varGamma(c.B, v.B)
}

func varGamma(color, variance float64) float64 {

	dev := math.Sqrt(float64(variance))
	gammaDev := dev * colorf.SRGBSlope(color)
	gammaVar := gammaDev * gammaDev
	if math.IsNaN(gammaVar) {
		return 0
	}
	return gammaVar
}

func (s *Adaptive) variance3(i, j int) colorf.Color {
	sum := s.sum[j][i]
	sumSq := s.sumSq[j][i]
	n := s.n[j][i]

	var v colorf.Color
	d := 1 / (n * (n - 1))
	v.R = (float64(n)*sumSq.R - sqr(sum.R)) * d
	v.G = (float64(n)*sumSq.G - sqr(sum.G)) * d
	v.B = (float64(n)*sumSq.B - sqr(sum.B)) * d
	return v
}

func (s *Adaptive) stdErr(i, j int) float64 {
	v := s.variance3(i, j)
	return math.Sqrt(float64((v.R + v.G + v.B) / s.n[j][i]))
}

func (s *Adaptive) TotalVariance() float64 {
	total := 0.
	for iy := range s.sum {
		for ix := range s.sum[0] {
			s := s.varianceGamma(ix, iy)
			if !math.IsNaN(s) {
				total += s
			}
		}
	}
	return total
}

func sqr(x float64) float64   { return x * x }
func sqr64(x float64) float64 { return x * x }
func sqrt(x float64) float64  { return float64(math.Sqrt(float64(x))) }

func (s *Adaptive) StdDevImage() image.Image { return stratifiedStdDev{s} }
func (s *Adaptive) StdErrImage() image.Image { return stratifiedStdErr{s} }

func (s *Adaptive) SamplingImage() image.Image {
	img := s.n
	var max float64
	for iy := range img {
		for ix := range img[iy] {
			if n := img[iy][ix]; n > max {
				max = n
			}
		}
	}
	cpy := imagef.MakeImageGray(s.Bounds().Dx(), s.Bounds().Dy())
	for iy := range img {
		for ix := range img[iy] {
			cpy[iy][ix] = img[iy][ix] / max
		}
	}
	return cpy
}

// Bounds implements image.Image
func (s *Adaptive) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(s.sum[0]), len(s.sum))
}

// ColorModel implements image.Image
func (s *Adaptive) ColorModel() color.Model {
	return nil
}

type stratifiedImage struct{ *Adaptive }
type stratifiedStdDev struct{ *Adaptive }
type stratifiedStdErr struct{ *Adaptive }

// At implements image.Image
func (s stratifiedImage) At(i, j int) color.Color {
	return s.Color64At(i, j)
}

func (s *Adaptive) Color64At(i, j int) colorf.Color {
	return s.sum[j][i].Mul(float64(1 / s.n[j][i]))
}

// At implements image.Image
func (s stratifiedStdDev) At(i, j int) color.Color {
	c := float64(math.Sqrt(s.varianceGamma(i, j)))
	return colorf.Color{c, c, c}
}

func (s stratifiedStdErr) At(i, j int) color.Color {
	c := float64(s.stdErr(i, j))
	return colorf.Color{c, c, c}
}
