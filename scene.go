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
	e.objs = append(e.objs, object{s, sh})
}

func (e *Env) Shade(r *Ray, N int) Color {
	if N == 0 {
		return e.Ambient(r)
	}
	t, obj := e.Hit(r)
	if obj != nil {
		return obj.Shade(e, r, t, N)
	} else {
		return e.Ambient(r)
	}
}

func (e *Env) Hit(r *Ray) (float64, Obj) {
	var (
		minT   float64 = 0
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

//func (e *Env) HitsAny(r *Ray) bool {
//	_, obj := e.Hit(r)
//	return obj != nil
//}

func (s *Env) Ambient(r *Ray) Color {
	if s.amb == nil {
		return 0
	}
	return s.amb(r.Dir)
}

type Obj interface {
	Shape
	Shade(e *Env, r *Ray, t float64, N int) Color
}

type Shape interface {
	Hit(r *Ray) float64
}

//type Shader interface{
//}
