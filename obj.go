package bruteray

// Object has a shape and material (i.e. "color"),
// and can be rendered.
type Obj interface {
	Hit(r *Ray) Surf // Returns the surface fragment where the ray first hits (T>0)
}

// -- Primitive Object (shape + material)
type prim struct {
	s SolidShape
	m Material
}

var _ Obj = (*prim)(nil)

func (o *prim) Inters(r *Ray) []BiSurf {

	i := o.s.Inters(r)
	var bi []BiSurf

	for _, i := range i {
		n1 := o.s.Normal(r.At(i.Min))
		n2 := o.s.Normal(r.At(i.Max))
		bi = append(bi, BiSurf{
			S1: Surf{T: i.Min, Norm: n1, Material: o.m},
			S2: Surf{T: i.Max, Norm: n2, Material: o.m},
		})
	}
	return bi

}

func (o *prim) Hit(r *Ray) Surf {
	t := o.s.Hit(r)
	if t <= 0 {
		return Surf{}
	}
	return Surf{T: t, Norm: o.s.Normal(r.At(t)), Material: o.m}
}
