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
	seed int64
}

func NewEnv() *Env {
	return &Env{
		amb:  func(Vec) Color { return 0 },
		seed: 1,
		Rand: *rand.New(rand.NewSource(1)),
	}
}

func (e *Env) AddObj(o Obj) {
	e.objs = append(e.objs, o)
}

func (e *Env) Add(s Shape, sh Shader) {
	e.AddObj(&object{s, sh})
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
		ival := o.Inters(r)
		if !ival.OK() || ival.Min < 0 {
			continue
		}
		t := ival.Min
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

func (s *Env) Ambient(dir Vec) Color {
	return s.amb(dir)
}
