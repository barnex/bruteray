package main

//func ABox(min, max Vec) Shape {
//	return ShapeFunc(func(r Ray) Inter {
//		tmin := min.Sub(r.Start).Div3(r.Dir)
//		tmax := max.Sub(r.Start).Div3(r.Dir)
//
//		txen := Min(tmin.X, tmax.X)
//		txex := Max(tmin.X, tmax.X)
//
//		tyen := Min(tmin.Y, tmax.Y)
//		tyex := Max(tmin.Y, tmax.Y)
//
//		tzen := Min(tmin.Z, tmax.Z)
//		tzex := Max(tmin.Z, tmax.Z)
//
//		ten := Max3(txen, tyen, tzen)
//		tex := Min3(txex, tyex, tzex)
//
//		if ten < tex {
//			return Inter{ten, tex}
//		}
//		return empty
//	})
//}
