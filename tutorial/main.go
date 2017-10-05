package main

import (
	"fmt"

	. "github.com/barnex/bruteray"
)

func main() {

	{
		e := env(
			Flat(BLACK),
			Flat(RED),
		)
		e.Camera.AA = false
		e.Camera = Camera(1).Transl(0, 0, -3)
		render1(e)
	}

	{
		e := env(
			Flat(WHITE.EV(-1)),
			Flat(RED),
			Flat(BLUE),
		)
		e.Camera.AA = false
		e.Camera = Camera(1).Transl(0, 0, -3)
		render1(e)
	}

	{
		e := env(
			Diffuse00(WHITE.EV(-1)),
			Diffuse00(RED),
			Diffuse00(BLUE),
		)
		e.Camera.AA = false
		e.Camera = Camera(1).Transl(0, 0, -3)
		e.AddLight(
			PointLight(lightPos, lightCol),
		)
		render1(e)
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		e.Camera.AA = false
		e.Camera = Camera(1).Transl(0, 0, -3)
		e.AddLight(
			PointLight(lightPos, lightCol),
		)
		render1(e)
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		e.Camera.AA = false
		e.Camera = Camera(1).Transl(0, 0, -3)
		e.AddLight(
			SphereLight(lightPos, 2, lightCol),
		)
		render1(e)
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		e.Camera.AA = false
		e.Camera = Camera(1).Transl(0, 0, -3)
		e.AddLight(
			SphereLight(lightPos, 2, lightCol),
		)
		render2(e)
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		e.Camera = Camera(1).Transl(0, 0, -3)
		e.Camera.AA = true
		e.AddLight(
			SphereLight(lightPos, 2, lightCol),
		)
		render2(e)
	}

	{
		e := env(
			Diffuse1(WHITE.EV(-1)),
			Diffuse1(RED),
			Diffuse1(BLUE),
		)
		e.Camera = Camera(1).Transl(0, 0, -3)
		e.Camera.AA = true
		e.AddLight(
			SphereLight(lightPos, 2, lightCol),
		)
		render2(e)
	}

	{
		e := env(
			Diffuse1(WHITE.EV(-1)),
			Diffuse1(RED),
			Diffuse1(BLUE),
		)
		e.Camera = Camera(1).Transl(0, 0, -3)
		e.Camera.AA = true
		e.AddLight(
			SphereLight(lightPos, 2, lightCol),
		)
		e.SetAmbient(Flat(WHITE))
		render2(e)
	}

}

var (
	shapes = []func(Material) Obj{
		func(m Material) Obj { return Sheet(Ey, -1.0, m) },
		func(m Material) Obj { return Sphere(Vec{}, 1.0, m) },
		func(m Material) Obj { return Sphere(Vec{1.5, 0, 1}, 1.0, m) },
	}
	lightPos = Vec{4, 5, -1}
	lightCol = WHITE.Mul(EV(8))
)

func env(m ...Material) *Env {
	e := NewEnv()
	for i, m := range m {
		e.Add(shapes[i](m))
	}
	e.Camera = Camera(1).Transl(0, 1, -3)
	e.Camera.AA = true
	return e
}

var cnt = 0

func render1(e *Env) {
	cnt++
	name := fmt.Sprintf("rt%02d.jpg", cnt)
	img := MakeImage(600, 400)
	Render(e, img)
	Encode(img, name)
}

func render2(e *Env) {
	cnt++
	name := fmt.Sprintf("rt%02d.jpg", cnt)
	img := MakeImage(600, 400)
	MultiPass(e, img, 300)
	Encode(img, name)
}
