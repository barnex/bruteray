package geom

type Transformable interface {
	// CtrlPoints returns pointers to control points,
	// so that transfomrations can modify them.
	CtrlPoints() []*Vec
}

func Translate(x Transformable, delta Vec) {
	for _, p := range x.CtrlPoints() {
		*p = (*p).Add(delta)
	}
}

func TranslateTo(x Transformable, src, dst Vec) {
	Translate(x, dst.Sub(src))
}

func Scale(x Transformable, origin Vec, factor float64) {
	Linear(x, origin, UnitMatrix.Mulf(factor))
}

func Scale3(s Transformable, origin Vec, x, y, z float64) {
	m := Matrix{
		{x, 0, 0},
		{0, y, 0},
		{0, 0, z},
	}
	Linear(s, origin, m)
}

func Roll(x Transformable, origin Vec, angle float64) {
	c := cos(angle)
	s := sin(angle)
	m := Matrix{
		{c, -s, 0},
		{s, c, 0},
		{0, 0, 1},
	}
	Linear(x, origin, m)
}

func Yaw(x Transformable, origin Vec, angle float64) {
	c := cos(angle)
	s := sin(angle)
	m := Matrix{
		{c, 0, -s},
		{0, 1, 0},
		{s, 0, c},
	}
	Linear(x, origin, m)
}

func Pitch(x Transformable, origin Vec, angle float64) {
	c := cos(angle)
	s := sin(angle)
	m := Matrix{
		{1, 0, 0},
		{0, c, -s},
		{0, s, c},
	}
	Linear(x, origin, m)
}

// Linear performs a general linear transformation.
func Linear(x Transformable, origin Vec, m Matrix) {
	for _, p := range x.CtrlPoints() {
		*p = m.MulVec((*p).Sub(origin)).Add(origin)
	}
}
