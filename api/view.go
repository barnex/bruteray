package api

import (
	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/tracer/cameras"
)

// TODO: embed in spec
// TODO: rename spec.Camera -> spec.CameraType
type View struct {
	Width             int
	Height            int
	AntiAlias         bool
	CamPos            Vec
	CamYaw            float64
	CamPitch          float64
	DebugNormals      int
	DebugIsometric    bool
	DebugIsometricFOV float64
	DebugIsometricDir int
}

func (v *View) ApplyTo(s Spec) Spec {
	s.Width = v.Width
	s.Height = v.Height
	s.Camera = cameras.Transform(s.Camera, geom.YawPitchRoll(v.CamYaw, v.CamPitch, 0).A, v.CamPos)

	s.DebugNormals = v.DebugNormals

	if v.DebugIsometric {
		s.DebugIsometricFOV = v.DebugIsometricFOV
		s.DebugIsometricDir = v.DebugIsometricDir
	} else {

		s.DebugIsometricFOV = 0
		s.DebugIsometricDir = 0
	}

	s.InitDefaults() // aaargh
	return s
}
