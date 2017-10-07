// +build ignore

package main

import . "github.com/barnex/bruteray"
import "github.com/barnex/bruteray/serve"

func main() {
	e := NewEnv()

	w := Diffuse(WHITE.EV(-.6))

	e.Add(Sheet(Ey, -1, w))

	e.Add(SurfaceAnd(
		Quad(Vec{}, Vec{1, 0, 1}, .75, w),
		Slab(Ey, -1, 1.3, w),
	))
	e.Add(SurfaceAnd(
		Quad(Vec{}, Vec{1, 0, 1}, 1.5, w),
		Slab(Ey, -1, 1.3, w),
	))
	e.Add(SurfaceAnd(
		Cube(Vec{2, 0, -2}, 1, w),
		Slab(Ey, -1, .9, w),
	))
	e.Add(SurfaceAnd(
		Cube(Vec{2, .3, -2}, .75, w),
		Slab(Ey, -2, .9, w),
	))
	e.Add(SurfaceAnd(
		Sphere(Vec{-2, 0.5, -2.5}, 1.5, w),
		Slab(Ey, -2, .9, w),
	))
	e.Add(SurfaceAnd(
		Sphere(Vec{-2, 0.7, -2.5}, 1.2, w),
		Slab(Ey, -1, 1, w),
	))

	e.AddLight(
		SphereLight(Vec{16, 16, -8}.Mul(2), 16, WHITE.EV(15.6)),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-2))))

	e.Camera = Camera(1).Transl(-.8, 5, -5).RotScene(10 * Deg).Transf(RotX4(55 * Deg))
	e.Camera.AA = true

	serve.Env(e)
}
