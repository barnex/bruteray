package builder

import (
	"reflect"
	"testing"

	. "github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/test"
	. "github.com/barnex/bruteray/v2/tracer"
)

func TestSphere(t *testing.T) {
	scene := NewSceneBuilder()

	s := NewSphere(material.Normal(), 2)
	s.Translate(Vec{-.5, 0, 4})

	scene.Add(s)
	scene.Camera.FocalLen = 1

	test.OnePass(t, scene.Build(), test.DefaultTolerance)
}

func TestSphere_Intersect(t *testing.T) {
	ctx := NewCtx(0)
	s := NewSphere(nil, 1)
	s.Init()

	cases := []struct {
		ray  *Ray
		want HitRecord
	}{
		{ray(Vec{+2, 0, 0}, Vec{-1, 0, 0}), HitRecord{T: 1.5, Normal: Vec{+1, 0, 0}}},
		{ray(Vec{-2, 0, 0}, Vec{+1, 0, 0}), HitRecord{T: 1.5, Normal: Vec{-1, 0, 0}}},
		{ray(Vec{+0, 0, 0}, Vec{+1, 0, 0}), HitRecord{T: 0.5, Normal: Vec{+1, 0, 0}}},
		{ray(Vec{+2, 0, 0}, Vec{+1, 0, 0}), HitRecord{T: 0.0, Normal: Vec{+0, 0, 0}}},
	}

	for i, c := range cases {
		got := s.Intersect(ctx, c.ray)
		if got.Normal != (Vec{}) {
			got.Normal = got.Normal.Normalized()
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("case %v: %v: got: %v, want: %v", i, c.ray, got, c.want)
		}
	}
}

func BenchmarkSphereIntersectHit(b *testing.B) {
	b.ReportAllocs()

	s := NewSphere(nil, 1)
	ctx := NewCtx(123)
	r := ray(Vec{}, Vec{1, 0, 0})
	f := s.Intersect(ctx, r)
	if f.T != 0.5 {
		b.Fatal(f)
	}
	for i := 0; i < b.N; i++ {
		s.Intersect(ctx, r)
	}
}

func BenchmarkSphereIntersectMiss(b *testing.B) {
	b.ReportAllocs()

	s := NewSphere(nil, 1)
	ctx := NewCtx(123)
	r := ray(Vec{2, 0, 0}, Vec{1, 0, 0})
	f := s.Intersect(ctx, r)
	if f.T != 0 {
		b.Fatal(f)
	}
	for i := 0; i < b.N; i++ {
		s.Intersect(ctx, r)
	}
}

func checkDeepEq(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
