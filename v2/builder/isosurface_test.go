package builder

import (
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
)

func TestIsoSurface_Sphere(t *testing.T) {
	//scene := NewSceneBuilder()

	//tex := texture.ScalarFunc3D(func(p Vec) float64 {
	//	return p.Dot(p) - 1
	//})

	////bb :=
	//s := NewIsoSurface(Normal(), bb, tex)

	//built := scene.Build()
	//built.Camera.FocalLen = 1
	//test.OnePass(t, built, test.DefaultTolerance)
}

func TestIsoSurface_Intersect(t *testing.T) {
	ctx := NewCtx(0)
	s := NewSphere(nil, 1)
	s.Init()

	cases := []struct {
		ray    *Ray
		t1, t2 float64
	}{
		{ray(Vec{+2, 0, 0}, Vec{-1, 0, 0}), 1.5, 2.5},
		{ray(Vec{-2, 0, 0}, Vec{+1, 0, 0}), 1.5, 2.5},
		{ray(Vec{+0, 0, 0}, Vec{+1, 0, 0}), -0.5, 0.5},
		{ray(Vec{+2, 0, 0}, Vec{+1, 0, 0}), 0.0, 0.0},
	}

	for i, c := range cases {
		got1, got2 := intersect2(s, ctx, c.ray)
		if !(got1 == c.t1 && got2 == c.t2) {
			t.Errorf("case %v: %v: got: %v,%v, want: %v,%v", i, c.ray, got1, got2, c.t1, c.t2)
		}
	}
}
