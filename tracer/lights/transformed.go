package lights

import (
	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/tracer/objects"
	. "github.com/barnex/bruteray/tracer/types"
)

// TODO: we might want to expose a lights.Interface
// which is broader than tracer.Light. See "here be dragons".

// Transformed returns a transformed instance of the original light.
// The transform may be a rotation and/or scale (equal in all directions),
// but not an anistropic scale or shear, etc.
//
// Multiple transformed instances of the same object may be made,
// they efficiently share underlying state.
//
// TODO: if object is already an instance of *transformed,
// combine both transforms into one.
//
// TODO: arguments: first transform, then object
func Transformed(l Light, t *geom.AffineTransform) Light {
	return &transformed{
		forward: *t,
		inverse: *t.Inverse(),
		orig:    l,
		// here be dragons
		object: objects.Transformed(l.Object().(objects.Interface), t),
	}
}

type transformed struct {
	forward geom.AffineTransform
	inverse geom.AffineTransform
	orig    Light
	object  Object
}

func (l *transformed) Object() Object {
	return l.object
}

func (l *transformed) Sample(ctx *Ctx, target Vec) (Vec, Color) {
	ttarget := l.inverse.TransformPoint(target)
	tpos, bright := l.orig.Sample(ctx, ttarget)
	pos := l.forward.TransformPoint(tpos)
	return pos, bright
}
