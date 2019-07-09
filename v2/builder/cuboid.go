package builder

import (
	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
)

// TODO: handle material uniformly
func CuboidFaces(m Material, o, a, b, c Vec) *Tree {
	faces := []Builder{
		NewRectangle(m, o, a, b),
		NewRectangle(m, o.Add(c), a.Add(c), b.Add(c)),
		NewRectangle(m, o, b, c),
		NewRectangle(m, o.Add(a), b.Add(a), c.Add(a)),
		NewRectangle(m, o, c, a),
		NewRectangle(m, o.Add(b), c.Add(b), a.Add(b)),
	}
	t := NewTree(faces...)
	t.NoDivide = true
	return t
}
