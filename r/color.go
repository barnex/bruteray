package r

var (
	BLACK = Color{0, 0, 0}
	WHITE = Color{1, 1, 1}
	RED   = Color{1, 0, 0}
	GREEN = Color{0, 1, 0}
	BLUE  = Color{0, 0, 1}
)

type Color struct {
	R, G, B float64
}
