package bruteray

import "math"

var (
	BLACK = Color{0, 0, 0}
	WHITE = Color{1, 1, 1}
	RED   = Color{1, 0, 0}
	GREEN = Color{0, 1, 0}
	BLUE  = Color{0, 0, 1}
)

// Color or light intensity with float64 precision.
type Color struct {
	R, G, B float64
}

func (c Color) Mul(s float64) Color {
	return Color{s * c.R, s * c.G, s * c.B}
}

// linear to sRGB gamma curve
// https://en.wikipedia.org/wiki/SRGB
func srgb(c float64) float64 {
	c = clip(c)
	if c <= 0.0031308 {
		return 12.92 * c
	}
	c = 1.055*math.Pow(float64(c), 1./2.4) - 0.05
	if c > 1 {
		return 1
	}
	return c
}

// clip color value between 0 and 1
func clip(v float64) float64 {
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	return v
}
