// Package tutorial explains some ray-tracing basics.
package main

import (
	"fmt"

	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/light"
	. "github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
	"github.com/barnex/bruteray/shape"
)

func main() {

	cam := raster.Camera(1).Transl(0, 0, -3)

	//flat
	{
		e := env(
			Flat(WHITE.EV(-1)),
			Flat(RED),
			Flat(BLUE),
		)
		e.SetAmbient(Flat(WHITE))
		render1(cam, e)
	}

	// light
	{
		e := env(
			Diffuse00(WHITE.EV(-1)),
			Diffuse00(RED),
			Diffuse00(BLUE),
		)
		e.AddLight(
			light.PointLight(lightPos, lightCol),
		)
		e.SetAmbient(Flat(WHITE))
		render1(cam, e)
	}

	// shadow
	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		e.AddLight(
			light.PointLight(lightPos, lightCol),
		)
		e.SetAmbient(Flat(WHITE))
		render1(cam, e)
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		e.AddLight(
			light.Sphere(lightPos, 2, lightCol),
		)
		e.SetAmbient(Flat(WHITE))
		render1(cam, e)
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		e.AddLight(
			light.Sphere(lightPos, 2, lightCol),
		)
		e.SetAmbient(Flat(WHITE))
		render2(cam, e, 2)
	}

	// soft shadow
	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		cam.AA = true
		e.AddLight(
			light.Sphere(lightPos, 2, lightCol),
		)
		e.SetAmbient(Flat(WHITE))
		for _, n := range []int{1, 2, 3, 10} {
			render2(cam, e, n)
		}
	}

	// indirect lighting
	{
		e := env(
			Diffuse(WHITE.EV(-1)),
			Diffuse(RED),
			Diffuse(BLUE),
		)
		cam.AA = true
		e.AddLight(
			light.Sphere(lightPos, 2, lightCol),
		)
		e.SetAmbient(Flat(WHITE))
		e.Recursion = 1
		render2(cam, e, 10)
	}

	{
		e := env(
			Shiny(WHITE.EV(-1), .1),
			Shiny(RED, .1),
			Shiny(BLUE, .1),
		)
		cam.AA = true
		e.AddLight(
			light.Sphere(lightPos, 2, lightCol),
		)
		e.SetAmbient(Flat(WHITE))
		render2(cam, e, 10)
	}
}

var (
	shapes = []func(Material) Obj{
		func(m Material) Obj { return shape.NewSheet(Ey, -1.0, m) },
		func(m Material) Obj { return shape.NewSphere(2.0, m) },
		func(m Material) Obj { return shape.NewSphere(2.0, m).Transl(Vec{1.5, 0, 1}) },
	}
	lightPos = Vec{4, 5, -1}
	lightCol = WHITE.Mul(EV(8))
)

func env(m ...Material) *Env {
	e := NewEnv()
	for i, m := range m {
		e.Add(shapes[i](m))
	}
	return e
}

var cnt = 0

func render1(cam *raster.Cam, e *Env) {
	cnt++
	name := fmt.Sprintf("rt%02d.jpg", cnt)
	fmt.Println(name)
	img := raster.MakeImage(1920, 1080)
	raster.SinglePass(cam, e, img)
	raster.Encode(img, name)
}

func render2(cam *raster.Cam, e *Env, passes int) {
	cnt++
	name := fmt.Sprintf("rt%02d.jpg", cnt)
	fmt.Println(name)
	img := raster.MakeImage(1920, 1080)
	raster.MultiPass(cam, e, img, passes)
	raster.Encode(img, name)
}
