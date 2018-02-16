package br

type Pool struct {
	New func() interface{}
	p   []interface{}
}

func (p *Pool) Get() interface{} {
	if len(p.p) == 0 {
		return p.New()
	}
	fb := p.p[len(p.p)-1]
	p.p = p.p[:len(p.p)-1]
	return fb
}

func (p *Pool) Put(v interface{}) {
	p.p = append(p.p, v)
}
