package shape

import . "github.com/barnex/bruteray/v1/br"

// Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).
// TODO: Transl
func NewCylinder(dir int, center Vec, diam, h float64, m Material) CSGObj {
	r := diam / 2
	coeff := Vec{1, 1, 1}
	coeff[dir] = 0
	infCyl := &quad{center, coeff, r * r, m}
	capped := And(infCyl, Slab(Unit[dir], center[dir]-h/2, center[dir]+h/2, m))
	return capped
}
