package test

import (
	"github.com/barnex/bruteray/color"
	. "github.com/barnex/bruteray/tracer/types"
)

var (
	Blue    = WithShadows(color.Blue)
	Cyan    = WithShadows(color.Cyan)
	Gray    = WithShadows(color.Gray(0.5))
	Green   = WithShadows(color.Green)
	Magenta = WithShadows(color.Magenta)
	White   = WithShadows(color.White)
	Yellow  = WithShadows(color.Yellow)
	// No Red, used to show back faces

	Checkers1 = Checkers(White, Magenta)
	Checkers2 = Checkers(White, Cyan)
	Checkers3 = Checkers(White, Green)
	Checkers4 = Checkers(WithShadows(color.Gray(0.4)), WithShadows(color.Gray(0.7)))
)

// Flat is a minimal implementation of material.Flat,
// a non-physical material that always evaluates to the same color.
// Intended for tests that should not depend on package materials.
type Flat Color

// Eval implements Material.
func (m Flat) Eval(_ *Ctx, _ *Scene, r *Ray,  h HitCoords) Color {
	return Color(m)
}

// Normal is a non-physical material that reveals the component of the normal vector
// that points towards the camera.
// The normal pointing perpendicular to the direction of view is rendered as black.
// Away from the viewing direction is rendered as red to draw attention to visible backfaces.
var Normal Material = &normal{1, 1, 1}

type normal Color

// Eval implements Material.
func (m *normal) Eval(ctx *Ctx, s *Scene, r *Ray,  h HitCoords) Color {
	checkRay(r)
	v := h.Normal.Dot(r.Dir)
	if v < 0 {
		return Color(*m).Mul(-v) // towards cam
	} else {
		return color.Red.Mul(v) // away from cam
	}
}

var Normal2 = normal2{}

type normal2 struct{}

// Eval implements Material.
func (normal2) Eval(ctx *Ctx, s *Scene, r *Ray,  h HitCoords) Color {
	checkRay(r)
	v := h.Normal.Dot(r.Dir)
	if v < 0 {
		v *= -1
	}
	return Color{0.5, 0.5, 0.5}.Mul(v).Add(Color{0.5, 0.5, 0.5})
}

func Transparent(trans, add Color) Material {
	return &transparent{trans, add}
}

// minimal re-implementation of materials.Transparent to avoid dependency.
type transparent struct {
	trans Color
	add   Color
}

func (m *transparent) Eval(ctx *Ctx, s *Scene, r *Ray,  h HitCoords) Color {
	pos := r.At(h.T + Tiny)
	r2 := ctx.Ray()
	r2.Start = pos
	r2.Dir = r.Dir
	defer ctx.PutRay(r2)
	return s.Eval(ctx, r2,).Mul3(m.trans).Add(m.add)
}

// WithShadows returns a non-physical material similar to Normal,
// but overlayed with shadows.
func WithShadows(c Color) Material {
	n := normal(c)
	return &matte{&n}
}

// matte is a non-physical material similar to real matte white,
// but without indirect illumination.
type matte struct {
	reflectivity Material
}

// Eval implements tracer.Material.
func (m matte) Eval(ctx *Ctx, s *Scene, r *Ray, h HitCoords) Color {
	var acc Color

	//normal := flipTowards(h.Normal, r.Dir)
	normal := h.Normal
	if normal.Dot(r.Dir) > 0 {
		normal = normal.Mul(-1)
	}
	p := r.At(h.T).MAdd(Tiny, normal)
	secundary := ctx.Ray()
	secundary.Start = p

	for _, l := range s.Lights() {
		lpos, intens := l.Sample(ctx, p)
		lDelta := lpos.Sub(p)
		lDir := lDelta.Normalized()
		cosTheta := lDir.Dot(normal)
		if cosTheta <= 0 {
			continue
		}
		lDist := lDelta.Len()
		secundary.Dir = lDir
		intens = s.Occlude(secundary, lDist, intens)
		acc = acc.Add(intens.Mul(cosTheta))
	}
	acc = acc.Mul(0.6).Add(Color{0.4, 0.4, 0.4})
	return m.reflectivity.Eval(ctx, s, r, h).Mul3(acc)
}

// blend evaluates to 50% a plus 50% b.
type blend struct {
	a, b Material
}

// Eval implements Material.
func (m *blend) Eval(ctx *Ctx, s *Scene, r *Ray,  h HitCoords) Color {
	a := m.a.Eval(ctx, s, r,  h)
	b := m.b.Eval(ctx, s, r,  h)
	return a.Mul(0.5).MAdd(0.5, b)
}

// Checkers alternates between materials a and b
// in a checkerboard pattern with period 1.
func Checkers(a, b Material) Material {
	return &checkers{a, b}
}

type checkers struct {
	a, b Material
}

// Eval implements material
func (m *checkers) Eval(ctx *Ctx, s *Scene, r *Ray,  h HitCoords) Color {
	checkRay(r)
	u := h.Local[0]
	v := h.Local[1]
	if (int(u*2+10000)+int(v*2+10000))%2 == 0 {
		return m.a.Eval(ctx, s, r,  h)
	} else {
		return m.b.Eval(ctx, s, r,  h)
	}
}
