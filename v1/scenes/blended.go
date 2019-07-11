// +build ignore

package main

import (
	. "github.com/barnex/bruteray/v1/br"
	"github.com/barnex/bruteray/v1/light"
	. "github.com/barnex/bruteray/v1/mat"
	"github.com/barnex/bruteray/v1/raster"
	"github.com/barnex/bruteray/v1/serve"
	. "github.com/barnex/bruteray/v1/shape"
)

func main() {
	e := NewEnv()

	ch1 := Diffuse(WHITE.EV(-1))
	ch2 := Diffuse(WHITE.EV(-2))
	w := Shiny(WHITE.EV(-.6), EV(-2))
	//w = DebugShape(WHITE)

	round := 0.1

	e.Add(
		NewSheet(Ey, 0, Checkboard(1, ch1, ch2)),

		NewTile(round, w).Transl(Vec{0, 1, 0}),
	)

	e.AddLight(
		light.Sphere(Vec{4, 5, -9}, 1, WHITE.EV(10)),
	)

	cam := raster.Camera(1).Transl(0, 1, -3).Transf(RotX4(0 * Deg))
	cam.AA = true
	cam.Focus = 3.2
	cam.Diaphragm = DiaCircle
	//cam.Aperture = 1. / 16.

	serve.Env(cam, e)
}

func NewTile(round float64, m Material) *Tile {
	p0 := Vec{round, round, 0}
	pu := Vec{1 - round, round, 0}
	pv := Vec{round, 1 - round, 0}

	rect := NewRect(p0, pu, pv, m)
	top := pv[Y]
	bot := p0[Y]
	lft := p0[X]
	rgt := pu[X]

	t1 := NewCylinder(X, Vec{0.5, top, round}, 2*round, 1-2*round, m)
	t2 := NewCylinder(X, Vec{0.5, bot, round}, 2*round, 1-2*round, m)
	t3 := NewCylinder(Y, Vec{lft, 0.5, round}, 2*round, 1-2*round, m)
	t4 := NewCylinder(Y, Vec{rgt, 0.5, round}, 2*round, 1-2*round, m)
	return &Tile{
		R: rect,
		S: [4]*Sphere{
			NewSphere(2*round, m).Transl(rect.Pos(0, 0).Add(Vec{0, 0, round})),
			NewSphere(2*round, m).Transl(rect.Pos(1, 0).Add(Vec{0, 0, round})),
			NewSphere(2*round, m).Transl(rect.Pos(0, 1).Add(Vec{0, 0, round})),
			NewSphere(2*round, m).Transl(rect.Pos(1, 1).Add(Vec{0, 0, round})),
		},
		C: [4]CSGObj{t1, t2, t3, t4},
	}
}

type Tile struct {
	R *Rectangle
	S [4]*Sphere
	C [4]CSGObj
}

func (s *Tile) Hit1(r *Ray, f *[]Fragment) {
	s.R.Hit1(r, f)
	for _, s := range s.S {
		s.Hit1(r, f)
	}
	for _, s := range s.C {
		s.Hit1(r, f)
	}
}

func (s *Tile) Transl(d Vec) *Tile {
	s.R.Transl(d)
	for _, s := range s.S {
		s.Transl(d)
	}
	for _, s := range s.C {
		s.(Translate).Translate(d)
	}
	return s
}

func ctrl(p Vec, c Color) Obj {
	return NewSphere(0.1, DebugShape(c)).Transl(p)
}
