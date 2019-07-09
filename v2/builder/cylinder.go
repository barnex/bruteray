package builder

/*
import (
	"math"

	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
)

type Cylinder struct {
	Frame
	Texture Material
	//r2      float64
}

func NewCylinder(m Material, diam, height float64) *Cylinder {
	r := diam / 2
	return &Cylinder{
		Frame:   Frame{O, Vec{r, 0, 0}, Vec{0, r, 0}, Vec{0, 0, height / 2}},
		Texture: m,
	}
}

var _ Builder = (*Cylinder)(nil)

func (s *Cylinder) Init() {}

func (s *Cylinder) BoundingBox() BoundingBox {
	return BoundingBoxFromHull(s.Frame.MaxHull())
}

func (s *Cylinder) Intersect(ctx *Ctx, r *Ray) (f HitRecord) {
	v := r.Start.Sub(s.Origin())
	d := r.Dir
	vd := v.Dot(d)
	D := vd*vd - (v.Len2() - s.r2)
	if D < 0 {
		return HitRecord{}
	}
	sqrtD := math.Sqrt(D)
	t1 := (-vd - sqrtD)
	t2 := (-vd + sqrtD)
	t := frontSolution(t1, t2, r.Len)
	n := r.At(t).Sub(s.Origin())
	if t > 0 {
		return HitRecord{T: t, Normal: n, Material: s.Texture}
	}
	return HitRecord{}
}

func (s *Sphere) Inside(p Vec) bool {
	v := p.Sub(s.Origin())
	return v.Len2() < s.r2
}
*/
