package bruteray

type Light interface {
	Sample(e *Env, target Vec) (dir Vec, intens Color)
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

func (l *dirLight) Sample(e *Env, target Vec) (Vec, Color) {
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

func (l *pointLight) Sample(e *Env, target Vec) (Vec, Color) {
	return l.pos, l.c.Mul((1 / (1)) / target.Sub(l.pos).Len2()) // TODO: 1-> 4*pi
}

// Smooth light source
func SmoothLight(pos Vec, radius float64, intensity Color) Light {
	return &smoothLight{pos, radius, intensity}
}

type smoothLight struct {
	pos Vec
	r   float64
	c   Color
}

func (l *smoothLight) Sample(e *Env, target Vec) (Vec, Color) {
	pos := l.pos.MAdd(l.r, RandVec(e))
	return pos, l.c.Mul((1 / (1)) / target.Sub(pos).Len2()) // TODO: 1->4*pi
}

// returns a cheaper, lower quality light used for quick preview.
func toPreview(l Light) Light {
	switch l := l.(type) {
	default:
		return l
	case *smoothLight:
		return &pointLight{l.pos, l.c}
	}
}
