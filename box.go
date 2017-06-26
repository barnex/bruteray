package main

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
