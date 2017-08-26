package bruteray

import "math"

var (
	BLACK   = Color{0, 0, 0}
	WHITE   = Color{1, 1, 1}
	RED     = Color{1, 0, 0}
	GREEN   = Color{0, 1, 0}
	YELLOW  = Color{1, 1, 0}
	BLUE    = Color{0, 0, 1}
	MAGENTA = Color{1, 0, 1}
)

// Color or light intensity with float64 precision.
type Color struct {
	R, G, B float64
}

// Implements color.Color.
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(srgb(c.R) * 0xffff)
	g = uint32(srgb(c.G) * 0xffff)
	b = uint32(srgb(c.B) * 0xffff)

	return r, g, b, 0xffff
}

// Multiplies the color, making it darker (s<1) or brighter (s>1).
func (c Color) Mul(s float64) Color {
	return Color{s * c.R, s * c.G, s * c.B}
}

// Multiplies the color by 2^ev.
func (c Color) EV(ev float64) Color {
	return c.Mul(EV(ev))
}

// Point-wise multiplication of two colors.
// E.g.: light reflecting off a colored surface.
func (c Color) Mul3(b Color) Color {
	return Color{c.R * b.R, c.G * b.G, c.B * b.B}
}

// Adds two colors (i.e. blends them).
func (c Color) Add(b Color) Color {
	return Color{c.R + b.R, c.G + b.G, c.B + b.B}
}

// Adds s*b to color c.
func (c Color) MAdd(s float64, b Color) Color {
	return Color{c.R + s*b.R, c.G + s*b.G, c.B + s*b.B}
}

// Exposure value, 2^exp.
func EV(exp float64) float64 {
	return math.Pow(2, exp)
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
