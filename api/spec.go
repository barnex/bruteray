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
	Camera  Camera

	Recursion int
	NumPass   int

	Width  int
	Height int

	DebugNormals      bool
	DebugIsometricFOV float64
	DebugIsometricDir int
}

func (s *Spec) imageFunc() tracer.ImageFunc {
	objs := make([]tracer.Object, len(s.Objects))
	for i := range objs {
		objs[i] = s.Objects[i]
	}
	scene := tracer.NewSceneWithMedia(s.Media, s.Lights, objs...)
	return scene.ImageFunc(s.Camera, s.Recursion)
}

func initDefaults(s *Spec) {
	if s.Recursion == 0 {
		s.Recursion = 1
	}
	if s.NumPass == 0 {
		s.NumPass = 1
	}
	if s.Camera == nil {
		s.Camera = Projective(90*Deg, Vec{0, 1, 0}, 0, 0, 0)
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
		for i, o := range s.Objects {
			s.Objects[i] = o.WithMaterial(test.Normal)
		}
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
