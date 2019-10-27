package geom

import "fmt"

func ExampleRotate() {
	angle := 30 * Deg
	inputs := [3]Vec{Ex, Ey, Ez}

	Rx := Rotate(O, Vec{1, 0, 0}, angle)
	for _, v := range inputs {
		fmt.Printf("Rx: % .2f -> % .2f\n", v, Rx.TransformPoint(v))
	}
	fmt.Println()

	Ry := Rotate(O, Vec{0, 1, 0}, angle)
	for _, v := range inputs {
		fmt.Printf("Ry: % .2f -> % .2f\n", v, Ry.TransformPoint(v))
	}
	fmt.Println()

	Rz := Rotate(O, Vec{0, 0, 1}, angle)
	for _, v := range inputs {
		fmt.Printf("Rz: % .2f -> % .2f\n", v, Rz.TransformPoint(v))
	}
	fmt.Println()

	//Output:
	// Rx: [ 1.00  0.00  0.00] -> [ 1.00  0.00  0.00]
	// Rx: [ 0.00  1.00  0.00] -> [ 0.00  0.87  0.50]
	// Rx: [ 0.00  0.00  1.00] -> [ 0.00 -0.50  0.87]
	//
	// Ry: [ 1.00  0.00  0.00] -> [ 0.87  0.00 -0.50]
	// Ry: [ 0.00  1.00  0.00] -> [ 0.00  1.00  0.00]
	// Ry: [ 0.00  0.00  1.00] -> [ 0.50  0.00  0.87]
	//
	// Rz: [ 1.00  0.00  0.00] -> [ 0.87  0.50  0.00]
	// Rz: [ 0.00  1.00  0.00] -> [-0.50  0.87  0.00]
	// Rz: [ 0.00  0.00  1.00] -> [ 0.00  0.00  1.00]
}

func ExampleScale() {
	tf := Scale(O, 2)
	v := Vec{1, 2, 3}
	fmt.Printf("point: % g -> % g\n", v, tf.TransformPoint(v))
	fmt.Printf("dir:   % g -> % g\n", v, tf.TransformDir(v))

	//Output:
	// point: [ 1  2  3] -> [ 2  4  6]
	// dir:   [ 1  2  3] -> [ 2  4  6]
}

func ExampleTranslate() {
	tf := Translate(Vec{1, 0, 0})
	v := Vec{1, 2, 3}
	fmt.Printf("point: % g -> % g\n", v, tf.TransformPoint(v))
	fmt.Printf("dir:   % g -> % g\n", v, tf.TransformDir(v))

	//Output:
	// point: [ 1  2  3] -> [ 2  2  3]
	// dir:   [ 1  2  3] -> [ 1  2  3]
}

func ExampleAffineTransform_After_2() {
	// Make various compositions of a rotation and flip,
	// which do not commute.
	R := Rotate(O, Ez, 90*Deg)
	F := &AffineTransform{A: Matrix{{-1, 0, 0}, {0, 1, 0}, {0, 0, 1}}}

	RF := R.After(F)
	FR := F.After(R)

	for _, v := range []Vec{{1, 0, 0}, {0, 1, 0}} {
		fmt.Printf("R        : % g -> % g\n", v, R.TransformPoint(v))
		fmt.Printf("F        : % g -> % g\n", v, F.TransformPoint(v))
		fmt.Printf("RF       : % g -> % g\n", v, RF.TransformPoint(v))
		fmt.Printf("R(F)     : % g -> % g\n", v, R.TransformPoint(F.TransformPoint(v)))
		fmt.Printf("Comp(F,R): % g -> % g\n", v, ComposeLR(F, R).TransformPoint(v))
		fmt.Printf("FR       : % g -> % g\n", v, FR.TransformPoint(v))
		fmt.Printf("F(R)     : % g -> % g\n", v, F.TransformPoint(R.TransformPoint(v)))
		fmt.Printf("Comp(R,F): % g -> % g\n", v, ComposeLR(R, F).TransformPoint(v))
	}

	//Output:
	// R        : [ 1  0  0] -> [ 0  1  0]
	// F        : [ 1  0  0] -> [-1  0  0]
	// RF       : [ 1  0  0] -> [ 0 -1  0]
	// R(F)     : [ 1  0  0] -> [ 0 -1  0]
	// Comp(F,R): [ 1  0  0] -> [ 0 -1  0]
	// FR       : [ 1  0  0] -> [ 0  1  0]
	// F(R)     : [ 1  0  0] -> [ 0  1  0]
	// Comp(R,F): [ 1  0  0] -> [ 0  1  0]
	// R        : [ 0  1  0] -> [-1  0  0]
	// F        : [ 0  1  0] -> [ 0  1  0]
	// RF       : [ 0  1  0] -> [-1  0  0]
	// R(F)     : [ 0  1  0] -> [-1  0  0]
	// Comp(F,R): [ 0  1  0] -> [-1  0  0]
	// FR       : [ 0  1  0] -> [ 1  0  0]
	// F(R)     : [ 0  1  0] -> [ 1  0  0]
	// Comp(R,F): [ 0  1  0] -> [ 1  0  0]
}

func ExampleAffineTransform_After() {

	R := Rotate(O, Ez, 90*Deg)
	S := Scale(O, 2)
	T := Translate(Vec{1, 0, 0})

	ST := S.After(T)
	TR := T.After(R)
	RS := R.After(S)
	TS := T.After(S)
	RT := R.After(T)
	SR := S.After(R)
	RST := R.After(S).After(T)
	STR := S.After(T).After(R)
	TRS := T.After(R).After(S)

	x := Vec{1, 0, 0}

	fmt.Printf("R               : % g -> % g\n", x, R.TransformPoint(x))
	fmt.Printf("S               : % g -> % g\n", x, S.TransformPoint(x))
	fmt.Printf("T               : % g -> % g\n", x, T.TransformPoint(x))
	fmt.Printf("ST              : % g -> % g\n", x, ST.TransformPoint(x))
	fmt.Printf("S(T)            : % g -> % g\n", x, S.TransformPoint(T.TransformPoint(x)))
	fmt.Printf("TR              : % g -> % g\n", x, TR.TransformPoint(x))
	fmt.Printf("T(R)            : % g -> % g\n", x, T.TransformPoint(R.TransformPoint(x)))
	fmt.Printf("RS              : % g -> % g\n", x, RS.TransformPoint(x))
	fmt.Printf("R(S)            : % g -> % g\n", x, R.TransformPoint(S.TransformPoint(x)))
	fmt.Printf("TS              : % g -> % g\n", x, TS.TransformPoint(x))
	fmt.Printf("T(S)            : % g -> % g\n", x, T.TransformPoint(S.TransformPoint(x)))
	fmt.Printf("RT              : % g -> % g\n", x, RT.TransformPoint(x))
	fmt.Printf("R(T)            : % g -> % g\n", x, R.TransformPoint(T.TransformPoint(x)))
	fmt.Printf("SR              : % g -> % g\n", x, SR.TransformPoint(x))
	fmt.Printf("S(R)            : % g -> % g\n", x, S.TransformPoint(R.TransformPoint(x)))
	fmt.Printf("RST             : % g -> % g\n", x, RST.TransformPoint(x))
	fmt.Printf("R(S(T))         : % g -> % g\n", x, R.TransformPoint(S.TransformDir(T.TransformPoint(x))))
	fmt.Printf("ComposeLR(T,S,R): % g -> % g\n", x, ComposeLR(T, S, R).TransformPoint(x))
	fmt.Printf("STR             : % g -> % g\n", x, STR.TransformPoint(x))
	fmt.Printf("S(T(R))         : % g -> % g\n", x, S.TransformPoint(T.TransformPoint(R.TransformPoint(x))))
	fmt.Printf("ComposeLR(R,T,S): % g -> % g\n", x, ComposeLR(R, T, S).TransformPoint(x))
	fmt.Printf("TRS             : % g -> % g\n", x, TRS.TransformPoint(x))
	fmt.Printf("T(R(S))         : % g -> % g\n", x, T.TransformPoint(R.TransformPoint(S.TransformPoint(x))))
	fmt.Printf("ComposeLR(S,R,T): % g -> % g\n", x, ComposeLR(S, R, T).TransformPoint(x))

	//Output:
	// R               : [ 1  0  0] -> [ 0  1  0]
	// S               : [ 1  0  0] -> [ 2  0  0]
	// T               : [ 1  0  0] -> [ 2  0  0]
	// ST              : [ 1  0  0] -> [ 4  0  0]
	// S(T)            : [ 1  0  0] -> [ 4  0  0]
	// TR              : [ 1  0  0] -> [ 1  1  0]
	// T(R)            : [ 1  0  0] -> [ 1  1  0]
	// RS              : [ 1  0  0] -> [ 0  2  0]
	// R(S)            : [ 1  0  0] -> [ 0  2  0]
	// TS              : [ 1  0  0] -> [ 3  0  0]
	// T(S)            : [ 1  0  0] -> [ 3  0  0]
	// RT              : [ 1  0  0] -> [ 0  2  0]
	// R(T)            : [ 1  0  0] -> [ 0  2  0]
	// SR              : [ 1  0  0] -> [ 0  2  0]
	// S(R)            : [ 1  0  0] -> [ 0  2  0]
	// RST             : [ 1  0  0] -> [ 0  4  0]
	// R(S(T))         : [ 1  0  0] -> [ 0  4  0]
	// ComposeLR(T,S,R): [ 1  0  0] -> [ 0  4  0]
	// STR             : [ 1  0  0] -> [ 2  2  0]
	// S(T(R))         : [ 1  0  0] -> [ 2  2  0]
	// ComposeLR(R,T,S): [ 1  0  0] -> [ 2  2  0]
	// TRS             : [ 1  0  0] -> [ 1  2  0]
	// T(R(S))         : [ 1  0  0] -> [ 1  2  0]
	// ComposeLR(S,R,T): [ 1  0  0] -> [ 1  2  0]
}

func ExampleTransform_WithOrigin() {
	R000 := Rotate(O, Ez, 90*Deg)         // rotate around [0 0 0]
	R010 := R000.WithOrigin(Vec{0, 1, 0}) // rotate around [0 1 0]

	x := Vec{1, 0, 0}
	fmt.Printf("R000: % g -> % g\n", x, R000.TransformPoint(x))
	fmt.Printf("R010: % g -> % g\n", x, R010.TransformPoint(x))

	//Output:
	// R000: [ 1  0  0] -> [ 0  1  0]
	// R010: [ 1  0  0] -> [ 1  2  0]
}

func ExampleTransform_Inverse() {
	T := Rotate(O, Ez, 30*Deg).WithOrigin(Vec{1, 0, 0})
	Inv := T.Inverse()
	TInv := T.After(Inv)
	InvT := Inv.After(T)

	x := Vec{0, 1, 0}
	fmt.Printf("T:      % .2f -> % .2f\n", x, T.TransformPoint(x))
	fmt.Printf("Inv:    % .2f -> % .2f\n", x, Inv.TransformPoint(x))
	fmt.Printf("T(Inv): % .2f -> % .2f\n", x, T.TransformPoint(Inv.TransformPoint(x)))
	fmt.Printf("Inv(T): % .2f -> % .2f\n", x, Inv.TransformPoint(T.TransformPoint(x)))
	fmt.Printf("TInv:   % .2f -> % .2f\n", x, TInv.TransformPoint(x))
	fmt.Printf("InvT:   % .2f -> % .2f\n", x, InvT.TransformPoint(x))

	//Output:
	// T:      [ 0.00  1.00  0.00] -> [-0.37  0.37  0.00]
	// Inv:    [ 0.00  1.00  0.00] -> [ 0.63  1.37  0.00]
	// T(Inv): [ 0.00  1.00  0.00] -> [ 0.00  1.00  0.00]
	// Inv(T): [ 0.00  1.00  0.00] -> [-0.00  1.00  0.00]
	// TInv:   [ 0.00  1.00  0.00] -> [ 0.00  1.00  0.00]
	// InvT:   [ 0.00  1.00  0.00] -> [ 0.00  1.00  0.00]
}
