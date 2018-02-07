package bruteray

import "math"

func (e *Env) NewRay(start, dir Vec) *Ray {
	r := &Ray{Start: start}
	r.SetDir(dir)
	return r
}

func (e *Env) RRay(r *Ray) {

}

func (r *Ray) Dir() Vec {
	return r.d
}

func (r *Ray) SetDir(dir Vec) {
	r.d = dir
	r.InvDir = Vec{1 / dir[X], 1 / dir[Y], 1 / dir[Z]}
}

// A Ray is a half-line,
// starting at the Start point (exclusive) and extending in direction Dir.
type Ray struct {
	Start  Vec
	d      Vec
	InvDir Vec // pre-calculated inverse direction for marginal speed improvements
}

// Returns point Start + t*Dir.
// t must be > 0 for the point to lie on the Ray.
func (r *Ray) At(t float64) Vec {
	if math.IsNaN(t) {
		panic(t)
	}
	return r.Start.Add(r.d.Mul(t))
}

func (r *Ray) Transf(t *Matrix4) {
	r.Start = t.TransfPoint(r.Start)
	r.SetDir(t.TransfDir(r.d))
}
