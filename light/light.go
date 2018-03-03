// Package light implements various types of light sources.
// They all implement br.Light.
package light

import (
	"math"
	"math/rand"

	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/shape"
)

type noIntersection struct{}

func (noIntersection) Hit1(*Ray, *[]Fragment) { return }

// Directed light source without fall-off.
// Position should be far away from the scene (indicates a direction)
func DirLight(pos Vec, intensity Color) Light {
	return &dirLight{pos: pos, c: intensity}
}

type dirLight struct {
	pos Vec
	c   Color
	noIntersection
}

func (l *dirLight) Sample(ctx *Ctx, target Vec) (Vec, Color) {
	return l.pos, l.c
}

// Point light source (with fall-off).
func PointLight(pos Vec, intensity Color) Light {
	return &pointLight{pos: pos, c: intensity}
}

type pointLight struct {
	pos Vec
	c   Color
	noIntersection
}

func (l *pointLight) Sample(ctx *Ctx, target Vec) (Vec, Color) {
	return l.pos, l.c.Mul((1 / (4 * Pi)) / target.Sub(l.pos).Len2())
}

// Spherical light source.
// Throws softer shadows than an point source and is visible in specular reflections.
// TODO: nearby samples must limit their intensity to the analytical value for that limit.
func Sphere(pos Vec, radius float64, intensity Color) Light {
	mat := mat.Flat(intensity.Mul(1 / (4 * Pi * radius * radius)))
	return &sphereLight{
		//sphere: sphere{pos, radius, mat},
		Obj: shape.Sphere(pos, radius, mat),
		r:   radius,
		c:   intensity,
		pos: pos,
	}
}

type sphereLight struct {
	Obj
	r   float64
	c   Color
	pos Vec
}

func (l *sphereLight) Sample(ctx *Ctx, target Vec) (Vec, Color) {
	pos := l.pos.MAdd(l.r, sphereVec(ctx.Rng))
	return pos, l.c.Mul((1 / (4 * Pi)) / target.Sub(pos).Len2())
}

// Samples a vector from inside the volume of a unit sphere.
func sphereVec(rng *rand.Rand) Vec {
	v := cubeVec(rng)
	for v.Len2() > 1 {
		v = cubeVec(rng)
	}
	return v
}

// Samples a vector form inside a cube with edge 2.
func cubeVec(rng *rand.Rand) Vec {
	return Vec{
		2.0*rng.Float64() - 1,
		2.0*rng.Float64() - 1,
		2.0*rng.Float64() - 1,
	}
}

func RectLight(pos Vec, rx, ry, rz float64, c Color) Light {
	var dir Vec
	surf := 1.0
	R := Vec{rx, ry, rz}
	for i, r := range R {
		if r == 0 {
			dir = Unit[i]
			R[i] = 1 // rect does not work with r=0
		} else {
			surf *= r
		}
	}

	intens := c.Mul(1 / surf)
	return &rectLight{
		Obj:   shape.Rect(pos, dir, R[X], R[Y], R[Z], mat.Flat(intens)),
		color: c,
		pos:   pos,
		rx:    rx,
		ry:    ry,
		rz:    rz,
		dir:   dir,
	}
}

type rectLight struct {
	Obj        // TODO: replace by shape.Rect
	color      Color
	rx, ry, rz float64
	pos        Vec
	dir        Vec
}

func (l *rectLight) Sample(ctx *Ctx, target Vec) (Vec, Color) {
	rnd := cubeVec(ctx.Rng).Mul3(Vec{l.rx, l.ry, l.rz})
	pos := l.pos.Add(rnd)
	delta := target.Sub(pos)
	lamb := math.Abs(l.dir.Dot(delta.Normalized()))
	return pos, l.color.Mul((lamb) / delta.Len2())
}
