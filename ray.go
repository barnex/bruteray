package bruteray

type Ray struct {
	Start Vec
	Dir   Vec
}

func (r *Ray) At(t float64) Vec {
	return r.Start.Add(r.Dir.Mul(t))
}

func (r *Ray) Transf(t *Matrix4) {
	r.Start = t.TransfPoint(r.Start)
	r.Dir = t.TransfDir(r.Dir)
}

//func ray(start, dir Vec) Ray {
//	return Ray{start, dir.Normalized()}
//}
