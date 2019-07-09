package texture

import (
	"github.com/barnex/bruteray/v2/color"
	"github.com/barnex/bruteray/v2/geom"
)

// Texture maps each point in 3D space to a color.
type Texture interface {
	At(p geom.Vec) color.Color
}
