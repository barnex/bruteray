package geom

import (
	"fmt"
	"math"
	"testing"
)

func ExampleVec() {
	a := Vec{1, 2, 3}
	b := Vec{0, 1, -2}

	fmt.Printf("% g.Add(% g)     = % g\n", a, b, a.Add(b))
	fmt.Printf("% g.Cross(% g)   = % g\n", a, b, a.Cross(b))
	fmt.Printf("% g.Dot(% g)     = % g\n", a, b, a.Dot(b))
	fmt.Printf("% g.MAdd(2, % g) = % g\n", a, b, a.MAdd(2, b))
	fmt.Printf("% g.Sub(% g)     = % g\n", a, b, a.Sub(b))

	fmt.Printf("% g.IsNaN()     = %v\n", a, a.IsNaN())
	fmt.Printf("% g.Mul(2)      = % g\n", a, a.Mul(2))
	fmt.Printf("% g.Len()       = % g\n", a, a.Len())
	fmt.Printf("% g.Len2()     = % 3g\n", a, a.Len2())
	fmt.Printf("% g.Normalized()= % g\n", a, a.Normalized())

	//Output:
	// [ 1  2  3].Add([ 0  1 -2])     = [ 1  3  1]
	// [ 1  2  3].Cross([ 0  1 -2])   = [-7  2  1]
	// [ 1  2  3].Dot([ 0  1 -2])     = -4
	// [ 1  2  3].MAdd(2, [ 0  1 -2]) = [ 1  4 -1]
	// [ 1  2  3].Sub([ 0  1 -2])     = [ 1  1  5]
	// [ 1  2  3].IsNaN()     = false
	// [ 1  2  3].Mul(2)      = [ 2  4  6]
	// [ 1  2  3].Len()       =  3.7416573867739413
	// [ 1  2  3].Len2()     =  14
	// [ 1  2  3].Normalized()= [ 0.2672612419124244  0.5345224838248488  0.8017837257372732]
}

func ExampleVec_Normalize() {
	a := Vec{3, 4, 0}
	a.Normalize()
	fmt.Printf("% .1f\n", a)

	//Output:
	// [ 0.6  0.8  0.0]
}

func ExampleTriangleNormal() {
	a := Vec{1, 0, 0}
	b := Vec{2, 0, 0}
	c := Vec{1, 1, 0}
	fmt.Println(TriangleNormal(a, b, c))

	//Output:
	//[0 0 1]
}

func BenchmarkVec_Add_Accumulate(b *testing.B) {
	var a Vec
	v := Vec{1, 2, 3}
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = a.Add(v)
	}
	_ = a
}

func BenchmarkVec_Add_Accumulate_Inlined(b *testing.B) {
	var a Vec
	v := Vec{1, 2, 3}
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a[0] += v[0]
		a[1] += v[1]
		a[2] += v[2]
	}
	_ = a
}

func BenchmarkVec_Add_Accumulate_Registerized(b *testing.B) {
	var a Vec
	ax, ay, az := a[X], a[Y], a[Z]
	v := Vec{1, 2, 3}
	vx, vy, vz := v[X], v[Y], v[Z]
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ax += vx
		ay += vy
		az += vz
	}
	a[X], a[Y], a[Z] = ax, ay, az
	_ = a
}

func BenchmarkVec_Add_Streaming(b *testing.B) {
	v := Vec{1, 2, 3}
	var a [256]Vec
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := byte(i)
		a[j] = a[j].Add(v)
	}
	_ = a
}

func BenchmarkVec_Assign_Streaming(b *testing.B) {
	v := Vec{1, 2, 3}
	var a [256]Vec
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := byte(i)
		a[j] = v
	}
	_ = a
}

func TestVec(t *testing.T) {
	check := func(got, want Vec) {
		t.Helper()
		const tol = 1e-6
		fail :=
			(math.Abs(got[X]-want[X]) > tol) ||
				(math.Abs(got[Y]-want[Y]) > tol) ||
				(math.Abs(got[Z]-want[Z]) > tol)
		if fail {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	check(Vec{1, 2, 3}.Add(Vec{4, 5, 6}), Vec{5, 7, 9})
	check(Vec{1, 2, 3}.MAdd(2, Vec{4, 5, 6}), Vec{9, 12, 15})
	check(Vec{1, 2, 3}.Sub(Vec{1, 3, 2}), Vec{0, -1, 1})
	check(Vec{1, 2, 3}.Mul(2), Vec{2, 4, 6})
	check(Vec{0, 3, 4}.Normalized(), Vec{0, 3. / 5., 4. / 5.})
	check(Vec{0, 0, 0}.Normalized(), Vec{0, 0, 0})
	check(Vec{1, 0, 0}.Cross(Vec{0, 1, 0}), Vec{0, 0, 1})
	check(Vec{0, 1, 0}.Cross(Vec{1, 0, 0}), Vec{0, 0, -1})
	check(Vec{-1, 7, 4}.Cross(Vec{-5, 8, 4}), Vec{-4, -16, 27})
}
