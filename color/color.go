package color

import (
	"math"

	"github.com/barnex/bruteray/geom"
)

var (
	Black   = Color{0, 0, 0}
	White   = Color{1, 1, 1}
	Red     = Color{1, 0, 0}
	Yellow  = Color{1, 1, 0}
	Green   = Color{0, 1, 0}
	Cyan    = Color{0, 1, 1}
	Blue    = Color{0, 0, 1}
	Magenta = Color{1, 0, 1}
)

// Color represents either a reflectivity or intensity.
//
// In case of reflectivity, values should be [0..1],
// 1 meaning 100% reflectivity for that color.
//
// The color space is linear.
type Color struct {
	R, G, B float64
}

func Gray(v float64) Color {
	return Color{v, v, v}
}

// At implements tracer.Texture3D.
func (c Color) At(_ geom.Vec) Color {
	return c
}

// AtUV implements tracer.Texture2D.
func (c Color) AtUV(u, v float64) Color {
	return c
}

// Implements color.Color.
// Converts from float64 linear space to 16-bit srgb.
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(LinearToSRGB(c.R) * 0xffff)
	g = uint32(LinearToSRGB(c.G) * 0xffff)
	b = uint32(LinearToSRGB(c.B) * 0xffff)

	return r, g, b, 0xffff
}

// Multiplies the color, making it darker (s<1) or brighter (s>1).
// E.g.:
// 	RED.Mul(0.5) // 50% reflective red, i.e. dark red.
func (c Color) Mul(a float64) Color {
	s := float64(a)
	return Color{s * c.R, s * c.G, s * c.B}
}

// Multiplies the color by 2^ev.
// E.g.:
// 	RED.EV(-1) // 50% reflective red, i.e. dark red.
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
func (c Color) MAdd(a float64, b Color) Color {
	s := float64(a)
	return Color{c.R + s*b.R, c.G + s*b.G, c.B + s*b.B}
}

// Exposure value, 2^exp.
func EV(exp float64) float64 {
	return math.Pow(2, exp)
}

func (c *Color) Max() float64 {
	max := c.R
	if c.G > max {
		max = c.G
	}
	if c.B > max {
		max = c.B
	}
	return max
}

func (c Color) Gray() float64 {
	return (c.R + c.G + c.B) / 3
}

// linear to sRGB conversion
// https://en.wikipedia.org/wiki/SRGB
func LinearToSRGB(c float64) float64 {
	c = clip(c)
	if c <= 0.0031308 {
		return 12.92 * c
	}
	c = float64(1.055*math.Pow(float64(c), 1./2.4) - 0.05)
	if c > 1 {
		return 1
	}
	return c
}

// Slope of linear to SRGB curve (without clipping).
func SRGBSlope(c float64) float64 {
	if c <= 0.0031308 {
		return 12.92
	}
	return 1.055 * (1. / 2.4) * math.Pow(float64(c), (1./2.4)-1)
}

// sRGB to linear conversion
// https://en.wikipedia.org/wiki/SRGB
// TODO: Look-up table for bytes
func SRGBToLinear(s float64) float64 {
	if s <= 0.04045 {
		return s / 12.92
	}
	const a = 0.055
	return math.Pow((s+a)/(1+a), 2.4)
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

func (a Color) IsNaN() bool {
	return a.R != a.R || a.G != a.G || a.B != a.B
}
