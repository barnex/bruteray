package tracer_test

import (
	"bufio"
	"encoding/gob"
	"math"
	"math/rand"
	"os"
	"path"
	"testing"

	"github.com/barnex/bruteray/imagef"
	"github.com/barnex/bruteray/imagef/colorf"
	"github.com/barnex/bruteray/tracer"
	"github.com/barnex/bruteray/tracer/cameras"
	"github.com/barnex/bruteray/tracer/lights"
	"github.com/barnex/bruteray/tracer/materials"
	"github.com/barnex/bruteray/tracer/objects"
	"github.com/barnex/bruteray/tracer/test"
	. "github.com/barnex/bruteray/tracer/types"
)

func TestSampler_Convergence(t *testing.T) {
	s := tracer.NewSampler(haltonGauss, 300, 200, false)
	s.Sample(100)
	test.Compare(t, 0.06, s.Image())
}

func TestSampler_StdDev(t *testing.T) {
	s := tracer.NewSampler(haltonGauss, 300, 200, false)
	s.Sample(100)
	test.Compare(t, 0.06, s.StdDev())
}

func TestSampler_Cornell(t *testing.T) {
	f := cornelli().ImageFunc(
		cameras.Projective(70 * Deg).Translate(Vec{.250, .250001, 0.97}),
	)

	w, h := 100, 75

	aa := false

	golden := tracer.Uniform(f, 1000, w, h, aa)

	smplr := test.Convergence{
		Sampler: tracer.NewSampler(f, w, h, aa),
		Golden:  golden,
	}

	for i := 0; i < 300; i++ {
		smplr.Sample(1)
	}
	//fname := path.Join(test.Testdata(), test.TestName()+".gob")
	//test.Check(t, gencode(smplr.Image(), fname))
	//test.Save(t, smplr.Image(), path.Join("got", test.TestName()+".png"))
	test.Save(t, golden, path.Join("got", test.TestName()+".png"))
	smplr.MustWrite("cornell.txt")
}

func gencode(img imagef.Image, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	defer b.Flush()
	return gob.NewEncoder(b).Encode(img)
}

func gdecode(fname string) (Image, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b := bufio.NewReader(f)
	var img Image
	err = gob.NewDecoder(b).Decode(&img)
	return img, err
}

func loadImg(t *testing.T, fname string) Image {
	t.Helper()
	img, err := imagef.Load(fname, colorf.SRGBToLinear)
	test.Check(t, err)
	return img
}

func cornelli() *Scene {
	s := .500
	L := .150
	dy := 0.0001
	white := materials.Matte(Color{1, 1, 1}.EV(-1))
	brightness := colorf.White.EV(5)
	return NewScene(
		2,
		[]Light{
			lights.RectangleLight(brightness, L, L, Vec{s / 2, s - dy, s / 2}),
		},
		objects.RectangleWithVertices(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, 0, s}),
		objects.RectangleWithVertices(white, Vec{0, s, 0}, Vec{s, s, 0}, Vec{0, s, s}),
		objects.RectangleWithVertices(white, Vec{0, 0, 0}, Vec{s, 0, 0}, Vec{0, s, 0}),
		objects.RectangleWithVertices(white, Vec{0, 0, 0}, Vec{0, 0, s}, Vec{0, s, 0}),
		objects.RectangleWithVertices(white, Vec{s, 0, 0}, Vec{s, 0, s}, Vec{s, s, 0}),
		objects.Sphere(white, 0.3, Vec{0.3, 0.15, 0.3}),
	)
}

func TestSampler_Convergence_Rate(t *testing.T) {
	w := 300
	h := 200
	golden := tracer.Uniform(cleanGauss, 1, w, h, false)

	s := test.Convergence{
		Sampler: tracer.NewSampler(haltonGauss, 300, 200, false),
		Golden:  golden,
	}
	for i := 0; i < 100; i++ {
		s.Sample(1)
	}
	s.MustWrite("convergence_rate.txt")
}

func BenchmarkSampler(b *testing.B) {
	s := tracer.NewSampler(haltonGauss, 300, 200, false)
	b.SetBytes(200 * 300)
	s.Sample(b.N)
}

func haltonGauss(c *Ctx, u, v float64) Color {
	c.CurrentRecursionDepth = 1
	rnd, _ := c.Generate2()
	return gauss(rnd, u, v)
}

func randomGauss(c *Ctx, u, v float64) Color {
	rnd := rand.Float64()
	return gauss(rnd, u, v)
}

func cleanGauss(c *Ctx, u, v float64) Color {
	return gauss(0.5, u, v)
}

func gauss(rnd, u, v float64) Color {
	u -= 0.5
	v -= 0.5
	r2 := u*u + v*v
	a := math.Exp(-r2 * 200)
	x := 2 * a * rnd
	return Color{x, x, x}
}
