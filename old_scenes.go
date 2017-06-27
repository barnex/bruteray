package main

func spheresInARoom() *Env {
	scene := &Env{}
	scene.amb = func(v Vec) Color { return Color(0.2*v.Y + 0.2) }
	scene.Add(Sheet(-3, Ey), Diffuse2(0.5))  // floor
	scene.Add(Sheet(8, Ey), Diffuse2(0.5))   // ceiling
	scene.Add(Sheet(20, Ez), Diffuse2(0.8))  // back
	scene.Add(Sheet(-5, Ez), Diffuse2(0.8))  // front
	scene.Add(Sheet(10, Ex), Diffuse2(0.8))  // left
	scene.Add(Sheet(-10, Ex), Diffuse2(0.8)) // right
	scene.Add(Sphere(Vec{1, -2, 8}, 1), Reflective(0.3))
	scene.Add(Sphere(Vec{-1, -2, 6}, 1), Diffuse2(0.95))
	scene.AddLight(SmoothLight(Vec{0, 7, 1}, 100, 2))
	return scene
}

func checkboard() *Env {
	s := &Env{}
	s.amb = func(dir Vec) Color { return 0.5 }

	s.Add(Sheet(0, Ey), Diffuse2(0.7))                                                      // floor
	s.Add(Box(Vec{0, 0, 8}, 5, 0.45, 5), Diffuse2(0.1))                                     // base
	s.Add(Rect(Vec{0, 0.5, 8}, Ey, 4, inf, 4), CheckBoard(Reflective(0.05), Diffuse2(0.9))) // checkboard

	// walls
	s.Add(Sheet(20, Ez), Diffuse2(0.7)) // back
	s.Add(Sheet(20, Ey), Diffuse2(0.6)) // ceiling

	slab := Slab(0, 0.7)

	for j := 0.; j < 2; j++ {
		for i := j; i < 8; i += 2 {
			cyl := Cylinder(Vec{j - 3.5, -1.5, i + 4.5}, 0.37)
			s.Add(ShapeAnd(cyl, slab), Reflective(0.05))

			cyl = Cylinder(Vec{j + 2.5, -1.5, i + 4.5}, 0.37)
			s.Add(ShapeAnd(cyl, slab), Diffuse2(0.95))
		}
	}

	s.AddLight(SmoothLight(Vec{3, 12, 6}, 130, 2))

	return s
}
