package r

type Ray struct {
	Start Vec
	Dir   Vec
}

func ray(start, dir Vec) Ray {
	return Ray{start, dir.Normalized()}
}

func (r *Ray) At(t float64) Vec {
	return r.Start.Add(r.Dir.Mul(t))
}
