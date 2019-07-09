package builder

import (
	"log"

	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
	"github.com/barnex/bruteray/v2/util"
)

type Mesh struct {
	tree     Tree
	Texture  Material
	faces    []Face
	vertices []Vertex
}

var _ Transformable = (*Mesh)(nil)

func PlyFile(m Material, file string) *Mesh {
	v, f, err := ParseFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return NewMesh(m, v, f)
}

func NewMesh(m Material, pos []Vec, faceIdx [][3]int) *Mesh {
	vertex := make([]Vertex, len(pos))
	for i, p := range pos {
		vertex[i].Pos = p
	}
	faces := make([]Face, len(faceIdx))
	for i, idxs := range faceIdx {
		for c := 0; c < 3; c++ {
			faces[i][c] = &vertex[idxs[c]]
		}
	}
	return &Mesh{Texture: m, faces: faces, vertices: vertex}
}

func (m *Mesh) Init() {
	f := m.faces
	calcNormals(f)
	for i := range f {
		m.tree.Add(&f[i])
	}
	m.tree.Init()
}

func (m *Mesh) Intersect(c *Ctx, r *Ray) HitRecord {
	f := m.tree.Intersect(c, r)
	f.Material = m.Texture
	return f
}

func (m *Mesh) CtrlPoints() []*Vec {
	c := make([]*Vec, len(m.vertices))
	for i := range c {
		c[i] = &m.vertices[i].Pos
	}
	return c
}

func (m *Mesh) Bounds() BoundingBox {
	return BoundingBoxFromHull(m.Hull())
}

// TODO: BoundingBox without copying:
// for all vertices: BB.Add(*p)
func (m *Mesh) Hull() []Vec {
	hull := make([]Vec, len(m.vertices))
	for i, v := range m.vertices {
		hull[i] = v.Pos
	}
	return hull
}

// TODO: unexport
type Face [3]*Vertex

// TODO: unexport
type Vertex struct {
	Pos    Vec
	Normal Vec
	U, V   float64
}

func (f *Face) Init() {}

func (f *Face) Intersect(c *Ctx, r *Ray) HitRecord {
	o := (*f[0]).Pos
	a := (*f[1]).Pos.Sub(o)
	b := (*f[2]).Pos.Sub(o)

	s := r.Start.Sub(o)
	d := r.Dir
	n := a.Cross(b)

	t := -n.Dot(s) / n.Dot(d)
	if t < 0 {
		return HitRecord{}
	}
	// TODO: use tricks like if !(t>0)... to handle NaNs
	if util.IsBad(t) {
		return HitRecord{}
	}

	p := r.At(t).Sub(o)
	x, y, _ := p[X], p[Y], p[Z]

	r2 := a
	r3 := b

	x2, y2, _ := r2[X], r2[Y], r2[Z]
	x3, y3, _ := r3[X], r3[Y], r3[Z]

	x1, y1, _ := 0., 0., 0.
	D := (y2-y3)*(x1-x3) + (x3-x2)*(y1-y3)

	l1 := ((y2-y3)*(x-x3) + (x3-x2)*(y-y3)) / D
	l2 := ((y3-y1)*(x-x3) + (x1-x3)*(y-y3)) / D
	l3 := 1 - l1 - l2

	if util.IsBad(l1) || util.IsBad(l2) {
		return HitRecord{}
	}

	if l1 < 0 || l2 < 0 || l3 < 0 {
		return HitRecord{}
	}

	v1 := f.Vertex(0)
	v2 := f.Vertex(1)
	v3 := f.Vertex(2)
	shadingNormal := (v1.Normal.Mul(l1)).Add(v2.Normal.Mul(l2)).Add(v3.Normal.Mul(l3))
	//_=shadingNormal
	//geomNormal := TriangleNormal(o,a,b)
	normal := shadingNormal

	u := v1.U*l1 + v2.U*l2 + v3.U*l3
	v := v1.V*l1 + v2.V*l2 + v3.V*l3
	return HitRecord{T: t, Normal: normal, Local: Vec{u, v, 0}}
}

func (f *Face) Bounds() BoundingBox {
	return BoundingBoxFromHull([]Vec{(*f[0]).Pos, (*f[1]).Pos, (*f[2]).Pos})
}

func (f *Face) Vertex(i int) *Vertex {
	return f[i]
}

func (f *Face) Normal() Vec {
	return TriangleNormal(f.Vertex(0).Pos, f.Vertex(1).Pos, f.Vertex(2).Pos)
}

func calcNormals(f []Face) {
	for _, f := range f {
		for i := range f {
			f.Vertex(i).Normal = Vec{}
		}
	}

	for _, f := range f {
		n := f.Normal()
		if util.IsBadVec(n) {
			continue
		}
		for i := range f {
			v := f.Vertex(i)
			v.Normal = v.Normal.Add(n) // TODO: .Towards, handling zero gracefully
		}
	}

	for _, f := range f {
		for i := range f {
			v := f.Vertex(i)
			v.Normal = v.Normal.Normalized()
		}
	}
}
