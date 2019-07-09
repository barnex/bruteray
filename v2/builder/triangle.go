package builder

import . "github.com/barnex/bruteray/v2/geom"
import . "github.com/barnex/bruteray/v2/tracer"

type Triangle struct {
	Face
	Texture Material
}

var _ Builder = (*Triangle)(nil)

func NewTriangle(t Material, a, b, c Vec) *Triangle {
	n := TriangleNormal(a, b, c)
	return &Triangle{
		Face:    Face{&Vertex{a, n, 0, 0}, &Vertex{b, n, 1, 0}, &Vertex{c, n, 0, 1}},
		Texture: t,
	}
}

func (t *Triangle) Intersect(c *Ctx, r *Ray) HitRecord {
	f := t.Face.Intersect(c, r)
	f.Material = t.Texture
	return f
}

func (t *Triangle) CtrlPoints() []*Vec {
	return []*Vec{&t.Face[0].Pos, &t.Face[1].Pos, &t.Face[2].Pos}
}
