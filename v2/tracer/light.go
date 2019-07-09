package tracer

import (
	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
)

type Light interface {

	// Sample returns a position on the light's surface
	// (used to determine shadows for extended sources),
	// and the intensity at given target position.
	Sample(ctx *Ctx, target Vec) (pos Vec, intens Color)

	// Object is rendered when the camera looks at the light directly.
	Object
}
