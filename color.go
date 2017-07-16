package bruteray

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
