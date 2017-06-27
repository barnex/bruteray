package main

func Sheet(pos float64, dir Vec) Shape {
	return &shape{
		inters: func(r *Ray) Inter {
			rs := r.Start.Dot(dir)
			rd := r.Dir.Dot(dir)
			t := (pos - rs) / rd
			return Inter{t, t}
		},
		normal: func(r *Ray, t float64) Vec {
			return dir.Towards(r.Dir)
		},
	}
}

func Rect(pos, dir Vec, rx, ry, rz float64) Shape {
	return &shape{
		inters: func(r *Ray) Inter {
			rs := r.Start.Dot(dir)
			rd := r.Dir.Dot(dir)
			t := (pos.Dot(dir) - rs) / rd
			if t < 0 {
				return empty
			}
			p := r.At(t).Sub(pos)
			if p.X < -rx || p.X > rx ||
				p.Y < -ry || p.Y > ry ||
				p.Z < -rz || p.Z > rz {
				return empty
			}
			return Inter{t, t}
		},
		normal: func(r *Ray, t float64) Vec {
			return dir.Towards(r.Dir)
		},
	}
}
