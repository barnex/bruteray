package objects

import (
	"log"

	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/tracer/objects/ply"
	. "github.com/barnex/bruteray/tracer/types"
	"github.com/barnex/bruteray/util"
)

// Mesh returns a smooth triangle mesh with given vertices.
// The faces are defined by a list vertex indices.
//
// Normal vectors are calculated per vertex as the average
// over the faces sharing that vertex. The normals are
// also smoothly interpolated over inside each face.
//
// TODO: allow to (partially) turn smoothing off.
func Mesh(m Material, vertices []Vec, faceIdx [][3]int) Interface {
	return MeshWithUV(m, vertices, faceIdx, nil)
}

// MeshWithUV is like Mesh, but also attaches a (u,v) coordiante to each vertex.
// This allows the mesh to be textured.
//
// len(UV) must be equal to len(vertices) (exactly one UV coordinate per vertex).
func MeshWithUV(m Material, vertices []Vec, faceIdx [][3]int, UV []Vec2) Interface {
	// TODO: check slice lengths.

	// Convert positions to vertices
	Vertices := make([]vertex, len(vertices))
	for i, p := range vertices {
		Vertices[i].Pos = p
		if UV != nil {
			Vertices[i].U = UV[i][0]
			Vertices[i].V = UV[i][1]
		}
	}

	// Convert face indices to face structs.
	faces := make([]face, len(faceIdx))
	for i, idxs := range faceIdx {
		for c := 0; c < 3; c++ {
			faces[i][c] = &Vertices[idxs[c]]
		}
	}
	return mesh(m, faces)
}

// mesh calculates normal vectors and constructs a BHV tree containing the faces.
func mesh(m Material, faces []face) Interface {
	// Set each vertex's normal to the average normal of the faces sharing it.
	calcNormals(faces)

	// Convert faces to interface type so they can be use in a Tree.
	faceIf := make([]Interface, len(faces))
	for i := range faceIf {
		faceIf[i] = &faces[i]
	}

	// Return Tree containing the faces, wrapped with the desired material.
	return &withMaterial{
		orig: Tree(faceIf...),
		mat:  m,
	}
}

// Triangle constructs a triangle with vertices a, b, c.
// The orientation of the normal vector is determined by the right-hand rule.
// The respective vertices get UV coordinates
// 	(0,0), (1,0), (0,1)
func Triangle(m Material, a, b, c Vec) Interface {
	return MeshWithUV(m,
		[]Vec{a, b, c},
		[][3]int{{0, 1, 2}},
		[]Vec2{{0, 0}, {1, 0}, {0, 1}},
	)
}

// RectangleFromVertices constructs a rectangle with vertices o, a, b,
// and an automatically determined fourth vertex a+b-o.
// The vertices are ordered as shown below.
// The vertices' UV coordinates are shown in parentheses.
//
//      b *-------------* a+b-o (automatic)
// 	 (0,1)|             | (1,1)
//        |             |
// 	    o *-------------* a
//    (0,0)               (0,1)
//
// It is OK for the sides not to be perpendicular.
// In that case a parallelogram is retrurned.
func RectangleWithVertices(m Material, o, a, b Vec) Interface {
	c := a.Add(b.Sub(o))
	return Quadrilateral(m, o, a, c, b)
}

// Rectangle constructs a rectangle in the XZ (horizontal) plane
// with given width (along X), depth (along Z) and center.
func Rectangle(m Material, width, depth float64, center Vec) Interface {
	return RectangleWithVertices(m,
		Vec{-width / 2, 0, -depth / 2}.Add(center),
		Vec{+width / 2, 0, -depth / 2}.Add(center),
		Vec{-width / 2, 0, +depth / 2}.Add(center),
	)
}

// Quadrilateral constructs a general quadrilateral with given vertices
// ordered as shown below.
// The vertices' UV coordinates are shown in parentheses.
//
//      d *-------------* c
// 	 (0,1)|             | (1,1)
//        |             |
// 	    a *-------------* b
//    (0,0)               (0,1)
//
// TODO: currently, a quadrilateral is composed of two triagles which each do their own UV mapping.
// While this mapping is continuous (seamless), it does not map straight lines to straight lines:
// there is a "kink" when crossing from one triangle to another.
func Quadrilateral(m Material, a, b, c, d Vec) Interface {
	return MeshWithUV(m,
		[]Vec{a, b, c, d},
		[][3]int{{0, 1, 3}, {2, 3, 1}},
		[]Vec2{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
	)
}

// Parametric constructs a mesh approximating the parametric surface defined by
// 	x,y,z = f(u,v)
// where u, v are varied from 0 to 1 (inclusive) in numU, numV steps respectively.
//
// TODO: automatically make normals seamless.
func Parametric(m Material, numU, numV int, f func(u, v float64) Vec) Interface {
	// E.g.: numU: 3, numV: 2 => vertices: 2*3 = 6, faces: 2*((3-1)*(2-1)) = 4
	//  [0,0]***********[1,0]*************[2,0]
	//    *  *            *   *             *
	//    *      *        *       *         *
	//    *         *     *          *      *
	//    *            *  *              *  *
	//  [0,1]***********[1,1]*************[2,1]
	vertices := make([]vertex, numU*numV)
	faces := make([]face, 2*(numU-1)*(numV-1))
	I := func(iu, iv int) int {
		return iu*numV + iv
	}
	V := func(iu, iv int) *vertex {
		return &vertices[I(iu, iv)]
	}
	maxU := float64(numU - 1)
	maxV := float64(numV - 1)
	for iu := 0; iu < numU; iu++ {
		u := float64(iu) / maxU
		for iv := 0; iv < numV; iv++ {
			v := float64(iv) / maxV
			p := f(u, v)
			vertices[I(iu, iv)] = vertex{Pos: p, U: u, V: v}
			if iu < numU-1 && iv < numV-1 {
				faces[2*(iu*(numV-1)+iv)+0] = face{V(iu, iv), V(iu+1, iv), V(iu+1, iv+1)}
				faces[2*(iu*(numV-1)+iv)+1] = face{V(iu, iv), V(iu+1, iv+1), V(iu, iv+1)}
			}
		}
	}
	return mesh(m, faces)
}

// TODO: withUVmapping
func PlyFile(m Material, file string, transf ...*geom.AffineTransform) Interface {
	v, f, err := ply.ParseFile(file)
	if err != nil {
		log.Fatal(err)
	}
	applyTransform(geom.Compose(transf), v)
	return Mesh(m, v, f)
}

func applyTransform(t *geom.AffineTransform, v []Vec) {
	for i := range v {
		v[i] = t.TransformPoint(v[i])
	}
}

func WithMaterial(m Material, obj Interface) Interface {
	return &withMaterial{m, obj}
}

// withMaterial wraps an object with an other material.
// In case of a mesh, the faces do not individually store their material,
// we wrap it around the mesh in its entirety afterwards.
type withMaterial struct {
	mat  Material
	orig Interface
}

func (o *withMaterial) Intersect(r *Ray) HitRecord {
	f := o.orig.Intersect(r)
	f.Material = o.mat
	return f
}

func (o *withMaterial) Bounds() BoundingBox {
	return o.orig.Bounds()
}

func (o *withMaterial) Inside(p Vec) bool {
	return o.orig.Inside(p)
}

type face [3]*vertex

// TODO: store all in float32 precision
type vertex struct {
	Pos    Vec
	Normal Vec
	U, V   float64
}

func (f *face) Inside(Vec) bool { return false }

func (f *face) Intersect(r *Ray) HitRecord {
	//o := f[0].Pos
	ox := f[0].Pos[X]
	oy := f[0].Pos[Y]
	oz := f[0].Pos[Z]

	//a := f[1].Pos.Sub(o)
	ax := f[1].Pos[X] - ox
	ay := f[1].Pos[Y] - oy
	az := f[1].Pos[Z] - oz

	//b := f[2].Pos.Sub(o)
	bx := f[2].Pos[X] - ox
	by := f[2].Pos[Y] - oy
	bz := f[2].Pos[Z] - oz

	//n := a.Cross(b)
	nx := ay*bz - az*by
	ny := az*bx - ax*bz
	nz := ax*by - ay*bx

	//d := &r.Dir
	dx := r.Dir[X]
	dy := r.Dir[Y]
	dz := r.Dir[Z]

	// s := r.Start.Sub(o)
	sx := r.Start[X] - ox
	sy := r.Start[Y] - oy
	sz := r.Start[Z] - oz

	//nDotD := n.Dot(d)
	nDotD := nx*dx + ny*dy + nz*dz
	// nDotS := n.Dot(s)
	nDotS := nx*sx + ny*sy + nz*sz
	t := -nDotS / nDotD

	if !(t > 0) { // handles NaN gracefully
		return HitRecord{}
	}

	//p := r.At(t).Sub(o)
	px := sx + t*dx // Sub(o) already included in s
	py := sy + t*dy
	pz := sz + t*dz

	// Barycentric coordinates for 3D triangle, after
	// Peter Shirley, Fundamentals of Computer Graphics, 2nd Edition.

	//nc := a.Cross(p)
	ncx := ay*pz - az*py
	ncy := az*px - ax*pz
	ncz := ax*py - ay*px

	// p_a := p.Sub(a)
	p_ax := px - ax
	p_ay := py - ay
	p_az := pz - az

	// b_a := b.Sub(a)
	b_ax := bx - ax
	b_ay := by - ay
	b_az := bz - az

	//na := b.Sub(a).Cross(p.Sub(a))
	nax := b_ay*p_az - b_az*p_ay
	nay := b_az*p_ax - b_ax*p_az
	naz := b_ax*p_ay - b_ay*p_ax

	//n2 := n.Len2()
	n2 := nx*nx + ny*ny + nz*nz

	//l1 := n.Dot(na) / n2
	//l3 := n.Dot(nc) / n2
	l1 := (nx*nax + ny*nay + nz*naz) / n2
	l3 := (nx*ncx + ny*ncy + nz*ncz) / n2
	l2 := 1 - l1 - l3

	//if util.IsBad(l1) || util.IsBad(l2) {
	//	return HitRecord{}
	//}
	if l1 < 0 || l2 < 0 || l3 < 0 {
		return HitRecord{}
	}

	v1 := f.Vertex(0)
	v2 := f.Vertex(1)
	v3 := f.Vertex(2)
	shadingNormal := (v1.Normal.Mul(l1)).Add(v2.Normal.Mul(l2)).Add(v3.Normal.Mul(l3))
	//geomNormal := Vec{nx, ny, nz}
	// TODO: clamp shadingNormal

	u := v1.U*l1 + v2.U*l2 + v3.U*l3
	v := v1.V*l1 + v2.V*l2 + v3.V*l3
	return HitRecord{T: t, Normal: shadingNormal, Local: Vec{u, v, 0}}
}

func (f *face) Bounds() BoundingBox {
	return boundingBoxFromHull([]Vec{f[0].Pos, f[1].Pos, f[2].Pos})
}

func (f *face) Vertex(i int) *vertex {
	return f[i]
}

func (f *face) Normal() Vec {
	return geom.TriangleNormal(f.Vertex(0).Pos, f.Vertex(1).Pos, f.Vertex(2).Pos)
}

func calcNormals(f []face) {
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
			v.Normal.Normalize()
		}
	}
}
