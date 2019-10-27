package geom

// Matrix is a 3x3 matrix intended for linear transformations.
type Matrix [3]Vec

// Mul performs a Matrix-Matrix multiplication
func (m *Matrix) Mul(b *Matrix) Matrix {
	var c Matrix
	for i := range c {
		for j := range c[i] {
			for k := range b {
				c[i][j] += b[i][k] * m[k][j]
			}
		}
	}
	return c
}

// MulVec performs a Matrix-Vector multiplication
// 	m . vT
func (m *Matrix) MulVec(v Vec) Vec {
	return Vec{
		m[0][0]*v[0] + m[1][0]*v[1] + m[2][0]*v[2],
		m[0][1]*v[0] + m[1][1]*v[1] + m[2][1]*v[2],
		m[0][2]*v[0] + m[1][2]*v[1] + m[2][2]*v[2],
	}
}

// Mulf multiplies element-wise by a scalar.
func (m *Matrix) Mulf(f float64) Matrix {
	return Matrix{m[0].Mul(f), m[1].Mul(f), m[2].Mul(f)}
}

// Inverse returns the inverse matrix.
func (m *Matrix) Inverse() Matrix {
	a := m[0][0]
	b := m[1][0]
	c := m[2][0]
	d := m[0][1]
	e := m[1][1]
	f := m[2][1]
	g := m[0][2]
	h := m[1][2]
	i := m[2][2]

	A := e*i - f*h
	B := f*g - d*i
	C := d*h - e*g
	inv := Matrix{
		{e*i - f*h, f*g - d*i, d*h - e*g},
		{c*h - b*i, a*i - c*g, b*g - a*h},
		{b*f - c*e, c*d - a*f, a*e - b*d},
	}
	det := a*A + b*B + c*C
	return inv.Mulf(1 / det)
}

//func (a *Matrix) Transpose() Matrix {
//	return Matrix{
//		{a[0][0], a[1][0], a[2][0]},
//		{a[0][1], a[1][1], a[2][1]},
//		{a[0][2], a[1][2], a[2][2]},
//	}
//}
