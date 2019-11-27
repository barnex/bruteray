package api

import (
	"github.com/barnex/bruteray/imagef/post"
	"github.com/barnex/bruteray/tracer"
	"github.com/barnex/bruteray/tracer/cameras"
	"github.com/barnex/bruteray/tracer/materials"
	"github.com/barnex/bruteray/tracer/test"
)

// A Spec specifies a BruteRay scene. Usage:
//	Render(Spec{
// 		...
// 	})
type Spec struct {
	Lights  []Light
	Objects []Object
	Media   []Medium
	Camera  tracer.Camera

	Recursion int
	NumPass   int

	Width  int
	Height int

	DebugNormals      int
	DebugIsometricFOV float64
	DebugIsometricDir int

	PostProcess post.Params
}

const SpecMaxDebugNormals = 2

// TODO: this should honour DebugNormals, etc, not InitDefaults
func (s *Spec) ImageFunc() tracer.ImageFunc {
	return s.Scene().ImageFunc(s.Camera)
}

func (s *Spec) Scene() *tracer.Scene {
	objs := make([]tracer.Object, len(s.Objects))
	for i := range objs {
		objs[i] = s.Objects[i].Interface
	}
	return tracer.NewSceneWithMedia(s.Recursion, s.Media, s.Lights, objs...)
}

// TODO: remove!!!! aargh
func (s *Spec) InitDefaults() {
	if s.Recursion == 0 {
		s.Recursion = 1
	}
	if s.NumPass == 0 {
		s.NumPass = 1
	}
	if s.Camera == nil {
		s.Camera = Projective(90 * Deg).Translate(V(0, 1, 0))
	}
	if s.Width == 0 && s.Height == 0 {
		s.Width = defaultImageWidth
		s.Height = defaultImageHeight
	}
	if s.Width == 0 {
		s.Width = (s.Height * defaultImageWidth) / defaultImageHeight
	}
	if s.Height == 0 {
		s.Height = (s.Width * defaultImageHeight) / defaultImageWidth
	}
	if s.DebugNormals == 1 {
		s.applyMaterial(test.Normal)
	}
	if s.DebugNormals == 2 {
		s.applyMaterial(materials.Blend(
			0.3, test.Normal2,
			0.7, materials.Transparent(C(1, 1, 1), true),
		))
		s.Recursion = 20
	}
	if s.DebugIsometricFOV != 0 {
		s.Camera = cameras.Isometric(s.DebugIsometricDir, s.DebugIsometricFOV)
	}
	if s.Recursion == 0 {
		s.Recursion = defaultRecursion
	}
}

func (s *Spec) applyMaterial(m tracer.Material) {
	orig := s.Objects
	s.Objects = make([]Object, 0, len(orig))
	for _, o := range orig {
		if _, ok := o.Interface.(interface{ IsBackdrop() }); ok {
			s.Objects = append(s.Objects, o.WithMaterial(Flat(Black)))
		} else {
			s.Objects = append(s.Objects, o.WithMaterial(m))
		}
	}
	s.Media = nil
}

const (
	defaultImageWidth  = 1920 / 2
	defaultImageHeight = 1080 / 2
	defaultRecursion   = 3
)
