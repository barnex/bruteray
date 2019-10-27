package objects

import (
	"math"

	. "github.com/barnex/bruteray/tracer/types"
	"github.com/barnex/bruteray/util"
)

// BoundingBox is an Axis Aligned box, used to accelerate intersection tests with groups of objects.
// See https://en.wikipedia.org/wiki/Minimum_bounding_box#Axis-aligned_minimum_bounding_box.
type BoundingBox struct {
	Min, Max Vec
}

// boundingBoxFromHull constructs a bounding box that contains the given points.
func boundingBoxFromHull(hull []Vec) BoundingBox {
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

// translated returns a translated copy of b.
func (b BoundingBox) translated(delta Vec) BoundingBox {
	return BoundingBox{
		Min: b.Min.Add(delta),
		Max: b.Max.Add(delta),
	}
}

// withMargin returns a copy of b that is enlarged by distance delta
// in all directions.
func (b BoundingBox) withMargin(delta float64) BoundingBox {
	return BoundingBox{
		Min: b.Min.Add(Vec{-delta, -delta, -delta}),
		Max: b.Max.Add(Vec{+delta, +delta, +delta}),
	}
}

func (b BoundingBox) Size() Vec   { return b.Max.Sub(b.Min) }
func (b BoundingBox) Dx() float64 { return b.Size()[X] }
func (b BoundingBox) Dy() float64 { return b.Size()[Y] }
func (b BoundingBox) Dz() float64 { return b.Size()[Z] }

// inside returns true if point p lies inside b.
func (b *BoundingBox) inside(p Vec) bool {
	return p[0] > b.Min[0] && p[0] < b.Max[0] &&
		p[1] > b.Min[1] && p[1] < b.Max[1] &&
		p[2] > b.Min[2] && p[2] < b.Max[2]
}

func (b BoundingBox) Center() Vec {
	return safeAverage(b.Min, b.Max)
}

func (b BoundingBox) CenterBottom() Vec {
	c := safeAverage(b.Min, b.Max)
	c[Y] = b.Min[Y]
	return c
}

func (b BoundingBox) CenterTop() Vec {
	c := safeAverage(b.Min, b.Max)
	c[Y] = b.Max[Y]
	return c
}

func (b BoundingBox) CenterBack() Vec {
	c := safeAverage(b.Min, b.Max)
	c[Z] = b.Min[Z]
	return c
}

func (b BoundingBox) CenterFront() Vec {
	c := safeAverage(b.Min, b.Max)
	c[Z] = b.Max[Z]
	return c
}

// average of a and b, but return 0 for inf - inf.
func safeAverage(a, b Vec) Vec {
	avg := a.Add(b).Mul(0.5)
	for i := range avg {
		if math.IsInf(a[i], -1) && math.IsInf(b[i], 1) {
			avg[i] = 0
		}
	}
	return avg
}

func (s *BoundingBox) intersect(r *Ray) float64 {
	idirx := 1 / r.Dir[X]
	idiry := 1 / r.Dir[Y]
	idirz := 1 / r.Dir[Z]

	startx := r.Start[X]
	starty := r.Start[Y]
	startz := r.Start[Z]

	tminx := (s.Min[X] - startx) * idirx
	tmaxx := (s.Max[X] - startx) * idirx

	tminy := (s.Min[Y] - starty) * idiry
	tmaxy := (s.Max[Y] - starty) * idiry

	tminz := (s.Min[Z] - startz) * idirz
	tmaxz := (s.Max[Z] - startz) * idirz

	txen := util.Min(tminx, tmaxx)
	txex := util.Max(tminx, tmaxx)
	tyen := util.Min(tminy, tmaxy)
	tyex := util.Max(tminy, tmaxy)
	tzen := util.Min(tminz, tmaxz)
	tzex := util.Max(tminz, tmaxz)
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

func (s *BoundingBox) intersects(r *Ray) bool {
	return s.intersect(r) > 0
}

func max3(x, y, z float64) float64 {
	max := x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}

func min3(x, y, z float64) float64 {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}

var (
	infBox = BoundingBox{Min: Vec{-inf, -inf, -inf}, Max: Vec{inf, inf, inf}}
	inf    = math.Inf(1)
)
