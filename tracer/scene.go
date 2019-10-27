package tracer

import (
	"fmt"
	"math"
	"os"

	. "github.com/barnex/bruteray/color"
)

//var Check = true

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
}

// NewScene constructs a scene storing objects and light sources.
func NewScene(lights []Light, objs ...Object) *Scene {
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
	}
}

// TODO: unify
func NewSceneWithMedia(media []Medium, lights []Light, objs ...Object) *Scene {
	s := NewScene(lights, objs...)
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
// recDepth controls the maximum recursion depth,
// i.e., the number of reflections a ray may undergo before being truncated.
//
// TODO: when testing indirect lighting: ImageFunc for non-luminous, lights only
func (s *Scene) ImageFunc(c Camera, recDepth int) ImageFunc {
	return func(ctx *Ctx, u, v float64) Color {
		r := c.RayFrom(ctx, u, v)
		c := s.Eval(ctx, r, recDepth)
		ctx.PutRay(r)
		return c
	}
}

type ImageFunc func(c *Ctx, u, v float64) Color

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

func (b *black) Eval(*Ctx, *Scene, *Ray, int, HitCoords) Color {
	return Color{0, 0, 0}
}

// Eval returns the light field of all objects in scene.
func (s *Scene) Eval(ctx *Ctx, r *Ray, recDepth int) Color {
	return s.eval(s.objectsAndLights, ctx, r, recDepth)
}

// EvalMinusLights returns the light field the Scene's objects
// but replaces the light sources by black.
// This is used by Bi-directional path tracing which adds light sources separately.
func (s *Scene) EvalMinusLights(ctx *Ctx, r *Ray, recDepth int) Color {
	return s.eval(s.objectsMinusLights, ctx, r, recDepth)
}

// eval returns the light field of an arbitrary group of objects.
func (s *Scene) eval(who []Object, ctx *Ctx, r *Ray, recDepth int) Color {
	ctx.RayCount++
	if recDepth == 0 {
		return Color{}
	}

	front := intersectFrontmost(who, r)
	if front.Material == nil { // TODO: this test should not be neccesary
		return Color{}
	}

	brightness := front.Material.Eval(ctx, s, r, recDepth-1, HitCoords{
		T:      front.T,
		Normal: front.Normal.Normalized(), // Scale surface normal to unit length now that we are sure we are going to use it
		Local:  front.Local,
	})

	for _, m := range s.media {
		brightness = m.Filter(ctx, s, r, front.T, brightness)
	}

	return brightness
}

func (s *Scene) Lights() []Light {
	return s.lights
}

func (s *Scene) ObjectsAndLights() []Object {
	return s.objectsAndLights
}

//func (s *Scene) IsOccluded(r *Ray, len float64) bool {
//	for _, o := range s.objects { // range over objects only, lights are considered transparent
//		f := o.Intersect(r)
//		if f.T > 0 && f.T < len {
//			return true
//		}
//	}
//	return false
//}

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
		if hit.T > 0 && !(hit.T > front.T) { // handles inf correctly
			front = hit
		}
	}
	return front
}

func log(x ...interface{}) {
	fmt.Fprintln(os.Stderr, x...)
}

func Frontmost(a, b *HitRecord) HitRecord {
	if b.T < a.T {
		a, b = b, a
	}
	if a.T > 0 {
		return *a
	}
	if b.T > 0 {
		return *b
	}
	return HitRecord{}
}

func FrontT(t1, t2 float64) float64 {
	if t2 < t1 {
		t1, t2 = t2, t1
	}
	if t1 > 0 {
		return t1
	}
	if t2 > 0 {
		return t2
	}
	return 0
}

func CheckHit(o Object, r *Ray, h *HitRecord) {
	if h.T == 0 {
		return
	}
	if math.IsNaN(h.T +
		h.Normal[0] + h.Normal[1] + h.Normal[2] +
		h.Local[0] + h.Local[1] + h.Local[2]) { //|| h.Local == (Vec{}) {
		panic(fmt.Sprintf("Scene: Intersect: bad HitRecord\nObject: %#v\nRay: %v\nHitRecord: %v", o, r, *h))
		// Note: *h, otherwise escape analysis thinks h escpaes, causing an alloc per ray, 3x overall performance penalty.
	}
}

func CheckRay(r *Ray) {
	const tol = 1e-9
	if math.Abs(1-r.Dir.Len2()) > tol {
		panic(fmt.Sprintf("unnormalized Ray dir: %v (len=%v)", r.Dir, r.Dir.Len()))
	}
}
