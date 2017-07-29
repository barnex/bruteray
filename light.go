package bruteray

type Light interface {
	Sample(target Vec) (dir Vec, intens Color)
}

// Directed light source without fall-off.
// Position should be far away from the scene (indicates a direction)
func DirLight(pos Vec, intensity Color) Light {
	return &dirLight{pos, intensity}
}

type dirLight struct {
	pos Vec
	c   Color
}

func (l *dirLight) Sample(target Vec) (Vec, Color) {
	return l.pos, l.c
}

// Point light source (with fall-off).
func PointLight(pos Vec, intensity Color) Light {
	return &pointLight{pos, intensity}
}

type pointLight struct {
	pos Vec
	c   Color
}

func (l *pointLight) Sample(target Vec) (Vec, Color) {
	return l.pos, l.c.Mul(1 / target.Sub(l.pos).Len2()) // TODO: 4 pi
}
