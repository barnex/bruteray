package objects

/*
import (
	"github.com/barnex/bruteray/geom"
	. "github.com/barnex/bruteray/tracer/types"
)

type T func(interface{}) func(Vec) Vec

func Translate(delta Vec) T {
	return func(interface{}) func(Vec) Vec {
		return geom.Translate(delta).TransformPoint
	}
}

func Rotate(axis Vec, angle float64) T {
	return func(obj interface{}) func(Vec) Vec {
		o := origin(obj)
		return geom.Rotate(o, axis, angle).TransformPoint
	}
}

func origin(obj interface{}) Vec {
	switch obj := obj.(type) {
	default:
		return Vec{}
	case originer:
		return obj.Origin()
	case bounded:
		return obj.Bounds().Center()
	}
}

//type bounded interface {
//	Bounds() BoundingBox
//}

type originer interface {
	Origin() Vec
}
*/
