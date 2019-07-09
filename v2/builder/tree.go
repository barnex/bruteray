package builder

import (
	"sort"

	. "github.com/barnex/bruteray/v2/color"
	. "github.com/barnex/bruteray/v2/geom"
	. "github.com/barnex/bruteray/v2/tracer"
)

type Tree struct {
	NoDivide bool
	Children []Builder
	Lights   []LightBuilder
	root     node
}

type LightBuilder interface {
	Builder
	Sample(ctx *Ctx, target Vec) (pos Vec, intens Color) // Light
}

func NewTree(o ...Builder) *Tree {
	t := &Tree{}
	for _, o := range o {
		t.Add(o)
	}
	return t
}

func (n *Tree) Add(ch ...Builder) {
	for _, ch := range ch {
		switch ch := ch.(type) {
		case LightBuilder:
			n.Lights = append(n.Lights, ch)
		case *Tree:
			n.Lights = append(n.Lights, ch.Lights...)
			n.Children = append(n.Children, ch)
			ch.Lights = nil // hack so we don't transform lights twice
		default:
			n.Children = append(n.Children, ch)
		}
	}
}

func (n *Tree) Bounds() BoundingBox {
	var h []Vec
	for _, c := range n.Children {
		cbb := c.Bounds()
		h = append(h, cbb.Min, cbb.Max)
	}
	return BoundingBoxFromHull(h)
}

func (n *Tree) CtrlPoints() []*Vec {
	var p []*Vec
	for _, c := range n.Children {
		p = append(p, c.(Transformable).CtrlPoints()...)
	}
	for _, c := range n.Lights {
		p = append(p, c.(Transformable).CtrlPoints()...)
	}
	return p
}

func (n *Tree) Init() {
	if n.NoDivide {
		n.root = makeLeaf(n.Children)
	} else {
		n.root = buildTree(n.Children)
	}
	for _, ch := range n.Lights {
		ch.Init()
	}
	for _, ch := range n.Children {
		ch.Init()
	}
}

func buildTree(ch []Builder) node {
	if len(ch) <= 4 { // TODO: tune me?
		return makeLeaf(ch)
	}

	bb := MakeBoundingBox(ch)
	bbSize := bb.Max.Sub(bb.Min)
	splitDir := argMax(bbSize)
	// TODO: this is rudimentary and slow
	sort.Slice(ch, func(i, j int) bool { return ch[i].Bounds().Min[splitDir] < ch[j].Bounds().Min[splitDir] })
	N := len(ch)
	left := ch[:N/2]
	right := ch[N/2:]
	return node{
		children:    &[2]node{buildTree(left), buildTree(right)},
		boundingBox: bb,
	}
}

func makeLeaf(ch []Builder) node {
	l := make([]Object, len(ch))
	for i, ch := range ch {
		l[i] = ch
	}
	return node{
		leafs:       l,
		boundingBox: MakeBoundingBox(ch),
	}
}

func (t *Tree) Intersect(ctx *Ctx, r *Ray) HitRecord {
	if t.root.boundingBox.Intersect(r) <= 0 { // TODO: and ray Len
		return HitRecord{}
	}
	return t.root.Intersect(ctx, r)
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

type node struct {
	boundingBox BoundingBox
	children    *[2]node
	leafs       []Object
}

var _ Object = (*node)(nil)

// TODO: ray len
func (n *node) Intersect(ctx *Ctx, r *Ray) HitRecord {

	//if n.boundingBox.Intersect(r) <= 0 { // TODO: and ray Len
	//	return HitRecord{}
	//}

	front := HitRecord{T: 9e99}

	for _, o := range n.leafs {
		frag := o.Intersect(ctx, r)
		if frag.T > 0 && frag.T < front.T {
			front = frag
		}
	}

	if n.children == nil {
		return front
	}

	ch0, ch1 := &(n.children[0]), &(n.children[1])
	t0 := ch0.boundingBox.Intersect(r)
	t1 := ch1.boundingBox.Intersect(r)

	//if t1 < t0 {
	//	ch0, ch1 = ch1, ch0
	//	t0, t1 = t1, t0
	//}

	if t0 > 0 {
		frag := ch0.Intersect(ctx, r)
		if frag.T > 0 && frag.T < front.T {
			front = frag
		}
	}
	if t1 > 0 { //&& t1 <= front.T+1e-6 {
		frag := ch1.Intersect(ctx, r)
		if frag.T > 0 && frag.T < front.T {
			front = frag
		}
	}

	if front.T == 9e99 {
		front.T = 0
	}

	return front
}
