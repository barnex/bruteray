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

var (
	Width   = 1920
	Height  = 1080
	MaxIter = 500
)

func main() {

	cam := raster.Camera(1).Transl(0, 0.5, -3).Transf(RotX4(12 * Deg))
	floor := WHITE.EV(-1.3)
	red := RED.EV(-.3)
	blue := BLUE.EV(-.3)
	amb := Flat(WHITE.EV(-1))

	{
		e := env(
			Flat(floor),
			Flat(red),
			Flat(blue),
		)
		e.SetAmbient(amb)
		render(cam, e, 1, "01-flat.jpg")
	}
	{
		e := env(
			Diffuse00(floor),
			Diffuse00(RED),
			Diffuse00(blue),
		)
		e.AddLight(
			light.PointLight(lightPos, lightCol),
		)
		e.SetAmbient(amb)
		render(cam, e, 1, "02-light.jpg")
	}

	{
		e := env(
			Diffuse0(floor),
			Diffuse0(red),
			Diffuse0(blue),
		)
		e.AddLight(
			light.PointLight(lightPos, lightCol),
		)
		e.SetAmbient(amb)
		render(cam, e, 1, "03-shadow.jpg")
	}

	{
		e := env(
			Diffuse0(WHITE.EV(-1)),
			Diffuse0(red),
			Diffuse0(blue),
		)
		e.AddLight(
			light.Sphere(lightPos, 2, lightCol),
		)
		e.SetAmbient(amb)
		cam.AA = true
		for _, iter := range []int{1, 4, 16, 64, 256} {
			name := fmt.Sprintf("04-softshadow-%v.jpg", iter)
			render(cam, e, iter, name)
		}
	}

	// indirect lighting
	{
		e := env(
			Diffuse(WHITE.EV(-1)),
			Diffuse(red),
			Diffuse(blue),
		)
		cam.AA = true
		e.AddLight(
			light.Sphere(lightPos, 2, lightCol),
		)
		e.SetAmbient(amb)
		for _, rec := range []int{1, 2, 3, 4} {
			name := fmt.Sprintf("05-indirect-%v.jpg", rec)
			e.Recursion = rec
			render(cam, e, MaxIter, name)
		}
	}
	// reflection
	{
		e := env(
			Reflective(WHITE.EV(-3)),
			Diffuse(red),
			Diffuse(blue),
		)
		cam.AA = true
		e.AddLight(
			light.Sphere(lightPos, 2, lightCol),
		)
		e.SetAmbient(amb)
		render(cam, e, MaxIter, "07-reflect.jpg")
	}
	//// reflection2
	//{

	//	shapes = []func(Material) Obj{
	//		func(m Material) Obj { return shape.NewSheet(Ey, -1.0, m) },
	//		func(m Material) Obj { return shape.NewSphere(2.0, m).Transl(Vec{-1.1, 0, 1}) },
	//		func(m Material) Obj { return shape.NewSphere(2.0, m).Transl(Vec{1.1, 0, 1}) },
	//	}

	//	e := env(
	//		Checkboard(.8, Diffuse(WHITE.EV(-.6)), WHITE.EV(-.9)),
	//		Reflective(WHITE.EV(-1.2)),
	//		Reflective(WHITE.EV(-1.2)),
	//	)
	//	sky := MustLoad("../assets/sky1.jpg").Mul(2)
	//	e.SetAmbient(SkyDome(sky, -90*Deg))
	//	cam.AA = true
	//	e.AddLight(
	//	//light.Sphere(lightPos, 2, lightCol),
	//	)
	//	render(cam, e, MaxIter, "08-reflect.jpg")
	//}
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

func render(cam *raster.Cam, e *Env, passes int, name string) {
	fmt.Println(name)
	img := raster.MakeImage(Width, Height)
	raster.MultiPass(cam, e, img, passes)
	raster.Encode(img, name)
}
