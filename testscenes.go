package main

import "math"

// some shapes for testing

func coolSphere() Shape {
	const (
		R = 2
		H = 2
		D = 0.85
	)
	base := Slab(8, 0.1, 8).Transl(0, -H, 0)
	s := Sphere(R)
	s = s.Sub(CylinderZ(R-D, H))
	s = s.Sub(CylinderZ(R-D, H).RotX(90 * deg))
	s = s.Sub(CylinderZ(R-D, H).RotY(90 * deg))
	s = s.RotY(-20 * deg)
	return s.Add(base)
}

func cubeFrame() Shape {
	const (
		X = 1
		Y = 0.5
		Z = 1
		D = 0.2
	)
	frame := Slab(X, Y, Z).Sub(Slab(X, Y-D, Z-D)).Sub(Slab(X-D, Y, Z-D)).Sub(Slab(X-D, Y-D, Z))
	return frame.RotY(-0.5).Transl(0, -0.2, 2)
}

func sinc() Shape {
	sinc := Shape(func(r Vec) bool {
		R := math.Sqrt(r.X*r.X+r.Z*r.Z) * 5
		return r.Y < 2*math.Sin(R)/R
	})
	return sinc.Intersect(Slab(4, 2, 4)).RotY(-0.4).RotX(-0.5).Transl(0, 0, 8).RotX(-0.2)
}
