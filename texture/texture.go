package texture

import (
	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/imagef/colorf"
)

// Texture maps each point in 3D space to a color.
type Texture interface {
	At(p geom.Vec) colorf.Color
}
