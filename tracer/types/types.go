// Package types provides constants and type aliases that are considered bruteray primitives.
// It is intended to be dot-imported:
//	import . "github.com/barnex/bruteray/tracer/types"
package types

import (
	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/imagef"
	"github.com/barnex/bruteray/imagef/colorf"
	"github.com/barnex/bruteray/tracer"
)

type Camera = tracer.Camera
type Color = colorf.Color
type Ctx = tracer.Ctx
type HitCoords = tracer.HitCoords
type HitRecord = tracer.HitRecord
type Image = imagef.Image
type ImageFunc = tracer.ImageFunc
type Light = tracer.Light
type Material = tracer.Material
type Medium = tracer.Medium
type Object = tracer.Object
type Ray = tracer.Ray
type Scene = tracer.Scene
type Vec = geom.Vec
type Vec2 = geom.Vec2

var (
	O                 = geom.O
	Ex                = geom.Ex
	Ey                = geom.Ey
	Ez                = geom.Ez
	Inf               = geom.Inf
	NewScene          = tracer.NewScene
	NewSceneWithMedia = tracer.NewSceneWithMedia
	MakeImage         = imagef.MakeImage
)

const (
	Deg  = geom.Deg
	Pi   = geom.Pi
	Tiny = tracer.Tiny
	X    = geom.X
	Y    = geom.Y
	Z    = geom.Z
)
