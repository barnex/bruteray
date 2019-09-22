package test

import (
	"math"

	"github.com/barnex/bruteray/geom"
	. "github.com/barnex/bruteray/tracer/types"
)

func Sheet(m Material, height float64) Object {
	return &sheet{mat: m, height: height}
}

type sheet struct {
	height float64
	mat    Material
}

func (s *sheet) Intersect(r *Ray) HitRecord {
	origin := Vec{0, s.height, 0}
	rs := r.Start.Sub(origin)[1]
	rd := r.Dir[1]
	t := -rs / rd
	if t < 0 || math.IsInf(t, 1) {
		return HitRecord{}
	}
	p := r.At(t)
	return HitRecord{T: t, Normal: Vec{0, 1, 0}, Material: s.mat, Local: Vec{p[0], p[2], 0}}
}

// Sphere returns a minimal implementation of a sphere.
// Intended for tests that should not depend on package objects/
func Sphere(m Material, diam float64, origin geom.Vec) Object {
	return &sphere{
		origin: origin,
		mat:    m,
		r:      diam / 2,
	}
}

type sphere struct {
	origin geom.Vec
	r      float64
	mat    Material
}

func (s *sphere) Intersect(r *Ray) HitRecord {
	checkRay(r)
	r2 := s.r * s.r
	v := r.Start.Sub(s.origin)
	d := r.Dir
	vd := v.Dot(d)
	D := vd*vd - (v.Len2() - r2)
	if D < 0 {
		return HitRecord{}
	}
	sqrtD := math.Sqrt(D)
	t1 := (-vd - sqrtD)
	t2 := (-vd + sqrtD)
	t := t1
	if t < 0 {
		t = t2
	}
	if t > 0 {
		n := r.At(t).Sub(s.origin)
		return HitRecord{T: t, Normal: n, Material: s.mat, Local: r.At(t).Sub(s.origin)}
	}
	return HitRecord{}
}

type pointLight struct {
	origin Vec
	sphere sphere
}

func PointLight(pos Vec) Light {
	return &pointLight{
		origin: pos,
		sphere: sphere{
			origin: pos,
			mat:    Flat{1, 1, 0},
			r:      0.1,
		},
	}
}

func (p *pointLight) Object() Object {
	return &p.sphere
}

func (p *pointLight) Sample(_ *Ctx, dst Vec) (Vec, Color) {
	return p.origin.MAdd(-0.101, p.origin.Sub(dst).Normalized()), Color{1, 1, 1}
}
