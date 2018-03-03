package shape

import . "github.com/barnex/bruteray/br"

// TODO: remove
var CsgAnd_ func(a, b CSGObj) CSGObj

// Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).
// TODO: Transl
func NewCylinder(dir int, center Vec, diam, h float64, m Material) CSGObj {
	r := diam / 2
	coeff := Vec{1, 1, 1}
	coeff[dir] = 0
	infCyl := &quad{center, coeff, r * r, m}
	if CsgAnd_ == nil {
		panic("csgand==nil")
	}
	capped := CsgAnd_(infCyl, Slab(Unit[dir], center[dir]-h/2, center[dir]+h/2, m))
	return capped
}
