// Package types provides constants and type aliases that are considered bruteray primitives.
// It is intended to be dot-imported:
//	import . "github.com/barnex/bruteray/tracer/types"
package types

import (
	"github.com/barnex/bruteray/color"
	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/tracer"
)

type Camera = tracer.Camera
type Color = color.Color
type Ctx = tracer.Ctx
type HitCoords = tracer.HitCoords
type HitRecord = tracer.HitRecord
type ImageFunc = tracer.ImageFunc
type Light = tracer.Light
type Material = tracer.Material
type Object = tracer.Object
type Medium = tracer.Medium
type Ray = tracer.Ray
type Scene = tracer.Scene
type Vec = geom.Vec
type Vec2 = geom.Vec2

var (
	NewScene          = tracer.NewScene
	NewSceneWithMedia = tracer.NewSceneWithMedia
	O                 = geom.O
	Ex                = geom.Ex
	Ey                = geom.Ey
	Ez                = geom.Ez
)

const (
	Tiny = tracer.Tiny
	Deg  = geom.Deg
	Pi   = geom.Pi
	X    = geom.X
	Y    = geom.Y
	Z    = geom.Z
)
