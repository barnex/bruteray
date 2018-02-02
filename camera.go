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

func (c *Cam) RayFrom(e *Env, i, j int, W, H int) *Ray {
	focalPoint := Vec{0, 0, -c.FocalLen}

	r := new(Ray)

	// ray start point
	y0 := (-float64(i) + c.aa(e) + float64(H)/2) / float64(H)
	x0 := (float64(j) + c.aa(e) - float64(W)/2) / float64(H)
	r.Start = Vec{x0, y0, 0}

	// ray direction
	if c.FocalLen != 0 {
		r.SetDir(r.Start.Sub(focalPoint).Normalized())
	} else {
		r.SetDir(Vec{0, 0, 1})
	}

	// camera transform
	r.Transf(&(c.transf))

	return r
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

// Anti-aliasing jitter
func (c *Cam) aa(e *Env) float64 {
	if c.AA {
		return random(e)
	} else {
		return 0.5
	}
}
