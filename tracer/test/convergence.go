package test

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/barnex/bruteray/imagef"
	"github.com/barnex/bruteray/tracer"
	. "github.com/barnex/bruteray/tracer/types"
	"github.com/barnex/bruteray/util"
)

type Convergence struct {
	*tracer.Sampler
	Golden    imagef.Image
	Output    []Row
	totalPass int
}

type Row struct {
	Pass  int
	Error float64
}

func (c *Convergence) Sample(nPass int) {
	c.Sampler.Sample(nPass)
	c.totalPass += nPass
	c.Output = append(c.Output, c.Error())
}

func (c *Convergence) MustWrite(fname string) {
	f, err := os.Create(fname)
	util.Check(err)
	defer f.Close()
	b := bufio.NewWriter(f)
	defer b.Flush()
	util.Check(c.Write(b))
}

func (c *Convergence) Write(w io.Writer) error {
	for _, r := range c.Output {
		if _, err := fmt.Fprintf(w, "%v \t %v\n", r.Pass, r.Error); err != nil {
			return err
		}
	}
	return nil
}

func (c *Convergence) Error() Row {
	return Row{c.totalPass, c.rmsError()}
}

func (c *Convergence) rmsError() float64 {
	errSq := 0.0
	img := c.Sampler.Image()
	for iy := range img {
		errSqx := 0.0
		for ix := range img[iy] {
			a := clamp(img[iy][ix])
			b := clamp(c.Golden[iy][ix])
			errSqx += Vec{a.R, a.G, a.B}.Sub(Vec{b.R, b.G, b.B}).Len2()
		}
		errSq += errSqx
	}
	return math.Sqrt(errSq) / float64(c.Golden.NumPixels())
}

func clamp(c Color) Color {
	c.R = util.Min(1, c.R)
	c.G = util.Min(1, c.G)
	c.B = util.Min(1, c.B)
	return c
}
