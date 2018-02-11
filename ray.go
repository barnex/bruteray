package bruteray

func (e *Env) NewRay(start, dir Vec) *Ray {
	//r := &Ray{Start: start}
	//r.SetDir(dir)
	//return r
	r := e.rayPool.Get().(*Ray)
	r.Start = start
	r.SetDir(dir)
	return r
}

func (e *Env) RRay(r *Ray) {
	e.rayPool.Put(r)
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
	//return r.Start.Add(r.d.Mul(t))

	// pprof shows this is where we spend most of our time.
	// manually inlined for ~10% overall performance improvement.
	return Vec{
		r.Start[X] + t*r.d[X],
		r.Start[Y] + t*r.d[Y],
		r.Start[Z] + t*r.d[Z],
	}
}

func (r *Ray) Transf(t *Matrix4) {
	r.Start = t.TransfPoint(r.Start)
	r.SetDir(t.TransfDir(r.d))
}
