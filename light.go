package bruteray

type Light interface {
}

func PointLight(pos Vec, intensity Color) Light {
	return &pointLight{pos, intensity}
}

type pointLight struct {
	pos Vec
	c   Color
}
