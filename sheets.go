package main

func Sheet(pos float64, dir Vec) Shape {
	return &shape{
		hit: func(r *Ray) float64 {
			rs := r.Start.Dot(dir)
			rd := r.Dir.Dot(dir)
			t := (pos - rs) / rd
			return Max(t, 0)
		},
		normal: func(r *Ray, t float64) Vec {
			return dir.Towards(r.Dir)
		},
	}
}

func Rect(pos, dir Vec, rx, ry, rz float64) Shape {
	return &shape{
		hit: func(r *Ray) float64 {
			rs := r.Start.Dot(dir)
			rd := r.Dir.Dot(dir)
			t := (pos.Dot(dir) - rs) / rd
			if t < 0 {
				return 0
			}
			p := r.At(t).Sub(pos)
			if p.X < -rx || p.X > rx ||
				p.Y < -ry || p.Y > ry ||
				p.Z < -rz || p.Z > rz {
				return 0
			}
			return t
		},
		normal: func(r *Ray, t float64) Vec {
			return dir.Towards(r.Dir)
		},
	}
}
