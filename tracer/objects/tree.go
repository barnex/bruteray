package objects

import (
	"sort"

	"github.com/barnex/bruteray/geom"
	. "github.com/barnex/bruteray/tracer/types"
	"github.com/barnex/bruteray/util"
)

const maxFacePerLeaf = 7

// Tree returns a Bounding Volume Hierarchy containing the given objects,
// which as an efficient Intersect method.
func Tree(objects ...Interface) Interface {
	return &tree{
		root: buildTree(objects),
	}
}

type tree struct {
	root node
}

type node struct {
	boundingBox boundingBoxf
	children    *[2]node
	leafs       []Interface
}

// TODO: Currently we divide along the longest dimension
// and always cut right in the middle.
// This should be replaced by the Surface Area Heuristic.
func buildTree(ch []Interface) node {
	if len(ch) <= maxFacePerLeaf {
		return node{
			leafs:       ch,
			boundingBox: boundingBoxToF(makeBoundingBox(ch)),
		}
	}

	bb := makeBoundingBox(ch)
	bbSize := bb.Max.Sub(bb.Min)
	splitDir := argMax(bbSize)
	sort.Slice(ch, func(i, j int) bool {
		bi := ch[i].Bounds()
		bj := ch[j].Bounds()
		return bi.Min[splitDir]+bi.Max[splitDir] < bj.Min[splitDir]+bj.Max[splitDir]
	})
	N := len(ch)
	left := ch[:N/2]
	right := ch[N/2:]
	return node{
		children:    &[2]node{buildTree(left), buildTree(right)},
		boundingBox: boundingBoxToF(bb),
	}
}

func (t *tree) Intersect(r *Ray) HitRecord {
	if intersectAABBf(&t.root.boundingBox, r) <= 0 { // TODO: and ray Len
		return HitRecord{}
	}
	return t.root.Intersect(r)
}

func (t *tree) Inside(p Vec) bool {
	return t.root.Inside(p)
}

func (t *node) Inside(p Vec) bool {
	if t.children != nil {
		if t.children[0].Inside(p) {
			return true
		}
		if t.children[1].Inside(p) {
			return true
		}
	}
	for _, l := range t.leafs {
		if l.Inside(p) {
			return true
		}
	}
	return false
}

func (n *node) Intersect(r *Ray) HitRecord {
	// TODO: use tracer.FrontSolution
	front := HitRecord{T: 9e99}
	if n.children == nil {
		for _, o := range n.leafs {
			frag := o.Intersect(r)
			if frag.T > 0 && frag.T < front.T {
				front = frag
			}
		}
		return front
	}

	ch0, ch1 := &(n.children[0]), &(n.children[1])
	t0 := intersectAABBf(&ch0.boundingBox, r)
	t1 := intersectAABBf(&ch1.boundingBox, r)

	//if t1 < t0 {
	//	ch0, ch1 = ch1, ch0
	//	t0, t1 = t1, t0
	//}

	if t0 > 0 {
		frag := ch0.Intersect(r)
		if frag.T > 0 { //&& frag.T < front.T {
			front = frag
		}
	}
	if t1 > 0 { //&& !(t1 > front.T+0*Tiny) {
		frag := ch1.Intersect(r)
		if frag.T > 0 && frag.T < front.T {
			front = frag
		}
	}

	if front.T == 9e99 {
		front.T = 0
	}
	return front
}

func intersectAABB(s *BoundingBox, r *Ray) float64 {
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

// makeBoundingBox constructs the minimal axis-aligned bounding box
// that countains all points in hull.
// TODO: do not use intermediate memory
func makeBoundingBox(children []Interface) BoundingBox {
	var hull []Vec
	for _, c := range children {
		cbb := c.Bounds()
		hull = append(hull, cbb.Min, cbb.Max)
	}
	return boundingBoxFromHull(hull)
}

func (n *tree) Bounds() BoundingBox {
	return n.root.boundingBox.to64()
}

type boundingBoxf struct {
	Min, Max [3]float32
}

func boundingBoxToF(b BoundingBox) boundingBoxf {
	return boundingBoxf{
		Min: vecf(&b.Min),
		Max: vecf(&b.Max),
	}
}

func (b *boundingBoxf) to64() BoundingBox {
	return BoundingBox{
		Min: toVecf(b.Min),
		Max: toVecf(b.Max),
	}
}

func vecf(a *geom.Vec) [3]float32 {
	return [3]float32{
		float32(a[0]),
		float32(a[1]),
		float32(a[2]),
	}
}

func toVecf(a [3]float32) Vec {
	return Vec{
		float64(a[0]),
		float64(a[1]),
		float64(a[2]),
	}
}

func intersectAABBf(s *boundingBoxf, r *Ray) float64 {
	idirx := 1 / r.Dir[X]
	idiry := 1 / r.Dir[Y]
	idirz := 1 / r.Dir[Z]

	startx := r.Start[X]
	starty := r.Start[Y]
	startz := r.Start[Z]

	minx := float64(s.Min[X])
	miny := float64(s.Min[Y])
	minz := float64(s.Min[Z])
	maxx := float64(s.Max[X])
	maxy := float64(s.Max[Y])
	maxz := float64(s.Max[Z])

	tminx := (minx - startx) * idirx
	tmaxx := (maxx - startx) * idirx
	tminy := (miny - starty) * idiry
	tmaxy := (maxy - starty) * idiry
	tminz := (minz - startz) * idirz
	tmaxz := (maxz - startz) * idirz

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

// TODO: move to isosurface
func intersectAABB2(s *BoundingBox, r *Ray) (float64, float64) {
	invdirx := 1 / r.Dir[X]
	invdiry := 1 / r.Dir[Y]
	invdirz := 1 / r.Dir[Z]

	startx := r.Start[X]
	starty := r.Start[Y]
	startz := r.Start[Z]

	tminx := (s.Min[X] - startx) * invdirx
	tminy := (s.Min[Y] - starty) * invdiry
	tminz := (s.Min[Z] - startz) * invdirz

	tmaxx := (s.Max[X] - startx) * invdirx
	tmaxy := (s.Max[Y] - starty) * invdiry
	tmaxz := (s.Max[Z] - startz) * invdirz

	txen := util.Min(tminx, tmaxx)
	txex := util.Max(tminx, tmaxx)
	tyen := util.Min(tminy, tmaxy)
	tyex := util.Max(tminy, tmaxy)
	tzen := util.Min(tminz, tmaxz)
	tzex := util.Max(tminz, tmaxz)
	ten := max3(txen, tyen, tzen)
	tex := min3(txex, tyex, tzex)

	if ten > tex {
		return 0, 0
	}
	return ten, tex
}

func argMax(v Vec) int {
	I := 0
	max := v[0]
	for i, v := range v {
		if v > max {
			max = v
			I = i
		}
	}
	return I
}
