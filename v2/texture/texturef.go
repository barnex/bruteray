package texture

import "github.com/barnex/bruteray/v2/geom"

// Texturef maps each point in 3D space to a number.
// TODO: remove?
type Texturef interface {
	At(p geom.Vec) float64
}
