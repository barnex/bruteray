// Package tutorial explains some ray-tracing basics.
package main

import (
	"fmt"

	. "github.com/barnex/bruteray/br"
	. "github.com/barnex/bruteray/light"
	. "github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
	"github.com/barnex/bruteray/shape"
)

func main() {

	{
		e := env(
			Flat(BLACK),
			Flat(RED),
		)
		cam := raster.Camera(1).Transl(0, 0, -3)
		render1(cam, e)
	}

	{
		e := env(
			Flat(WHITE.EV(-1)),
			Flat(RED),
			Flat(BLUE),
		)
		cam := raster.Camera(1).Transl(0, 0, -3)
		render1(cam, e)
	}

	{
		e := env(
			Diffuse00(WHITE.EV(-1)),
			Diffuse00(RED),
			Diffuse00(BLUE),
		)
		cam := raster.Camera(1).Transl(0, 0, -3)
		e.AddLight(
			PointLight(lightPos, lightCol),
		)
		render1(cam, e)
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		cam := raster.Camera(1).Transl(0, 0, -3)
		e.AddLight(
			PointLight(lightPos, lightCol),
		)
		render1(cam, e)
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		cam := raster.Camera(1).Transl(0, 0, -3)
		e.AddLight(
			Sphere(lightPos, 2, lightCol),
		)
		render1(cam, e)
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		cam := raster.Camera(1).Transl(0, 0, -3)
		e.AddLight(
			Sphere(lightPos, 2, lightCol),
		)
		render2(cam, e)
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(RED),
			Diffuse0(BLUE),
		)
		cam := raster.Camera(1).Transl(0, 0, -3)
		cam.AA = true
		e.AddLight(
			Sphere(lightPos, 2, lightCol),
		)
		render2(cam, e)
	}

	{
		e := env(
			Diffuse(WHITE.EV(-1)),
			Diffuse(RED),
			Diffuse(BLUE),
		)
		cam := raster.Camera(1).Transl(0, 0, -3)
		cam.AA = true
		e.AddLight(
			Sphere(lightPos, 2, lightCol),
		)
		render2(cam, e)
	}

	{
		e := env(
			Diffuse(WHITE.EV(-1)),
			Diffuse(RED),
			Diffuse(BLUE),
		)
		cam := raster.Camera(1).Transl(0, 0, -3)
		cam.AA = true
		e.AddLight(
			Sphere(lightPos, 2, lightCol),
		)
		e.SetAmbient(Flat(WHITE))
		render2(cam, e)
	}

	{
		e := env(
			Reflective(WHITE.EV(-4)),
			Diffuse(RED),
			Diffuse(BLUE),
		)
		cam := raster.Camera(1).Transl(0, 0, -3)
		cam.AA = true
		e.AddLight(
			SphereLight(lightPos, 2, lightCol),
		)
		e.SetAmbient(Flat(WHITE))
		render2(cam, e)
	}
}

var (
	shapes = []func(Material) Obj{
		func(m Material) Obj { return shape.Sheet(Ey, -1.0, m) },
		func(m Material) Obj { return shape.Sphere(Vec{}, 1.0, m) },
		func(m Material) Obj { return shape.Sphere(Vec{1.5, 0, 1}, 1.0, m) },
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
	img := raster.MakeImage(600, 400)
	raster.SinglePass(cam, e, img)
	raster.Encode(img, name)
}

func render2(cam *raster.Cam, e *Env) {
	cnt++
	name := fmt.Sprintf("rt%02d.jpg", cnt)
	fmt.Println(name)
	img := raster.MakeImage(600, 400)
	raster.MultiPass(cam, e, img, 300)
	raster.Encode(img, name)
}
