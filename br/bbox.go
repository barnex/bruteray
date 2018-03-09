package br

func BoundBox(orig Obj, min, max Vec) Obj {
	return &bbox{orig, min, max}
}

type bbox struct {
	orig     Obj
	Min, Max Vec
}

func (s *bbox) Hit1(r *Ray, f *[]Fragment) {
	if !s.hit(r) {
		return
	}
	s.orig.Hit1(r, f)
}

func (s *bbox) hit(r *Ray) bool {
	min_ := s.Min
	max_ := s.Max

	tmin := min_.Sub(r.Start).Mul3(r.InvDir)
	tmax := max_.Sub(r.Start).Mul3(r.InvDir)

	txen := min(tmin[X], tmax[X])
	txex := max(tmin[X], tmax[X])

	tyen := min(tmin[Y], tmax[Y])
	tyex := max(tmin[Y], tmax[Y])

	tzen := min(tmin[Z], tmax[Z])
	tzex := max(tmin[Z], tmax[Z])

	ten := max3(txen, tyen, tzen)
	tex := min3(txex, tyex, tzex)

	return ten < tex
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func min3(x, y, z float64) float64 {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func max3(x, y, z float64) float64 {
	max := x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}
