package main

// Two flat-shaded spheres, partially overlapping.
func Scene1() *Scene {
	const r = 0.25
	objects = []Obj{
		&Shaded{&Sphere{C: Vec{-r / 2, 0, 3}, R: r}, 1},
		&Shaded{&Sphere{C: Vec{r / 2, 0, 3}, R: r}, 0.5},
	}

	return &Scene{
		objs: objects,
	}
}

// A sphere behind the camera, should not be visible
func Scene2() *Scene {
	const r = 0.25
	objects = []Obj{
		&Shaded{&Sphere{C: Vec{0, 0, -3}, R: r}, 1},
	}

	return &Scene{
		objs: objects,
	}
}

// Intersection of flat-shaded spheres
func Scene3() *Scene {
	const r = 0.25
	s1 := &Shaded{&Sphere{C: Vec{-r / 2, 0, 3}, R: r}, 1}
	s2 := &Shaded{&Sphere{C: Vec{r / 2, 0, 3}, R: r}, 0.5}

	return &Scene{
		objs: []Obj{
			&ObjAnd{s1, s2},
		},
	}
}

// Intersection of spheres, as shapes (not objects)
func Scene4() *Scene {
	const r = 0.25
	s1 := &Sphere{C: Vec{-r / 2, 0, 3}, R: r}
	s2 := &Sphere{C: Vec{r / 2, 0, 3}, R: r}
	s := ShapeAnd{s1, s2}

	return &Scene{
		objs: []Obj{
			&Shaded{s, 1},
		},
	}
}

// Minus of spheres, as shapes (not objects)
func Scene5() *Scene {
	const r = 0.5
	s1 := &Sphere{C: Vec{0, 0, 3}, R: r}
	s2 := &Sphere{C: Vec{0, 0, 2 + r/2}, R: r}
	s := ShapeMinus{s1, s2}

	return &Scene{
		objs: []Obj{
			&Shaded{s, 1},
		},
	}
}
