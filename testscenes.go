package main

// Two flat-shaded spheres, partially overlapping.
func Scene1() *Scene {
	const r = 0.25
	objects = []Obj{
		Flat(Sphere(Vec{-r / 2, 0, 3}, r), 1),
		Flat(Sphere(Vec{r / 2, 0, 3}, r), 0.5),
	}

	return &Scene{
		objs: objects,
	}
}

// A sphere behind the camera, should not be visible
func Scene2() *Scene {
	const r = 0.25
	objects = []Obj{
		Flat(Sphere(Vec{0, 0, -3}, r), 1),
	}

	return &Scene{
		objs: objects,
	}
}

// Intersection of flat-shaded spheres
func Scene3() *Scene {
	const r = 0.25
	s1 := Flat(Sphere(Vec{-r / 2, 0, 3}, r), 1)
	s2 := Flat(Sphere(Vec{r / 2, 0, 3}, r), 0.5)

	return &Scene{
		objs: []Obj{
			&ObjAnd{s1, s2},
		},
	}
}

// Intersection of spheres, as shapes (not objects)
func Scene4() *Scene {
	const r = 0.25
	s1 := Sphere(Vec{-r / 2, 0, 3}, r)
	s2 := Sphere(Vec{r / 2, 0, 3}, r)
	s := ShapeAnd{s1, s2}

	return &Scene{
		objs: []Obj{
			Flat(s, 1),
		},
	}
}

// Minus of spheres, as shapes (not objects)
func Scene5() *Scene {
	const r = 0.5
	s1 := Sphere(Vec{0, 0, 3}, r)
	s2 := Sphere(Vec{0, 0, 2 + r/2}, r)
	s := ShapeMinus{s1, s2}

	return &Scene{
		objs: []Obj{
			Flat(s, 1),
		},
	}
}

// Intersection of normal.z-shaded spheres
func Scene6() *Scene {
	const r = 0.25
	s1 := ShadeNormal(Sphere(Vec{-r / 2, 0, 3}, r))
	s2 := ShadeNormal(Sphere(Vec{r / 2, 0, 3}, r))

	return &Scene{
		objs: []Obj{
			s1,
			s2,
		},
	}
}
