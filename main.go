package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {

	const W = 300
	const H = 200

	img := make([][]float64, H)
	for i := range img {
		img[i] = make([]float64, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {

			//y0 := (i - H/2) / H
			//x0 := (j - W/2) / W

			img[i][j] = 0.5

		}
	}

	f, err := os.Create("a.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = png.Encode(f, Gray(img))
	if err != nil {
		log.Fatal(err)
	}
}

type Gray [][]float64

func (g Gray) At(i, j int) color.Color {
	return color.Gray16{uint16(g[j][i] * (1 << 15))}
}

func (g Gray) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(g[0]), len(g))
}

func (g Gray) ColorModel() color.Model {
	return nil
}

type Shape interface {
	Inside(r Vec) bool
}

type ShapeFn func(Vec) bool

func (f ShapeFn) Inside(r Vec) bool {
	return f(r)
}

func Sphere(r float64) Shape {
	return ShapeFn(func(x Vec) bool {
		return x.Dot(x) < r*r
	})
}

type Vec struct {
	X, Y, Z float64
}

func (a Vec) Dot(b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}
