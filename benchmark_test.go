package bruteray

import "testing"

func BenchmarkSphere(b *testing.B) {
	e := NewEnv()
	e.Add(Sphere(Vec{0, 0, 1}, 0.25, Flat(WHITE)))
	c := Camera(0)

	benchmark(b, e, c)
}

func Benchmark9Spheres(b *testing.B) {
	e := NewEnv()
	r := 0.5

	//nz := ShadeNormal(Ez)
	e.Add(Sphere(Vec{0, 0, 0}, r, Flat(WHITE)))
	e.Add(Sphere(Vec{0, 0, 2}, r, Flat(WHITE)))
	e.Add(Sphere(Vec{0, 0, 4}, r, Flat(WHITE)))

	e.Add(Sphere(Vec{2, 0, 0}, r, Flat(WHITE)))
	e.Add(Sphere(Vec{2, 0, 2}, r, Flat(WHITE)))
	e.Add(Sphere(Vec{2, 0, 4}, r, Flat(WHITE)))

	e.Add(Sphere(Vec{-2, 0, 0}, r, Flat(WHITE)))
	e.Add(Sphere(Vec{-2, 0, 2}, r, Flat(WHITE)))
	e.Add(Sphere(Vec{-2, 0, 4}, r, Flat(WHITE)))

	c := Camera(1).Transl(0, 4, -4).Transf(RotX4(Pi / 5))
	benchmark(b, e, c)
}

func benchmark(b *testing.B, e *Env, c *Cam) {
	b.SetBytes((testW + 1) * (testH + 1))
	img := MakeImage(testW, testH)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Render(e, img)
	}
}
