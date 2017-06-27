package main

func Box(center Vec, rx, ry, rz float64) Shape {
	min := center.Sub(Vec{rx, ry, rz})
	max := center.Add(Vec{rx, ry, rz})
	s := &shape{
		inters: func(r *Ray) Inter {
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

			if ten < tex && ten > 0 {
				return Inter{ten, tex}
			}
			return empty
		}}

	s.normal = func(r *Ray, t float64) Vec {
		return NumNormal(s, r, t)
	}
	return s
}
