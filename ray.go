package bruteray

// A Ray is a half-line,
// starting at the Start point (exclusive) and extending in direction Dir.
type Ray struct {
	Start Vec
	Dir   Vec
}

// Returns point Start + t*Dir.
// t must be > 0 for the point to lie on the Ray.
func (r *Ray) At(t float64) Vec {
	// TODO
	//if t <= 0 {
	//	panic(fmt.Sprintf("Ray.At: invalid t: %v", t))
	//}
	return r.Start.Add(r.Dir.Mul(t))
}

func (r *Ray) Transf(t *Matrix4) {
	r.Start = t.TransfPoint(r.Start)
	r.Dir = t.TransfDir(r.Dir)
}

//func ray(start, dir Vec) Ray {
//	return Ray{start, dir.Normalized()}
//}
