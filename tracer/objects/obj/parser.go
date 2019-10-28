package obj

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/barnex/bruteray/geom"
)

// TODO: unify with PLY
// one package (objects/parsers)
// standordply.go, wavefrontobj.go

func ParseFile(fname string) (o Obj, e error) {
	defer func() {
		if err, ok := recover().(error); ok {
			o = Obj{}
			e = err
		}
	}()
	p := &parser{obj: Obj{Faces: make(map[string][][]int32)}}
	p.parseFile(fname)
	return p.obj, nil
}

type Obj struct {
	Vertices []geom.Vec
	Faces    map[string][][]int32
}

type parser struct {
	usemtl string // current material set by usemtl
	obj    Obj
}

func (p *parser) parseFile(fname string) {
	f, err := os.Open(fname)
	p.check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		p.parseLine(scanner.Text())
	}
	p.check(scanner.Err())
}

func (p *parser) parseLine(l string) {
	fields := strings.Fields(l)
	if len(fields) == 0 {
		return
	}

	command, args := fields[0], fields[1:]
	switch command {
	default:
		p.errorf("line has unrecognized prefix: %q", l)
	case "v":
		p.parseV(args)
	case "f":
		p.parseF(args)
	case "usemtl":
		p.parseUsemtl(args)
	case "#":
	case "o":
	case "vn":
	case "vt":
	case "s":
	case "l":
	case "g":
	case "mtllib":
	}
}

func (p *parser) parseV(l []string) {
	if !(len(l) == 3 || len(l) == 4) {
		p.errorf("need 3 or 4 vertex coordinates, got %v", len(l))
	}
	var v geom.Vec
	for i := range v {
		v[i] = p.parseFloat64(l[i])
	}
	// we ignore the 4th coordinate, if present
	// in principle we should divide by it.
	p.obj.Vertices = append(p.obj.Vertices, v)
}

func (p *parser) parseF(l []string) {
	if !(len(l) == 3 || len(l) == 4) {
		p.errorf("need 3 or 4 face indices, got %v", len(l))
	}
	// TODO: texture indices are ignored
	f := make([]int32, len(l))
	for i, w := range l {

		words := strings.Split(w, "/")
		idx := int32(p.parseInt(words[0]))
		// TODO: also parse texture and normal indices

		if idx < 1 {
			p.errorf("invalid face index: %v", idx)
		}
		f[i] = idx - 1 // 1-based to 0-based indexing
	}
	m := p.usemtl
	p.obj.Faces[m] = append(p.obj.Faces[m], f)
}

func (p *parser) parseUsemtl(l []string) {
	p.needArgs(l, 1)
	p.usemtl = l[0]
}

func (p *parser) parseFloat32(x string) float32 {
	v, err := strconv.ParseFloat(x, 32)
	p.check(err)
	return float32(v)
}

func (p *parser) parseFloat64(x string) float64 {
	v, err := strconv.ParseFloat(x, 32)
	p.check(err)
	return v
}

func (p *parser) parseInt(x string) int {
	v, err := strconv.ParseInt(x, 10, 32)
	p.check(err)
	return int(v)
}

func (p *parser) needArgs(l []string, n int) {
	if len(l) != n {
		p.errorf("need %v argements, got %v: %q", n, len(l), strings.Join(l, " "))
	}
}

func (p *parser) check(err error) {
	if err != nil {
		p.errorf("%v", err)
	}
}

func (p *parser) errorf(format string, x ...interface{}) {
	panic(fmt.Errorf(format, x...))
}
