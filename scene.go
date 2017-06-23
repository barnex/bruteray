package main

import (
	"fmt"
	"math/rand"
)

type Env struct {
	objs    []Obj
	sources []Source
	amb     func(Vec) Color
	rand.Rand
}

func (e *Env) Add(s Shape, sh Shader) {
	e.objs = append(e.objs, &object{s, sh})
}

func (e *Env) Shade(r *Ray, N int) Color {
	if N == 0 {
		return e.Ambient(r.Dir)
	}
	t, obj := e.Hit(r)
	if obj != nil {
		return obj.Shade(e, r, t, N)
	} else {
		return e.Ambient(r.Dir)
	}
}

func (e *Env) Hit(r *Ray) (float64, Obj) {
	var (
		minT   float64 = inf
		shader Obj     = nil
	)

	for i, o := range e.objs {
		t := o.Hit(r)
		if t < 0 {
			panic(fmt.Sprintf("object %v: %#v: t=%v", i, o, t))
		}
		if t < minT && t > 0 {
			minT = t
			shader = o
		}
	}
	return minT, shader
}

func (s *Env) Ambient(dir Vec) Color {
	if s.amb == nil {
		return 0
	}
	return s.amb(dir)
}

type Obj interface {
	Shape
	Shade(e *Env, r *Ray, t float64, N int) Color
}

type Shape interface {
	Hit(r *Ray) float64
	Normal(r *Ray, t float64) Vec
}

type shape struct {
	hit    func(r *Ray) float64
	normal func(r *Ray, t float64) Vec
}

func (s *shape) Hit(r *Ray) float64 {
	return s.hit(r)
}

func (s *shape) Normal(r *Ray, t float64) Vec {
	return s.normal(r, t)
}

func Sheet(pos float64, dir Vec) Shape {
	return &shape{
		hit: func(r *Ray) float64 {
			rs := r.Start.Dot(dir)
			rd := r.Dir.Dot(dir)
			t := (pos - rs) / rd
			return Max(t, 0)
		},
		normal: func(r *Ray, t float64) Vec {
			return dir
		},
	}
}

type object struct {
	Shape
	shader Shader
}

func (o *object) Shade(e *Env, r *Ray, t float64, N int) Color {
	if N == 0 {
		return e.Ambient(r.Dir)
	}
	n := o.Shape.Normal(r, t)
	return o.shader.Shade(e, r, t, n)
}

type Shader interface {
	Shade(e *Env, r *Ray, t float64, n Vec) Color
}

func Flat(v Color) Shader {
	return shadeFn(func(*Env, *Ray, float64, Vec) Color {
		return v
	})
}

type shadeFn func(e *Env, r *Ray, t float64, n Vec) Color

func (f shadeFn) Shade(e *Env, r *Ray, t float64, n Vec) Color {
	return f(e, r, t, n)
}

//type Shader interface{
//}
