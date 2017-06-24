package main

type Shape interface {
	Hit(r *Ray) float64
	Normal(r *Ray, t float64) Vec
}

type shape struct {
	hit    func(r *Ray) float64
	normal func(r *Ray, t float64) Vec
}

func (s *shape) Hit(r *Ray) float64 {
	return s.hit(r)
}

func (s *shape) Normal(r *Ray, t float64) Vec {
	return s.normal(r, t)
}

func Sheet(pos float64, dir Vec) Shape {
	return &shape{
		hit: func(r *Ray) float64 {
			rs := r.Start.Dot(dir)
			rd := r.Dir.Dot(dir)
			t := (pos - rs) / rd
			return Max(t, 0)
		},
		normal: func(r *Ray, t float64) Vec {
			return dir.Towards(r.Dir)
		},
	}
}

func Rect(pos, dir Vec, rx, ry, rz float64) Shape {
	return &shape{
		hit: func(r *Ray) float64 {
			rs := r.Start.Dot(dir)
			rd := r.Dir.Dot(dir)
			t := (pos.Dot(dir) - rs) / rd
			if t < 0 {
				return 0
			}
			p := r.At(t).Sub(pos)
			if p.X < -rx || p.X > rx ||
				p.Y < -ry || p.Y > ry ||
				p.Z < -rz || p.Z > rz {
				return 0
			}
			return t
		},
		normal: func(r *Ray, t float64) Vec {
			return dir.Towards(r.Dir)
		},
	}
}

func Box(center Vec, rx, ry, rz float64) Shape {
	min := center.Sub(Vec{rx, ry, rz})
	max := center.Add(Vec{rx, ry, rz})
	s := &shape{
		hit: func(r *Ray) float64 {
			tmin := min.Sub(r.Start).Div3(r.Dir)
			tmax := max.Sub(r.Start).Div3(r.Dir)

			txen := Min(tmin.X, tmax.X)
			txex := Max(tmin.X, tmax.X)

			tyen := Min(tmin.Y, tmax.Y)
			tyex := Max(tmin.Y, tmax.Y)

			tzen := Min(tmin.Z, tmax.Z)
			tzex := Max(tmin.Z, tmax.Z)

			ten := Max3(txen, tyen, tzen)
			tex := Min3(txex, tyex, tzex)

			if ten < tex {
				return Max(0, ten)
			}
			return 0
		}}

	s.normal = func(r *Ray, t float64) Vec {
		return NumNormal(s, r, t)
	}
	return s
}

// Numerical approximation of normal vector
func NumNormal(s Shape, r *Ray, t float64) Vec {
	t0 := s.Hit(r)
	c := r.At(t0)

	perp1 := Vec{-r.Dir.Z, 0, r.Dir.X}.Normalized()
	perp2 := Vec{0, -r.Dir.Z, r.Dir.Y}.Normalized()

	const diff = 1. / (1024 * 1024)
	ra := r
	ra.Dir = r.Dir.MAdd(t*diff, perp1).Normalized()
	t1 := s.Hit(ra)
	a := ra.At(t1)

	rb := r
	rb.Dir = r.Dir.MAdd(t*diff, perp2).Normalized()
	t2 := s.Hit(rb)
	b := rb.At(t2)

	a = a.Sub(c)
	b = b.Sub(c)
	n := b.Cross(a).Normalized()

	return n.Towards(r.Dir)
}

func (n Vec) Towards(d Vec) Vec {
	if n.Dot(d) > 0 {
		return n.Mul(-1)
	}
	return n
}

//type ShapeAnd struct {
//	a, b Shape
//}
//
//func (s ShapeAnd) Intersect(r Ray) Inter {
//	a := s.a.Intersect(r)
//	if !a.OK() {
//		return a
//	}
//	b := s.b.Intersect(r)
//
//	return a.And(b)
//}
//
//func (s ShapeAnd) Normal(r Ray, t float64) Vec {
//	return Normal(s, r, t)
//}
//
//func (s ShapeAnd) Hit(r Ray) float64 {
//	return s.Intersect(r).Min
//}
//
//type ShapeMinus struct {
//	a, b Shape
//}
//
//func (s ShapeMinus) Intersect(r Ray) Inter {
//	a := s.a.Intersect(r)
//	if !a.OK() {
//		return a
//	}
//	b := s.b.Intersect(r)
//
//	return a.Minus(b)
//}
//
//func (s ShapeMinus) Normal(r Ray, t float64) Vec {
//	return Normal(s, r, t)
//}
//
//func (s ShapeMinus) Hit(r Ray) float64 {
//	return s.Intersect(r).Min
//}
