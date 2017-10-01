package bruteray

// -- quadric

//func Quad(center Vec, a Vec, b float64, m Material) Obj {
//	return &prim{&quad{center, a, b}, m}
//}
//
//type quad struct {
//	c Vec // unused
//	a Vec
//	b float64
//}
//
//func (s *quad) Inters(r *Ray) []Interval {
//	a0 := s.a[0]
//	a1 := s.a[1]
//	a2 := s.a[2]
//
//	s0 := r.Start[0]
//	s1 := r.Start[1]
//	s2 := r.Start[2]
//
//	d0 := r.Dir[0]
//	d1 := r.Dir[1]
//	d2 := r.Dir[2]
//
//	A := a0*d0*d0 + a1*d1*d1 + a2*d2*d2
//	B := 2 * (a0*d0*s0 + a1*d1*s1 + a2*d2*s2)
//	C := a0*s0*s0 + a1*s1*s1 + a2*s2*s2 - s.b
//
//	V := math.Sqrt(B*B - 4*A*C)
//
//	if math.IsNaN(V) {
//		return nil
//	}
//
//	t1 := (-B - V) / (2 * A)
//	t2 := (-B + V) / (2 * A)
//	t1, t2 = sort(t1, t2)
//	return []Interval{Interval{t1, t2}}
//}
//
//func (s *quad) Hit(r *Ray) float64 {
//	i := s.Inters(r)
//	if len(i) == 0 {
//		return 0
//	}
//	return i[0].Front()
//}
//
//func (s *quad) Normal(x Vec) Vec {
//	//return x.Normalized().check()
//	return s.a.Mul3(x).Normalized()
//}
//
////func(s*quad) Hit(r*Ray) float64{ }
