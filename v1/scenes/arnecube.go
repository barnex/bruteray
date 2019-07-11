// +build ignore

package main

import . "github.com/barnex/bruteray"

func main() {
	e := NewEnv()

	const (
		Sx = 10
		Sy = 4
		Sz = 6.5

		Rx = Sx / 2
		Ry = Sy / 2
		Rz = Sz / 2
		h  = 0.999
	)
	w := Diffuse1(Color{0.88, 0.88, 1}.EV(-.6))
	//w := Flat(WHITE.EV(-.3))

	e.Add(
		Sheet(Ey, -h-Sy, Diffuse1(WHITE.Mul(EV(-.6)))),

		Box(Vec{0, +Sy, +Sz}, Sx+h, h, h, w),
		Box(Vec{0, +Sy, -Sz}, Sx+h, h, h, w),
		Box(Vec{0, -Sy, +Sz}, Sx+h, h, h, w),
		Box(Vec{0, -Sy, -Sz}, Sx+h, h, h, w),

		Box(Vec{+Sx, 0, +Sz}, h, Sy+h, h, w),
		Box(Vec{+Sx, 0, -Sz}, h, Sy+h, h, w),
		Box(Vec{-Sx, 0, +Sz}, h, Sy+h, h, w),
		Box(Vec{-Sx, 0, -Sz}, h, Sy+h, h, w),

		Box(Vec{+Sx, +Sy, 0}, h, h, Sz+h, w),
		Box(Vec{+Sx, -Sy, 0}, h, h, Sz+h, w),
		Box(Vec{-Sx, +Sy, 0}, h, h, Sz+h, w),
		Box(Vec{-Sx, -Sy, 0}, h, h, Sz+h, w),
	)

	e.AddLight(
		SphereLight(Vec{18, 17, 5}.Mul(2), 16, WHITE.EV(15.3)),
		SphereLight(Vec{-6, 6, -25}.Mul(10), 160, WHITE.EV(18)),
	)
	e.SetAmbient(Flat(WHITE.Mul(EV(-6))))

	const (
		cx, cy, cz = -1.5, 8, -26
	)
	e.Camera = Camera(.9).Transl(cx, cy, cz).RotScene(45 * Deg).Transf(RotX4(25 * Deg))
	e.Camera.AA = true

	Serve(e)
}
