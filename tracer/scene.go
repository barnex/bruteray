package tracer

import (
	"math"

	. "github.com/barnex/bruteray/imagef/colorf"
)

// Scene is the ray tracer's central data type.
// A Scene stores a collection of objects and lights,
// and provides methods to evaluate the brightness seen by an arbitrary ray
// (I.e. the light field).
type Scene struct {
	lights             []Light
	objects            []Object
	objectsAndLights   []Object
	objectsMinusLights []Object
	media              []Medium
	RecursionDepth     int
}

// NewScene constructs a scene from objects and light sources.
// At render time, recursion will be limited to recursionDepth at maximum.
func NewScene(recursionDepth int, lights []Light, objs ...Object) *Scene {
	objMinusLights := make([]Object, len(objs), len(objs)+len(lights))
	objAndLights := make([]Object, len(objs), len(objs)+len(lights))
	copy(objMinusLights, objs)
	copy(objAndLights, objs)
	for _, l := range lights {
		o := l.Object()
		objMinusLights = append(objMinusLights, &black{o})
		objAndLights = append(objAndLights, o)
	}
	return &Scene{
		lights:             lights,
		objects:            objs,
		objectsAndLights:   objAndLights,
		objectsMinusLights: objMinusLights,
		RecursionDepth:     recursionDepth,
	}
}

// TODO: unify
func NewSceneWithMedia(recDepth int, media []Medium, lights []Light, objs ...Object) *Scene {
	s := NewScene(recDepth, lights, objs...)
	s.media = media
	return s
}

// ImageFunc returns an image function that can be sampled
// to obtain an image of the Scene seen through the given Camera.
//
// The image function expects to be sampled at (u,v)
// coordinates in the interval [0..1].
//
// Evaluating the image function returns potentially noisy sample,
// and typically many samples must be averaged to obtain a high-quality image.
//
// The context (*Ctx) is used to generate random numbers deterministically.
//
// TODO: when testing indirect lighting: ImageFunc for non-luminous, lights only
func (s *Scene) ImageFunc(c Camera) ImageFunc {
	return func(ctx *Ctx, u, v float64) Color {
		r := c.RayFrom(ctx, u, v)
		c := s.LightField(ctx, r)
		ctx.PutRay(r)
		return c
	}
}

type ImageFunc func(c *Ctx, u, v float64) Color

// LightField returns the light field of all objects in scene.
func (s *Scene) LightField(ctx *Ctx, r *Ray) Color {
	return s.lightField(s.objectsAndLights, ctx, r)
}

// LightFieldIndirect returns the light field the Scene's objects
// but replaces the light sources by black.
// This is used by Bi-directional path tracing which adds light sources separately.
func (s *Scene) LightFieldIndirect(ctx *Ctx, r *Ray) Color {
	return s.lightField(s.objectsMinusLights, ctx, r)
}

// lightField returns the light field of an arbitrary group of objects.
func (s *Scene) lightField(who []Object, ctx *Ctx, r *Ray) Color {
	if s.RecursionDepth == ctx.CurrentRecursionDepth {
		return Color{} // reached maximum recursion depth
	}
	ctx.Stats.NumRays++
	ctx.CurrentRecursionDepth++ // enter recursive evaluation

	front := intersectFrontmost(who, r)
	//if front.Material == nil { // TODO: this test should not be neccesary
	//return Color{}
	//}

	brightness := front.Material.Shade(ctx, s, r, HitCoords{
		T:      front.T,
		Normal: front.Normal.Normalized(), // Scale surface normal to unit length now that we are sure we are going to use it
		Local:  front.Local,
	})

	for _, m := range s.media {
		brightness = m.Filter(ctx, s, r, front.T, brightness)
	}

	ctx.CurrentRecursionDepth-- // exit recursive evaluation
	return brightness
}

func (s *Scene) Lights() []Light {
	return s.lights
}

func (s *Scene) ObjectsAndLights() []Object {
	return s.objectsAndLights
}

func (s *Scene) Occlude(r *Ray, len float64, orig Color) Color {
	for _, o := range s.objects { // range over objects only, lights are considered transparent
		f := o.Intersect(r)
		if f.T > 0 && f.T < len {
			if trans, ok := f.Material.(TransparentMaterial); ok {
				orig = trans.Filter(r, f, orig)
			} else {
				return Color{}
			}
		}
	}
	return orig
}

var blackMat Material = (*black)(nil)

// TODO: likewise for lights: per object: cache an occluding object for early return on shadows
func intersectFrontmost(objs []Object, r *Ray) HitRecord {
	front := HitRecord{T: math.Inf(1), Normal: r.Dir, Material: blackMat}
	//if Check {
	//	CheckRay(r)
	//}
	for _, o := range objs {
		hit := o.Intersect(r)
		//if Check {
		//	CheckHit(o, r, &hit) // DEBUG
		//}
		if hit.T > 0 && hit.T <= front.T { // handles inf and NaN correctly
			front = hit
		}
	}
	return front
}

// black wraps an object in a flat black material.
// Used to "remove" lights fromt the Scene for bidirectional path tracing.
type black struct {
	orig Object
}

func (b *black) Intersect(r *Ray) HitRecord {
	h := b.orig.Intersect(r)
	h.Material = b
	return h
}

func (b *black) Shade(*Ctx, *Scene, *Ray, HitCoords) Color {
	return Color{}
}
