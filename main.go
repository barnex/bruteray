package main

const (
	W = 300
	H = 200
)

var (
	Focal = Vec{0, 0, -1}
	Horiz = 10.0
)

func main() {

	img := make([][]float64, H)
	for i := range img {
		img[i] = make([]float64, W)
	}

	scene := Sphere(0.5)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {

			y0 := (float64(i) - H/2 + 0.5) / H
			x0 := (float64(j) - W/2 + 0.5) / H

			start := Vec{x0, y0, 0}
			r := Ray{start, start.Sub(Focal).Mul(Horiz)}

			if Inters(r, scene) {
				img[i][j] = 1
			}

		}
	}

	Encode(img, "a.out")
}

func Inters(r Ray, s Shape) bool {
	d := 0.01
	for t := 0.0; t < Horiz; t += d {
		if s.Inside(r.At(t)) {
			return true
		}
	}
	return false
}

type Ray struct {
	Start Vec
	Dir   Vec
}

func (r *Ray) At(t float64) Vec {
	return r.Start.Add(r.Dir.Mul(t))
}
