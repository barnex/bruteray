package bruteray

type Light interface {
	Sample(e *Env, target Vec) (dir Vec, intens Color)
	Obj
}

// embed to get a Hit that returns no intersection.
type noIntersection struct{}

func (noIntersection) Hit(*Ray, *[]Surf) { return }
func (noIntersection) Inside(Vec) bool   { return false }

// Directed light source without fall-off.
// Position should be far away from the scene (indicates a direction)
func DirLight(pos Vec, intensity Color) Light {
	return &dirLight{pos: pos, c: intensity}
}

type dirLight struct {
	pos Vec
	c   Color
	noIntersection
}

func (l *dirLight) Sample(e *Env, target Vec) (Vec, Color) {
	return l.pos, l.c
}

// Point light source (with fall-off).
func PointLight(pos Vec, intensity Color) Light {
	return &pointLight{pos: pos, c: intensity}
}

type pointLight struct {
	pos Vec
	c   Color
	noIntersection
}

func (l *pointLight) Sample(e *Env, target Vec) (Vec, Color) {
	return l.pos, l.c.Mul((1 / (4 * Pi)) / target.Sub(l.pos).Len2()) // TODO: 1-> 4*pi
}

//func (l *pointLight) Inters2(*Ray) BiSurf {
//	return BiSurf{}
//}

// Spherical light source.
// Throws softer shadows than an point source and is visible in specular reflections.
//func SphereLight(pos Vec, radius float64, intensity Color) Light {
//	return &smoothLight{
//		sph: sphere{pos, radius},
//		r:   radius,
//		c:   intensity,
//		mat: Flat(intensity.Mul(1 / (4 * Pi * radius * radius))),
//	}
//}

type smoothLight struct {
	sph sphere
	r   float64
	c   Color
	mat Material
}

func (l *smoothLight) Sample(e *Env, target Vec) (Vec, Color) {
	pos := l.sph.c.MAdd(l.r, sphereVec(e))
	return pos, l.c.Mul((1 / (4 * Pi)) / target.Sub(pos).Len2())
}

//func (l *smoothLight) Inters2(r *Ray) BiSurf {
//	inter := l.sph.Inters2(r)
//	return BiSurf{
//		S1: Surf{inter.Min, r.At(inter.Min), l.mat},
//		S2: Surf{inter.Max, r.At(inter.Max), l.mat},
//	}
//}

func (l *smoothLight) Hit(r *Ray, s *[]Surf) {
	l.sph.Hit(r, s)
	for i := range *s {
		(*s)[i].Material = l.mat
	}
}

func sphereVec(e *Env) Vec {
	v := cubeVec(e)
	for v.Len2() > 1 {
		v = cubeVec(e)
	}
	return v
}

func cubeVec(e *Env) Vec {
	return Vec{
		e.rng.Float64() - 0.5,
		e.rng.Float64() - 0.5,
		e.rng.Float64() - 0.5,
	}
}
