package main

// Two flat-shaded spheres, partially overlapping.
func Scene1() *Scene {
	const r = 0.25
	objects = []Obj{
		&Sphere{C: Vec{-r / 2, 0, 3}, R: r, Color: 1},
		&Sphere{C: Vec{r / 2, 0, 3}, R: r, Color: 0.5},
	}

	return &Scene{
		objs: objects,
	}
}

// A sphere behind the camera, should not be visible
func Scene2() *Scene {
	const r = 0.25
	objects = []Obj{
		&Sphere{C: Vec{0, 0, -3}, R: r, Color: 1},
	}

	return &Scene{
		objs: objects,
	}
}

//func Scene3() {
//	const r = 0.25
//	sp := Sphere(Vec{0, 0, 3}, r)
//	objects = []*Obj{
//		{Shape: sp.Transl(r/2, 0, 0)},
//		{Shape: sp.Transl(-r/2, 0, 0)},
//	}
//}
//
//func Scene4() {
//	const r = 0.25
//	sp := Sphere(Vec{0, 0, 3}, r)
//	objects = []*Obj{
//		{Shape: And(sp.Transl(r/2, 0, 0), sp.Transl(-r/2, 0, 0))},
//	}
//}
