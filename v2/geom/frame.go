package geom

import "fmt"

type Frame [4]Vec

func (f *Frame) CtrlPoints() []*Vec { return addresses(f[:]) }

func (f *Frame) SetCtrlVec(i int, v Vec) {
	f[i+1] = f.Origin().Add(v)
}

func (f *Frame) TransformToAbsolute(v Vec) Vec {
	m := f.Matrix()
	return m.MulVec(v).Add(f.Origin())
}

func (f *Frame) TransformToFrame(v Vec) Vec {
	v = v.Sub(f.Origin())

	m := f.Matrix()
	i := m.Inv()
	return i.MulVec(v)
}

func (f *Frame) Origin() Vec {
	return f[0]
}

func (f *Frame) CtrlVecs() []Vec {
	return []Vec{
		f.CtrlVec(0),
		f.CtrlVec(1),
		f.CtrlVec(2),
	}
}

func (f *Frame) CtrlVec(i int) Vec {
	return f[i+1].Sub(f.Origin())
}

func (f *Frame) Matrix() Matrix {
	return Matrix{
		f.CtrlVec(0),
		f.CtrlVec(1),
		f.CtrlVec(2),
	}
}

// constructs a hull symmetrically around the origin.
//  x  ^   x
//     |
//     O -->
//
//  x      x
func (f *Frame) MaxHull() []Vec {
	c := f.CtrlVecs()
	o := f.Origin()
	a := [8][3]float64{
		{+1, +1, +1},
		{+1, +1, -1},
		{+1, -1, +1},
		{+1, -1, -1},
		{-1, +1, +1},
		{-1, +1, -1},
		{-1, -1, +1},
		{-1, -1, -1},
	}
	h := make([]Vec, len(a))
	for i := range h {
		a := a[i]
		h[i] = o.MAdd(a[0], c[0]).MAdd(a[1], c[1]).MAdd(a[2], c[2])
	}
	return h
}

func (f *Frame) AHull() []Vec {
	c := f.CtrlVecs()
	o := f.Origin()
	a := [8][3]float64{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 1},
	}
	h := make([]Vec, len(a))
	for i := range h {
		a := a[i]
		h[i] = o.MAdd(a[0], c[0]).MAdd(a[1], c[1]).MAdd(a[2], c[2])
	}
	return h
}

func (f *Frame) Scale(factor float64)   { Scale(f, f.Origin(), factor) }
func (f *Frame) Scale3(x, y, z float64) { Scale3(f, f.Origin(), x, y, z) }
func (f *Frame) Translate(delta Vec)    { Translate(f, delta) }

func (f *Frame) Pitch(angle float64) { Pitch(f, f.Origin(), angle) }
func (f *Frame) Yaw(angle float64)   { Yaw(f, f.Origin(), angle) }
func (f *Frame) Roll(angle float64)  { Roll(f, f.Origin(), angle) }

func (f *Frame) Check() {
	for _, v := range f.CtrlVecs() {
		if v == (Vec{}) {
			panic(fmt.Sprintf("invalid frame: %v", f))
		}
	}
}

var Ex = Vec{1, 0, 0}
var Ey = Vec{0, 1, 0}
var Ez = Vec{0, 0, 1}
var O = Vec{}
var XYZ = Frame{O, Ex, Ey, Ez}
var ZXY = Frame{O, Ez, Ex, Ey}
var YZX = Frame{O, Ey, Ez, Ex}

func addresses(v []Vec) []*Vec {
	a := make([]*Vec, len(v))
	for i := range v {
		a[i] = &v[i]
	}
	return a
}
