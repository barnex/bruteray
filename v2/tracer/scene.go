package tracer

import (
	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
)

type Scene struct {
	nonLuminous []Object
	Lights      []Light
	allObjects  []Object
	Camera      Camera // TODO: remove
}

func (s *Scene) AddObject(o ...Object) {
	for _, o := range o {
		s.allObjects = append(s.allObjects, o)
		s.nonLuminous = append(s.nonLuminous, o)
	}
}

func (s *Scene) AddLight(o ...Light) {
	for _, o := range o {
		s.allObjects = append(s.allObjects, o)
		s.Lights = append(s.Lights, o)
	}
}

// Light field of all objects in scene.
func (s *Scene) Eval(ctx *Ctx, r *Ray, recDepth int) Color {
	return s.lightField(s.allObjects, ctx, r, recDepth)
}

// TODO: LightFieldDebug(..., Material) replaces materials by Grid, etc

// Light field of objects that or not light sources.
// Bi-directional tracing adds sources separately.
func (s *Scene) EvalNonLuminous(ctx *Ctx, r *Ray, recDepth int) Color {
	return s.lightField(s.nonLuminous, ctx, r, recDepth)
}

// Light field of arbitrary group of objects.
func (s *Scene) lightField(who []Object, ctx *Ctx, r *Ray, recDepth int) Color {
	ctx.RayCount++
	if recDepth == 0 {
		return Color{}
	}
	front := IntersectFrontmost(ctx, who, r)
	if front.Material == nil {
		return Color{}
	}
	// Hack: if local coordinates are not set, use absolute.
	//if front.Local == O {
	//	front.Local = r.At(front.T)
	//}
	return front.Material.Eval(ctx, s, r, recDepth-1, HitCoords{front.T, front.Normal, front.Local})
}

func (s *Scene) IsOccluded(ctx *Ctx, r *Ray) bool {
	for _, o := range s.allObjects {
		f := o.Intersect(ctx, r)
		if f.T > 0 && f.T < r.Len {
			return true
		}
	}
	return false
}

// Light field projected through camera onto 2D texture.
func (s *Scene) ImageFunc(recDepth int) ImageFunc {
	return func(ctx *Ctx, x, y float64) Color {
		cam := &s.Camera
		r := cam.RayFrom(ctx, x, y)
		c := s.Eval(ctx, r, recDepth)
		ctx.PutRay(r)
		return c
	}
}

// TODO: cache last frontmost object in Ctx, use for early return based on culled ray Len
// TODO: likewise for lights: per object: cache an occluding object for early return on shadows
func IntersectFrontmost(ctx *Ctx, objs []Object, r *Ray) HitRecord {
	front := HitRecord{T: Inf}
	tBackup := r.Len

	for _, o := range objs {
		frag := o.Intersect(ctx, r)
		if frag.T > 0 && frag.T < front.T {
			front = frag
			r.Len = frag.T
		}
	}

	if front.Material == nil { // !ok
		return HitRecord{}
	}

	front.Normal = front.Normal.Normalized()
	r.Len = tBackup
	return front
}
