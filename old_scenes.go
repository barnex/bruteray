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
