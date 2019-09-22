package geom

type Frame struct {
	origin  Vec
	axes    Matrix
	inverse Matrix
}

func NewFrame(origin Vec, axes [3]Vec) Frame {
	f := Frame{
		origin: origin,
		axes:   axes,
	}
	f.initInverse()
	return f
}

func UnitFrame() Frame {
	return NewFrame(O, Matrix{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
}

func (f *Frame) TransformToLocal(x Vec) Vec {
	return f.inverse.MulVec(x.Sub(f.origin))
}

func (f *Frame) TransformToGlobal(x Vec) Vec {
	return f.axes.MulVec(x).Add(f.origin)
}

func (f *Frame) TransformDirToGlobal(x Vec) Vec {
	return f.axes.MulVec(x)
}

func (f *Frame) Origin() Vec {
	return f.origin
}

func (f *Frame) Axes() [3]Vec {
	return f.axes
}

func (f *Frame) ApplyTransform(tr func(Vec) Vec) { // TODO: *geom.AffineTransform
	f.origin = tr(f.origin)
	for i, a := range f.axes {
		f.axes[i] = applyToVec(tr, a)
	}
	f.initInverse()
}

func (f *Frame) initInverse() {
	f.inverse = f.axes.Inverse()
}

func applyToVec(f func(Vec) Vec, x Vec) Vec {
	return f(x).Sub(f(O))
}

// These should go through Apply()
//func (f *Frame) Translate(delta Vec) {
//	f.origin = f.origin.Add(delta)
//}
//
//func (f *Frame) Rotate(axis Vec, angle float64) {
//	m := rotationMatrix(axis, angle)
//	f.axes = f.axes.Mul(&m)
//}
//
//func (f *Frame) Scale(factor float64) {
//	f.axes = f.axes.Mulf(factor)
//}
