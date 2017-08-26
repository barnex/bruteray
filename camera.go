package bruteray

// Camera renders a scene into a raw intensity image.
type Cam struct {
	FocalLen float64
	transf   Matrix4
	AA       bool
}

// Constructs a camera with given focal length.
// Focal length 0 means infinity (orthogonal projection).
// Camera is at the origin, looking in the +z direction,
// and can be transformed later.
func Camera(focalLen float64) *Cam {
	return &Cam{
		FocalLen: focalLen,
		transf:   *UnitMatrix4(),
	}
}

// Translates the camera.
func (c *Cam) Transl(dx, dy, dz float64) *Cam {
	c.Transf(Transl4(Vec{dx, dy, dz}))
	return c
}

// Transforms the camera direction,
// e.g. rotating the camera.
func (c *Cam) Transf(T *Matrix4) *Cam {
	c.transf = *((&c.transf).Mul(T))
	return c
}

func (c *Cam) RotScene(theta float64) *Cam {
	d := Vec{c.transf[X][W], c.transf[Y][W], c.transf[Z][W]}
	T := Transl4(d.Mul(-1)).Mul(RotY4(theta)).Mul(Transl4(d))
	c.transf = *((&c.transf).Mul(T))
	return c
}
