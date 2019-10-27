package api

import (
	"github.com/barnex/bruteray/tracer"
	"github.com/barnex/bruteray/tracer/cameras"
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

	DebugNormals      bool
	DebugIsometricFOV float64
	DebugIsometricDir int
}

// TODO: this should honour DebugNormals, etc, not InitDefaults
func (s *Spec) ImageFunc() tracer.ImageFunc {
	objs := make([]tracer.Object, len(s.Objects))
	for i := range objs {
		objs[i] = s.Objects[i].Interface
	}
	scene := tracer.NewSceneWithMedia(s.Media, s.Lights, objs...)
	return scene.ImageFunc(s.Camera, s.Recursion)
}

func (s *Spec) InitDefaults() {
	if s.Recursion == 0 {
		s.Recursion = 1
	}
	if s.NumPass == 0 {
		s.NumPass = 1
	}
	if s.Camera == nil {
		s.Camera = Projective(90*Deg, Vec{0, 1, 0}, 0, 0)
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
	if s.DebugNormals {
		orig := s.Objects
		s.Objects = make([]Object, len(orig))
		for i := range s.Objects {
			s.Objects[i] = orig[i].WithMaterial(test.Normal)
		}
		s.Media = nil
	}
	if s.DebugIsometricFOV != 0 {
		s.Camera = cameras.NewIsometric(s.DebugIsometricDir, s.DebugIsometricFOV)
	}
	if s.Recursion == 0 {
		s.Recursion = defaultRecursion
	}
}

const (
	defaultImageWidth  = 1920 / 2
	defaultImageHeight = 1080 / 2
	defaultRecursion   = 3
)
