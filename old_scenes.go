package main

import "math"

func spheresInARoom() *Env {
	scene := NewEnv()
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
	s := NewEnv()
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

func onedice() *Env {

	scene := NewEnv()
	scene.amb = func(dir Vec) Color { return Color(0.4 * dir.Y) }
	scene.Add(Sheet(-1, Ey), Diffuse2(0.5))
	//scene.Add(Sheet(25, Ex), Diffuse2(0.7))  // right wall
	scene.Add(Rect(Vec{-25, 0, 0}, Ex, 0, 10, 10), Diffuse2(0.7)) // left wall

	diec := Diffuse2(0.9)
	cube := &object{Box(Vec{0, 0, 0}, -1, -1, -1), diec}
	var die Obj = cube

	const r = 0.175
	pipc := Reflective(0.2)
	die = &objMinus{die, &object{Sphere(Vec{0, 0, -0.9}, r), pipc}}
	die = &objMinus{die, &object{Sphere(Vec{0.5, 0.5, -0.9}, r), pipc}}
	die = &objMinus{die, &object{Sphere(Vec{-0.5, 0.5, -0.9}, r), pipc}}
	die = &objMinus{die, &object{Sphere(Vec{0.5, -0.5, -0.9}, r), pipc}}
	die = &objMinus{die, &object{Sphere(Vec{-0.5, -0.5, -0.9}, r), pipc}}

	die = &objMinus{die, &object{Sphere(Vec{0.4, 1.1, -0.4}, r), pipc}}
	die = &objMinus{die, &object{Sphere(Vec{-0.4, 1.1, 0.4}, r), pipc}}
	die = &objMinus{die, &object{Sphere(Vec{0, 1.1, 0}, r), pipc}}

	die = &objAnd{die, &object{Sphere(Vec{}, 0.98*math.Sqrt(2.)), diec}}

	scene.objs = append(scene.objs, die)

	lp := Vec{4, 8, -6}
	scene.AddLight(SmoothLight(lp, 50, 1))
	hilight := &object{Sphere(lp.Mul(2), 4), Flat(2)}
	scene.objs = append(scene.objs, hilight)

	//cam := Camera(*focalLen)
	//cam.Transl(Vec{0, 4, -6})
	//cam.Transf(RotX(-15 * deg))
	//cam.AA = true
	return scene
}
