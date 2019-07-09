package tracer

type Object interface {
	Intersect(ctx *Ctx, r *Ray) HitRecord
}
