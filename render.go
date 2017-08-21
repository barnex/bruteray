package bruteray

import (
	"runtime"
	"sync"
)

func Render1Thread(e *Env, maxRec int, img Image) {
	H := img.Bounds().Dy()
	const stride = 1
	for i := 0; i < H; i += stride {
		renderLine(e, maxRec, img, i, i+stride)
	}
}

func Render(e *Env, maxRec int, img Image) {
	H := img.Bounds().Dy()
	var wg sync.WaitGroup
	const stride = 1
	for i := 0; i < H; i += stride {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			renderLine(e.Copy(), maxRec, img, i, i+stride)
		}(i)
	}
	wg.Wait()
}

type Loop struct {
	env     *Env
	r, w, h int
	acc     Image
	n       float64
	mu      sync.Mutex
}

func (l *Loop) run() {
	for {
		l.iter()
	}
}

func (l *Loop) iter() {
	img := MakeImage(l.w, l.h)
	Render(l.env, l.r, img)
	l.mu.Lock()
	l.acc.Add(img)
	l.n++
	l.mu.Unlock()
}

func (l *Loop) Current() Image {
	l.mu.Lock()
	defer l.mu.Unlock()
	//log.Println("current")
	img := MakeImage(l.w, l.h)
	for i := range img {
		for j := range img[i] {
			img[i][j] = l.acc[i][j].Mul(1 / l.n)
		}
	}
	return img
}

func RenderLoop(e *Env, maxRec int, w, h int) *Loop {
	l := &Loop{env: e, r: maxRec, w: w, h: h, acc: MakeImage(w, h)}
	l.iter()   // make sure we have 1 pass at least
	go l.run() // refine in the background
	return l
}

func MultiPass(e *Env, maxRec int, img Image, N int) {
	que := make(chan struct{}, N)
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		que <- struct{}{}
	}
	close(que)

	imgCh := make(chan Image, 1)
	imgCh <- img

	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range que {
				img2 := MakeImage(w, h)
				Render1Thread(e.Copy(), maxRec, img2)
				img := <-imgCh

				// TODO: img.Add()
				for i := range img {
					for j := range img[i] {
						img[i][j] = img[i][j].Add(img2[i][j])
					}
				}
				imgCh <- img
			}
		}()
	}
	wg.Wait()

	// TODO: img.Scale()
	for i := range img {
		for j := range img[i] {
			img[i][j] = img[i][j].Mul(1 / float64(N))
		}
	}
}

func renderLine(e *Env, maxRec int, img Image, hMin, hMax int) {
	c := e.Camera
	focalPoint := Vec{0, 0, -c.FocalLen}
	W, H := img.Bounds().Dx(), img.Bounds().Dy()
	r := &Ray{}
	for i := hMin; i < hMax; i++ {
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
			v := e.ShadeAll(r, maxRec)
			img[i][j] = v
		}
	}
}

// Anti-aliasing jitter
func (c *Cam) aa(e *Env) float64 {
	if c.AA {
		return Rand(e)
	} else {
		return 0.5
	}
}
