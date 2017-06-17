package main

// Two flat-shaded spheres, partially overlapping.
func Scene1() *Scene {
	const r = 0.25
	objects = []Obj{
		&Sphere{C: Vec{-r / 2, 0, 3}, R: r, Col: 1},
		&Sphere{C: Vec{r / 2, 0, 3}, R: r, Col: 0.5},
	}

	return &Scene{
		objs: objects,
	}
}
