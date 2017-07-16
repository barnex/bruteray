package main

func ray(start, dir Vec) Ray {
	return Ray{start, dir.Normalized()}
}

type Ray struct {
	Start Vec
	Dir   Vec
}

func (r *Ray) At(t float64) Vec {
	return r.Start.Add(r.Dir.Mul(t))
}
