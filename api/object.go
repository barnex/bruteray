package api

import (
	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/tracer/objects"
)

type Object struct {
	objects.Interface
}

func Backdrop(m Material) Object {
	return Object{objects.Backdrop(m)}
}

func Box(m Material, dx, dy, dz float64, center Vec) Object {
	return Object{objects.Box(m, dx, dy, dz, center)}
}

func Cylinder(m Material, diam, height float64, center Vec) Object {
	return Object{objects.Cylinder(m, diam, height, center)}
}

func CylinderWithCaps(m Material, diam, height float64, center Vec) Object {
	return Object{objects.CylinderWithCaps(m, diam, height, center)}
}

func Parametric(m Material, numU, numV int, f func(u, v float64) Vec) Object {
	return Object{objects.Parametric(m, numU, numV, f)}
}

func PlyFile(m Material, file string, transf ...*geom.AffineTransform) Object {
	return Object{objects.PlyFile(m, file, transf...)}
}

func ObjFile(m map[string]Material, file string, transf ...*geom.AffineTransform) Object {
	return Object{objects.ObjFile(m, file, transf...)}
}

func RectangleWithVertices(m Material, o, a, b Vec) Object {
	return Object{objects.RectangleWithVertices(m, o, a, b)}
}

func Rectangle(m Material, dx, dz float64, center Vec) Object {
	return Object{objects.Rectangle(m, dx, dz, center)}
}

func Sphere(m Material, diam float64, center Vec) Object {
	return Object{objects.Sphere(m, diam, center)}
}

func Tree(children ...Object) Object {
	ch := make([]objects.Interface, len(children))
	for i := range children {
		ch[i] = children[i]
	}
	return Object{objects.Tree(ch...)}
}

func IsoSurface(m Material, dx, dy, dz float64, f func(u, v float64) float64) Object {
	return Object{objects.IsoSurface(m, dx, dy, dz, f)}
}

func Difference(a, b Object) Object {
	return Object{objects.Difference(a.Interface, b.Interface)}
}

func (o Object) Center() Vec {
	return o.Bounds().Center()
}

func (o Object) WithCenter(pos Vec) Object {
	delta := pos.Sub(o.Bounds().Center())
	return o.Translate(delta)
}

func (o Object) WithCenterBottom(pos Vec) Object {
	delta := pos.Sub(o.Bounds().CenterBottom())
	return o.Translate(delta)
}

func (o Object) Rotate(axis Vec, radians float64) Object {
	c := o.Bounds().Center()
	return o.Transform(geom.Rotate(c, axis, radians))
}

func (o Object) Translate(delta Vec) Object {
	return o.Transform(geom.Translate(delta))
}

func (o Object) Scale(s float64) Object {
	return o.Transform(geom.Scale(o.Center(), s))
}

func (o Object) Transform(tr *geom.AffineTransform) Object {
	return Object{objects.Transformed(o.Interface, tr)}
}

func (o Object) WithMaterial(m Material) Object {
	return Object{objects.WithMaterial(m, o.Interface)}
}

func (o Object) ScaleToSize(maxSize float64) Object {
	s := o.Bounds().Size()[0]
	return o.Scale(maxSize / s)
}
