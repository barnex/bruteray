package builder

import (
	"math"

	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
	"github.com/barnex/bruteray/v2/util"
)

//TODO: move to geom, Pitch, Roll, Yaw, MoveTo should act on box

type BoundingBox struct {
	Min, Max Vec
}

//const BBOffset = 1. / (4096.)

// MakeBoundingBox constructs the minimal axis-aligned bounding box
// that countains all points in hull.
func MakeBoundingBox(children []Builder) BoundingBox {
	var hull []Vec
	for _, c := range children {
		cbb := c.Bounds()
		hull = append(hull, cbb.Min, cbb.Max)
	}
	return BoundingBoxFromHull(hull)
}

func BoundingBoxFromHull(hull []Vec) BoundingBox {
	Min := hull[0]
	Max := hull[0]
	for _, v := range hull {
		util.CheckNaNVec(v)
		for i := 0; i < 3; i++ {
			if v[i] < Min[i] {
				Min[i] = v[i]
			}
			if v[i] > Max[i] {
				Max[i] = v[i]
			}
		}
	}
	util.CheckNaNVec(Min)
	util.CheckNaNVec(Max)
	return BoundingBox{Min, Max}
}

func infBox() BoundingBox {
	return BoundingBox{Vec{-Inf, -Inf, -Inf}, Vec{Inf, Inf, Inf}}
}

func (b BoundingBox) Center() Vec {
	return b.Min.Add(b.Max).Mul(0.5)
}

func (b BoundingBox) CenterBottom() Vec {
	c := b.Min.Add(b.Max).Mul(0.5)
	c[Y] = b.Min[Y]
	return c
}

func nan(x float64) bool {
	return math.IsNaN(x) || math.IsInf(x, 1) || math.IsInf(x, -1)
}

func (s *BoundingBox) Intersect(r *Ray) float64 {
	invdirx := 1 / r.Dir[X]
	invdiry := 1 / r.Dir[Y]
	invdirz := 1 / r.Dir[Z]

	// manual inlining for ~3x speed-up
	tminx := (s.Min[X] - r.Start[X]) * invdirx
	tminy := (s.Min[Y] - r.Start[Y]) * invdiry
	tminz := (s.Min[Z] - r.Start[Z]) * invdirz
	tmaxx := (s.Max[X] - r.Start[X]) * invdirx
	tmaxy := (s.Max[Y] - r.Start[Y]) * invdiry
	tmaxz := (s.Max[Z] - r.Start[Z]) * invdirz

	txen := min(tminx, tmaxx)
	txex := max(tminx, tmaxx)
	tyen := min(tminy, tmaxy)
	tyex := max(tminy, tmaxy)
	tzen := min(tminz, tmaxz)
	tzex := max(tminz, tmaxz)
	ten := max3(txen, tyen, tzen)
	tex := min3(txex, tyex, tzex)

	if ten > tex {
		return 0
	}
	if ten < 0 {
		return tex
	}
	return ten
}
