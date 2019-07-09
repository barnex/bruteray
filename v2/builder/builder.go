package builder

import (
	"github.com/barnex/bruteray/v2/geom"
	"github.com/barnex/bruteray/v2/tracer"
	"github.com/barnex/bruteray/v2/util"
)

type Builder interface {
	tracer.Object
	Init()
	Bounds() BoundingBox // TODO: rename: Bounds
}

func ScaleToSize(s interface {
	geom.Transformable
	Bounds() BoundingBox
}, size float64) {
	bounds := s.Bounds()
	bs := bounds.Max.Sub(bounds.Min)
	max := util.Max3(bs[0], bs[1], bs[2])
	geom.Scale(s, bounds.Center(), 1/max)
}
