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

// Intersection of flat-shaded spheres
func Scene3() *Scene {
	const r = 0.25
	s1 := &Sphere{C: Vec{-r / 2, 0, 3}, R: r, Color: 1}
	s2 := &Sphere{C: Vec{r / 2, 0, 3}, R: r, Color: 0.5}

	return &Scene{
		objs: []Obj{
			&And{s1, s2},
		},
	}
}
