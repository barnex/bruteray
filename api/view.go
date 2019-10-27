package api

import (
	"github.com/barnex/bruteray/tracer/cameras"
)

// TODO: embed in spec
// TODO: rename spec.Camera -> spec.CameraType
type View struct {
	Width        int
	Height       int
	AntiAlias    bool
	CamPos       Vec
	CamYaw       float64
	CamPitch     float64
	DebugNormals bool
}

func (v *View) ApplyTo(s Spec) Spec {
	s.Width = v.Width
	s.Height = v.Height
	s.Camera = cameras.YawPitchRoll(s.Camera, v.CamYaw, v.CamPitch, 0)
	s.Camera = cameras.Translate(s.Camera, v.CamPos)
	s.DebugNormals = v.DebugNormals
	s.InitDefaults() // aaargh
	return s
}
