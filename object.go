package main

type Obj interface {
	Shape
	Shade(e *Env, r *Ray, t float64, N int) Color
}

type object struct {
	Shape
	shader Shader
}

func (o *object) Shade(e *Env, r *Ray, t float64, N int) Color {
	if N == 0 {
		return e.Ambient(r.Dir)
	}
	n := o.Shape.Normal(r, t)
	return o.shader.Shade(e, r, t, n, N)
}
