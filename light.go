package bruteray

type Light interface {
	Sample(target Vec) (Vec, Color)
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

// Source with parallel rays,
// like an infinitely far away point source.
//func DirLight(dir Vec, intensity Color) Light {
//	return &dirLight{dir, intensity}
//}
//
//type dirLight struct {
//	dir Vec
//	c   Color
//}
