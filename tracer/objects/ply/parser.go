package ply

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/barnex/bruteray/geom"
)

func ParseFile(fname string) ([]geom.Vec, [][3]int, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	return Parse(f)
}

func Parse(r io.Reader) (vertices []geom.Vec, faces [][3]int, e error) {
	defer func() {
		if p := recover(); p != nil {
			e = errors.New(fmt.Sprint(p))
			p = nil
		}
	}()

	parser := parser{b: bufio.NewReader(r)}
	v, f := parser.parse()
	return v, f, nil
}

func (p *parser) parse() ([]geom.Vec, [][3]int) {
	if magic := p.readLine(); magic != "ply" {
		p.panicf("bad header: %q", magic)
	}
	if format := p.readLine(); format != "format ascii 1.0" {
		p.panicf("bad header: %q", format)
	}

	var numVertex, numFace int
	for l := p.readLine(); l != "end_header"; l = p.readLine() {
		switch {
		case strings.HasPrefix(l, "element vertex "):
			numVertex = p.atoi(l[len("element vertex "):])
		case strings.HasPrefix(l, "element face "):
			numFace = p.atoi(l[len("element face "):])
		case l == "end_header":
			break
		}
	}

	vertex := make([]geom.Vec, numVertex)
	for i := 0; i < numVertex; i++ {
		fields := strings.Fields(p.readLine())
		for c := 0; c < 3; c++ {
			vertex[i][c] = p.atof(fields[c])
		}
	}

	faces := make([][3]int, numFace)
	for i := 0; i < numFace; i++ {
		fields := strings.Fields(p.readLine())
		if n := p.atoi(fields[0]); n != 3 {
			p.panicf("need 3 vertices, have: %v", n)
		}
		for c := 0; c < 3; c++ {
			idx := p.atoi(fields[c+1])
			faces[i][c] = idx
		}
	}
	return vertex, faces
}

type parser struct {
	b        *bufio.Reader
	lastLine string
}

func (p *parser) readLine() string {
	line, err := p.b.ReadString('\n')
	p.check(err)
	p.lastLine = line
	return strings.TrimSpace(line)
}

func (p *parser) check(e error) {
	if e != nil {
		p.panicf("readline: %v", e)
	}
}

func (p *parser) atoi(s string) int {
	i, err := strconv.Atoi(s)
	p.check(err)
	return i
}

func (p *parser) atof(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	p.check(err)
	return i
}

func (p *parser) panicf(format string, x ...interface{}) {
	panic(fmt.Sprintf(p.lastLine+": "+format, x...))
}
