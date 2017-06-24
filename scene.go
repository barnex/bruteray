package main

import (
	"fmt"
	"math"
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
func (e *Env) AddLight(s Source) {
	e.sources = append(e.sources, s)
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
	if math.IsInf(minT, 0) {
		minT = 0
	}
	return minT, shader
}

//func (e *Env) HitAny(r *Ray) bool {
//	t, _ := e.Hit(r)
//	return t > 0 //obj != nil
//}

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

type object struct {
	Shape
	shader Shader
}

func (o *object) Shade(e *Env, r *Ray, t float64, N int) Color {
	if N == 0 {
		return e.Ambient(r.Dir)
	}
	n := o.Shape.Normal(r, t)
	return o.shader.Shade(e, r, t, n, N)
}

//type Shader interface{
//}
